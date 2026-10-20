# Паттерны отказоустойчивости и Graceful Shutdown в Go

## О чём эта тема (краткий контекст)

В распределённой системе внешние зависимости неизбежно тормозят и падают, а
клиентский трафик приходит неравномерными всплесками. Чтобы локальный сбой
одного микросервиса не вызвал лавинообразное падение всей платформы (cascading
failure), сервис защищают таймаутами, ограничителями скорости (Rate Limiting),
предохранителями (Circuit Breaker), умными повторами (Retries с Backoff &
Jitter) и корректным завершением работы (Graceful Shutdown).

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [Балансировка нагрузки, Health Checks, Rate Limiting, Circuit Breaker, Retries и Graceful Shutdown](../../../../5_devops_sre_knowledgebase/14_system_design/03_distributed_components_and_data/01_theory/03_load_balancing_and_resilience_patterns.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

## 1. Rate Limiting в Go: Token Bucket и `golang.org/x/time/rate`

Стандарт де-факто для локального (in-process) ограничения RPS в Go — пакет
`golang.org/x/time/rate`, реализующий алгоритм **Token Bucket** со сложностью
$O(1)$ и без фоновых горутин-таймеров (токены пересчитываются лениво по разнице
времени `time.Now()` в момент вызова). См. собственную реализацию в примере
`02_examples/07_token_bucket/main.go`.

У `rate.Limiter` три режима работы:
- `limiter.Allow()` — мгновенный неблокирующий ответ `bool` (используется во
  входном HTTP middleware, чтобы сразу ответить `429 Too Many Requests`).
- `limiter.Wait(ctx)` — блокирует текущую горутину до появления токена, но
  мгновенно возвращает ошибку, если `ctx` уже истёк или время ожидания
  превышает дедлайн `ctx`! (Идеально для исходящих клиентов к внешнему API).
- `limiter.Reserve()` — резервирует слот и сообщает, сколько времени нужно
  подождать.

> ⚠️ **Ловушка утечки памяти в per-IP / per-User Rate Limiter:** если хранить
> лимитеры в `map[string]*rate.Limiter` по IP-адресу клиента, то за сутки
> миллионы уникальных IP навсегда осядут в памяти. Обязательно храните рядом
> `lastSeen time.Time` и раз в минуту чистите неактивные записи фоновой
  горутиной (или используйте LRU-кэш фиксированного размера).

## 2. Circuit Breaker в Go (`Closed` → `Open` → `Half-Open`)

Когда внешний сервис упал или отвечает по таймауту 5 секунд, бесполезно держать
сотни горутин заблокированными в ожидании заведомо мёртвого сокета. Паттерн
**Circuit Breaker** (см. полную реализацию конечного автомата на `sync.Mutex` в
примере `02_examples/08_circuit_breaker/main.go`, а в продакшене — библиотеки
`sony/gobreaker` или `failsafe-go`) размыкает цепь после $N$ ошибок подряд и
мгновенно возвращает `ErrCircuitOpen` за наносекунды (`fail-fast`), периодически
пропуская пробный запрос в состоянии `Half-Open`.

## 3. Контекстно-осведомлённые Retries с Exponential Backoff и Full Jitter

Повторять неудавшийся сетевой запрос можно **только для идемпотентных операций**
и **только для временных (transient) ошибок** (`502/503/504`, сетевой reset), но
никогда — для `400 Bad Request`, `401 Unauthorized` или отмены контекста.

Чтобы все клиенты после сбоя не ударили в оживающий сервер синхронными волнами
(Retry Storm), экспоненциальная пауза обязательно размывается случайным шумом
(**Full Jitter**), а ожидание между попытками слушает `<-ctx.Done()`:

```go
func RetryWithJitter(ctx context.Context, maxAttempts int, base time.Duration, fn func(context.Context) error) error {
	var err error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if err = fn(ctx); err == nil {
			return nil
		}
		if ctx.Err() != nil {
			return ctx.Err() // Контекст уже отменён — не ретраим!
		}
		if attempt == maxAttempts-1 {
			break
		}
		// Exponential backoff: base * 2^attempt + Full Jitter [0, sleep)
		maxSleep := base * (1 << attempt)
		jitter := time.Duration(rand.Int63n(int64(maxSleep)))

		timer := time.NewTimer(jitter)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
	return err
}
```

## 4. Эталонный Graceful Shutdown Go-сервиса в Kubernetes

При деплое новой версии или скейлинге вниз Kubernetes посылает поду сигнал
`SIGTERM` и одновременно начинает удалять IP пода из Endpoints балансировщика
(эти два процесса идут **параллельно**!). Если Go-сервис просто завершится по
`SIGTERM` в ту же миллисекунду, он оборвёт все летящие запросы (`in-flight`) и
транзакции в БД.

Правильная последовательность остановки Go-сервиса (см. пример
`02_examples/10_graceful_shutdown/main.go`):

```go
// 1. Подписываемся на SIGINT и SIGTERM через signal.NotifyContext.
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

go func() {
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server failed", "err", err)
	}
}()

// 2. Ждём сигнала остановки от ОС / Kubernetes.
<-ctx.Done()
slog.Info("shutting down gracefully...")

// 3. Переводим readiness-пробу (/readyz) в 503, чтобы балансировщик перестал слать трафик.
isShuttingDown.Store(true)

// 4. Даём серверу ограниченное время (например, 15s < terminationGracePeriodSeconds)
// на завершение текущих HTTP-запросов; новые соединения уже не принимаются.
shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()
_ = srv.Shutdown(shutdownCtx)

// 5. Дожидаемся завершения фоновых воркеров (WaitGroup) и только В КОНЦЕ закрываем пул БД!
close(jobsChan)
workersWG.Wait()
_ = db.Close()
```

> ⚠️ **Критическая ошибка порядка остановки:** никогда не закрывайте
> `db.Close()`
> или клиент Redis **до** того, как завершился `srv.Shutdown(shutdownCtx)` и
> остановились фоновые воркеры (`workersWG.Wait()`)! Иначе запросы, которые
> честно пытаются доработать во время graceful shutdown, упадут с ошибкой `sql:
> database is closed`.

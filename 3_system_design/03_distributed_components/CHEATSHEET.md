# Распределённые компоненты, очереди и отказоустойчивость в Go — шпаргалка

Подробности — в [01_theory/](./01_theory/) и каноническом разделе
[DevOps System Design](../../../5_devops_sre_knowledgebase/14_system_design/03_distributed_components_and_data/README.md).
⚠️ — частая ловушка.

## Пул соединений `database/sql` и шардирование

Подробно:
[01_db_pools_and_sharding_in_go.md](./01_theory/01_db_pools_and_sharding_in_go.md).

| Настройка / Метод `*sql.DB` | Дефолт | Рекомендация для продакшена |
| :--- | :--- | :--- |
| `db.SetMaxOpenConns(n)` | `0` (без лимита!) | `20–50` на под (с учётом числа реплик пода и PgBouncer) |
| `db.SetMaxIdleConns(n)` | **`2` (!)** | Равно `MaxOpenConns` (например, `25`), чтобы избежать пересоздания TCP-сессий |
| `db.SetConnMaxLifetime(d)` | `0` | `15–30m` (плавная ротация соединений) |
| `db.SetConnMaxIdleTime(d)` | `0` | `5m` |
| `defer rows.Close()` | — | Обязательно после проверки `err == nil`, иначе соединение навсегда утечёт из пула |
| `defer tx.Rollback()` | — | Безопасный `no-op` после `tx.Commit()`, гарантирует освобождение транзакции при ошибке |

## Очереди, Outbox и Backpressure в Go

Подробно:
[02_messaging_and_backpressure_in_go.md](./01_theory/02_messaging_and_backpressure_in_go.md).

| Задача | Идиоматичное решение в Go |
| :--- | :--- |
| Атомарно записать в БД и отправить событие | **Transactional Outbox:** запись в таблицу `outbox` в том же `*sql.Tx` + фоновый воркер `FOR UPDATE SKIP LOCKED`. |
| Защитить консьюмер от дублей (`at-least-once`) | Таблица `processed_events(event_id PRIMARY KEY)` в одной транзакции с бизнес-действием. |
| Ограничить параллелизм и память воркера | Буферизованный канал `jobs := make(chan Job, N)` + фиксированный пул воркеров (`Worker Pool`). |
| Мгновенно отбросить лишнюю нагрузку (Load Shedding) | `select { case jobs <- job: default: return ErrOverloaded }`. |
| Параллельный вызов с отменой при первой ошибке | `g, ctx := errgroup.WithContext(ctx)` + `g.SetLimit(k)`. |

## Отказоустойчивость и Graceful Shutdown

Подробно:
[03_resilience_and_graceful_shutdown_in_go.md](./01_theory/03_resilience_and_graceful_shutdown_in_go.md).

| Паттерн | Реализация в Go |
| :--- | :--- |
| **Rate Limiting (Token Bucket)** | `golang.org/x/time/rate`: `limiter.Allow()` (в HTTP middleware) или `limiter.Wait(ctx)` (в клиенте). |
| **Circuit Breaker** | Автомат `Closed → Open → Half-Open` под `sync.Mutex` / `atomic` (`sony/gobreaker`). |
| **Retry + Exponential Backoff & Jitter** | Цикл с `base * (1 << attempt)` + случайный `jitter`, ожидание через `select` по `time.NewTimer` и `<-ctx.Done()`. |
| **Graceful Shutdown** | `signal.NotifyContext(ctx, SIGINT, SIGTERM)` → `/readyz = 503` → `srv.Shutdown(ctx)` → `wg.Wait()` → `db.Close()`. |

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func CallWithTimeout(parent context.Context, d time.Duration, fn func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(parent, d)
	defer cancel()
	done := make(chan error, 1) // Буфер 1 защищает горутину от утечки при таймауте!
	go func() { done <- fn(ctx) }()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	err := CallWithTimeout(context.Background(), 50*time.Millisecond, func(ctx context.Context) error {
		return nil
	})
	fmt.Println(err) // <nil>
}
```

## Ловушки

- ⚠️ Оставить `db.SetMaxIdleConns` по умолчанию (`2`) при `MaxOpenConns(50)` —
  постоянный шторм открытий/закрытий соединений к БД под нагрузкой.
- ⚠️ Делать `go handle(msg)` на каждое сообщение из Kafka без семафора или
  Worker Pool — неконтролируемый рост памяти и DDoS собственной БД.
- ⚠️ Ретраить неидемпотентные запросы или ретраить без jitter — синхронный
  Retry Storm добьёт восстанавливающийся сервис.
- ⚠️ Закрывать `db.Close()` сразу по `SIGTERM` до завершения `srv.Shutdown(ctx)`
  и `wg.Wait()` — обрыв всех текущих запросов с ошибкой
  `sql: database is closed`.

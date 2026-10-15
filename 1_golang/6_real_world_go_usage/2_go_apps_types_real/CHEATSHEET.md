# Типы Go-приложений — шпаргалка

Подробности и объяснения — в `01_theory/` каждой темы (ссылки в разделах). ⚠️ — частая ловушка.
Код со сторонними библиотеками (`cobra`, `grpc`, `robfig/cron`, `controller-runtime`, клиенты Kafka) не запускался.

## Что выбрать

| Задача | Тип приложения | Ключевое |
| :--- | :--- | :--- |
| Утилита для людей и скриптов | CLI | stdout — данные, stderr — ошибки, exit code |
| REST API для фронта / партнёров | Backend | слои, graceful shutdown, probes |
| Фоновый агент ОС | System app / daemon | сигналы, atomic write, systemd |
| Межсервисный RPC, streaming, строгий контракт | gRPC | proto, `codes.*`, дедлайны |
| Очередь задач, расписание | Worker / cron | идемпотентность, backoff + jitter, DLQ |
| Kubernetes-native управление сложной системой | Operator | идемпотентный reconcile, finalizers |
| Непрерывный поток событий | Streaming consumer | at-least-once, commit после обработки |

## 1. CLI tools

Подробно: [01_overview.md](./1_cli_tools/01_theory/01_overview.md).

| Задача | Код / команда |
| :--- | :--- |
| Флаги без подкоманд | `flag` (stdlib) |
| Подкоманды, help, completion | `cobra` (`&cobra.Command{Use: "build", RunE: ...}`, `rootCmd.AddCommand`), `urfave/cli` |
| Ошибка и выход | `fmt.Fprintln(os.Stderr, "error:", err)`; `os.Exit(1)` |
| Приоритет конфига | defaults → file (`--config`) → env с префиксом утилиты → flags |
| Ресурсы в бинарнике | `import _ "embed"` + `//go:embed templates/default.yaml` над `var defaultConfig []byte` |
| Кросс-компиляция | `GOOS=linux GOARCH=amd64 go build -o tool-linux-amd64` |
| Все платформы | `go tool dist list`; релизы — `goreleaser` |
| Тест cobra-команды | `cmd.SetArgs([]string{...})` + `cmd.Execute()`; E2E — собрать бинарник и `os/exec` |
| Интерактив | только по флагу (`--interactive`); скрипты работают без TTY |

## 2. Backend

Подробно: [01_overview.md](./2_backend/01_theory/01_overview.md).

| Каталог | Содержимое |
| :--- | :--- |
| `cmd/api/main.go` | сборка зависимостей, запуск, shutdown |
| `internal/http`, `service`, `repository`, `domain` | слои (`internal/` — недоступен другим модулям) |
| `migrations/`, `api/`, `deployments/` | SQL-миграции, OpenAPI/proto, Dockerfile и манифесты |

| Шаг `main` | Код |
| :--- | :--- |
| Конфиг | `cfg := config.Load()`; `if err := cfg.Validate(); err != nil { log.Fatal(err) }` |
| Сигналы | `ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)`; `defer stop()` |
| Запуск | `go func() { if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {...} }()` |
| Graceful shutdown | `<-ctx.Done()` → `sdCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)` → `srv.Shutdown(sdCtx)` |

Стек по умолчанию: `net/http` + `chi`, PostgreSQL + `pgx`/`sqlc`, `log/slog`, `prometheus/client_golang`, `opentelemetry-go`, `envconfig`/`viper`, `testcontainers-go`.

## 3. System apps

Подробно: [01_overview.md](./3_system_apps/01_theory/01_overview.md).

| Сигнал | Обычное значение |
| :--- | :--- |
| `SIGINT` / `SIGTERM` | завершиться (graceful) |
| `SIGHUP` | перечитать конфиг |
| `SIGUSR1` / `SIGUSR2` | свои действия (dump диагностики) |

| Задача | Код |
| :--- | :--- |
| Ловить сигналы (1.16+) | `signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)` |
| Атомарная запись файла | `tmp, _ := os.CreateTemp(filepath.Dir(target), ".tmp-*")` → `Write` → `Close` → `os.Rename(tmp.Name(), target)` (атомарно в пределах одной ФС) |
| Внешний процесс с таймаутом | `cmd := exec.CommandContext(ctx, "bash", "-c", "...")`; `cmd.Run()` — при отмене ctx процесс убивается |
| Пути | `path/filepath`, не ручная склейка через `/` |
| Следить за файлом | `fsnotify` |
| Автозапуск | systemd unit: `ExecStart=`, `Restart=on-failure`, `User=`, `WantedBy=multi-user.target` |
| Демонизация | не нужна: foreground-процесс, демонизирует systemd / Docker / K8s |
| Логи | stderr / journal, не свой файл |

## 4. gRPC services

Подробно: [01_overview.md](./4_grpc_services/01_theory/01_overview.md).

| Выбирать gRPC | Выбирать REST/JSON |
| :--- | :--- |
| service-to-service внутри компании | публичный API, веб-фронт |
| строгий типизированный контракт, streaming | простые CRUD-случаи |
| высокая нагрузка | нет HTTP/2 по пути (старые прокси, CDN) |

| Задача | Код / инструмент |
| :--- | :--- |
| Генерация | `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc`; `buf` (lint, breaking-change detector) |
| Сервер | `s := grpc.NewServer(grpc.UnaryInterceptor(logging))`; `pb.RegisterUserServiceServer(s, srv)`; `s.Serve(lis)` |
| Совместимость | встраивать `pb.UnimplementedUserServiceServer` в свой тип |
| Ошибка | `status.Errorf(codes.NotFound, "user %d not found", id)`, не голый `error` |
| Разобрать ошибку | `st, ok := status.FromError(err)`; `st.Code()` |
| Дедлайн клиента | `ctx, cancel := context.WithTimeout(ctx, 2*time.Second)`; `defer cancel()` |
| Остаток времени на сервере | `ctx.Deadline()` |
| Middleware | `grpc.UnaryInterceptor` / `grpc.StreamInterceptor` |
| Версионирование | `package user.v1;` и `user.v2` рядом |

| Код | Когда |
| :--- | :--- |
| `InvalidArgument` | невалидный запрос |
| `NotFound` | нет ресурса |
| `PermissionDenied` | нет прав |
| `DeadlineExceeded` | истёк дедлайн |
| `Unavailable` | временно недоступно, можно повторить |
| `Internal` | ошибка сервера |

## 5. Workers and cron

Подробно: [01_overview.md](./5_workers_cron/01_theory/01_overview.md).

| Задача | Как |
| :--- | :--- |
| Пул воркеров | N горутин читают `jobs <-chan Job` в `select` с `<-ctx.Done()`; `wg.Wait()` |
| Очередь in-process | `chan Job` — если сервис не масштабируется горизонтально |
| Очередь внешняя | Redis lists, RabbitMQ, NATS, SQS, Kafka — persistence, несколько инстансов |
| Retry | только retriable-ошибки, лимит попыток, exponential backoff + jitter, ждать в `select` с `ctx.Done()` |
| Финальный провал | DLQ |
| Дубли | idempotency key, проверка «уже сделано», компенсации |
| Простой cron | `t := time.NewTicker(5 * time.Minute)`; `defer t.Stop()`; `select { case <-ctx.Done(): case <-t.C: }` |
| Cron-выражения | `robfig/cron/v3`: `c := cron.New()`; `c.AddFunc("*/5 * * * *", job)`; `c.Start()` |
| Cron с секундами (6 полей) | `cron.New(cron.WithSeconds())` |
| Несколько инстансов | leader election (etcd, k8s Lease, Redis lock), `SELECT ... FOR UPDATE SKIP LOCKED`, K8s CronJob |
| Shutdown | перестать брать задачи → доделать текущие → вернуть упавшие в очередь/DLQ → закрыть ресурсы |

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

var errTemporary = errors.New("temporary")

func retry(ctx context.Context, attempts int, base time.Duration, do func() error) error {
	var err error
	for i := range attempts { // range over int — Go 1.22+
		if err = do(); err == nil || !errors.Is(err, errTemporary) {
			return err // успех или fatal-ошибка — не повторяем
		}
		backoff := base << i                                      // экспонента: base, 2·base, 4·base…
		backoff += time.Duration(rand.Int64N(int64(backoff) / 2)) // jitter
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}
	}
	return fmt.Errorf("после %d попыток: %w", attempts, err)
}

func main() {
	calls := 0
	err := retry(context.Background(), 5, time.Millisecond, func() error {
		calls++
		if calls < 3 {
			return errTemporary
		}
		return nil
	})
	fmt.Println(calls, err) // 3 <nil>

	err = retry(context.Background(), 2, time.Millisecond, func() error { return errTemporary })
	fmt.Println(err) // после 2 попыток: temporary
}
```

## 6. K8s operators

Подробно: [01_overview.md](./6_k8s_operators/01_theory/01_overview.md).

| Понятие | Суть |
| :--- | :--- |
| CRD / CR | схема своего ресурса / объект (`kubectl get mydb`) |
| Controller | reconcile-цикл: привести реальное состояние к желаемому |
| Status + Conditions | `Type` / `Status` / `Reason` / `Message`: Ready, Progressing, Degraded |
| Finalizer | блокирует удаление CR, пока не вычищены внешние ресурсы |
| Стек | kubebuilder (стандарт), operator-sdk, controller-runtime |

| Задача | Код |
| :--- | :--- |
| Сигнатура | `Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error)` |
| Объект удалён | `return ctrl.Result{}, client.IgnoreNotFound(err)` |
| Повторить позже | `ctrl.Result{RequeueAfter: 30 * time.Second}` |
| Удаление с cleanup | `DeletionTimestamp != nil` → cleanup → `controllerutil.RemoveFinalizer` → `r.Update` |
| Тесты | fake client (unit), envtest (integration), kind / k3d (E2E) |

Operator нужен, только если Deployment / StatefulSet / Helm не покрывают автоматизацию (failover, backup, scaling).

## 7. Streaming data

Подробно: [01_overview.md](./7_streaming_data/01_theory/01_overview.md).

| Семантика | Когда коммит | Результат |
| :--- | :--- | :--- |
| At-most-once | до обработки | возможна потеря |
| At-least-once (стандарт) | после успешной обработки | возможны дубли → идемпотентный consumer |
| Exactly-once | Kafka transactions | дорого, часто маркетинг |

| Задача | Как |
| :--- | :--- |
| Идемпотентность | idempotency key + таблица обработанных, upsert, `ON CONFLICT DO NOTHING` |
| Порядок | гарантирован только в пределах partition → ключ сообщения (например, `userID`) |
| Backpressure | мониторить lag, worker pool, отдельный topic для тяжёлых, больше инстансов |
| Rebalance | обработка встаёт; в хуке зафиксировать offset и очистить state |
| Poison message | retry с лимитом → DLQ-topic + alert; не крутить бесконечно |
| Схема | JSON (без схемы) vs Avro / Protobuf + schema registry |
| Клиенты | Kafka: `segmentio/kafka-go` (чистый Go), `confluent-kafka-go` (cgo); `nats.go`; `go-redis` (Streams) |

## Ловушки

- ⚠️ CLI пишет данные в stderr или логи в stdout → ломаются pipe'ы (`tool list | grep`).
- ⚠️ Exit code 0 при ошибке → скрипты и CI не видят провала.
- ⚠️ `srv.ListenAndServe()` без проверки `http.ErrServerClosed` → штатный shutdown выглядит как ошибка.
- ⚠️ Нет graceful shutdown → при rolling deploy теряются in-flight запросы и задачи.
- ⚠️ `os.WriteFile` поверх конфига → при падении частично записанный файл; нужен temp + rename.
- ⚠️ `exec.Command` без `Context` → зависший дочерний процесс живёт вечно.
- ⚠️ Самодельный double fork вместо systemd.
- ⚠️ Смена номера поля в `.proto` — breaking change.
- ⚠️ proto3 без `optional` не отличает «не задано» от нуля.
- ⚠️ gRPC-вызов без дедлайна на клиенте → висит бесконечно.
- ⚠️ Retry без jitter → thundering herd; retry неидемпотентной операции → двойное списание.
- ⚠️ Линейный `attempt * base` — не экспоненциальный backoff.
- ⚠️ Ticker не ждёт окончания job: длинная задача «съедает» тики, интервал плывёт.
- ⚠️ `robfig/cron/v3` по умолчанию 5 полей; 6-польное выражение с секундами без `cron.WithSeconds()` не распарсится.
- ⚠️ Cron на N инстансах без lock → задача выполняется N раз.
- ⚠️ Reconcile не идемпотентен → каждое событие создаёт ресурсы заново.
- ⚠️ Внешние ресурсы без finalizer → «сиротские» S3-бакеты и DNS-записи после удаления CR.
- ⚠️ Commit offset до обработки → потеря сообщений при падении.
- ⚠️ Poison message без DLQ → весь pipeline стоит.

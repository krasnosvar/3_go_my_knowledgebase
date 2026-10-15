# Workers and Cron

Worker — фоновое приложение, которое разгребает задачи из очереди. Cron — задачи по расписанию. Часто они в одном бинарнике с backend, иногда — отдельный сервис.

## Worker pattern

```go
func runWorker(ctx context.Context, jobs <-chan Job) {
    var wg sync.WaitGroup
    for i := 0; i < N; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                select {
                case <-ctx.Done():
                    return
                case j, ok := <-jobs:
                    if !ok { return }
                    if err := process(ctx, j); err != nil {
                        slog.Error("job failed", "id", j.ID, "err", err)
                        scheduleRetry(j, err)
                    }
                }
            }
        }()
    }
    wg.Wait()
}
```

См. `3_concurrency/07_patterns/01_worker_pool` для деталей.

## Очереди

- **Внутренняя**: `chan Job` для in-process. Подходит, если сервис не масштабируется горизонтально.
- **Внешняя**: Redis lists, RabbitMQ, NATS, AWS SQS, Kafka. Гарантии, persistence, multi-instance.

## Retry strategy

```go
for attempt := 1; attempt <= maxAttempts; attempt++ {
    err := do()
    if err == nil { return nil }
    if !isRetriable(err) { return err }
    backoff := baseBackoff<<(attempt-1) + jitter() // base, 2·base, 4·base…
    select {
    case <-ctx.Done(): return ctx.Err()
    case <-time.After(backoff): // continue
    }
}
return ErrMaxRetries
```

Ключевое:

- **Exponential backoff** + **jitter** (чтобы не было thundering herd).
- Различение **retriable** и **fatal** ошибок.
- Лимит количества попыток.
- DLQ (dead-letter queue) для финального промаха.

## Idempotency

Задача может быть выполнена дважды (retry, дубликат в очереди). Сервис обязан уметь:

- Использовать idempotency key (хеш задачи или внешний ID).
- Проверять «уже сделано» перед выполнением.
- При side-effect'ах — компенсирующие действия (refund, отмена).

## Cron

Простой ticker:

```go
t := time.NewTicker(5 * time.Minute)
defer t.Stop()
for {
    select {
    case <-ctx.Done(): return
    case <-t.C: runJob()
    }
}
```

Ticker не учитывает время выполнения предыдущей итерации — если runJob занял 10 минут, ты будешь стартовать новую раз в 10 минут, не 5.

Для cron-выражений (`* * * * *`):

- `robfig/cron/v3` — стандарт de-facto.

```go
c := cron.New() // 5 полей; для 6 полей с секундами — cron.New(cron.WithSeconds())
c.AddFunc("*/5 * * * *", func() { runJob() })
c.Start()
defer c.Stop()
```

## Distributed cron

Если сервис масштабирован — несколько инстансов могут запустить одну и ту же задачу одновременно. Решения:

- **Leader election** (через etcd, k8s lease, Redis lock).
- **Database lock**: `SELECT ... FOR UPDATE SKIP LOCKED`.
- **Внешний scheduler**: Kubernetes CronJob, Airflow, Temporal.

## Graceful shutdown

Worker должен:

1. Перестать брать новые задачи (`close(jobs)` или ctx.Done).
2. Доделать текущие.
3. Зафейленные — вернуть в очередь / DLQ.
4. Закрыть БД, метрики, логгер.

Без этого при rolling deploy теряются in-flight задачи.

## Правило

- Идемпотентность — фундамент. Без неё retry опасен.
- Backoff + jitter в retry.
- DLQ для финальных провалов, чтобы их не потерять.
- Distributed cron требует locking механизма.

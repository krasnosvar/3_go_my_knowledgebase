# Use-cases goroutines

## 1. Конкурентные независимые операции

Параллельно делаем то, что не связано друг с другом:

```go
g, ctx := errgroup.WithContext(ctx)
var user User
var orders []Order

g.Go(func() (err error) { user, err = fetchUser(ctx, id); return })
g.Go(func() (err error) { orders, err = fetchOrders(ctx, id); return })

if err := g.Wait(); err != nil { return err }
```

Сокращаем latency до max из всех запросов.

## 2. Обработка очереди сообщений

```go
for msg := range inbox {
    go handle(msg)
}
```

Один producer, много consumer'ов. На практике — с WaitGroup и ограничением.

## 3. Фоновые задачи сервиса

```go
go runMetricsReporter(ctx)
go runScheduledJobs(ctx)
```

Контекст — механизм graceful shutdown: `ctx` получит `Done` при остановке сервиса, фоновые задачи корректно завершатся.

## 4. Время ожидания с fallback

```go
resultCh := make(chan Result, 1)
go func() { resultCh <- slowOp() }()

select {
case r := <-resultCh:
    return r
case <-time.After(timeout):
    return fallback
}
```

Buffered-канал здесь важен — иначе горутина заблокируется и утечёт.

## 5. Batch-обработка

Разбиваем большой объём на куски и обрабатываем параллельно:

```go
chunks := partition(items, 100)
g, ctx := errgroup.WithContext(ctx)
for _, c := range chunks {
    c := c
    g.Go(func() error { return processChunk(ctx, c) })
}
return g.Wait()
```

## 6. Pipeline

Цепочка stage'ов, соединённых каналами. Каждый stage — отдельная горутина. См. `07_patterns/03_pipeline.md`.

## 7. Симуляция многопользовательского доступа

В тестах и бенчмарках goroutines — простой способ сымитировать параллельный трафик.

## Когда goroutine НЕ нужен

- Последовательный код, где нет блокирующих операций.
- Одна небольшая функция, результата которой ждут сразу.
- Когда overhead синхронизации превысит выигрыш от параллелизма.

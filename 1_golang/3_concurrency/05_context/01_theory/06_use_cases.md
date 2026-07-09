# Use-cases context

## 1. HTTP handler → нижние слои

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
    defer cancel()

    u, err := service.GetUser(ctx, id)
    // ...
}
```

Весь стек знает: «лимит 2 секунды, клиент мог отключиться».

## 2. gRPC / network-клиенты

Все официальные клиенты (`database/sql`, `net/http`, `grpc`, `redis`) умеют `XxxContext(ctx, ...)`. Их отмена мгновенно прерывает операцию.

## 3. Работа с БД

```go
rows, err := db.QueryContext(ctx, "SELECT ...")
```

Если ctx отменён — запрос будет прерван БД-сервером (для поддерживающих драйверов).

## 4. Фоновые задачи / воркеры

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return
        case j := <-jobs:
            process(ctx, j)
        }
    }
}
```

Graceful shutdown: сервис вызывает `cancel()`, и все воркеры штатно завершаются.

## 5. Отмена цепочки параллельных операций

```go
g, ctx := errgroup.WithContext(ctx)
for _, u := range urls {
    u := u
    g.Go(func() error { return fetch(ctx, u) })
}
return g.Wait()
```

Первая ошибка → `ctx` отменяется → остальные запросы отменяются.

## 6. Пробрасывание request-scoped значений

Request ID, trace ID, user — через `context.WithValue` + типизированные аксессоры. Middleware кладёт, handler / service читает.

## 7. Дедлайн на часть операции

```go
parentCtx := r.Context()

// внешний вызов — ограничим 500 мс, не больше
extCtx, cancel := context.WithTimeout(parentCtx, 500*time.Millisecond)
defer cancel()

resp, err := httpClient.Do(req.WithContext(extCtx))
```

Если запрос превысит 500 мс — прерываем, но основной ctx продолжает жить.

## 8. Graceful shutdown в main

```go
ctx, cancel := signal.NotifyContext(context.Background(),
    syscall.SIGINT, syscall.SIGTERM)
defer cancel()

server.ListenAndServeCtx(ctx)
```

Удобный идиом c Go 1.16+: SIGINT/SIGTERM сразу превращается в отмену контекста.

## Когда context НЕ нужен

- Чистые функции без I/O и долгих вычислений.
- Статические утилиты, которые мгновенно возвращаются.
- Внутренние вспомогательные типы, где никто не ожидает отмены.

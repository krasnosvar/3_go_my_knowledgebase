# Use-cases select

## 1. Ожидание «первого из»

Когда нужно выбрать, что пришло раньше: результат или таймаут/отмена.

```go
select {
case r := <-resultCh:
    return r, nil
case <-ctx.Done():
    return zero, ctx.Err()
}
```

Это самый частый паттерн в Go-сервисах.

## 2. Мультиплексирование источников

Несколько каналов обрабатываются одной горутиной:

```go
for {
    select {
    case <-ctx.Done():
        return
    case cmd := <-commands:
        handleCommand(cmd)
    case tick := <-ticker.C:
        handleTick(tick)
    case err := <-errors:
        handleError(err)
    }
}
```

Удобно для event loop / координаторов.

## 3. Non-blocking send/recv

```go
select {
case metrics <- event:
default:
    droppedCounter.Inc()
}
```

Если буфер метрик забит — считаем «дропов», не тормозим основной поток.

## 4. Heartbeat

Воркер периодически сообщает «я жив»:

```go
hb := time.NewTicker(heartbeat)
defer hb.Stop()

for {
    select {
    case <-ctx.Done():
        return
    case <-hb.C:
        status <- "alive"
    case j := <-jobs:
        process(j)
    }
}
```

## 5. Fan-in

Собрать значения из нескольких каналов в один поток:

```go
func merge(a, b <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for a != nil || b != nil {
            select {
            case v, ok := <-a:
                if !ok { a = nil; continue }
                out <- v
            case v, ok := <-b:
                if !ok { b = nil; continue }
                out <- v
            }
        }
    }()
    return out
}
```

`nil`-канал в select «отключает» ветку — удобный трюк.

## 6. Rate limiting через ticker

```go
limiter := time.Tick(100 * time.Millisecond)
for req := range requests {
    <-limiter
    go handle(req)
}
```

Простой леакируемый rate limit — на prod-уровне лучше `golang.org/x/time/rate`.

## 7. Graceful shutdown

Главный цикл сервера реагирует на сигналы OS:

```go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

select {
case <-sigCh:
    cancel()
case err := <-serverErr:
    return err
}
```

## Когда select — лишний

- Одна операция с одним каналом — просто `<-ch`.
- Синхронная цепочка — обычные вызовы без каналов.

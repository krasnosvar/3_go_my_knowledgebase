# Worker pool

Worker pool — паттерн, где N горутин параллельно берут задачи из общего канала и обрабатывают их.

## Зачем

- Ограничить параллелизм (не создавать тысячу горутин под каждую задачу).
- Сгладить burst-нагрузку: задачи накапливаются, worker'ы разгребают.
- Единая точка control-flow (остановить можно, закрыв канал jobs).

## Минимальная реализация

```go
func workerPool(ctx context.Context, n int, jobs <-chan Job) {
    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                select {
                case <-ctx.Done():
                    return
                case j, ok := <-jobs:
                    if !ok {
                        return
                    }
                    process(ctx, j)
                }
            }
        }()
    }
    wg.Wait()
}
```

## С результатами

```go
func workerPool(ctx context.Context, n int, jobs <-chan Job) <-chan Result {
    out := make(chan Result)
    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := range jobs {
                out <- process(ctx, j)
            }
        }()
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

Один producer → канал jobs → N worker'ов → общий канал out.

## Размер пула

- Для CPU-bound: `runtime.NumCPU()`.
- Для IO-bound: обычно больше (от десятков до сотен), но не бесконечно — ограничено backend'ом (DB connection pool, API rate limits).
- Неправильный размер маскируется мониторингом: если latency растёт — пул мал, если упёрлись в ресурс — пул велик.

## Graceful shutdown

Два пути сигнала:

1. **Закрыть jobs** — worker'ы домолотят оставшееся и выйдут.
2. **Отменить ctx** — worker'ы прерываются сразу, даже если jobs ещё есть.

Оба важны: shutdown обычно «больше не принимаем новые» + «доделываем текущие с дедлайном».

## Ошибки

- Паника в job убивает worker и весь pool. Всегда ловите `recover` в обработчике.
- Без `ctx.Done()` в цикле worker не отреагирует на отмену.
- Если result channel без читателя — worker'ы заблокируются на отправке. Либо buffered, либо гарантированный consumer.

## Когда НЕ worker pool

- Разовые задачи — простая goroutine.
- Последовательные операции — обычный цикл.
- Fan-out с разной природой задач — лучше отдельные pool'ы.

# Use-cases concurrency patterns

## 1. Параллельные HTTP-запросы → errgroup

```go
g, ctx := errgroup.WithContext(ctx)
var user User
var orders []Order

g.Go(func() (err error) { user, err = fetchUser(ctx, id); return })
g.Go(func() (err error) { orders, err = fetchOrders(ctx, id); return })

if err := g.Wait(); err != nil {
    return err
}
```

Сокращаем latency до максимума из веток.

## 2. Обработка очереди → worker pool

N воркеров читают из одного канала jobs. Ограничиваем параллелизм, выравниваем нагрузку.

## 3. Расщепление задачи → fan-out

Большой набор → параллельная обработка → объединение результатов (fan-in).

## 4. ETL / streaming → pipeline

Цепочка стадий на каналах. Каждая стадия масштабируется независимо.

## 5. Ограничение ресурсов → semaphore

Не больше M одновременных обращений к внешнему API. Обычно buffered-канал размера M.

## 6. Таймауты / отмена → context

Каждая операция в сервисе подчиняется ctx от вышестоящего вызова.

## 7. Graceful shutdown → signal.NotifyContext

Сервис должен уметь корректно завершаться по SIGTERM.

## 8. Rate limiting

Для ограничения RPS — `golang.org/x/time/rate`. Не изобретай велосипед на ticker'ах для prod.

## 9. Circuit breaker

Когда downstream падает, временно перестаём его дёргать, чтобы не заспамить. Реализации: `sony/gobreaker`, свои обёртки вокруг счётчиков.

## 10. Batching

Агрегировать запросы в пакеты, отправлять реже. Часто реализуется через ticker + буфер:

```go
batch := make([]Item, 0, maxBatch)
t := time.NewTicker(flushInterval)
for {
    select {
    case item := <-in:
        batch = append(batch, item)
        if len(batch) >= maxBatch {
            flush(batch)
            batch = batch[:0]
        }
    case <-t.C:
        if len(batch) > 0 {
            flush(batch)
            batch = batch[:0]
        }
    case <-ctx.Done():
        flush(batch)
        return
    }
}
```

## 11. Heartbeat

Периодический пинг от воркера, чтобы мониторинг видел «жив».

## Когда не применять эти паттерны

- Последовательный код быстрее пишется и проще тестируется.
- Overhead горутин/каналов превысит профит для очень коротких операций.
- Нет реальной параллельной нагрузки — параллелизм «впрок» только добавляет багов.

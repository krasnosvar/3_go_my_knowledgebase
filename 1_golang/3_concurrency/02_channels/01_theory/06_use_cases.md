# Use-cases channels

## 1. Producer–Consumer

Классика: одна горутина кладёт в канал, другая забирает.

```go
jobs := make(chan Job, 100)
go producer(jobs)
go consumer(jobs)
```

Плюс: разрыв связанности, естественный backpressure.

## 2. Worker pool

Набор worker'ов читают из одного канала:

```go
jobs := make(chan Job)
for i := 0; i < n; i++ {
    go func() {
        for j := range jobs {
            process(j)
        }
    }()
}
```

Параллелизм ограничен числом worker'ов.

## 3. Fan-out / fan-in

- **Fan-out**: один source, несколько читателей.
- **Fan-in**: несколько источников, один потребитель (через select или merge-функцию).

## 4. Pipeline

Цепочка каналов: `stage1 → stage2 → stage3`. Каждая стадия — горутина, читает из предыдущего канала, пишет в следующий.

```go
nums := generate(1, 2, 3, 4)
sq := square(nums)
for v := range sq {
    fmt.Println(v)
}
```

## 5. Broadcast через close

```go
done := make(chan struct{})
// множество читателей
for i := 0; i < n; i++ {
    go func() {
        <-done
        cleanup()
    }()
}
close(done) // все читатели проснутся
```

Один close разбудит всех — дешёвый multicast.

## 6. Timeout / Deadline

Через `select`:

```go
select {
case v := <-ch:
    use(v)
case <-time.After(timeout):
    return ErrTimeout
}
```

Или через `context.Context` — современный стандарт.

## 7. Semaphore / ограничение concurrency

Буферизованный канал фиксированной ёмкости работает как semaphore:

```go
sem := make(chan struct{}, 3)
for _, j := range jobs {
    sem <- struct{}{}
    go func(j Job) {
        defer func() { <-sem }()
        process(j)
    }(j)
}
```

## 8. Event bus (осторожно)

Небольшая шина событий на канале:

```go
events := make(chan Event, 100)
```

Годится, пока это внутри одного процесса. Для долгоживущих / межсервисных сценариев — Kafka, NATS.

## Когда НЕ канал

- Защитить shared state → `sync.Mutex`.
- Атомарный счётчик → `sync/atomic`.
- Строгая межсервисная коммуникация → очереди, не go-каналы.

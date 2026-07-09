# Fan-out / Fan-in

## Fan-out

Одна задача «раздувается» на несколько параллельных обработчиков. Типичная ситуация: пришёл пакет данных — обработать каждый элемент параллельно.

```go
in := make(chan Item)
go producer(in)

// 5 параллельных обработчиков
results := fanOut(in, 5, process)
```

`fanOut` запускает N горутин, каждая из которых читает из общего входа.

## Fan-in

Несколько источников сводятся в один канал. Полезно, когда у тебя несколько продьюсеров, а консьюмер один.

```go
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range cs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(c)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## Вместе: fan-out → fan-in

Типичный pipeline обработки:

```
input
  |
  v
[fan-out to N workers]  -> worker1 -> resultCh
                          worker2 -> resultCh
                          ...
                          workerN -> resultCh
                               |
                               v
                         [fan-in merge]
                               |
                               v
                         consumer
```

## Контроль параллелизма

- Небольшой fan-out (2-8) — часто оптимум для IO.
- Большой fan-out без нужды → перегруз downstream.
- Динамический fan-out → semaphore через buffered-канал:

```go
sem := make(chan struct{}, maxConcurrent)
for _, j := range jobs {
    sem <- struct{}{}
    go func(j Job) {
        defer func() { <-sem }()
        process(j)
    }(j)
}
```

## Ошибки и отмена

Когда один worker вернул ошибку, остальных часто надо остановить. Для этого — `errgroup.WithContext`:

```go
g, ctx := errgroup.WithContext(ctx)
for _, j := range jobs {
    j := j
    g.Go(func() error { return process(ctx, j) })
}
return g.Wait()
```

Первая ошибка → ctx отменяется → остальные прерываются.

## Частые ошибки

- Не закрыть merged-канал. Consumer зависнет на `for v := range out`.
- Забыть увеличить WaitGroup до go func.
- Неограниченный fan-out — «на каждое сообщение горутина» → утечки при задержке downstream.

## Когда применять

- CPU/IO можно распараллелить.
- Источников несколько, потребитель один.
- Хочется простой backpressure через ёмкость каналов.

# Ожидание завершения goroutine

Go не встраивает «await». Ожидание реализуется через примитивы синхронизации.

## sync.WaitGroup

Самый распространённый способ дождаться набора горутин.

```go
var wg sync.WaitGroup

for _, job := range jobs {
    wg.Add(1)
    go func(j Job) {
        defer wg.Done()
        process(j)
    }(job)
}

wg.Wait()
```

Правила:

- `Add(n)` — до запуска горутин.
- `Done()` — в defer, чтобы не пропустить при panic.
- Нельзя вызывать `Add` после `Wait`.

## Канал-сигнал

```go
done := make(chan struct{})
go func() {
    work()
    close(done)
}()
<-done // ждём
```

Подходит для одной горутины. `close(done)` транслирует сигнал всем, кто читает.

## errgroup

`golang.org/x/sync/errgroup` — WaitGroup + первая ошибка + контекст.

```go
g, ctx := errgroup.WithContext(ctx)
for _, u := range urls {
    url := u
    g.Go(func() error {
        return fetch(ctx, url)
    })
}
if err := g.Wait(); err != nil {
    return err
}
```

Это идиоматический вариант для бизнес-кода с множественными параллельными операциями.

## Результат и ошибка

Если нужно вернуть результат — делай канал или собственные структуры. `WaitGroup` результата не даёт.

```go
results := make(chan int, len(jobs))
for _, j := range jobs {
    go func(j Job) { results <- process(j) }(j)
}
for i := 0; i < len(jobs); i++ {
    fmt.Println(<-results)
}
```

## Антипаттерны

- `time.Sleep(X)` «чтобы точно успели» — гарантии нет, только флaки.
- Бесконечный `for { }` в main вместо правильного ожидания — чревато утечками.
- `runtime.Gosched()` как замена синхронизации — не работает так, как кажется.

## Правило

Всегда имей явное место ожидания. Любая горутина, которую запускает код — должна контролироваться каким-то из способов выше.

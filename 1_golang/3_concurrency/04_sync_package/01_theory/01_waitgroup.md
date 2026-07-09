# sync.WaitGroup

`WaitGroup` — счётчик незавершённых горутин. Основной способ дождаться окончания группы задач.

## Основное API

```go
var wg sync.WaitGroup

wg.Add(1)        // увеличить счётчик перед запуском
go func() {
    defer wg.Done() // уменьшить при выходе
    work()
}()

wg.Wait()        // ждём, пока счётчик не станет 0
```

- `Add(n)` — увеличить на n.
- `Done()` — эквивалент `Add(-1)`.
- `Wait()` — блокирует, пока счётчик > 0.

## Правила

- `Add` — **до** запуска горутины, не внутри.
- `Done` — в `defer`, чтобы не пропустить при panic.
- Нельзя вызывать `Add` одновременно с `Wait` — гонка.
- Нельзя переиспользовать `WaitGroup` после `Wait`, если `Add` ещё в процессе.

## Типовой паттерн

```go
var wg sync.WaitGroup
for _, j := range jobs {
    wg.Add(1)
    go func(j Job) {
        defer wg.Done()
        process(j)
    }(j)
}
wg.Wait()
```

## Ошибки

### 1. Add внутри горутины

```go
go func() {
    wg.Add(1) // плохо: main может уже запустить Wait()
    defer wg.Done()
}()
```

Увеличивай счётчик в той же горутине, где `Wait`.

### 2. Забыли Done при ошибке

```go
go func() {
    if err := doA(); err != nil {
        return // Done не вызван, Wait никогда не завершится
    }
    defer wg.Done()
}()
```

Всегда `defer wg.Done()` сразу в начале горутины.

### 3. Копирование WaitGroup

```go
func run(wg sync.WaitGroup) { // плохо: копия
    wg.Add(1)
}
```

Передавай `*sync.WaitGroup`. Копирование ломает учёт.

## Когда WaitGroup недостаточно

- Нужен результат от каждой горутины — используй каналы или errgroup.
- Нужна первая ошибка — `errgroup.Group`.
- Нужен таймаут — `errgroup` + контекст или обёртка с `select`.

## Пример с errgroup (рекомендовано)

```go
g, ctx := errgroup.WithContext(ctx)
for _, u := range urls {
    u := u
    g.Go(func() error { return fetch(ctx, u) })
}
return g.Wait()
```

В 90% бизнес-кода это удобнее голого `WaitGroup`.

# Goroutine leaks

Утечка горутины — ситуация, когда горутина запущена, но никогда не завершится. Последствия: удерживается память, переменные из замыкания не GC-ются, ресурсы продолжают жить.

## Типичные сценарии

### 1. Отправка в канал без получателя

```go
ch := make(chan int)
go func() {
    ch <- compute() // блок навсегда, если никто не читает
}()
```

### 2. Бесконечный цикл без условия выхода

```go
go func() {
    for {
        work()
    }
}()
```

### 3. Забытый контекст

```go
go func() {
    for {
        select {
        case <-ticker.C:
            work()
        }
        // нет ctx.Done() — нельзя остановить
    }
}()
```

### 4. Горутина ждёт, потребитель ушёл

```go
resultCh := make(chan Result)
go func() {
    resultCh <- slowOp()
}()

select {
case r := <-resultCh:
    use(r)
case <-time.After(timeout):
    return // горутина останется висеть, пока slowOp не вернётся
}
```

## Как предотвращать

- Всегда иметь путь выхода: `context.Context`, канал `done`, таймер.
- Буферизированный канал размера 1, если результат может быть «никому не нужен».
- `select` с `ctx.Done()` в долгоживущих циклах.
- Тесты и профайлер.

## Поиск утечек

- `runtime.NumGoroutine()` — мониторить.
- `pprof.Lookup("goroutine").WriteTo(w, 2)` — получить stack trace всех горутин.
- `net/http/pprof` в деве: `curl localhost:6060/debug/pprof/goroutine?debug=1`.
- `go.uber.org/goleak` для тестов — прерывает тест, если в конце остались лишние горутины.

## Паттерн «горутина с контекстом»

```go
func run(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        case job := <-jobs:
            process(job)
        }
    }
}
```

Правило: каждая «долгоживущая» горутина должна уметь завершаться по `ctx.Done()`.

## Паттерн «buffered result»

```go
resultCh := make(chan Result, 1) // буфер = 1
go func() { resultCh <- slowOp() }()

select {
case r := <-resultCh:
    use(r)
case <-ctx.Done():
    // горутина завершится, когда slowOp вернётся, и не заблокируется
}
```

Буфер 1 позволяет горутине завершиться, даже если получателя уже нет.

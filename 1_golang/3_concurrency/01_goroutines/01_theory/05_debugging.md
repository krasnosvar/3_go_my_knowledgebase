# Отладка goroutines

Goroutines не имеют ID (в публичном API), но рантайм даёт достаточно инструментов.

## runtime.NumGoroutine

```go
fmt.Println(runtime.NumGoroutine())
```

Быстрый способ проверить, не растёт ли число горутин в неожиданном месте (индикатор утечек).

## Stack dump всех горутин

```go
buf := make([]byte, 64*1024)
n := runtime.Stack(buf, true)
os.Stderr.Write(buf[:n])
```

Даёт полные стек-трейсы всех живых горутин. Полезно в diagnostic endpoint.

Или через pprof:

```go
import _ "net/http/pprof"

go http.ListenAndServe("localhost:6060", nil)
```

Затем:

```
curl localhost:6060/debug/pprof/goroutine?debug=2
```

## Race detector

```
go test -race ./...
go run -race ./...
```

Находит data race в рантайме. Замедляет исполнение в 5-10 раз, но это лучший детектор гонок в экосистеме.

## Execution tracer

```go
import "runtime/trace"

f, _ := os.Create("trace.out")
trace.Start(f)
defer trace.Stop()
```

Потом `go tool trace trace.out` — визуализация планирования, блокировок, GC.

## GODEBUG

Несколько полезных:

- `GODEBUG=schedtrace=1000,scheddetail=1` — каждую секунду планировщик печатает состояние.
- `GODEBUG=gctrace=1` — логи сборщика мусора.

## Проверка утечек в тестах

`go.uber.org/goleak`:

```go
func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}
```

Падает, если после теста остались горутины.

## Общие советы

- Логируй момент старта и завершения долгоживущих горутин — это даёт видимость в логах.
- Помечай их осмысленными именами в `slog`/`zap`: `slog.Info("worker started", "id", n)`.
- Любая горутина, которая не завершается естественно — кандидат на диагностику.

## Признаки, на которые стоит реагировать

- `NumGoroutine()` растёт в мониторинге линейно с временем.
- pprof-дамп показывает сотни одинаковых стек-трейсов — типовой leak.
- Race-детектор стабильно даёт репорты в CI — чинить сразу, а не «потом».

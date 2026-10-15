# net/http: server

## Минимальный сервер

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello")
    })

    srv := &http.Server{
        Addr:              ":8080",
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
    }
    log.Fatal(srv.ListenAndServe())
}
```

`http.ServeMux` — встроенный роутер. С Go 1.22 у него появился синтаксис методов и паттернов: `mux.HandleFunc("GET /users/{id}", ...)`.

## Handler / HandlerFunc

```go
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
```

`http.HandlerFunc(f)` — функциональная обёртка. Любой `func(w, r)` можно превратить в `http.Handler`.

## Ответ

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(resp)
```

Важно: после первого `Write` заголовки уже ушли, `WriteHeader` повторно не сработает.

## Middleware

```go
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

srv := &http.Server{Handler: logging(mux)}
```

Цепочка: `auth(recovery(logging(mux)))`.

## Timeouts

Почти всегда нужно выставить:

```go
srv := &http.Server{
    ReadHeaderTimeout: 5 * time.Second,
    ReadTimeout:       20 * time.Second,
    WriteTimeout:      20 * time.Second,
    IdleTimeout:       60 * time.Second,
}
```

Без таймаутов сервер уязвим к медленным клиентам (slowloris) и утечкам соединений.

## Graceful shutdown

```go
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()

go func() { _ = srv.ListenAndServe() }()

<-ctx.Done()
sdCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
_ = srv.Shutdown(sdCtx)
```

`Shutdown` перестаёт принимать новые соединения, дожидается текущих.

## Context в Request

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    res, err := svc.Find(ctx, id)
    // ...
}
```

Контекст запроса отменяется, если клиент отключился. Пробрасывай его в БД и внешние вызовы.

## Частые ошибки

- Не ставить таймауты — ждём DoS.
- Писать в `w` после `WriteHeader` с ошибкой — поздно, заголовки уже отправлены.
- Забыть `r.Body.Close()` в handler — обычно не критично (сервер сам закрывает), но если читаешь body частично и уходишь — закрой.
- `panic` в handler: `net/http` перехватывает её и закрывает только это соединение (в логе `http: panic serving ...`), клиент получает обрыв вместо ответа; `panic` в горутине, запущенной из handler, убивает весь процесс — нужен recovery middleware (ответ 500 + лог).

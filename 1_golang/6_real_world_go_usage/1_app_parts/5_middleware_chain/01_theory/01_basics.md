# HTTP middleware

Middleware — функция-обёртка вокруг handler'а. Добавляет cross-cutting поведение: логирование, аутентификацию, recovery, метрики, CORS.

## Базовая форма

```go
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("request",
            "method", r.Method,
            "path", r.URL.Path,
            "duration", time.Since(start),
        )
    })
}
```

Тип: `func(http.Handler) http.Handler`. Каждое middleware принимает следующий handler и возвращает новый.

## Цепочка

```go
h := logging(recovery(auth(coreHandler)))

http.Handle("/", h)
```

Порядок исполнения — снаружи внутрь:

1. `logging` стартует таймер.
2. `recovery` ставит defer на panic.
3. `auth` проверяет токен.
4. `coreHandler` — собственно логика.
5. Назад: `auth` → `recovery` → `logging`.

## Типичная цепочка backend'а

Снаружи внутрь:

1. **request_id** — генерим уникальный ID, кладём в context.
2. **logging** — структурный лог запроса.
3. **recovery** — defer recover, превращает panic в 500.
4. **metrics** — Prometheus counter / histogram.
5. **tracing** — OpenTelemetry span.
6. **timeout** — `http.TimeoutHandler` или `context.WithTimeout`.
7. **CORS** — выставляем заголовки.
8. **auth** — валидация токена, кладём пользователя в context.
9. **rate limit** — на endpoint или per-user.
10. **handler** — наш бизнес-код.

## Чтение через context

Middleware кладёт значение в context, handler читает:

```go
type ctxKey int
const userKey ctxKey = iota

func auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        u, err := authenticate(r)
        if err != nil { http.Error(w, "unauthorized", 401); return }
        ctx := context.WithValue(r.Context(), userKey, u)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// в handler
u, ok := r.Context().Value(userKey).(User)
```

Используй типизированный ключ (свой `ctxKey`), не `string` — избежишь коллизий.

## Recovery

Должен быть как можно ближе к handler'у:

```go
func recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if rec := recover(); rec != nil {
                slog.Error("panic", "err", rec, "stack", debug.Stack())
                http.Error(w, "internal error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

Без recovery panic убьёт процесс (Go HTTP-сервер ловит panic, но это не повод не иметь явный middleware с логированием).

## Когда писать своё, когда брать готовое

- `chi/middleware`, `gorilla/handlers` — готовые качественные реализации.
- Своё — когда хочется специфики (например, кастомный формат логов, специальная trace propagation).

## Правило

- Middleware — короткие, по одной ответственности.
- Цепочка чёткая, в `main.go` явная.
- Cross-cutting — middleware. Бизнес — handler.

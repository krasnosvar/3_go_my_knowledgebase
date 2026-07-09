# Backend / REST API на Go

Стандартное приложение Go-разработчика. HTTP-сервер с handler'ами, бизнес-логикой, БД, авторизацией, observability.

## Типичная структура

```
service/
├── go.mod
├── cmd/
│   └── api/main.go
├── internal/
│   ├── http/         (handlers, middleware, router)
│   ├── service/      (бизнес-логика)
│   ├── repository/   (БД-доступ)
│   ├── domain/       (типы User, Order и т.п.)
│   ├── auth/
│   ├── config/
│   └── observability/
├── migrations/       (SQL-миграции)
├── api/              (OpenAPI / Proto)
├── deployments/      (Dockerfile, k8s manifests)
└── README.md
```

`internal/` — пакеты, недоступные другим модулям. Защита публичного API.

## Слои

1. **Transport (handler)** — HTTP-специфика, decode/encode, status mapping.
2. **Service** — бизнес-правила, координация.
3. **Repository** — БД, внешние API.
4. **Domain** — типы, не зависящие от слоёв (User, Money, Status).

См. `1_app_parts` для детальных слоёв.

## Tech-стек по умолчанию

- HTTP: stdlib `net/http` + `chi` для удобного роутинга/middleware.
- БД: PostgreSQL + `pgx` или `sqlc`.
- Логи: `log/slog` (Go 1.21+).
- Метрики: `prometheus/client_golang`.
- Traces: `opentelemetry-go`.
- Конфиг: `envconfig` или `viper`.
- Тесты: stdlib + `testcontainers-go` для интеграции.

## main.go pattern

```go
func main() {
    cfg := config.Load()
    if err := cfg.Validate(); err != nil { log.Fatal(err) }

    db := mustOpen(cfg.DB.DSN)
    defer db.Close()

    repo := repository.New(db)
    svc := service.New(repo, deps...)

    router := chi.NewRouter()
    router.Use(recovery, logging, metrics)
    api.RegisterRoutes(router, svc)

    srv := &http.Server{
        Addr:    cfg.HTTP.Addr,
        Handler: router,
        ReadTimeout:  cfg.HTTP.ReadTimeout,
        WriteTimeout: cfg.HTTP.WriteTimeout,
    }

    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    go func() {
        if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            slog.Error("server", "err", err)
        }
    }()

    <-ctx.Done()
    sdCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _ = srv.Shutdown(sdCtx)
}
```

## Что отдельно учесть

- **API contract** — OpenAPI spec, сгенерированный или вручную написанный.
- **Pagination** — cursor-based для больших списков.
- **Versioning** — `/v1/users` или header. Не «сломать клиентов» — must.
- **Rate limit** — на endpoint и/или на user.
- **Timeouts на каждый внешний вызов** — БД, API, Redis.

## Правило

- Слои разделены, тесты на каждом уровне.
- Конфиг через env, секреты через secret manager.
- Observability с первого дня.
- Graceful shutdown.
- Health/ready probes для kubernetes.

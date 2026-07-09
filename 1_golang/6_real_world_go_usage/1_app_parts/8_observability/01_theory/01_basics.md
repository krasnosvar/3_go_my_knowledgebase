# Observability: три кита

Observability — способность понять, **что** и **почему** происходит в приложении в проде. Базируется на трёх типах данных:

- **Logs** — что случилось.
- **Metrics** — сколько / как часто / как быстро.
- **Traces** — путь конкретного запроса через систему.

## Logs

С Go 1.21 в stdlib есть `log/slog`:

```go
import "log/slog"

slog.Info("user registered",
    "user_id", u.ID,
    "email", u.Email,
)

slog.Error("save failed",
    "err", err,
    "user_id", u.ID,
)
```

Структурный лог = key-value пары → легко парсится в Loki/ELK/CloudWatch.

Уровни:

- DEBUG — детали для разработки.
- INFO — нормальные события.
- WARN — что-то пошло необычно, но не сломалось.
- ERROR — операция провалилась.

Не логируй каждый успешный запрос на INFO — только бизнес-события и аномалии. Иначе лог становится swap-ом.

## Metrics

Стандарт de-facto — Prometheus client `prometheus/client_golang`:

```go
var (
    requests = prometheus.NewCounterVec(
        prometheus.CounterOpts{Name: "http_requests_total"},
        []string{"method", "path", "status"},
    )
    duration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{Name: "http_request_duration_seconds"},
        []string{"method", "path"},
    )
)
```

Типы:

- **Counter** — только растёт (количество запросов, ошибок).
- **Gauge** — может расти и падать (число активных коннектов, использование памяти).
- **Histogram** — распределение значений (latency, размер запроса).
- **Summary** — то же что histogram, но с предрассчитанными квантилями.

Endpoint `/metrics` отдаёт всё — Prometheus сам собирает.

## Traces

OpenTelemetry — стандарт. Trace — это дерево spans, каждый span — один шаг обработки запроса.

```go
ctx, span := tracer.Start(ctx, "service.GetUser")
defer span.End()
span.SetAttributes(attribute.Int("user.id", id))
```

Контекст несёт trace_id через все вызовы. На дашборде видно, что запрос провёл 80% времени в Postgres → найдена точка боли.

## Минимум для production

- `slog` + структурный JSON в stdout.
- Prometheus `/metrics` endpoint.
- HTTP middleware: `request_id` + время ответа в каждом логе.
- OpenTelemetry — если есть несколько сервисов или distributed system.

## Healthchecks

Не путай с observability, но рядом:

- `/healthz` — liveness: «процесс жив». Простое 200.
- `/readyz` — readiness: «готов принимать трафик» (БД доступна, миграции применены).

K8s по этим endpoint'ам решает, рестартить ли pod и слать ли в него трафик.

## Правило

- Logs — для понимания, что произошло.
- Metrics — для алертов и трендов.
- Traces — для поиска bottleneck'ов.
- Все три вместе — для нормальной диагностики прода.

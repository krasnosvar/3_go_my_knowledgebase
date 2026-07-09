# Observability: pitfalls

## 1. Logs из stdout без structure

```go
log.Println("user", u.ID, "logged in from", ip)
```

Парсить такое в Loki/ELK — ад. Используй `slog` или другой структурный логгер с JSON output.

## 2. INFO на каждый успешный запрос

```go
slog.Info("request handled", "method", r.Method, "path", r.URL.Path)
```

При 1000 RPS это миллионы строк лога в день. Стоимость хранения и поиска растёт.

Решения:

- INFO только на бизнес-события (registration, payment, error).
- DEBUG для детальных логов под нагрузкой.
- Сэмплинг: логировать 1% запросов.

## 3. Логирование секретов

```go
slog.Info("login", "password", input.Password)
slog.Info("api request", "headers", r.Header) // Authorization туда попадёт
```

Маскируй чувствительные поля. Используй `slog.HandlerOptions.ReplaceAttr` для глобального правила.

## 4. Метрики с unbounded labels

```go
requests.WithLabelValues(userID).Inc()
```

`userID` — миллионы значений → миллионы time series в Prometheus → OOM. Лимит — пара тысяч уникальных значений per metric.

Безопасные labels: `method`, `path` (нормализованный), `status_code`, `service`. Опасные: `user_id`, `request_id`, `email`.

## 5. Histogram bucket'ы по умолчанию

Дефолтные бакеты Prometheus подходят не всегда. Для latency 1-100 ms они слишком грубые. Настраивай:

```go
prometheus.HistogramOpts{
    Buckets: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0},
}
```

## 6. Алерты на голые counter'ы

«Сработало 100 ошибок» — это много или нормально? Без rate / процента — непонятно. Алерты:

```
rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m]) > 0.05
```

Доля 5xx > 5% за 5 мин.

## 7. Trace без context propagation

```go
ctx, span := tracer.Start(ctx, "fetch")
defer span.End()
go func() {
    // ctx не пробрасывается → trace разорван
    fetch(context.Background(), ...)
}()
```

Контекст с trace передавай везде, иначе картинка распадается.

## 8. Healthcheck тяжёлый

```go
func healthz(w, r) {
    if err := db.Ping(ctx); err != nil { ... }
    if err := redis.Ping(ctx); err != nil { ... }
    // ещё 5 проверок
}
```

Liveness probe вызывается часто (каждые 10s). Тяжёлая проверка с timeout'ами и зависимостями → ложные «pod dead» при кратковременной флуктуации.

Liveness — простой `200 OK` если процесс отвечает.
Readiness — здесь уместно проверить ключевые зависимости.

## 9. Metric leak

```go
counter := prometheus.NewCounter(...)
// в обработчике запроса
counter.Inc() // создаётся новый counter каждый раз
```

Метрика должна быть глобальной (или поле сервиса), не локальной. Иначе она исчезает с горутиной.

## 10. Только prod-логи без dev/staging

Тесты прошли локально, прод сломался — а у тебя нет логов в staging. Включай минимальный observability стек на staging — дешевле, чем долго чинить инциденты.

## Правило

- Структурный лог.
- Метрики bounded labels.
- Алерты на rate / процент, не на количество.
- Liveness простой, readiness осмысленный.
- Не логировать секреты.

# 09. Observability (Наблюдаемость)

Как понимать, что происходит внутри системы в проде, не влезая в код: три столпа наблюдаемости (логи, метрики, трейсы) плюс эксплуатационные детали — пробы здоровья, SLI/SLO/SLA и алертинг. Теория подкреплена runnable-примерами на стандартной библиотеке Go (включая `log/slog`) и упражнениями с самопроверкой.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_three_pillars.md](./01_theory/01_three_pillars.md) — три столпа (логи/метрики/трейсы): чем отличаются, что дают и чего стоят, когда какой использовать; мониторинг vs наблюдаемость.
    *   [02_logging.md](./01_theory/02_logging.md) — структурные логи vs текст, уровни, корреляция запросов (`request_id`/`trace_id`), `log/slog` в Go: JSON-хендлер, атрибуты, дочерний логгер через `With`.
    *   [03_metrics.md](./01_theory/03_metrics.md) — типы метрик (counter/gauge/histogram), методики RED и USE, перцентили vs среднее, pull-модель Prometheus, exposition-формат, кардинальность меток.
    *   [04_tracing.md](./01_theory/04_tracing.md) — распределённый трейсинг: trace/span, `trace_id`/`span_id`, context propagation (W3C `traceparent`), OpenTelemetry концептуально, сэмплирование head/tail.
    *   [05_health_alerting.md](./01_theory/05_health_alerting.md) — пробы liveness vs readiness и их роль в оркестраторе, SLI/SLO/SLA и error budget кратко, алертинг на симптомы, борьба с шумом.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практика: разбор инцидента по трём столпам, что логировать/измерять, SLO-ориентированный алертинг, стоимость телеметрии, диагностика от симптома к причине.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, детерминированные, завершаются; серверы на `127.0.0.1:0`, клиент с таймаутом):
    *   [01_structured_logging](./02_examples/01_structured_logging/main.go) — структурное логирование через `log/slog`: JSON с полями и уровнями, дочерний логгер с `request_id`, отбрасывание Debug ниже порога.
    *   [02_inprocess_metrics](./02_examples/02_inprocess_metrics/main.go) — самодельные counter/gauge на `sync/atomic` и эндпоинт `/metrics` в текстовом формате Prometheus; клиент забирает срез (pull) и печатает его.
    *   [03_health_readiness](./02_examples/03_health_readiness/main.go) — HTTP-сервер с `/healthz` (liveness) и `/readyz` (readiness): переключаем флаг готовности, дёргаем оба эндпоинта до и после, graceful shutdown.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_red_metrics` — посчитать RED (rate/errors/duration) по списку запросов.
    *   `02_prometheus_format` — отформатировать метрику в строку Prometheus-экспозиции.
    *   `03_readiness_handler` — readiness-хендлер, отвечающий 200/503 по флагу готовности.
    *   `04_parse_log_line` — распарсить logfmt-строку (`key=value`) в `map[string]string`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_three_pillars.md](./01_theory/01_three_pillars.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, без внешних зависимостей и без `go.mod`):
   ```bash
   go run ./02_examples/01_structured_logging/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/01_red_metrics/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

# Observability

Логи, метрики и трейсы: структурные логи `slog`, счётчики и гистограммы в формате Prometheus, проброс trace ID.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#8-observability) — раздел «8. Observability» общей шпаргалки частей приложения.

## Пререквизиты

[5_middleware_chain](../5_middleware_chain/), `sync` из [3_concurrency/04_sync_package](../../../3_concurrency/04_sync_package/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — Observability: три кита.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Observability: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_structured_log_shape](./02_examples/01_structured_log_shape/main.go) — структурный лог `log/slog`.
    *   [02_metrics_counter_shape](./02_examples/02_metrics_counter_shape/main.go) — счётчик запросов.
    *   [03_trace_id_propagation](./02_examples/03_trace_id_propagation/main.go) — проброс trace ID через `context`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_structured_log](./03_exercises/01_structured_log/main.go) — структурные логи через log/slog.
    *   [02_counter_exposition](./03_exercises/02_counter_exposition/main.go) — счётчик с лейблами и вывод в формате Prometheus.
    *   [03_latency_histogram](./03_exercises/03_latency_histogram/main.go) — гистограмма латентности (как у Prometheus).
    *   [04_trace_propagation](./03_exercises/04_trace_propagation/main.go) — проброс trace ID между сервисами.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/8_observability` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/8_observability/02_examples/01_structured_log_shape
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/8_observability/03_exercises/01_structured_log
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/8_observability/04_exercises_answers/01_structured_log_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [2_go_apps_types_real](../../2_go_apps_types_real/) — как эти части складываются в приложения разных типов.

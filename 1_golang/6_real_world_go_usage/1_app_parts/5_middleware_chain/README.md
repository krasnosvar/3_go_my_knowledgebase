# Middleware chain

Middleware — сквозная логика вокруг handler'ов: логирование, recovery, request ID, метрики. Важно понимать порядок цепочки.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#5-middleware-chain) — раздел «5. Middleware chain» общей шпаргалки частей приложения.

## Пререквизиты

[1_http_handlers](../1_http_handlers/), `context` из [3_concurrency/05_context](../../../3_concurrency/05_context/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — HTTP middleware.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Middleware: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_logging_middleware](./02_examples/01_logging_middleware/main.go) — logging middleware вокруг handler'а.
    *   [02_timeout_shape](./02_examples/02_timeout_shape/main.go) — таймаут через `context.WithTimeout`.
    *   [03_request_id_shape](./02_examples/03_request_id_shape/main.go) — request ID в `context`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_chain_order](./03_exercises/01_chain_order/main.go) — сборка цепочки middleware.
    *   [02_recovery](./03_exercises/02_recovery/main.go) — recovery middleware.
    *   [03_request_id](./03_exercises/03_request_id/main.go) — request ID через заголовок и context.
    *   [04_status_recorder](./03_exercises/04_status_recorder/main.go) — middleware, который видит статус ответа.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/5_middleware_chain` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/5_middleware_chain/02_examples/01_logging_middleware
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/5_middleware_chain/03_exercises/01_chain_order
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/5_middleware_chain/04_exercises_answers/01_chain_order_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [6_auth_and_permissions](../6_auth_and_permissions/) — Auth and permissions.

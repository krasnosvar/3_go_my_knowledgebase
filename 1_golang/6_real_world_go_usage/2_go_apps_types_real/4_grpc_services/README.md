# gRPC services

Ключевые идеи gRPC-сервиса без внешних зависимостей: status codes, deadlines, валидация, цепочка интерсепторов, metadata, стриминг, ретраи, health.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#4-grpc-services) — раздел «4. gRPC services» общей шпаргалки типов Go-приложений.

## Пререквизиты

[2_backend](../2_backend/), [3_concurrency/05_context](../../../3_concurrency/05_context/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — gRPC services.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_status_error](./02_examples/01_status_error/main.go) — ошибка со статус-кодом в стиле gRPC.
    *   [02_context_timeout](./02_examples/02_context_timeout/main.go) — deadline в `context`.
    *   [03_unary_interceptor_shape](./02_examples/03_unary_interceptor_shape/main.go) — форма unary-интерсептора.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_status_codes](./03_exercises/01_status_codes/main.go) — gRPC status codes.
    *   [02_deadline_handling](./03_exercises/02_deadline_handling/main.go) — бюджет времени из deadline клиента.
    *   [03_request_validation](./03_exercises/03_request_validation/main.go) — валидация запроса с перечнем нарушений.
    *   [04_interceptor_chain](./03_exercises/04_interceptor_chain/main.go) — цепочка unary-интерсепторов.
    *   [05_auth_metadata](./03_exercises/05_auth_metadata/main.go) — аутентификация по metadata.
    *   [06_stream_loop](./03_exercises/06_stream_loop/main.go) — server-streaming цикл.
    *   [07_retry_policy](./03_exercises/07_retry_policy/main.go) — политика ретраев клиента.
    *   [08_health_rpc](./03_exercises/08_health_rpc/main.go) — сервис здоровья (grpc.health.v1).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/4_grpc_services` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/4_grpc_services/02_examples/01_status_error
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/4_grpc_services/03_exercises/01_status_codes
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/4_grpc_services/04_exercises_answers/01_status_codes_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [5_workers_cron](../5_workers_cron/) — Workers and cron.

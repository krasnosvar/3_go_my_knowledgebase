# Backend API

HTTP backend целиком: маршруты, единый формат ответов и ошибок, middleware, абстракция хранилища, таймауты, пагинация и health-checks.

## Пререквизиты

весь [1_app_parts](../../1_app_parts/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — Backend / REST API на Go.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_http_handlers](./02_examples/01_http_handlers/main.go) — JSON health-handler на `ServeMux`, запрос через `httptest.NewServer`.
    *   [02_middleware](./02_examples/02_middleware/main.go) — logging middleware вокруг mux, два запроса через `httptest.NewServer`.
    *   [03_repository_layer](./02_examples/03_repository_layer/main.go) — слой репозитория за интерфейсом.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_routes](./03_exercises/01_routes/main.go) — маршруты REST-ресурса на stdlib ServeMux (Go 1.22+).
    *   [02_json_response](./03_exercises/02_json_response/main.go) — единые хелперы JSON-ответов.
    *   [03_error_handling](./03_exercises/03_error_handling/main.go) — handler, возвращающий error.
    *   [04_middleware_chain](./03_exercises/04_middleware_chain/main.go) — цепочка logging + auth с прерыванием.
    *   [05_repo_abstraction](./03_exercises/05_repo_abstraction/main.go) — handler зависит от интерфейса хранилища.
    *   [06_context_timeout](./03_exercises/06_context_timeout/main.go) — таймаут на вызов зависимости.
    *   [07_pagination](./03_exercises/07_pagination/main.go) — пагинация списка в API.
    *   [08_healthcheck](./03_exercises/08_healthcheck/main.go) — liveness и readiness.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/2_backend` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/2_backend/02_examples/01_http_handlers
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/2_backend/03_exercises/01_routes
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/2_backend/04_exercises_answers/01_routes_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [3_system_apps](../3_system_apps/) — System apps / daemons.

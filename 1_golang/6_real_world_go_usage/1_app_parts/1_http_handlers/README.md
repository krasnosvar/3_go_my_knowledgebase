# HTTP handlers

Handler — граница между HTTP и бизнес-логикой: декодирует запрос, вызывает сервис, превращает результат и ошибку в статус и тело ответа. Здесь — как держать его тонким и не терять ошибки.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#1-http-handlers) — раздел «1. HTTP handlers» общей шпаргалки частей приложения.

## Пререквизиты

[1_go_basics](../../../1_go_basics/) (функции, структуры, ошибки), [2_oop](../../../2_oop/) (интерфейсы), [4_standard_library/05_net_http_net](../../../4_standard_library/05_net_http_net/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — HTTP handlers: основы.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — HTTP handlers: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_health_handler](./02_examples/01_health_handler/main.go) — health-handler, отдающий JSON, проверка через `httptest.NewRecorder`.
    *   [02_query_params](./02_examples/02_query_params/main.go) — чтение query-параметров из `r.URL.Query()`.
    *   [03_json_input_shape](./02_examples/03_json_input_shape/main.go) — декодирование JSON-тела запроса в структуру.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_status_mapping](./03_exercises/01_status_mapping/main.go) — маппинг доменных ошибок в HTTP-статусы.
    *   [02_query_params](./03_exercises/02_query_params/main.go) — разбор query-параметров пагинации.
    *   [03_json_create](./03_exercises/03_json_create/main.go) — handler создания ресурса с JSON-телом.
    *   [04_method_routing](./03_exercises/04_method_routing/main.go) — маршрутизация по методу и path-параметру (Go 1.22+ ServeMux).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/1_http_handlers` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/1_http_handlers/02_examples/01_health_handler
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/1_http_handlers/03_exercises/01_status_mapping
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/1_http_handlers/04_exercises_answers/01_status_mapping_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [2_service_layer](../2_service_layer/) — Service layer.

# Service layer

Сервис хранит бизнес-правила и не знает про HTTP и SQL. Зависимости приходят через интерфейсы, поэтому сервис тестируется без базы и сети.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#2-service-layer) — раздел «2. Service layer» общей шпаргалки частей приложения.

## Пререквизиты

[1_http_handlers](../1_http_handlers/), интерфейсы из [2_oop](../../../2_oop/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — Service layer.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Service layer: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_service_validate](./02_examples/01_service_validate/main.go) — валидация входа на уровне сервиса.
    *   [02_service_with_repo](./02_examples/02_service_with_repo/main.go) — сервис, зависящий от интерфейса репозитория.
    *   [03_service_errors](./02_examples/03_service_errors/main.go) — доменные ошибки сервиса.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_validate_input](./03_exercises/01_validate_input/main.go) — валидация входа в service layer.
    *   [02_translate_errors](./03_exercises/02_translate_errors/main.go) — сервис зависит от интерфейса репозитория и переводит его ошибки.
    *   [03_business_rule](./03_exercises/03_business_rule/main.go) — бизнес-правило перевода денег.
    *   [04_unique_email](./03_exercises/04_unique_email/main.go) — правило уникальности и зависимость от часов.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/2_service_layer` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/2_service_layer/02_examples/01_service_validate
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/2_service_layer/03_exercises/01_validate_input
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/2_service_layer/04_exercises_answers/01_validate_input_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [3_repository_layer](../3_repository_layer/) — Repository layer.

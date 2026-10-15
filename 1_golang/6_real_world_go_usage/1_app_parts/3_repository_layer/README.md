# Repository layer

Репозиторий прячет хранилище за узким интерфейсом: CRUD, маппинг ошибок драйвера в доменные, пагинация.

## Пререквизиты

[2_service_layer](../2_service_layer/), основы SQL из [2_db_kafka_redis/1_postgres](../../../../2_db_kafka_redis/1_postgres/) — желательно.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — Repository layer.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Repository: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_repo_interface](./02_examples/01_repo_interface/main.go) — контракт репозитория и in-memory реализация.
    *   [02_not_found_pattern](./02_examples/02_not_found_pattern/main.go) — паттерн «не найдено» как доменная ошибка.
    *   [03_pagination_shape](./02_examples/03_pagination_shape/main.go) — форма offset-пагинации.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_inmem_repo](./03_exercises/01_inmem_repo/main.go) — in-memory репозиторий по контракту.
    *   [02_map_db_errors](./03_exercises/02_map_db_errors/main.go) — маппинг ошибок драйвера в доменные.
    *   [03_offset_pagination](./03_exercises/03_offset_pagination/main.go) — offset-пагинация в репозитории.
    *   [04_keyset_pagination](./03_exercises/04_keyset_pagination/main.go) — keyset (cursor) пагинация.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/3_repository_layer` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/3_repository_layer/02_examples/01_repo_interface
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/3_repository_layer/03_exercises/01_inmem_repo
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/3_repository_layer/04_exercises_answers/01_inmem_repo_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [4_config_and_env](../4_config_and_env/) — Config and env.

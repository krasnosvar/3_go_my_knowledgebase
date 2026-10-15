# Config and env

Конфигурация по 12-factor: переменные окружения, значения по умолчанию, типизированный разбор, валидация и приоритет источников.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#4-config-and-env) — раздел «4. Config and env» общей шпаргалки частей приложения.

## Пререквизиты

[1_go_basics](../../../1_go_basics/) (maps, ошибки), `strconv` и `time` из [4_standard_library](../../../4_standard_library/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — Configuration.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Configuration: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_env_read](./02_examples/01_env_read/main.go) — чтение переменной окружения.
    *   [02_defaults_merge](./02_examples/02_defaults_merge/main.go) — значение по умолчанию, если переменная не задана.
    *   [03_config_validation](./02_examples/03_config_validation/main.go) — простая проверка конфига.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_env_defaults](./03_exercises/01_env_defaults/main.go) — чтение конфига из окружения с дефолтами.
    *   [02_typed_parse](./03_exercises/02_typed_parse/main.go) — типизированный разбор переменных окружения.
    *   [03_validate_all](./03_exercises/03_validate_all/main.go) — валидация конфига со сбором ВСЕХ ошибок.
    *   [04_priority_merge](./03_exercises/04_priority_merge/main.go) — приоритет источников конфигурации.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/4_config_and_env` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/4_config_and_env/02_examples/01_env_read
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/4_config_and_env/03_exercises/01_env_defaults
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/4_config_and_env/04_exercises_answers/01_env_defaults_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [5_middleware_chain](../5_middleware_chain/) — Middleware chain.

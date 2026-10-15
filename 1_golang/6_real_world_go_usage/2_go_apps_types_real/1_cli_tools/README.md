# CLI tools

CLI на Go: свой `flag.FlagSet`, подкоманды, fallback на окружение, коды выхода, stdout/stderr, dry-run и help. Логика выносится в `run(...)`, чтобы её можно было тестировать.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#1-cli-tools) — раздел «1. CLI tools» общей шпаргалки типов Go-приложений.

## Пререквизиты

[1_app_parts/4_config_and_env](../../1_app_parts/4_config_and_env/), `io`/`os` из [4_standard_library](../../../4_standard_library/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — CLI tools на Go.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_flag_parsing](./02_examples/01_flag_parsing/main.go) — разбор флагов пакетом `flag`.
    *   [02_subcommand_style](./02_examples/02_subcommand_style/main.go) — подкоманды через `os.Args`.
    *   [03_config_priority](./02_examples/03_config_priority/main.go) — приоритет источников конфига.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_arg_parse](./03_exercises/01_arg_parse/main.go) — разбор аргументов через flag.FlagSet.
    *   [02_subcommands](./03_exercises/02_subcommands/main.go) — подкоманды без сторонних библиотек.
    *   [03_env_fallback](./03_exercises/03_env_fallback/main.go) — флаг с fallback на переменную окружения.
    *   [04_exit_codes](./03_exercises/04_exit_codes/main.go) — коды выхода.
    *   [05_stderr_output](./03_exercises/05_stderr_output/main.go) — stdout для данных, stderr для диагностики.
    *   [06_config_merge](./03_exercises/06_config_merge/main.go) — слияние конфига из файла и флагов.
    *   [07_dry_run](./03_exercises/07_dry_run/main.go) — режим --dry-run.
    *   [08_help_text](./03_exercises/08_help_text/main.go) — help и -h.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/1_cli_tools` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/1_cli_tools/02_examples/01_flag_parsing
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/1_cli_tools/03_exercises/01_arg_parse
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/1_cli_tools/04_exercises_answers/01_arg_parse_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [2_backend](../2_backend/) — Backend API.

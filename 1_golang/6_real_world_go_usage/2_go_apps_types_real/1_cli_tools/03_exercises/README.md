# Упражнения: CLI tools

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_arg_parse
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_arg_parse_answers
```

## Список задач (от простого к сложному)

1. [`01_arg_parse`](./01_arg_parse/main.go) — разбор аргументов через flag.FlagSet; решение: [`01_arg_parse_answers`](../04_exercises_answers/01_arg_parse_answers/main.go).
2. [`02_subcommands`](./02_subcommands/main.go) — подкоманды без сторонних библиотек; решение: [`02_subcommands_answers`](../04_exercises_answers/02_subcommands_answers/main.go).
3. [`03_env_fallback`](./03_env_fallback/main.go) — флаг с fallback на переменную окружения; решение: [`03_env_fallback_answers`](../04_exercises_answers/03_env_fallback_answers/main.go).
4. [`04_exit_codes`](./04_exit_codes/main.go) — коды выхода; решение: [`04_exit_codes_answers`](../04_exercises_answers/04_exit_codes_answers/main.go).
5. [`05_stderr_output`](./05_stderr_output/main.go) — stdout для данных, stderr для диагностики; решение: [`05_stderr_output_answers`](../04_exercises_answers/05_stderr_output_answers/main.go).
6. [`06_config_merge`](./06_config_merge/main.go) — слияние конфига из файла и флагов; решение: [`06_config_merge_answers`](../04_exercises_answers/06_config_merge_answers/main.go).
7. [`07_dry_run`](./07_dry_run/main.go) — режим --dry-run; решение: [`07_dry_run_answers`](../04_exercises_answers/07_dry_run_answers/main.go).
8. [`08_help_text`](./08_help_text/main.go) — help и -h; решение: [`08_help_text_answers`](../04_exercises_answers/08_help_text_answers/main.go).

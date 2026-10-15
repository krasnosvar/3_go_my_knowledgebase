# Упражнения: Config and env

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_env_defaults
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_env_defaults_answers
```

## Список задач (от простого к сложному)

1. [`01_env_defaults`](./01_env_defaults/main.go) — чтение конфига из окружения с дефолтами; решение: [`01_env_defaults_answers`](../04_exercises_answers/01_env_defaults_answers/main.go).
2. [`02_typed_parse`](./02_typed_parse/main.go) — типизированный разбор переменных окружения; решение: [`02_typed_parse_answers`](../04_exercises_answers/02_typed_parse_answers/main.go).
3. [`03_validate_all`](./03_validate_all/main.go) — валидация конфига со сбором ВСЕХ ошибок; решение: [`03_validate_all_answers`](../04_exercises_answers/03_validate_all_answers/main.go).
4. [`04_priority_merge`](./04_priority_merge/main.go) — приоритет источников конфигурации; решение: [`04_priority_merge_answers`](../04_exercises_answers/04_priority_merge_answers/main.go).

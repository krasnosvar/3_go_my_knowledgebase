# Упражнения: type_assertions_switches

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, всё детерминированно.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_assertion_ok/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_assertion_ok_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_assertion_ok`](./01_assertion_ok/main.go) — `asString` через comma-ok. Ответ: [`01_assertion_ok_answers`](../04_exercises_answers/01_assertion_ok_answers/main.go).
2. [`02_assertion_fail`](./02_assertion_fail/main.go) — `mustString` без паники, с ошибкой и `%T`. Ответ: [`02_assertion_fail_answers`](../04_exercises_answers/02_assertion_fail_answers/main.go).
3. [`03_type_switch_basic`](./03_type_switch_basic/main.go) — `kind` по базовым типам. Ответ: [`03_type_switch_basic_answers`](../04_exercises_answers/03_type_switch_basic_answers/main.go).
4. [`04_type_switch_struct`](./04_type_switch_struct/main.go) — площадь фигуры по типу структуры, включая `*Rect` и nil. Ответ: [`04_type_switch_struct_answers`](../04_exercises_answers/04_type_switch_struct_answers/main.go).
5. [`05_any_parser`](./05_any_parser/main.go) — `format` для int/string/[]int/map/nil. Ответ: [`05_any_parser_answers`](../04_exercises_answers/05_any_parser_answers/main.go).
6. [`06_safe_cast`](./06_safe_cast/main.go) — `toIntOr` — приведение «чего угодно» к int. Ответ: [`06_safe_cast_answers`](../04_exercises_answers/06_safe_cast_answers/main.go).
7. [`07_interface_branching`](./07_interface_branching/main.go) — ветвление по интерфейсам `error` и `fmt.Stringer`. Ответ: [`07_interface_branching_answers`](../04_exercises_answers/07_interface_branching_answers/main.go).
8. [`08_json_like`](./08_json_like/main.go) — `getPath`/`getInt` по JSON-подобной `map[string]any`. Ответ: [`08_json_like_answers`](../04_exercises_answers/08_json_like_answers/main.go).

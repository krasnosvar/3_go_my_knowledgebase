# Упражнения: generics

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, всё детерминированно.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_generic_func/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_generic_func_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_generic_func`](./01_generic_func/main.go) — `Map` и `Filter`. Ответ: [`01_generic_func_answers`](../04_exercises_answers/01_generic_func_answers/main.go).
2. [`02_constraints`](./02_constraints/main.go) — `Sum` со своим `Number` и `Max` с `cmp.Ordered`. Ответ: [`02_constraints_answers`](../04_exercises_answers/02_constraints_answers/main.go).
3. [`03_generic_struct`](./03_generic_struct/main.go) — `Optional[T]`: `Some`, `Get`, `OrElse`. Ответ: [`03_generic_struct_answers`](../04_exercises_answers/03_generic_struct_answers/main.go).
4. [`04_generic_interface`](./04_generic_interface/main.go) — `memRepo[T]` под интерфейс `Repository[T]`. Ответ: [`04_generic_interface_answers`](../04_exercises_answers/04_generic_interface_answers/main.go).
5. [`05_comparable`](./05_comparable/main.go) — `Index` и `Unique` с `comparable`. Ответ: [`05_comparable_answers`](../04_exercises_answers/05_comparable_answers/main.go).
6. [`06_type_set`](./06_type_set/main.go) — type set `Integer` с `~`, `Abs`, `Clamp`. Ответ: [`06_type_set_answers`](../04_exercises_answers/06_type_set_answers/main.go).
7. [`07_any_alias`](./07_any_alias/main.go) — `any` как алиас `interface{}`: `Join` и `Count`. Ответ: [`07_any_alias_answers`](../04_exercises_answers/07_any_alias_answers/main.go).
8. [`08_stack`](./08_stack/main.go) — обобщённый стек `Stack[T]`. Ответ: [`08_stack_answers`](../04_exercises_answers/08_stack_answers/main.go).

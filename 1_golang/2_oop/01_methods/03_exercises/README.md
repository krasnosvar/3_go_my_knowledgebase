# Упражнения: methods

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, всё детерминированно.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_value_receiver/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_value_receiver_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_value_receiver`](./01_value_receiver/main.go) — `Area`/`Perimeter` у `Rect` с value receiver. Ответ: [`01_value_receiver_answers`](../04_exercises_answers/01_value_receiver_answers/main.go).
2. [`02_pointer_receiver`](./02_pointer_receiver/main.go) — `Inc`/`Add`/`Reset` у счётчика с pointer receiver. Ответ: [`02_pointer_receiver_answers`](../04_exercises_answers/02_pointer_receiver_answers/main.go).
3. [`03_method_set`](./03_method_set/main.go) — исправить receiver'ы так, чтобы `T{}` реализовывал `Namer`, а `Setter` — только `*T`. Ответ: [`03_method_set_answers`](../04_exercises_answers/03_method_set_answers/main.go).
4. [`04_constructor_usage`](./04_constructor_usage/main.go) — конструктор `NewUser` с ролью по умолчанию и метод `Greeting`. Ответ: [`04_constructor_usage_answers`](../04_exercises_answers/04_constructor_usage_answers/main.go).
5. [`05_mutation`](./05_mutation/main.go) — баг потерянной мутации в `Deposit` + `Withdraw` с проверкой баланса. Ответ: [`05_mutation_answers`](../04_exercises_answers/05_mutation_answers/main.go).
6. [`06_copy_cost`](./06_copy_cost/main.go) — перевести методы «тяжёлой» структуры на pointer receiver. Ответ: [`06_copy_cost_answers`](../04_exercises_answers/06_copy_cost_answers/main.go).
7. [`07_fluent_style`](./07_fluent_style/main.go) — построитель SQL-запроса с цепочкой методов. Ответ: [`07_fluent_style_answers`](../04_exercises_answers/07_fluent_style_answers/main.go).
8. [`08_validation_method`](./08_validation_method/main.go) — `Validate() error` с порядком проверок. Ответ: [`08_validation_method_answers`](../04_exercises_answers/08_validation_method_answers/main.go).

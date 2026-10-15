# Упражнения: interfaces

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, всё детерминированно.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_basic_interface/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_basic_interface_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_basic_interface`](./01_basic_interface/main.go) — `Cat`/`Dog` реализуют `Speaker`, функция `chorus`. Ответ: [`01_basic_interface_answers`](../04_exercises_answers/01_basic_interface_answers/main.go).
2. [`02_pointer_receiver`](./02_pointer_receiver/main.go) — `*counter` реализует `Counter`, а значение — нет. Ответ: [`02_pointer_receiver_answers`](../04_exercises_answers/02_pointer_receiver_answers/main.go).
3. [`03_type_assertions`](./03_type_assertions/main.go) — безопасный `v.(int)` в форме comma-ok. Ответ: [`03_type_assertions_answers`](../04_exercises_answers/03_type_assertions_answers/main.go).
4. [`04_type_switch`](./04_type_switch/main.go) — `describe` через type switch, включая `nil` и `error`. Ответ: [`04_type_switch_answers`](../04_exercises_answers/04_type_switch_answers/main.go).
5. [`05_nil_interface`](./05_nil_interface/main.go) — исправить возврат typed nil как `error`. Ответ: [`05_nil_interface_answers`](../04_exercises_answers/05_nil_interface_answers/main.go).
6. [`06_small_interfaces`](./06_small_interfaces/main.go) — `copyAll` поверх маленьких `Reader`/`Writer`. Ответ: [`06_small_interfaces_answers`](../04_exercises_answers/06_small_interfaces_answers/main.go).
7. [`07_dependency_injection`](./07_dependency_injection/main.go) — `Service.Notify` через внедрённый `Sender`. Ответ: [`07_dependency_injection_answers`](../04_exercises_answers/07_dependency_injection_answers/main.go).
8. [`08_mocking`](./08_mocking/main.go) — `greet` с подменяемыми часами `Clock`. Ответ: [`08_mocking_answers`](../04_exercises_answers/08_mocking_answers/main.go).

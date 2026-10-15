# Упражнения: embedding

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, всё детерминированно.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_basic_embedding/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_basic_embedding_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_basic_embedding`](./01_basic_embedding/main.go) — встроить `A` в `B`, чтобы появился `Hello`. Ответ: [`01_basic_embedding_answers`](../04_exercises_answers/01_basic_embedding_answers/main.go).
2. [`02_method_promotion`](./02_method_promotion/main.go) — встроить `Logger` и собрать сервис конструктором. Ответ: [`02_method_promotion_answers`](../04_exercises_answers/02_method_promotion_answers/main.go).
3. [`03_override`](./03_override/main.go) — свой `Print` у `Child`, вызывающий `Parent.Print`. Ответ: [`03_override_answers`](../04_exercises_answers/03_override_answers/main.go).
4. [`04_conflict`](./04_conflict/main.go) — разрешить конфликт `Name()` двух встроенных типов. Ответ: [`04_conflict_answers`](../04_exercises_answers/04_conflict_answers/main.go).
5. [`05_pointer_embedding`](./05_pointer_embedding/main.go) — инициализировать встроенный `*Counter` в конструкторе. Ответ: [`05_pointer_embedding_answers`](../04_exercises_answers/05_pointer_embedding_answers/main.go).
6. [`06_composition`](./06_composition/main.go) — `API` из `Auth` и журнала аудита. Ответ: [`06_composition_answers`](../04_exercises_answers/06_composition_answers/main.go).
7. [`07_interface`](./07_interface/main.go) — реализовать `Runner` через встроенный `Base`. Ответ: [`07_interface_answers`](../04_exercises_answers/07_interface_answers/main.go).
8. [`08_constructor`](./08_constructor/main.go) — `NewChild` с заполнением встроенного `Base`. Ответ: [`08_constructor_answers`](../04_exercises_answers/08_constructor_answers/main.go).

# Упражнения: Тема 06. Указатели (Pointers)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`06_pointers/`):

```bash
go run ./03_exercises/01_basic_deref                 # заготовка
go run ./04_exercises_answers/01_basic_deref_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_basic_deref`](./01_basic_deref/main.go) | разыменование с безопасным возвратом дефолтного значения при nil. | [`01_basic_deref_answers`](../04_exercises_answers/01_basic_deref_answers/main.go) |
| 2 | [`02_mutation`](./02_mutation/main.go) | изменение полей переданной структуры по указателю. | [`02_mutation_answers`](../04_exercises_answers/02_mutation_answers/main.go) |
| 3 | [`03_pointer_receiver`](./03_pointer_receiver/main.go) | объявление метода структуры с получателем-указателем (`*Account`). | [`03_pointer_receiver_answers`](../04_exercises_answers/03_pointer_receiver_answers/main.go) |
| 4 | [`04_nil_pointer`](./04_nil_pointer/main.go) | безопасная перезапись конфигурации без паник на nil-указателях. | [`04_nil_pointer_answers`](../04_exercises_answers/04_nil_pointer_answers/main.go) |
| 5 | [`05_swap`](./05_swap/main.go) | написание функции `swap(a, b *int)` для обмена значениями двух переменных. | [`05_swap_answers`](../04_exercises_answers/05_swap_answers/main.go) |
| 6 | [`06_pointer_to_pointer`](./06_pointer_to_pointer/main.go) | чтение значения через двойной указатель `**int`. | [`06_pointer_to_pointer_answers`](../04_exercises_answers/06_pointer_to_pointer_answers/main.go) |
| 7 | [`07_slice_pointer`](./07_slice_pointer/main.go) | модификация элементов слайса vs добавление элементов через append по указателю. | [`07_slice_pointer_answers`](../04_exercises_answers/07_slice_pointer_answers/main.go) |
| 8 | [`08_constructor`](./08_constructor/main.go) | создание конструктора структуры, возвращающего указатель. | [`08_constructor_answers`](../04_exercises_answers/08_constructor_answers/main.go) |

Решения открывайте только после собственной попытки.

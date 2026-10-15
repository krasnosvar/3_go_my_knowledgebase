# Упражнения: Тема 08. Структуры (Structs)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`08_structs/`):

```bash
go run ./03_exercises/01_declaration                 # заготовка
go run ./04_exercises_answers/01_declaration_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_declaration`](./01_declaration/main.go) | объявление структуры `Movie` с экспортируемыми полями. | [`01_declaration_answers`](../04_exercises_answers/01_declaration_answers/main.go) |
| 2 | [`02_constructor`](./02_constructor/main.go) | написание конструктора `NewUser`, возвращающего указатель. | [`02_constructor_answers`](../04_exercises_answers/02_constructor_answers/main.go) |
| 3 | [`03_method_value_receiver`](./03_method_value_receiver/main.go) | вычисление площади прямоугольника методом с ресивером по значению. | [`03_method_value_receiver_answers`](../04_exercises_answers/03_method_value_receiver_answers/main.go) |
| 4 | [`04_method_pointer_receiver`](./04_method_pointer_receiver/main.go) | изменение счетчика методом с ресивером по указателю. | [`04_method_pointer_receiver_answers`](../04_exercises_answers/04_method_pointer_receiver_answers/main.go) |
| 5 | [`05_embedded_struct`](./05_embedded_struct/main.go) | создание структуры `Employee`, расширяющей встроенный `Person`. | [`05_embedded_struct_answers`](../04_exercises_answers/05_embedded_struct_answers/main.go) |
| 6 | [`06_struct_tags`](./06_struct_tags/main.go) | настройка json-тэгов структуры для сопоставления полей. | [`06_struct_tags_answers`](../04_exercises_answers/06_struct_tags_answers/main.go) |
| 7 | [`07_zero_values`](./07_zero_values/main.go) | проектирование структуры форматирования логов с полезным нулевым значением по умолчанию. | [`07_zero_values_answers`](../04_exercises_answers/07_zero_values_answers/main.go) |
| 8 | [`08_copy_vs_pointer`](./08_copy_vs_pointer/main.go) | **доказательство изоляции копий структур при передаче по значению против передачи по указателю**. | [`08_copy_vs_pointer_answers`](../04_exercises_answers/08_copy_vs_pointer_answers/main.go) |

Решения открывайте только после собственной попытки.

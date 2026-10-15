# Упражнения: Тема 09. Слайсы (Slices)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`09_slices/`):

```bash
go run ./03_exercises/01_slice_basics                 # заготовка
go run ./04_exercises_answers/01_slice_basics_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_slice_basics`](./01_slice_basics/main.go) | создание слайса заданной длины и емкости с заполнением значениями. | [`01_slice_basics_answers`](../04_exercises_answers/01_slice_basics_answers/main.go) |
| 2 | [`02_append_capacity`](./02_append_capacity/main.go) | использование вариативного append для расширения слайса. | [`02_append_capacity_answers`](../04_exercises_answers/02_append_capacity_answers/main.go) |
| 3 | [`03_copy_delete`](./03_copy_delete/main.go) | удаление элемента по индексу с сохранением порядка элементов через `copy()`. | [`03_copy_delete_answers`](../04_exercises_answers/03_copy_delete_answers/main.go) |
| 4 | [`04_subslice_pitfalls`](./04_subslice_pitfalls/main.go) | анализ побочных эффектов при append во вложенный подлайс. | [`04_subslice_pitfalls_answers`](../04_exercises_answers/04_subslice_pitfalls_answers/main.go) |
| 5 | [`05_sort_search`](./05_sort_search/main.go) | сортировка слайса чисел и бинарный поиск значения с помощью стандартного пакета `slices`. | [`05_sort_search_answers`](../04_exercises_answers/05_sort_search_answers/main.go) |
| 6 | [`06_in_place_ops`](./06_in_place_ops/main.go) | **фильтрация положительных чисел in-place (на месте) без выделения новой памяти (`nums[:0]`)**. | [`06_in_place_ops_answers`](../04_exercises_answers/06_in_place_ops_answers/main.go) |
| 7 | [`07_memory_retention`](./07_memory_retention/main.go) | предотвращение удержания исходного массива в памяти при извлечении первых элементов. | [`07_memory_retention_answers`](../04_exercises_answers/07_memory_retention_answers/main.go) |
| 8 | [`08_concurrency_safety`](./08_concurrency_safety/main.go) | **написание безопасного конкурентного накопителя слайса с помощью sync.Mutex**. | [`08_concurrency_safety_answers`](../04_exercises_answers/08_concurrency_safety_answers/main.go) |

Решения открывайте только после собственной попытки.

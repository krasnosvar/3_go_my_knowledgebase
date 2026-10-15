# Упражнения: Тема 11. Хеш-таблицы (Maps)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`11_maps/`):

```bash
go run ./03_exercises/01_maps_basics                 # заготовка
go run ./04_exercises_answers/01_maps_basics_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_maps_basics`](./01_maps_basics/main.go) | безопасный поиск значения по ключу через comma-ok. | [`01_maps_basics_answers`](../04_exercises_answers/01_maps_basics_answers/main.go) |
| 2 | [`02_maps_advanced`](./02_maps_advanced/main.go) | изменение полей структур, находящихся внутри мапы (решение проблемы неадресуемости). | [`02_maps_advanced_answers`](../04_exercises_answers/02_maps_advanced_answers/main.go) |
| 3 | [`03_maps_set`](./03_maps_set/main.go) | реализация структуры данных Set (множество) с оптимизацией памяти через `map[string]struct{}`. | [`03_maps_set_answers`](../04_exercises_answers/03_maps_set_answers/main.go) |
| 4 | [`04_maps_sorting`](./04_maps_sorting/main.go) | извлечение и сортировка элементов мапы по ключам. | [`04_maps_sorting_answers`](../04_exercises_answers/04_maps_sorting_answers/main.go) |
| 5 | [`05_maps_key_constraints`](./05_maps_key_constraints/main.go) | использование структур в качестве сравниваемых (comparable) ключей мапы. | [`05_maps_key_constraints_answers`](../04_exercises_answers/05_maps_key_constraints_answers/main.go) |
| 6 | [`06_maps_concurrency`](./06_maps_concurrency/main.go) | **потокобезопасный счетчик посещений с использованием sync.Mutex**. | [`06_maps_concurrency_answers`](../04_exercises_answers/06_maps_concurrency_answers/main.go) |
| 7 | [`07_maps_memory`](./07_maps_memory/main.go) | **динамическая усадка (shrink) мапы для освобождения неиспользуемых бакетов из памяти**. | [`07_maps_memory_answers`](../04_exercises_answers/07_maps_memory_answers/main.go) |
| 8 | [`08_maps_lru_cache`](./08_maps_lru_cache/main.go) | **реализация алгоритма вытеснения LRU Cache на базе map и двусвязного списка container/list (Хардкор)**. | [`08_maps_lru_cache_answers`](../04_exercises_answers/08_maps_lru_cache_answers/main.go) |

Решения открывайте только после собственной попытки.

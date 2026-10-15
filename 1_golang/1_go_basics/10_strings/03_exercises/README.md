# Упражнения: Тема 10. Строки (Strings)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`10_strings/`):

```bash
go run ./03_exercises/01_string_basics                 # заготовка
go run ./04_exercises_answers/01_string_basics_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_string_basics`](./01_string_basics/main.go) | подсчет длины строки в байтах и рунах. | [`01_string_basics_answers`](../04_exercises_answers/01_string_basics_answers/main.go) |
| 2 | [`02_utf8_runes`](./02_utf8_runes/main.go) | декодирование первого символа строки и определение его размера. | [`02_utf8_runes_answers`](../04_exercises_answers/02_utf8_runes_answers/main.go) |
| 3 | [`03_concat_builder`](./03_concat_builder/main.go) | сборка строки из массива слов через `strings.Builder` с оптимизацией `Grow()`. | [`03_concat_builder_answers`](../04_exercises_answers/03_concat_builder_answers/main.go) |
| 4 | [`04_immutability`](./04_immutability/main.go) | модификация символа строки через приведение к слайсу рун. | [`04_immutability_answers`](../04_exercises_answers/04_immutability_answers/main.go) |
| 5 | [`05_safe_substring`](./05_safe_substring/main.go) | **безопасное извлечение подстроки по индексам символов (рун), исключающее повреждение UTF-8 байт**. | [`05_safe_substring_answers`](../04_exercises_answers/05_safe_substring_answers/main.go) |
| 6 | [`06_unsafe_conversions`](./06_unsafe_conversions/main.go) | **реализация zero-copy (без выделения памяти) конвертации между `[]byte` и `string` с помощью пакета `unsafe`**. | [`06_unsafe_conversions_answers`](../04_exercises_answers/06_unsafe_conversions_answers/main.go) |
| 7 | [`07_normalization`](./07_normalization/main.go) | нормализация тегов (trim, lower, слияние пробелов в дефис). | [`07_normalization_answers`](../04_exercises_answers/07_normalization_answers/main.go) |
| 8 | [`08_parsing`](./08_parsing/main.go) | парсинг параметров конфигурации формата `key=value` с обрезкой пробелов. | [`08_parsing_answers`](../04_exercises_answers/08_parsing_answers/main.go) |

Решения открывайте только после собственной попытки.

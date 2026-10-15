# Упражнения: Тема 02. Типы данных (Data Types)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`02_data_types/`):

```bash
go run ./03_exercises/01_numeric                 # заготовка
go run ./04_exercises_answers/01_numeric_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_numeric`](./01_numeric/main.go) | сложение и деление с защитой от деления на ноль. | [`01_numeric_answers`](../04_exercises_answers/01_numeric_answers/main.go) |
| 2 | [`02_bool_logic`](./02_bool_logic/main.go) | проверка прав доступа пользователя (логическое И). | [`02_bool_logic_answers`](../04_exercises_answers/02_bool_logic_answers/main.go) |
| 3 | [`03_string_bytes`](./03_string_bytes/main.go) | безопасное получение первого байта строки (защита от panic). | [`03_string_bytes_answers`](../04_exercises_answers/03_string_bytes_answers/main.go) |
| 4 | [`04_rune_utf8`](./04_rune_utf8/main.go) | извлечение первой руны UTF-8 строки с использованием `range`. | [`04_rune_utf8_answers`](../04_exercises_answers/04_rune_utf8_answers/main.go) |
| 5 | [`05_type_conversion`](./05_type_conversion/main.go) | вычисление среднего значения через явное приведение типов. | [`05_type_conversion_answers`](../04_exercises_answers/05_type_conversion_answers/main.go) |
| 6 | [`06_zero_values`](./06_zero_values/main.go) | проверка параметров конфигурации (`bool`, `int`, `string`) на нулевые значения. | [`06_zero_values_answers`](../04_exercises_answers/06_zero_values_answers/main.go) |
| 7 | [`07_alias_types`](./07_alias_types/main.go) | использование Named Type и Type Alias. | [`07_alias_types_answers`](../04_exercises_answers/07_alias_types_answers/main.go) |
| 8 | [`08_composite_preview`](./08_composite_preview/main.go) | инициализация простых слайсов и мап. | [`08_composite_preview_answers`](../04_exercises_answers/08_composite_preview_answers/main.go) |

Решения открывайте только после собственной попытки.

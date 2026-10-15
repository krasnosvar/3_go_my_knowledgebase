# Упражнения: Тема 03. Форматированный вывод (fmt & output)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`03_fmt_and_output/`):

```bash
go run ./03_exercises/01_print                 # заготовка
go run ./04_exercises_answers/01_print_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_print`](./01_print/main.go) | вывод в переданный Writer. | [`01_print_answers`](../04_exercises_answers/01_print_answers/main.go) |
| 2 | [`02_printf`](./02_printf/main.go) | базовый `Fprintf` с `%s` и `%d`. | [`02_printf_answers`](../04_exercises_answers/02_printf_answers/main.go) |
| 3 | [`03_sprintf`](./03_sprintf/main.go) | создание форматированной строки лога. | [`03_sprintf_answers`](../04_exercises_answers/03_sprintf_answers/main.go) |
| 4 | [`04_errorf`](./04_errorf/main.go) | обёртывание ошибок через `%w`. | [`04_errorf_answers`](../04_exercises_answers/04_errorf_answers/main.go) |
| 5 | [`05_scan`](./05_scan/main.go) | **парсинг строк вида IP:Port через Sscanf**. | [`05_scan_answers`](../04_exercises_answers/05_scan_answers/main.go) |
| 6 | [`06_width_precision`](./06_width_precision/main.go) | вывод вещественного числа с заданной точностью. | [`06_width_precision_answers`](../04_exercises_answers/06_width_precision_answers/main.go) |
| 7 | [`07_struct_print`](./07_struct_print/main.go) | вывод структуры в Go-синтаксисе (`%#v`). | [`07_struct_print_answers`](../04_exercises_answers/07_struct_print_answers/main.go) |
| 8 | [`08_table_output`](./08_table_output/main.go) | **выстраивание колонок таблицы с фиксированной шириной (padding)**. | [`08_table_output_answers`](../04_exercises_answers/08_table_output_answers/main.go) |

Решения открывайте только после собственной попытки.

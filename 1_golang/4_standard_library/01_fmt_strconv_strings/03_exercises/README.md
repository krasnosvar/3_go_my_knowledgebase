# Упражнения: fmt, strconv, strings

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, без внешней сети.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_printf_table/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_printf_table_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_printf_table`](./01_printf_table/main.go) — строка таблицы фиксированной ширины через `%-10s`, `%5d`, `%8.2f`.
2. [`02_atoi_validation`](./02_atoi_validation/main.go) — `parsePort`: `TrimSpace` + `Atoi` + диапазон, ошибка оборачивает `*strconv.NumError`.
3. [`03_join_words`](./03_join_words/main.go) — склейка непустых слов через `strings.Join`.
4. [`04_trim_inputs`](./04_trim_inputs/main.go) — нормализация ввода: `TrimSpace` + `ToLower`, без изменения исходного среза.
5. [`05_replace_token`](./05_replace_token/main.go) — подстановка `{{key}}` в шаблон через `strings.NewReplacer`.
6. [`06_count_substring`](./06_count_substring/main.go) — подсчёт непересекающихся вхождений без учёта регистра, ловушка пустой подстроки.
7. [`07_split_csv`](./07_split_csv/main.go) — `Split` по запятой с обрезкой пробелов, ловушка `Split("", ",")`.
8. [`08_format_report`](./08_format_report/main.go) — многострочный отчёт с итогом через `strings.Builder` + `fmt.Fprintf`.

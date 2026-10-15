# Упражнения: regexp, sort, time

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, без внешней сети.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_regex_email/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_regex_email_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_regex_email`](./01_regex_email/main.go) — валидация email заякоренным шаблоном, скомпилированным один раз.
2. [`02_regex_extract`](./02_regex_extract/main.go) — извлечение номеров `#123` группой захвата и `FindAllStringSubmatch`.
3. [`03_sort_numbers`](./03_sort_numbers/main.go) — сортировка по убыванию копии среза, вход не меняется.
4. [`04_sort_structs`](./04_sort_structs/main.go) — `sort.Slice` по двум ключам (Age, затем Name).
5. [`05_parse_date`](./05_parse_date/main.go) — `time.Parse` с layout `2006-01-02` и формат `02.01.2006`.
6. [`06_duration_math`](./06_duration_math/main.go) — сумма длительностей из строк через `time.ParseDuration`.
7. [`07_ticker_loop`](./07_ticker_loop/main.go) — цикл на `time.Ticker` ровно на n тиков с `defer Stop()`.
8. [`08_deadline_check`](./08_deadline_check/main.go) — остаток до дедлайна на детерминированном времени (`Sub`, `Before`).

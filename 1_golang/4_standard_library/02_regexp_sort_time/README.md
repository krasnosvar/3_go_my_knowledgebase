# 02. regexp, sort, time

Регулярные выражения для валидации и извлечения данных, сортировка срезов и структур, работа со временем, длительностями и таймерами. Здесь же — типичные ловушки: незаякоренный шаблон, компиляция regexp на каждый вызов, layout `2006-01-02` вместо `YYYY-MM-DD`, утечка `Ticker` без `Stop`.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — методы `regexp` и синтаксис RE2, `slices`/`sort`, `time` (layout, таймеры), ловушки.

## 🧭 Пререквизиты

- [01_fmt_strconv_strings](../01_fmt_strconv_strings/) — строки и `strconv`
- [1_go_basics/09_slices](../../1_go_basics/09_slices/) — срезы и копирование
- [1_go_basics/07_functions](../../1_go_basics/07_functions/) — функции как значения (`less` для `sort.Slice`)

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_regexp.md](./01_theory/01_regexp.md) — `regexp`: `MustCompile`, методы `Match`/`Find*`, группы, синтаксис RE2, производительность.
    *   [02_sort.md](./01_theory/02_sort.md) — `sort` и `slices.Sort` (Go 1.21+), stable/unstable, `Reverse`, бинарный поиск.
    *   [03_time.md](./01_theory/03_time.md) — `time`: `Time`/`Duration`, форматирование и парсинг, `Timer`/`Ticker`, `time.After`.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_regex_match](./02_examples/01_regex_match/main.go) — проверка строки регулярным выражением.
    *   [02_sort_slice](./02_examples/02_sort_slice/main.go) — сортировка среза.
    *   [03_time_parse_duration](./02_examples/03_time_parse_duration/main.go) — разбор даты и длительности.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_regex_email](./03_exercises/01_regex_email/main.go) — валидация email заякоренным шаблоном, скомпилированным один раз.
    *   [02_regex_extract](./03_exercises/02_regex_extract/main.go) — извлечение номеров `#123` группой захвата и `FindAllStringSubmatch`.
    *   [03_sort_numbers](./03_exercises/03_sort_numbers/main.go) — сортировка по убыванию копии среза, вход не меняется.
    *   [04_sort_structs](./03_exercises/04_sort_structs/main.go) — `sort.Slice` по двум ключам (Age, затем Name).
    *   [05_parse_date](./03_exercises/05_parse_date/main.go) — `time.Parse` с layout `2006-01-02` и формат `02.01.2006`.
    *   [06_duration_math](./03_exercises/06_duration_math/main.go) — сумма длительностей из строк через `time.ParseDuration`.
    *   [07_ticker_loop](./03_exercises/07_ticker_loop/main.go) — цикл на `time.Ticker` ровно на n тиков с `defer Stop()`.
    *   [08_deadline_check](./03_exercises/08_deadline_check/main.go) — остаток до дедлайна на детерминированном времени (`Sub`, `Before`).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Команды выполняются из папки темы (`1_golang/4_standard_library/02_regexp_sort_time`).

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_regexp.md](./01_theory/01_regexp.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_regex_match/main.go
   ```
3. Решите задачи в `03_exercises/` и запустите самопроверку (в стартовом состоянии печатает `FAIL`):
   ```bash
   go run ./03_exercises/01_regex_email/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_regex_email_answers/main.go
   ```

## ➡️ Дальше

[03_encoding_json_csv_base64](../03_encoding_json_csv_base64/) — сериализация данных.

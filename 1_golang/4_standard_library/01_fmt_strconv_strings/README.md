# 01. fmt, strconv, strings

Форматированный вывод, преобразование строк в числа и обратно, поиск/разбиение/склейка строк — то, с чего начинается любой разбор ввода, логов, конфигов и параметров запроса. Тема показывает идиоматичные инструменты stdlib вместо ручных циклов и `fmt.Sscan`.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — функции и глаголы `fmt`, `strconv` (парсинг, ошибки), `strings`, ловушки.

## 🧭 Пререквизиты

- [1_go_basics/03_fmt_and_output](../../1_go_basics/03_fmt_and_output/) — базовый `fmt`
- [1_go_basics/10_strings](../../1_go_basics/10_strings/) — устройство строк, руны, `strings.Builder`
- [1_go_basics/12_errors](../../1_go_basics/12_errors/) — ошибки и обёртывание `%w`

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_fmt.md](./01_theory/01_fmt.md) — `fmt`: функции, глаголы формата, `Errorf` с `%w`, `Stringer`, частые ошибки.
    *   [02_strconv.md](./01_theory/02_strconv.md) — `strconv`: целые, float, bool, quoting, `*NumError`, почему не `Sscan`.
    *   [03_strings.md](./01_theory/03_strings.md) — `strings`: поиск, модификация, `Split`/`Join`, `Builder`, `Reader`, Unicode.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_format_and_parse](./02_examples/01_format_and_parse/main.go) — форматирование и разбор чисел.
    *   [02_string_pipeline](./02_examples/02_string_pipeline/main.go) — цепочка преобразований строк.
    *   [03_safe_parse](./02_examples/03_safe_parse/main.go) — безопасный разбор с проверкой ошибки.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_printf_table](./03_exercises/01_printf_table/main.go) — строка таблицы фиксированной ширины через `%-10s`, `%5d`, `%8.2f`.
    *   [02_atoi_validation](./03_exercises/02_atoi_validation/main.go) — `parsePort`: `TrimSpace` + `Atoi` + диапазон, ошибка оборачивает `*strconv.NumError`.
    *   [03_join_words](./03_exercises/03_join_words/main.go) — склейка непустых слов через `strings.Join`.
    *   [04_trim_inputs](./03_exercises/04_trim_inputs/main.go) — нормализация ввода: `TrimSpace` + `ToLower`, без изменения исходного среза.
    *   [05_replace_token](./03_exercises/05_replace_token/main.go) — подстановка `{{key}}` в шаблон через `strings.NewReplacer`.
    *   [06_count_substring](./03_exercises/06_count_substring/main.go) — подсчёт непересекающихся вхождений без учёта регистра, ловушка пустой подстроки.
    *   [07_split_csv](./03_exercises/07_split_csv/main.go) — `Split` по запятой с обрезкой пробелов, ловушка `Split("", ",")`.
    *   [08_format_report](./03_exercises/08_format_report/main.go) — многострочный отчёт с итогом через `strings.Builder` + `fmt.Fprintf`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Команды выполняются из папки темы (`1_golang/4_standard_library/01_fmt_strconv_strings`).

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_fmt.md](./01_theory/01_fmt.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_format_and_parse/main.go
   ```
3. Решите задачи в `03_exercises/` и запустите самопроверку (в стартовом состоянии печатает `FAIL`):
   ```bash
   go run ./03_exercises/01_printf_table/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_printf_table_answers/main.go
   ```

## ➡️ Дальше

[02_regexp_sort_time](../02_regexp_sort_time/) — регулярные выражения, сортировка и время.

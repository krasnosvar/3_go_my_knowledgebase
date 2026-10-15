# Тема 03. Форматированный вывод (fmt & output)

Этот раздел посвящён работе с пакетом `fmt`: базовые функции печати (`Print`/`Println`/`Printf`), глаголы форматирования, сборка строк и ошибок (`Sprintf`/`Errorf`), запись в произвольный `io.Writer` (`Fprintf`), чтение форматированного ввода (`Scan`), а также частые ошибки и практические сценарии применения.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — функции `Print*`/`Fprint*`/`Sprint*`, глаголы, ширина и точность, `Errorf` и `%w`, `Scan`, ловушки.

## 📋 Пререквизиты

*   Пройдена предыдущая тема: [02_data_types](../02_data_types/).

> ⏩ **Забегание вперёд.** Некоторые примеры и упражнения используют конструкции из следующих тем:
> *   Упражнение `04_errorf` и часть примеров используют ошибки — подробно в [12_errors](../12_errors/).
> *   `07_struct_print` и примеры печатают структуры — подробно в [08_structs](../08_structs/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_fmt_basics.md](./01_theory/01_fmt_basics.md) — **вводный файл: базовые функции вывода (`Print`/`Println`/`Printf`), куда попадает вывод (stdout/stderr/io.Writer), обзор темы и заметка о производительности**.
    *   [02_printf_verbs.md](./01_theory/02_printf_verbs.md) — глаголы форматирования (`%v`, `%+v`, `%#v`, `%d`, `%s`, `%q` и др.), ширина, точность, порядок аргументов.
    *   [03_sprintf_errorf.md](./01_theory/03_sprintf_errorf.md) — `Sprintf`, `Fprintf`, `Errorf` и оборачивание ошибок через `%w`.
    *   [04_scan.md](./01_theory/04_scan.md) — чтение и парсинг ввода (`Scan`, `Sscanf`, `Fscan`) и когда `fmt` не подходит.
    *   [05_pitfalls.md](./01_theory/05_pitfalls.md) — частые ошибки с `fmt` (неверный verb, забытый `\n`, format string, потоки вывода).
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии: CLI-вывод, отладка, ошибки с контекстом, `fmt.Stringer`.
*   **[02_examples/](./02_examples/)** — практические примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — разница вывода структуры: `%v`, `%+v`, `%#v`, `%T`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — форматирование метрик с точностью и вывод указателей (`%p`).
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — оборачивание ошибок через `%w` и бенчмарк `strings.Builder` vs `fmt.Sprintf` в циклах.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи для самопроверки с автопроверками `PASS`/`FAIL`:
    *   `01_print` — вывод в переданный Writer.
    *   `02_printf` — базовый `Fprintf` с `%s` и `%d`.
    *   `03_sprintf` — создание форматированной строки лога.
    *   `04_errorf` — обёртывание ошибок через `%w`.
    *   `05_scan` — **парсинг строк вида IP:Port через Sscanf**.
    *   `06_width_precision` — вывод вещественного числа с заданной точностью.
    *   `07_struct_print` — вывод структуры в Go-синтаксисе (`%#v`).
    *   `08_table_output` — **выстраивание колонок таблицы с фиксированной шириной (padding)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Прочитайте вводный файл: [01_fmt_basics.md](./01_theory/01_fmt_basics.md), затем подтемы 02–06 по порядку.
2. Запустите и изучите примеры из папки `02_examples/` через `go run`:
   ```bash
   go run ./02_examples/03_product_backend/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/05_scan/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

## ➡️ Что дальше

Когда все упражнения темы печатают `PASS`, переходите к теме [04_operators_bitwise](../04_operators_bitwise/).

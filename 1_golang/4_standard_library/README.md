# 4. Стандартная библиотека Go

Стандартная библиотека Go закрывает большую часть задач backend-сервиса без внешних зависимостей: форматирование и разбор строк, регулярные выражения, сортировка, время, JSON/CSV/base64, файлы и потоки, HTTP-сервер и клиент, TCP. Раздел учит пользоваться этими пакетами идиоматично и обходить их типичные ловушки.

## 🧭 Пререквизиты

- [1_go_basics](../1_go_basics/) — весь базовый синтаксис, особенно строки, срезы, структуры и ошибки.
- [2_oop/02_interfaces](../2_oop/02_interfaces/) — интерфейсы (на них построены `io`, `sort`, `net/http`).
- Для темы 05 — [3_concurrency](../3_concurrency/): горутины и `context`.

## 🗂 Темы (проходить по порядку)

1. [01_fmt_strconv_strings](./01_fmt_strconv_strings/) — `fmt`, `strconv`, `strings`: форматирование, разбор чисел, операции над строками.
2. [02_regexp_sort_time](./02_regexp_sort_time/) — `regexp`, `sort`, `time`: валидация и извлечение, сортировка, даты, длительности, `Ticker`.
3. [03_encoding_json_csv_base64](./03_encoding_json_csv_base64/) — `encoding/json`, `encoding/csv`, `encoding/base64`: теги, строгий разбор, потоковое чтение.
4. [04_io_bufio_os_filepath](./04_io_bufio_os_filepath/) — `io`, `bufio`, `os`, `path/filepath`: потоки, файлы, окружение, безопасные пути.
5. [05_net_http_net](./05_net_http_net/) — `net/http`, `net`: хендлеры, клиент с таймаутом, ретраи, graceful shutdown, TCP.

Каждая тема устроена одинаково: `01_theory/` → `02_examples/` → `03_exercises/` (8 задач с самопроверкой `PASS`/`FAIL`) → `04_exercises_answers/`.

## 🚀 Как запускать

Из папки темы:

```bash
cd 1_golang/4_standard_library/01_fmt_strconv_strings
go run ./02_examples/01_format_and_parse/main.go
go run ./03_exercises/01_printf_table/main.go                 # заготовка: FAIL, пока не решено
go run ./04_exercises_answers/01_printf_table_answers/main.go # эталон: PASS
```

Проверить, что весь раздел компилируется (из корня репозитория):

```bash
go vet ./1_golang/4_standard_library/...
```

## ➡️ Дальше

[5_testing](../5_testing/) — как покрывать такой код тестами; затем [6_real_world_go_usage](../6_real_world_go_usage/) — сборка сервисов из этих кирпичей.

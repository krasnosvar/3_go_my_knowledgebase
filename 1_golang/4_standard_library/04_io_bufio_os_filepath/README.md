# 04. io, bufio, os, path/filepath

Интерфейсы `io.Reader`/`io.Writer` — общий язык stdlib: файлы, сеть, буферы и HTTP-тела устроены одинаково. Тема учит копировать потоки без загрузки в память, читать построчно, корректно писать и закрывать файлы, работать с окружением и безопасно строить пути.

## 🧭 Пререквизиты

- [2_oop/02_interfaces](../../2_oop/02_interfaces/) — интерфейсы (`io.Reader` — интерфейс из одного метода)
- [1_go_basics/12_errors](../../1_go_basics/12_errors/) — `errors.Is`, обёртывание
- [1_go_basics/07_functions](../../1_go_basics/07_functions/) — `defer` и именованные результаты

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_io_and_bufio.md](./01_theory/01_io_and_bufio.md) — `io` и `bufio`: базовые интерфейсы, `Copy`/`LimitReader`, `Pipe`, `Scanner`, `Writer` и `Flush`.
    *   [02_os.md](./01_theory/02_os.md) — `os`: файлы, директории, `Stat` и проверка существования, временные файлы, окружение, сигналы.
    *   [03_filepath.md](./01_theory/03_filepath.md) — `path/filepath`: `path` vs `filepath`, `Join`/`Clean`/`Rel`, `WalkDir`, `Glob`, кроссплатформенность.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_bufio_scan](./02_examples/01_bufio_scan/main.go) — построчное чтение через `bufio.Scanner`.
    *   [02_os_file_temp](./02_examples/02_os_file_temp/main.go) — временный файл через `os`.
    *   [03_filepath_join_walk](./02_examples/03_filepath_join_walk/main.go) — `filepath.Join` и обход директорий.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_reader_copy](./03_exercises/01_reader_copy/main.go) — `io.Copy` + `io.LimitReader`: копирование с лимитом без чтения в память.
    *   [02_scanner_lines](./03_exercises/02_scanner_lines/main.go) — подсчёт непустых строк без комментариев, проверка `sc.Err()`.
    *   [03_write_file](./03_exercises/03_write_file/main.go) — запись через `bufio.Writer` с `Flush` и возвратом ошибки `Close`.
    *   [04_temp_file](./03_exercises/04_temp_file/main.go) — `os.CreateTemp` + функция cleanup.
    *   [05_walk_dir](./03_exercises/05_walk_dir/main.go) — `filepath.WalkDir`: поиск файлов по расширению, относительные пути.
    *   [06_env_read](./03_exercises/06_env_read/main.go) — `os.LookupEnv`: различие «не задана» и «задана пустой».
    *   [07_path_normalize](./03_exercises/07_path_normalize/main.go) — `safeJoin`: защита от path traversal через `Join` + `Rel`.
    *   [08_file_exists](./03_exercises/08_file_exists/main.go) — `os.Stat` + `errors.Is(err, fs.ErrNotExist)`, директория ≠ файл.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Команды выполняются из папки темы (`1_golang/4_standard_library/04_io_bufio_os_filepath`).

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_io_and_bufio.md](./01_theory/01_io_and_bufio.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_bufio_scan/main.go
   ```
3. Решите задачи в `03_exercises/` и запустите самопроверку (в стартовом состоянии печатает `FAIL`):
   ```bash
   go run ./03_exercises/01_reader_copy/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_reader_copy_answers/main.go
   ```

## ➡️ Дальше

[05_net_http_net](../05_net_http_net/) — HTTP-сервер, клиент и TCP.

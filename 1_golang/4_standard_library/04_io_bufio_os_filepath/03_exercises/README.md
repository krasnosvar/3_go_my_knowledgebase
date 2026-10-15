# Упражнения: io, bufio, os, path/filepath

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, без внешней сети.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_reader_copy/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_reader_copy_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_reader_copy`](./01_reader_copy/main.go) — `io.Copy` + `io.LimitReader`: копирование с лимитом без чтения в память.
2. [`02_scanner_lines`](./02_scanner_lines/main.go) — подсчёт непустых строк без комментариев, проверка `sc.Err()`.
3. [`03_write_file`](./03_write_file/main.go) — запись через `bufio.Writer` с `Flush` и возвратом ошибки `Close`.
4. [`04_temp_file`](./04_temp_file/main.go) — `os.CreateTemp` + функция cleanup.
5. [`05_walk_dir`](./05_walk_dir/main.go) — `filepath.WalkDir`: поиск файлов по расширению, относительные пути.
6. [`06_env_read`](./06_env_read/main.go) — `os.LookupEnv`: различие «не задана» и «задана пустой».
7. [`07_path_normalize`](./07_path_normalize/main.go) — `safeJoin`: защита от path traversal через `Join` + `Rel`.
8. [`08_file_exists`](./08_file_exists/main.go) — `os.Stat` + `errors.Is(err, fs.ErrNotExist)`, директория ≠ файл.

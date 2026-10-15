# io, bufio, os, path/filepath — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## io: интерфейсы и функции

Подробно: [01_io_and_bufio.md](./01_theory/01_io_and_bufio.md).

| Интерфейс | Метод | Кто реализует |
| :--- | :--- | :--- |
| `io.Reader` | `Read(p []byte) (n int, err error)` | `*os.File`, `net.Conn`, `*strings.Reader`, `*bytes.Buffer`, `http.Request.Body` |
| `io.Writer` | `Write(p []byte) (n int, err error)` | `*os.File`, `net.Conn`, `*bytes.Buffer`, `http.ResponseWriter` |
| `io.Closer` | `Close() error` | файлы, соединения, body |
| Композиции | `io.ReadWriter`, `io.ReadCloser`, `io.ReadWriteCloser` | — |

| Задача | Код |
| :--- | :--- |
| Скопировать поток до EOF | `n, err := io.Copy(dst, src)` |
| Первые N байт | `io.CopyN(dst, src, 1024)` |
| Весь поток в память | `data, err := io.ReadAll(r)` |
| Ограничить чтение | `io.LimitReader(r, 1<<20)`; для HTTP-body — `http.MaxBytesReader` |
| Связать producer и consumer | `pr, pw := io.Pipe()` (writer — в горутине, `defer pw.Close()`) |
| Данные в памяти вместо файла | `strings.NewReader(s)`, `bytes.NewBuffer(b)` |

## bufio

| Задача | Код |
| :--- | :--- |
| Построчное чтение | `sc := bufio.NewScanner(r); for sc.Scan() { sc.Text() }; sc.Err()` |
| Строки длиннее 64 KiB | `sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)` |
| Читать до разделителя | `br := bufio.NewReader(r); line, err := br.ReadString('\n')` / `ReadBytes` |
| Посмотреть без чтения | `br.Peek(4)` |
| Буферизованная запись | `bw := bufio.NewWriter(f); defer bw.Flush(); bw.WriteString("line\n")` |

| Нужно | Инструмент |
| :--- | :--- |
| Читать поток строками | `bufio.Scanner` |
| Читать байты / с разделителем / `Peek` | `bufio.Reader` |
| Писать поток | `bufio.Writer` + `Flush` |
| Копировать | `io.Copy` |

## os: файлы и каталоги

Подробно: [02_os.md](./01_theory/02_os.md).

| Задача | Код |
| :--- | :--- |
| Открыть на чтение | `f, err := os.Open(p); defer f.Close()` |
| Создать / обрезать | `os.Create(p)` |
| С флагами и правами | `os.OpenFile(p, os.O_RDWR\|os.O_CREATE\|os.O_APPEND, 0644)` |
| Прочитать файл целиком | `data, err := os.ReadFile(p)` |
| Записать целиком | `os.WriteFile(p, data, 0644)` (truncate + write) |
| Каталог / дерево | `os.Mkdir(d, 0755)` / `os.MkdirAll("a/b/c", 0755)` |
| Удалить файл / дерево | `os.Remove(p)` / `os.RemoveAll(d)` |
| Список каталога | `entries, err := os.ReadDir(".")` → `e.Name()`, `e.IsDir()` |
| Существует? | `_, err := os.Stat(p); errors.Is(err, os.ErrNotExist)` |
| Метаданные | `info.Size()`, `Mode()`, `ModTime()`, `IsDir()` |
| Временный файл / каталог | `os.CreateTemp("", "prefix-*.txt")` / `os.MkdirTemp("", "work-*")` + `defer os.Remove…` |

## os: процесс

| Задача | Код |
| :--- | :--- |
| Переменная окружения | `os.Getenv("APP_PORT")` (`""` если нет) |
| Отличить «нет» от «пусто» | `v, ok := os.LookupEnv("APP_PORT")` |
| Установить / удалить / все | `os.Setenv`, `os.Unsetenv`, `os.Environ()` (`"KEY=VALUE"`) |
| Аргументы | `os.Args` (`[0]` — имя программы) |
| Выход с кодом | `os.Exit(1)` — ⚠️ без `defer` |
| Потоки | `os.Stdin`, `os.Stdout`, `os.Stderr` |
| Graceful shutdown по сигналу | `ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM); defer stop(); <-ctx.Done()` |

## path/filepath

Подробно: [03_filepath.md](./01_theory/03_filepath.md).

| Задача | Код | Результат (Unix) |
| :--- | :--- | :--- |
| Склеить (с `Clean`) | `filepath.Join("a", "b", "../c")` | `a/c` |
| Каталог / имя / расширение | `filepath.Dir("/a/b.txt")`, `Base`, `Ext("f.tar.gz")` | `/a`, `b.txt`, `.gz` |
| Разделить | `filepath.Split("/a/b/file.txt")` | `"/a/b/"`, `"file.txt"` |
| Нормализовать | `filepath.Clean("/a//b/../c")` | `/a/c` |
| Абсолютный / относительный | `filepath.Abs(p)` / `filepath.Rel(base, target)` (проверять `err`) | — |
| Симлинки | `filepath.EvalSymlinks(p)` (`Abs` их не раскрывает) | — |
| Шаблон | `filepath.Match("*.go", "main.go")` / `Glob("*.go")` | `true`; `*` не матчит `/`, `**` нет |
| Обход дерева | `filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {...})` | `return filepath.SkipDir` — пропустить каталог |
| Обход абстрактной FS | `fs.WalkDir(fsys, ".", fn)`; прод — `os.DirFS("/data")`, тест — `fstest.MapFS{...}` | — |

| Пакет | Для чего |
| :--- | :--- |
| `path/filepath` | пути файловой системы (`/` или `\` по ОС) |
| `path` | URL-подобные пути, всегда `/` |

## Ловушки

- ⚠️ Забыли `Flush()` у `bufio.Writer` — данные остались в буфере.
- ⚠️ `bufio.Scanner` останавливается на строке > 64 KiB: `sc.Err()` → `token too long`; настройте `Buffer`.
- ⚠️ Не проверили `sc.Err()` после цикла — ошибка чтения выглядит как конец файла.
- ⚠️ `io.ReadAll` на недоверенном входе без лимита — DoS по памяти.
- ⚠️ Чтение без буфера — системный вызов на каждый мелкий `Read`.
- ⚠️ Нет `defer f.Close()` — утечка дескрипторов.
- ⚠️ `os.Exit` пропускает `defer`: не будет flush логов, удаления временных файлов; лучше `return` из `main`.
- ⚠️ `os.Getenv("X") == ""` не отличает «нет» от «пусто» — `LookupEnv`.
- ⚠️ `dir + "/" + file` — ломается на Windows; `filepath.Join`. И не путать с `path.Join`.
- ⚠️ `filepath.Walk` вызывает `stat` на каждом файле — `WalkDir` быстрее; `ioutil.ReadDir` устарел — `os.ReadDir`.

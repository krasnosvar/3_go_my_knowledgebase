# io и bufio

`io` — фундаментальные интерфейсы ввода/вывода. `bufio` — буферизованные обёртки над ними.

## Основные интерфейсы

```go
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }
type Closer interface { Close() error }
```

Почти вся Go stdlib (`os.File`, `net.Conn`, `strings.Reader`, `http.Request.Body`) реализует их.

Композиция: `io.ReadWriter`, `io.ReadCloser`, `io.ReadWriteCloser`.

## Полезные функции

```go
n, err := io.Copy(dst, src)         // копирует, пока не EOF
err = io.CopyN(dst, src, 1024)      // первые N байт

data, err := io.ReadAll(r)          // весь поток в []byte

limited := io.LimitReader(r, 1<<20) // защита: не больше 1 MiB
```

`io.ReadAll` — удобно, но опасно на недоверенном входе: может съесть память. Используй `io.LimitReader` или `http.MaxBytesReader`.

## io.Pipe

In-memory пара Writer/Reader:

```go
pr, pw := io.Pipe()
go func() {
    defer pw.Close()
    pw.Write([]byte("data"))
}()
io.Copy(os.Stdout, pr)
```

Полезно соединить producer и consumer без временного буфера.

## bufio.Reader

Обёртка, добавляющая буферизацию:

```go
br := bufio.NewReader(f)
line, err := br.ReadString('\n')   // одна строка
b, err := br.ReadBytes('\n')       // []byte
p, err := br.Peek(4)               // посмотреть без consume
```

## bufio.Scanner

Самый удобный способ построчного чтения:

```go
sc := bufio.NewScanner(f)
sc.Buffer(make([]byte, 0, 64*1024), 1024*1024) // для длинных строк
for sc.Scan() {
    line := sc.Text()
    process(line)
}
if err := sc.Err(); err != nil {
    return err
}
```

По умолчанию `Scanner` бросает строки длиннее `bufio.MaxScanTokenSize` (64 KB). Для длинных — вручную через `Buffer`.

## bufio.Writer

```go
bw := bufio.NewWriter(f)
defer bw.Flush() // обязательно!
bw.WriteString("line\n")
```

Без `Flush` часть данных останется в буфере.

## Когда использовать

- Любое чтение файла/сети — через буферизацию (`bufio.Scanner` или `bufio.Reader`). Без буфера — системные вызовы на каждый байт.
- Копирование больших файлов/потоков — `io.Copy`, а не свои циклы.
- In-memory данные — `bytes.Buffer` / `strings.Reader` вместо реальных файлов в тестах.

## Частые ошибки

- Забыть `Flush` у `bufio.Writer` / `csv.Writer` — потеря данных.
- `io.ReadAll` на HTTP-body без лимита — DoS-вектор.
- Читать без буфера — катастрофическая производительность на больших файлах.
- Scanner падает на длинных строках — конфигурируй `Buffer`.

## Правило

- Читаешь поток → `bufio.Scanner` (строки) или `bufio.Reader` (байты).
- Пишешь поток → `bufio.Writer` + `defer Flush`.
- Копируешь → `io.Copy`.

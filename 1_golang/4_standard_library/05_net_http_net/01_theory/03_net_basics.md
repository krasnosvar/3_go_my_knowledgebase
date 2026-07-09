# net (TCP/UDP basics)

Пакет `net` — нижний уровень сетевой работы. В бизнес-коде Go прикладной программист сюда залезает редко: `net/http`, `database/sql`, gRPC клиенты уже делают всё, что нужно. Но понимать — полезно.

## Dial / Listen

```go
// клиент
conn, err := net.Dial("tcp", "example.com:80")
defer conn.Close()

// сервер
ln, err := net.Listen("tcp", ":8080")
defer ln.Close()

for {
    conn, err := ln.Accept()
    if err != nil { continue }
    go handle(conn)
}
```

`conn` реализует `io.Reader`/`Writer`/`Closer`, так что обычная работа через bufio.

## Timeout'ы и контексты

```go
d := net.Dialer{Timeout: 5 * time.Second}
conn, err := d.DialContext(ctx, "tcp", addr)

_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
_ = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
```

`Deadline` — абсолютное время, до которого операции должны завершиться. Иначе `Read/Write` вернут timeout-ошибку.

## net.ResolveTCPAddr / Resolve

```go
addr, err := net.ResolveTCPAddr("tcp", "example.com:80")
```

Резолвинг DNS. Кастомный Resolver через `net.Resolver`.

## UDP

```go
conn, err := net.ListenPacket("udp", ":1053")
buf := make([]byte, 1500)
n, addr, err := conn.ReadFrom(buf)
conn.WriteTo(response, addr)
```

UDP — без соединений, без гарантии доставки. Для DNS, метрик, игр.

## Unix domain sockets

```go
ln, _ := net.Listen("unix", "/tmp/app.sock")
```

Быстрее TCP на одной машине, часто используется для sidecar-паттернов.

## Частые ошибки

- `Read` без таймаута — зависаем на неподвижном соединении.
- Не проверять `n` у `Read`: может вернуть данные И ошибку одновременно.
- Держать сотни `Dial`-ов без pool'а — исчерпание портов.
- Игнорировать `Close()` — leak дескрипторов.

## Когда лезть в net напрямую

- Свой протокол поверх TCP (бинарный).
- Low-level прокси/туннель.
- Load testing утилиты.
- Нестандартные сетевые задачи (raw packets — через `syscall.Raw` или сторонние пакеты).

В остальных случаях — высокоуровневые пакеты (`net/http`, `database/sql`, `grpc`).

# Use-cases: переменные и константы

## 1. Константы конфигурации

Неизменяемые значения, которые знает вся программа:

```go
const (
    DefaultPort    = 8080
    DefaultTimeout = 5 * time.Second
    MaxRetries     = 3
)
```

Использовать, когда значение **не** зависит от окружения. Если зависит — это уже конфиг из env/flags.

## 2. Enum-like статусы

```go
type OrderStatus int

const (
    OrderPending OrderStatus = iota
    OrderPaid
    OrderShipped
    OrderCancelled
)
```

Плюс метод `String()` — и будет читаемый вывод в логах.

## 3. Битовые флаги

```go
type Permission int
const (
    PermRead Permission = 1 << iota
    PermWrite
    PermAdmin
)

user.Perms = PermRead | PermWrite
```

Компактнее, чем несколько bool-полей.

## 4. Глобальные переменные

Редко и осознанно. Хорошие примеры: `http.DefaultClient`, `os.Stdin/Stdout/Stderr`. Плохой — «глобальный конфиг, который изменяется из разных мест».

```go
var Version = "dev" // подменяется линкером через -ldflags "-X main.Version=..."
```

## 5. Локальные переменные

Объявляй как можно **ближе к использованию**. Это улучшает читаемость и уменьшает риск shadowing.

```go
func handle(req *Request) error {
    v, err := parse(req)
    if err != nil {
        return err
    }
    // ... дальше используем v
}
```

## 6. Короткое объявление для результатов

`:=` делает цепочку вызовов читаемой:

```go
u, err := repo.Find(id)
if err != nil { return err }
n, err := svc.Count(u)
```

## 7. Variable block

Несколько связанных переменных — в одном блоке:

```go
var (
    mu    sync.Mutex
    cache = make(map[string]int)
)
```

## Чего избегать

- Констант, которые на самом деле параметры (переменная конфигурация).
- Общего `var X string` на уровне пакета с последующим глобальным изменением.
- Shadowing `err` в разных слоях одной функции.

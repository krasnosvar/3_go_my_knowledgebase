# Use-cases методов

## 1. Domain-объекты с поведением

```go
type Order struct{ Items []Item; Paid bool }

func (o *Order) Total() int { /* ... */ }
func (o *Order) Pay() error { /* ... */ }
```

Логика, которая оперирует именно данным объектом, — в его методах.

## 2. Value object с чистым поведением

```go
type Money struct{ Amount int64; Currency string }

func (m Money) Add(other Money) (Money, error) { /* ... */ }
```

Value receiver, ничего не меняется, результат — новый объект.

## 3. Инкапсуляция состояния

Структура + приватные поля + методы — классический способ скрыть детали:

```go
type counter struct{ n int; mu sync.Mutex }
func (c *counter) Inc()      { c.mu.Lock(); c.n++; c.mu.Unlock() }
func (c *counter) Value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.n }
```

## 4. Реализация интерфейсов

```go
type Logger interface{ Log(msg string) }

type stderrLogger struct{}
func (stderrLogger) Log(msg string) { fmt.Fprintln(os.Stderr, msg) }
```

Тип получает методы, подходящие под внешние интерфейсы. Так работают все `io.Reader`, `http.Handler`, и т.д.

## 5. Методы на именованных базовых типах

```go
type Duration time.Duration
func (d Duration) String() string { /* custom */ }

type Slug string
func (s Slug) IsValid() bool { /* ... */ }
```

Полезно, когда хочешь «умный primitive» с правилами.

## 6. Методы-фабрики

```go
func (c *Client) Users() *UserClient   { return &UserClient{c: c} }
func (c *Client) Orders() *OrderClient { return &OrderClient{c: c} }
```

Нарезаем API по подтемам. Основной объект остаётся точкой входа.

## 7. Fluent API / builder

```go
req := NewRequest().
    WithHeader("X-Token", token).
    WithTimeout(5 * time.Second)
```

Каждый метод возвращает `*T`, накапливая состояние. Используй осмотрительно: fluent API хорош, пока методы не начинают маскировать побочные эффекты.

## Когда метод — лишнее

- Простая функция-утилита без состояния — пусть остаётся обычной функцией, не навязывайся в чужой тип.
- «Помощник» получающий два `T` одинакового веса — часто это свободная функция, а не метод.

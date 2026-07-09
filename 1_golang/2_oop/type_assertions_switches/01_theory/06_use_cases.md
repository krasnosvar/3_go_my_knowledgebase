# Use-cases type assertions / switches

## 1. Обработка JSON `any`

```go
var raw interface{}
_ = json.Unmarshal(data, &raw)

switch v := raw.(type) {
case map[string]interface{}:
    // объект
case []interface{}:
    // массив
case float64:
    // число
}
```

Классический случай: JSON с неизвестной схемой.

## 2. Оптимистичное обогащение интерфейса

```go
type Flusher interface{ Flush() }

func write(w io.Writer, data []byte) error {
    _, err := w.Write(data)
    if f, ok := w.(Flusher); ok {
        f.Flush()
    }
    return err
}
```

Работает для всех, а если поддерживает Flush — использует его. Это идиома stdlib.

## 3. Анализ ошибок

```go
var pgErr *pgconn.PgError
if errors.As(err, &pgErr) && pgErr.Code == "23505" {
    return ErrDuplicate
}
```

`errors.As` = безопасный assertion по цепочке ошибок.

## 4. Универсальные логгеры

```go
func log(v any) {
    switch x := v.(type) {
    case error:
        slog.Error("err", "msg", x.Error())
    case fmt.Stringer:
        slog.Info(x.String())
    default:
        slog.Info(fmt.Sprintf("%+v", x))
    }
}
```

## 5. Event bus

Разные события — разные обработчики:

```go
func handle(ev Event) {
    switch e := ev.(type) {
    case *UserCreated:
        handleUserCreated(e)
    case *OrderPlaced:
        handleOrderPlaced(e)
    default:
        slog.Warn("unknown event", "type", fmt.Sprintf("%T", e))
    }
}
```

Удобно, когда события редко меняются. Если типов десятки — лучше registry + handler'ы.

## 6. Маппинг в HTTP статус

```go
func httpStatus(err error) int {
    switch {
    case errors.Is(err, ErrNotFound):
        return 404
    case errors.Is(err, ErrForbidden):
        return 403
    default:
        return 500
    }
}
```

Здесь — не `type switch`, а обычный `switch` + `errors.Is`, что часто удобнее.

## 7. Runtime reflection + assertion

Если пишешь фреймворк (DI, ORM, валидатор) — ассерт часто встроен в reflection-код. В приложениях это стараются избегать.

## Когда не стоит

- Если у тебя 5+ кейсов одноимённых методов — делай интерфейс.
- Если работаешь с ошибками — бери `errors.Is/As`.
- Если нужна типизация на месте вызова — generics.

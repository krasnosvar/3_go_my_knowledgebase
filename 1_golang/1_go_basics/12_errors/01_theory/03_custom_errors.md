# Custom errors

Своя ошибка нужна, когда:

- надо проверять её по значению (`errors.Is`) или по типу (`errors.As`);
- важен дополнительный контекст (код, поле, подробности);
- нужно отдавать разное поведение на разные случаи.

## Sentinel-ошибки

Простое именованное значение:

```go
var ErrNotFound = errors.New("not found")

if err := repo.Get(id); errors.Is(err, ErrNotFound) {
    return 404
}
```

Удобно для понятных, немногочисленных состояний.

## Типизированные ошибки

Структура с методом `Error()`:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

Плюс: можно достать поля через `errors.As`.

## Ошибки с оборачиванием

Если есть причина ниже, сохрани её:

```go
type DBError struct {
    Op  string
    Err error
}

func (e *DBError) Error() string { return fmt.Sprintf("%s: %v", e.Op, e.Err) }
func (e *DBError) Unwrap() error { return e.Err }
```

Метод `Unwrap` позволяет `errors.Is/As` работать по цепочке.

## Ловушка типизированного nil (typed nil)

⚠️ Частый вопрос на собеседовании.

Интерфейс в Go под капотом — это пара `(type, value)`. Он равен `nil` только когда **и тип, и значение равны nil**. Если вернуть типизированный указатель на ошибку со значением `nil` через интерфейс `error`, то интерфейс `nil` **не будет**:

```go
func checkUser() error {
    var myErr *NotFoundError = nil // типизированный nil
    return myErr                   // пара (type: *NotFoundError, value: nil)
}

func main() {
    if err := checkUser(); err != nil {
        fmt.Println("Произошла ошибка!") // ВЫВЕДЕТСЯ, хотя значение nil
    }
}
```

Решение — никогда не возвращать конкретный тип ошибки как «пустой» через интерфейс. Возвращай явный `nil`:

```go
func checkUser() error {
    if someCondition {
        return &NotFoundError{Resource: "User"}
    }
    return nil // нетипизированный nil (type: nil, value: nil)
}
```

## Чего избегать

- Возвращать `*SomeErr` как `error`, потом получать nil-интерфейс с nil-значением — это классический баг (см. выше). Возвращай `nil` напрямую, если ошибки нет.
- Писать user-facing тексты прямо в `Error()`. Оставь там технический контекст, локализацию делай на границе.
- Плодить ошибки под каждый каприз. Если достаточно одной `ErrInvalidInput` с полем — не делай 20 похожих.

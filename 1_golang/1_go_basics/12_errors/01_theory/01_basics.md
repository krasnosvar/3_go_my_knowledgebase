# Основы ошибок в Go

В Go нет исключений для бизнес-логики. Ошибка — это обычное значение, которое функция возвращает наравне с результатом, а вызывающий код обязан его проверить.

## Интерфейс error

Ошибка в Go — это любой тип, реализующий встроенный интерфейс `error`:

```go
type error interface {
    Error() string
}
```

Любой тип с методом `Error() string` является ошибкой. Значение `nil` означает «ошибки нет».

## Создание ошибок

Простую ошибку создают через `errors.New`:

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

Если нужен форматированный текст с подставленными значениями — `fmt.Errorf`:

```go
func processAge(age int) error {
    if age < 0 {
        return fmt.Errorf("invalid age: %d (must be >= 0)", age)
    }
    return nil
}
```

Текст ошибки пишут с маленькой буквы и без точки в конце — так его удобно вкладывать в цепочку контекста.

## Стандартный паттерн проверки

Основная идиома — проверять ошибку сразу и выходить (early return), оставляя основной путь плоским:

```go
val, err := doSomething()
if err != nil {
    return err // логируем, обогащаем контекстом или отдаём наверх
}
// используем val
```

При успехе возвращают `nil` в позиции ошибки. Никаких лестниц `if err == nil { ... }`.

## Освобождение ресурсов

Открытые ресурсы закрывают через `defer` сразу после успешного получения:

```go
func fileOperation() error {
    f, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    // работаем с f
    return nil
}
```

## Мини-чеклист

- Всегда проверяйте ошибку, возвращённую функцией; не глотайте её.
- `errors.New` — для простой ошибки, `fmt.Errorf` — когда нужен форматированный текст.
- В позиции ошибки при успехе возвращайте `nil`.
- Не сравнивайте текст ошибок строками.

## Куда дальше

- Добавление контекста и разбор цепочки (`%w`, `errors.Is`, `errors.As`) — [02_wrapping_is_as.md](./02_wrapping_is_as.md).
- Свои типы ошибок (sentinel и типизированные) — [03_custom_errors.md](./03_custom_errors.md).
- `panic` / `recover` и их границы применения — [04_panic_recover_boundary.md](./04_panic_recover_boundary.md).
- Идиоматические паттерны — [05_patterns.md](./05_patterns.md).
- Практические сценарии — [06_use_cases.md](./06_use_cases.md).

# Use-cases fmt

## 1. CLI output

Данные — в stdout, ошибки — в stderr:

```go
fmt.Println(result)                          // данные
fmt.Fprintln(os.Stderr, "warn:", warning)    // диагностика
```

Это позволяет пайпить stdout в другую команду, не мешая stderr.

## 2. Отладочные принты

`%+v` и `%#v` для отладки структур:

```go
fmt.Printf("%+v\n", user) // {ID:1 Name:den}
```

Временно оставлять в коде — плохо, в проде используй структурированные логи (`slog`).

## 3. Ошибки с контекстом

```go
return fmt.Errorf("save user %d: %w", u.ID, err)
```

`Errorf + %w` — каноничный способ оборачивать ошибки.

## 4. Сборка сложных сообщений пользователю

```go
return fmt.Sprintf("user %s blocked until %s", name, until.Format(time.RFC3339))
```

Для пары подстановок — норм. Для многих — используй `strings.Builder`.

## 5. Рендер для API (редко)

Если ответ — не JSON, а plain text (например, `/version`, `/metrics`), `Fprintf(w, "...")` достаточно. Для JSON всё же `encoding/json`.

## 6. Тестовая отладка

`t.Logf` уже принимает формат:

```go
t.Logf("got %+v, want %+v", got, want)
```

Не нужно промежуточных Sprintf.

## 7. Interoperable верхний уровень

Если тип реализует `fmt.Stringer`:

```go
type Status int
func (s Status) String() string { return [...]string{"pending","active"}[s] }
```

`Printf("%v", s)` автоматически использует `String()`. Удобно для `log`, `fmt.Println`, ошибок.

## Когда не использовать fmt

- Высокочастотные логи — `slog`, `zap`, `zerolog`.
- Сборка больших строк в цикле — `strings.Builder`.
- Сериализация в стабильный формат — `encoding/json`, `encoding/xml`, и т.д.

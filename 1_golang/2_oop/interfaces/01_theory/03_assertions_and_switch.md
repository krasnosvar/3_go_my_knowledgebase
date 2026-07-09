# Type assertion и type switch

## Type assertion

Извлечение конкретного типа из интерфейса:

```go
var i interface{} = "hello"

s := i.(string)          // paniс, если тип не совпадает
s, ok := i.(string)       // безопасная форма: ok == true/false
if !ok {
    // i — не строка
}
```

Всегда используй форму с `ok`, если не уверен в типе.

## Проверка реализации интерфейса

Assertion можно применять и к интерфейсам:

```go
var w io.Writer = bytes.NewBuffer(nil)

if rw, ok := w.(io.ReadWriter); ok {
    _ = rw
}
```

Удобно, чтобы **опционально** использовать расширенное поведение, если оно есть у переданной реализации.

## Type switch

Удобно, когда нужно разобрать несколько возможных типов:

```go
func describe(v interface{}) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("int: %d", x)
    case string:
        return fmt.Sprintf("string: %q", x)
    case fmt.Stringer:
        return "stringer: " + x.String()
    case nil:
        return "nil"
    default:
        return fmt.Sprintf("unknown: %T", x)
    }
}
```

- `x` уже имеет нужный тип внутри каждого `case`.
- Порядок важен: более узкие случаи — выше.
- `case nil` отлавливает nil-интерфейс отдельно.

## Когда использовать

- Разбор `any`-параметров (JSON, конфиги).
- Маршалинг ошибок в HTTP/gRPC статусы.
- Реализация «handle X types» паттерна для универсальных визуализаторов/логгеров.

## Когда НЕ использовать

- Если можно выразить через маленький интерфейс — сделай его и не занимайся switch по типам.
- Для пары кейсов проще `errors.As` / `errors.Is` (они сами работают по цепочке).
- В hot path: каждый assertion — дополнительная работа runtime.

## Anti-pattern

```go
switch x := v.(type) {
case *TypeA: /* много логики */
case *TypeB: /* много логики */
case *TypeC: /* много логики */
}
```

Это признак того, что поведение должно лежать внутри самих типов (метод интерфейса), а не в длинном switch.

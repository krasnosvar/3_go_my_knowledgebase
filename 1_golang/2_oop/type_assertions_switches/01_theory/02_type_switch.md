# Type switch

Type switch — расширенная форма `switch` для разбора типа, лежащего внутри интерфейса.

## Синтаксис

```go
func describe(v interface{}) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("int: %d", x)
    case string:
        return fmt.Sprintf("string: %q", x)
    case []int:
        return fmt.Sprintf("slice len %d", len(x))
    case nil:
        return "nil"
    default:
        return fmt.Sprintf("unknown: %T", x)
    }
}
```

- `x := v.(type)` — специальная форма, имеющая смысл только внутри switch.
- В каждом `case` переменная `x` уже имеет нужный тип.

## Множественные типы в одном case

```go
case int, int32, int64:
    // x — interface{}, так как типы разные
```

В таком случае `x` остаётся интерфейсным значением — если нужно отдельно, делай отдельные `case`.

## Case nil

Отдельно ловит nil-интерфейс. Полезно, чтобы отделить «значение не задано» от «значение неожиданное».

## Case с интерфейсами

```go
case fmt.Stringer:
    return x.String()
```

Соответствует любому типу, реализующему этот интерфейс.

## Порядок важен

Go идёт по `case` сверху вниз и выбирает первый подходящий. Более узкие типы ставь выше, общие интерфейсы — ниже.

## Когда использовать

- Обработка входа неизвестного формата (JSON `any`, конфиги).
- Маппинг ошибок в коды ответа.
- Полиморфный вывод/визуализация.

## Когда НЕ использовать

Если почти все `case` вызывают метод с одинаковой семантикой — это сигнал, что нужен интерфейс, а не switch:

```go
// плохо
switch x := v.(type) {
case *Cat: x.Speak()
case *Dog: x.Speak()
}

// хорошо
type Speaker interface{ Speak() }
v.(Speaker).Speak()
```

## Совместимость с fallthrough

`fallthrough` внутри type switch **не работает** — по дизайну. Каждый case самостоятелен.

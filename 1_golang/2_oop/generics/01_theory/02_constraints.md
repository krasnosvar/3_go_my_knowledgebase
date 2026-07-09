# Constraints (ограничения)

Constraint — это интерфейс, описывающий **набор типов**, которые разрешены для type parameter.

## Встроенные

- `any` — любой тип. Alias для `interface{}`.
- `comparable` — типы, которые можно сравнивать через `==` и `!=`.

```go
func Eq[T comparable](a, b T) bool { return a == b }

func Index[T comparable](s []T, v T) int {
    for i, x := range s {
        if x == v { return i }
    }
    return -1
}
```

## Пользовательские constraint-интерфейсы

Интерфейс с методами — обычный constraint:

```go
type Stringer interface { String() string }

func Join[T Stringer](xs []T, sep string) string {
    parts := make([]string, len(xs))
    for i, x := range xs {
        parts[i] = x.String()
    }
    return strings.Join(parts, sep)
}
```

## Type sets

Интерфейс может перечислять конкретные типы:

```go
type Number interface {
    int | int32 | int64 | float32 | float64
}

func Sum[T Number](xs []T) T {
    var s T
    for _, x := range xs { s += x }
    return s
}
```

Вертикальная черта — объединение (union).

## Аппроксимация `~`

`~int` — «любой тип, чей underlying тип — int». Это позволяет работать с пользовательскими type aliases:

```go
type Celsius float64

type Number interface {
    ~int | ~int64 | ~float64
}

Sum([]Celsius{1.0, 2.5}) // работает, потому что ~float64
```

Без `~` пользовательский тип не подошёл бы.

## constraints пакет

В `golang.org/x/exp/constraints` есть готовые:

- `Ordered` (числа + строки)
- `Integer`, `Float`, `Signed`, `Unsigned`
- `Complex`

Их легко использовать, не переизобретая велосипед. С Go 1.21+ часть переехала в `cmp.Ordered`.

## Ограничения constraint'ов

- Нельзя использовать `~` с интерфейсами (`~io.Reader` запрещено).
- Constraint — это интерфейс, его можно использовать и как обычный интерфейс, но чаще он живёт только ради generic.
- При объединении типов через `|`, методы можно указывать, но они должны быть общими.

## Правила

- Берём минимально достаточный constraint. `any` — самый слабый, `comparable` и свои union — точнее.
- Constraint описывает то, что **умеет** тип, а не «что это». Имя интерфейса — по возможности описательное (`Number`, `Stringer`).

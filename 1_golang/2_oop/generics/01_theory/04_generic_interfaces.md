# Generic interfaces

Интерфейсы тоже могут иметь type parameter.

## Простой пример

```go
type Container[T any] interface {
    Get() T
    Set(v T)
}
```

Теперь реализацию можно делать под разные типы:

```go
type IntBox struct{ v int }
func (b *IntBox) Get() int { return b.v }
func (b *IntBox) Set(v int) { b.v = v }

var c Container[int] = &IntBox{}
```

## Применение: функциональные типы

Часто удобно описать контракт обработчика:

```go
type Handler[Req, Resp any] interface {
    Handle(ctx context.Context, req Req) (Resp, error)
}
```

Плюс — типизированный контракт «запрос/ответ» для middleware-слоёв.

## Комбинация с constraint-интерфейсами

Generic-интерфейс — это **не то же самое**, что constraint-интерфейс:

```go
// constraint
type Number interface{ ~int | ~float64 }

// generic interface (runtime)
type Pair[K, V any] interface {
    Key() K
    Value() V
}
```

Constraint-интерфейс используется в `[K, V any]` местах, generic-интерфейс — как обычный интерфейс значения.

## Встраивание в generic

Generic-интерфейсы можно комбинировать:

```go
type Reader[T any] interface {
    Read(ctx context.Context) (T, error)
}

type Writer[T any] interface {
    Write(ctx context.Context, v T) error
}

type ReadWriter[T any] interface {
    Reader[T]
    Writer[T]
}
```

Работает как обычное объединение интерфейсов.

## Ограничение

Интерфейсы не поддерживают методы с собственным type parameter'ом — это то же ограничение, что у структур:

```go
// нельзя
type X interface {
    Foo[T any](v T)
}
```

## Когда использовать

- Контракт репозиториев, работающих с разными entity: `Repo[User]`, `Repo[Order]`.
- Потребители/producer-ы в шинах событий.
- Обёртки над внешними API, где ты не хочешь фиксировать конкретный тип.

## Когда избегать

- Если интерфейс фактически используется с одним типом — лишнее усложнение.
- Если дженеризация скрывает смысл — лучше несколько специализированных интерфейсов.

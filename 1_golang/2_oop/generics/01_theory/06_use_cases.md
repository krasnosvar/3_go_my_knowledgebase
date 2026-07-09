# Use-cases generics

## 1. Контейнеры / коллекции

Классический случай:

```go
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T)         { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool    { _, ok := s[v]; return ok }
func (s Set[T]) Remove(v T)      { delete(s, v) }
```

Также Stack, Queue, RingBuffer, LRUCache и т. д.

## 2. Утилиты на slices/maps

Go 1.21 официально добавил `slices` и `maps` пакеты — внутри они построены на generics:

```go
slices.Index(xs, target)
slices.Contains(xs, v)
maps.Clone(m)
maps.Equal(m1, m2)
```

Свои утилиты полезны только когда stdlib не покрывает случай.

## 3. Функции Map / Filter / Reduce

Если в проекте часто нужно «преобразовать слайс», делай их дженериками:

```go
func Map[T, U any](xs []T, f func(T) U) []U {
    res := make([]U, 0, len(xs))
    for _, x := range xs { res = append(res, f(x)) }
    return res
}
```

## 4. Result / Option

```go
type Result[T any] struct { V T; Err error }

func Wrap[T any](v T, err error) Result[T] { return Result[T]{v, err} }
```

Полезно при batch-обработке, если хочется собрать результаты и ошибки раздельно.

## 5. Generic репозитории

```go
type Repo[T any] interface {
    Save(ctx context.Context, v T) error
    Get(ctx context.Context, id int64) (T, error)
}
```

Удобно, если у тебя много сущностей с типовым CRUD. Не злоупотребляй — бизнес-логика всё равно лежит в конкретных сервисах.

## 6. Числовые алгоритмы

```go
func Min[T cmp.Ordered](a, b T) T { if a < b { return a }; return b }

func Sum[T Number](xs []T) T {
    var s T
    for _, x := range xs { s += x }
    return s
}
```

Одна функция вместо 5 специализированных.

## 7. Typed wrappers над any

До generics часто писали `interface{}` + assertion. Сейчас пишется так:

```go
type Box[T any] struct{ V T }

func (b Box[T]) Value() T { return b.V }
```

## Когда НЕ использовать

- Ровно один тип используется — лишнее.
- В сигнатуре generic читается сложнее, чем два специализированных метода — выбирай ясность.
- Высокая нагрузка на компилятор/бинарник при десятках инстанцирований — проверяй по метрикам.

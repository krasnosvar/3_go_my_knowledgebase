# Generic structs

Structs тоже могут иметь type parameter.

## Объявление

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(v T) {
    s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    v := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return v, true
}
```

`[T]` идёт после имени типа и после имени receiver.

## Инстанцирование

```go
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)

strStack := Stack[string]{}
strStack.Push("a")
```

Type parameter фиксируется при создании экземпляра.

## Методы на generic struct

Методы **не могут** иметь собственные type parameter'ы — только те, что объявлены у типа.

```go
// не работает
func (s *Stack[T]) Map[U any](fn func(T) U) []U
```

Обход: выноси в свободную функцию:

```go
func StackMap[T, U any](s *Stack[T], fn func(T) U) []U { /* ... */ }
```

Это ограничение языка, и в обозримом будущем не планируют снять.

## Generic структуры с несколькими параметрами

```go
type Pair[K, V any] struct {
    Key   K
    Value V
}
```

Хорошо для Map-like типов:

```go
type Cache[K comparable, V any] struct {
    m map[K]V
}
```

## Zero value

Пустое значение generic структуры — это нулевая структура, где поле `items []T` — nil. Обычно подходит по смыслу, но проверяй, если есть сложная инициализация.

## Когда делать свой generic тип

- Контейнер/коллекция, который нужен для разных типов (Stack, Queue, LRUCache, Set).
- Обертка над значением с метаданными (`Result[T]`, `Option[T]`).
- Конфигурируемый клиент, работающий с типизированными сущностями.

## Когда НЕ делать

- Если тип используется только с одним конкретным типом элемента — обычная структура проще.
- Если generic только ради generic — будет overengineering.

## Embedding

Generic структуры можно встраивать обычным способом, но при встраивании нужно фиксировать тип:

```go
type LoggedCache[K comparable, V any] struct {
    *Cache[K, V]
    log *slog.Logger
}
```

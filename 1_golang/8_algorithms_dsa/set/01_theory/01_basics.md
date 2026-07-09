# Set в Go

В Go нет встроенного типа Set. Идиома — `map[T]struct{}` или `map[T]bool`.

## map[T]struct{}

```go
s := map[string]struct{}{}

s["foo"] = struct{}{}
_, exists := s["foo"]
delete(s, "foo")

for k := range s {
    fmt.Println(k)
}
fmt.Println(len(s))
```

`struct{}` занимает 0 байт — экономия памяти на больших set'ах.

## map[T]bool — альтернатива

```go
s := map[string]bool{}
s["foo"] = true
if s["foo"] { ... }
```

Чуть короче в записи, +1 байт на элемент. Удобно, когда уже привык к bool-паттерну.

## Generic Set

```go
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T)       { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool  { _, ok := s[v]; return ok }
func (s Set[T]) Delete(v T)    { delete(s, v) }
func (s Set[T]) Len() int      { return len(s) }
```

С Go 1.18+ — стандартный способ типизированного set'а без сторонних библиотек.

## Операции

```go
// Union
func Union[T comparable](a, b Set[T]) Set[T] {
    res := Set[T]{}
    for v := range a { res.Add(v) }
    for v := range b { res.Add(v) }
    return res
}

// Intersection
func Intersect[T comparable](a, b Set[T]) Set[T] {
    res := Set[T]{}
    for v := range a {
        if b.Has(v) { res.Add(v) }
    }
    return res
}

// Difference: a - b
func Diff[T comparable](a, b Set[T]) Set[T] {
    res := Set[T]{}
    for v := range a {
        if !b.Has(v) { res.Add(v) }
    }
    return res
}
```

## Ограничения ключей

Тип должен быть `comparable`:

- примитивы (int, string, bool, float);
- pointers;
- channels;
- interfaces (если не содержат noncomparable внутри);
- arrays из comparable;
- structs из comparable полей.

**Нельзя**: slice, map, func.

Композитный ключ — массив:

```go
type Key [2]int
s := map[Key]struct{}{}
s[Key{1, 2}] = struct{}{}
```

Для сложных композитов часто проще: канонизированная строка-ключ или отдельная struct.

## Concurrency

Обычный `map` не потокобезопасен. Параллельный доступ → race / panic. Варианты:

- `map` + `sync.Mutex` (или `sync.RWMutex`);
- `sync.Map` для специфических сценариев (read-heavy с редкими записями ключей-долгожителей);
- per-shard sharded map для high-load случаев.

## Когда использовать

- Дедупликация: «уникальные посетители за час».
- Membership: «есть ли user среди admin'ов».
- Permissions / tags: «у роли есть права X, Y».

## Когда НЕ нужен

- Если данных мало (до ~10) и сравнение быстро — slice + linear search проще и часто быстрее (cache-friendly).
- Когда порядок важен — set его не сохраняет.

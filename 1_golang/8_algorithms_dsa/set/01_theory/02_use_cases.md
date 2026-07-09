# Set: use cases и patterns

## 1. Дедупликация

```go
seen := map[string]struct{}{}
out := []string{}
for _, x := range items {
    if _, ok := seen[x]; ok { continue }
    seen[x] = struct{}{}
    out = append(out, x)
}
```

Сохраняет порядок первого вхождения.

## 2. Visited tracking в графах

```go
visited := map[*Node]struct{}{}
var dfs func(n *Node)
dfs = func(n *Node) {
    if _, ok := visited[n]; ok { return }
    visited[n] = struct{}{}
    for _, neighbor := range n.Neighbours {
        dfs(neighbor)
    }
}
```

Без множества visited — бесконечная рекурсия в циклических графах.

## 3. Membership check в фильтрах

```go
allowedRoles := map[string]struct{}{"admin": {}, "moderator": {}}
if _, ok := allowedRoles[user.Role]; !ok {
    return ErrForbidden
}
```

Read-only set, инициализированный один раз — основной паттерн для конфигов.

## 4. Множественные операции

Используешь Union / Intersection / Difference, когда нужно сравнивать наборы:

```go
a := Set[string]{"go", "rust", "python"}
b := Set[string]{"rust", "java", "kotlin"}

both := Intersect(a, b)        // {"rust"}
onlyA := Diff(a, b)            // {"go", "python"}
onlyB := Diff(b, a)            // {"java", "kotlin"}
```

## 5. Permission flags (компактно)

Если значений мало (до ~64) и они известны заранее — битовая маска (см. `1_go_basics/operators_bitwise/03_bitwise_masks`) часто компактнее и быстрее set'а.

```go
type Perm uint8
const (
    PermRead  Perm = 1 << iota
    PermWrite
    PermDelete
)

user.Perms = PermRead | PermWrite
if user.Perms&PermDelete != 0 { ... }
```

## 6. Ordered set

Set не гарантирует порядок. Если нужен ordered set:

- `sorted slice` + `sort.SliceIsSorted` + binary search;
- `treemap` через стороннюю библиотеку;
- maintain set + slice одновременно.

Стандартной библиотеки для ordered set в Go нет.

## 7. Concurrent set

`sync.Map` подходит для специфических сценариев:

- много read, редкие write новых ключей;
- ключи долго живут после первой записи;
- значения могут обновляться часто.

В остальных случаях `map[T]struct{}` + `sync.Mutex` понятнее и не медленнее.

## Правило

- map[T]struct{} — default.
- Generic Set — когда нужны переиспользуемые операции.
- Битовая маска — для маленького фиксированного набора.
- Concurrent — mutex или sync.Map по ситуации.

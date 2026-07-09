# sort и slices.Sort

Пакет `sort` — сортировка в Go. С Go 1.21 появился `slices`, который чаще удобнее.

## Быстрый путь (slices пакет, Go 1.21+)

```go
import "slices"

slices.Sort(nums)                          // ints
slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Age, b.Age)
})
slices.IsSorted(nums)
slices.BinarySearch(nums, 42)
slices.Reverse(nums)
```

Это type-safe generic API, работает с любыми слайсами.

## Классический sort

Если Go < 1.21 или нужно работать с типом, реализующим свой `sort.Interface`:

```go
sort.Ints(nums)
sort.Strings(names)
sort.Float64s(vals)

sort.Slice(users, func(i, j int) bool {
    return users[i].Age < users[j].Age
})

sort.Sort(sortableType)    // полный контроль через Len/Less/Swap
sort.Stable(sortableType)  // сохраняет порядок равных
```

### sort.Interface

```go
type ByAge []User
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

sort.Sort(ByAge(users))
```

Нужен, если ты делаешь свой специальный порядок или хочешь повторно использовать.

## Stable vs unstable

- `sort.Sort` — быстрее, но не гарантирует порядок для равных элементов.
- `sort.Stable` — сохраняет исходный порядок при равенстве.

Для «сортировать по age, потом по имени» — либо Stable + двойная сортировка (сначала по имени, потом Stable по age), либо composite Less.

## Reverse

```go
sort.Sort(sort.Reverse(sort.IntSlice(nums)))
// или просто: slices.SortFunc с обратным сравнением
```

## Binary search

```go
idx, found := slices.BinarySearch(sorted, 42)
```

На отсортированном слайсе. Если нужен свой предикат — `slices.BinarySearchFunc` или `sort.Search`.

## Частые ошибки

- Сортировать слайс, который ещё используется в другой горутине — гонка.
- `sort.Slice` — медленнее, чем `sort.Sort(sort.Interface)`, потому что reflection. На миллионах элементов заметно.
- Забыть stable при sort-on-sort — порядок по первому ключу теряется.
- Для сложной сортировки использовать regex/string compare вместо явного `cmp.Compare` на типах.

## Правило

- Go 1.21+: `slices.Sort` / `slices.SortFunc` — идиоматично.
- Меньшие версии: `sort.Slice` для разовой сортировки, `sort.Interface` для переиспользуемого порядка.

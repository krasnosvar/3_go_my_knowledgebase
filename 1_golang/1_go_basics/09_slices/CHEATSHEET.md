# Слайсы — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии указаны там, где поведение зависит от них.

## Массив vs слайс, заголовок

Подробно: [01_basics_len_cap.md](./01_theory/01_basics_len_cap.md).

| | Массив `[N]T` | Слайс `[]T` |
| :--- | :--- | :--- |
| Размер | фиксирован, часть типа (`[4]int` ≠ `[5]int`) | меняется |
| Что копируется при присваивании/передаче | все элементы | заголовок 24 байта: `ptr`, `len`, `cap` |
| Нулевое значение | массив из нулей | `nil` (`len == cap == 0`) |
| Сравнение `==` | да, если элементы сравнимы | только с `nil` (`slices.Equal` — по элементам) |

| Задача | Код |
| :--- | :--- |
| Литерал | `s := []int{10, 20, 30}` → `len == cap == 3` |
| С длиной и ёмкостью | `s := make([]int, 2, 5)` → `[0 0]`, `cap 5` |
| Пустой с запасом под `append` | `s := make([]T, 0, n)` |
| Нарезка | `s[low:high]` → `len = high-low`, `cap = cap(s)-low` |
| Three-index (ограничить cap) | `s[low:high:max]` → `cap = max-low` |
| nil vs пустой | `var s []int` → `s == nil`; `[]int{}` → не `nil`; оба `len 0`, оба годны для `append`/`range` |
| Очистить элементы (Go 1.21+) | `clear(s)` — зануляет, `len` не меняется |

## append, рост, copy

Подробно: [02_append_and_copy.md](./01_theory/02_append_and_copy.md).

| Ситуация | Что делает `append` |
| :--- | :--- |
| `len < cap` | пишет в тот же backing array — изменения видят все слайсы на него |
| `len == cap` | выделяет новый массив, копирует, связь со старым слайсом рвётся |
| Всегда | `s = append(s, v)` — результат присваивать обратно |
| Несколько / слайс | `s = append(s, 1, 2, 3)`, `s = append(s, other...)` |

| Рост ёмкости (Go 1.18+) | |
| :--- | :--- |
| Нужно больше `2 × oldcap` | берётся нужная ёмкость |
| `oldcap < 256` | удвоение |
| `oldcap ≥ 256` | `newcap += (newcap + 3*256) / 4` — коэффициент плавно падает от 2x к ~1.25x |
| Затем | округление вверх до size class аллокатора (например, 512 → 848 для `[]int`, а не 832) |

| Задача | Код |
| :--- | :--- |
| Скопировать | `n := copy(dst, src)` → `n = min(len(dst), len(src))` |
| Независимая копия | `c := make([]T, len(s)); copy(c, s)` или `slices.Clone(s)` (1.21+) |
| Удалить `i` с сохранением порядка | `copy(s[i:], s[i+1:]); s = s[:len(s)-1]` или `s = slices.Delete(s, i, i+1)` |
| Удалить `i` без порядка, O(1) | `s[i] = s[len(s)-1]; s = s[:len(s)-1]` |
| Фильтр in-place без аллокаций | `out := s[:0]; for _, v := range s { if keep(v) { out = append(out, v) } }` |

## Общий backing array и утечки

Подробно: [03_subslice_pitfalls.md](./01_theory/03_subslice_pitfalls.md), [04_memory_retention.md](./01_theory/04_memory_retention.md).

```go
s1 := []int{10, 20, 30, 40}
s2 := s1[1:3]       // [20 30], len 2, cap 3 — общий массив с s1
s2 = append(s2, 99) // len < cap: пишет в s1[3]
fmt.Println(s1)     // [10 20 30 99]

s3 := []int{10, 20, 30, 40}
s4 := s3[1:3:3]     // cap 2 → append обязан выделить новый массив
s4 = append(s4, 99)
fmt.Println(s3, s4) // [10 20 30 40] [20 30 99]
```

| Проблема | Решение |
| :--- | :--- |
| Изменение подслайса меняет исходный | `copy` в новый слайс перед изменением |
| `append` в подслайс затирает соседей | three-index `s[a:b:b]` или `copy` |
| Маленький подслайс держит огромный массив (GC не освободит) | `clip := make([]byte, n); copy(clip, huge[:n])` |
| Конкурентный `append` в общий слайс | `sync.Mutex` или канал; иначе data race |

## Сложность операций

| Операция | Сложность |
| :--- | :--- |
| `s[i]`, `s[i] = v`, `len`, `cap` | O(1) |
| `append` в конец | O(1) амортизированно (O(n) при переаллокации) |
| Нарезка `s[a:b]` | O(1), без копирования |
| Вставка / удаление в середине с порядком | O(n) |
| `copy`, `slices.Clone` | O(n) |
| Поиск `slices.Index` / `slices.Contains` | O(n) |
| `slices.Sort` | O(n log n) |
| `slices.BinarySearch` (по отсортированному) | O(log n) |

## Пакет slices (Go 1.21+)

| Задача | Код |
| :--- | :--- |
| Сортировка | `slices.Sort(s)`, `slices.SortFunc(s, func(a, b T) int { return cmp.Compare(a.X, b.X) })` |
| Поиск | `slices.Index(s, v)`, `slices.Contains(s, v)`, `i, ok := slices.BinarySearch(s, v)` |
| Сравнение | `slices.Equal(a, b)` |
| Мин/макс | `slices.Min(s)`, `slices.Max(s)` (panic на пустом) |
| Вставка / удаление | `s = slices.Insert(s, i, v...)`, `s = slices.Delete(s, i, j)` |
| Разное | `slices.Reverse`, `slices.Compact`, `slices.Clone`, `slices.Concat` (1.22+) |
| Из итератора (1.23+) | `slices.Collect(seq)`, `slices.Sorted(maps.Keys(m))` |

## Ловушки

- ⚠️ `append(s, v)` без присваивания — значение теряется (при переаллокации — и новые данные).
- ⚠️ `append` внутри функции не меняет `len` слайса у вызывающего — верните слайс.
- ⚠️ `b := append(a, x)` при свободной ёмкости делит массив с `a`: следующий `append(a, y)` перезапишет `b[len(a)]`.
- ⚠️ `make([]int, n)` + `append` → в начале `n` нулей; нужен `make([]int, 0, n)`.
- ⚠️ Выход за `len` → `panic: index out of range` (даже если `cap` больше).
- ⚠️ `for _, v := range s { v.X = 1 }` меняет копию элемента; изменять через `s[i].X`.
- ⚠️ Возврат подслайса большого буфера (файл, сетевой ответ) удерживает весь буфер в памяти.

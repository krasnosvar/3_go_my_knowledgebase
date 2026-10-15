# Алгоритмы и структуры данных — шпаргалка

Подробности и объяснения — в `01_theory/` каждой темы (ссылки в разделах). ⚠️ — частая ловушка.
`am` — амортизированно. Версии указаны там, где API зависит от них (`slices` — 1.21+).

## Что выбрать

| Нужно | Структура в Go | Операции |
| :--- | :--- | :--- |
| Последовательность, доступ по индексу | `[]T` (default) | append O(1) am, index O(1) |
| LIFO | `[]T`, push/pop в конце | O(1) am |
| FIFO | `[]T` + `q[1:]` (небольшие), ring buffer, `container/list`; между горутинами — канал | O(1) am |
| Вставка/удаление в середине по указателю, LRU | `container/list` + `map` | O(1) |
| Принадлежность, дедупликация | `map[T]struct{}` | O(1) am |
| Маленький фиксированный набор флагов (≤ 64) | битовая маска `uint64` | O(1) |
| Min/max, top-K, расписание по времени | `container/heap` | push/pop O(log n), peek O(1) |
| Порядок ключей, диапазоны | sorted slice + `slices.BinarySearch`, B-tree (`google/btree`) | O(log n) поиск |
| Связность при добавлении рёбер | DSU (union-find) | ≈ O(1) am |

## 01. Linked list

Подробно: [01_basics.md](./01_linked_list/01_theory/01_basics.md), [02_problems.md](./01_linked_list/01_theory/02_problems.md).

| Операция | Slice | Singly list | Doubly list |
| :--- | :--- | :--- | :--- |
| Push back | O(1) am | O(n), O(1) с tail | O(1) |
| Push front | O(n) | O(1) | O(1) |
| Pop back | O(1) | O(n) | O(1) |
| Pop front | O(n) | O(1) | O(1) |
| Insert в середину (есть указатель) | O(n) | O(1) | O(1) |
| Доступ по индексу | O(1) | O(n) | O(n) |

| Задача | Приём |
| :--- | :--- |
| Развернуть | `prev, cur`: `next := cur.Next; cur.Next = prev; prev = cur; cur = next` |
| Цикл (Floyd) | `slow` +1, `fast` +2; встретились → цикл |
| Середина | тот же slow/fast без цикла: `slow` на середине, когда `fast` дошёл до конца |
| Слить два отсортированных | dummy-узел + `tail` |
| N-й с конца | `fast` вперёд на N, затем оба до конца |
| Пересечение двух списков | пройти A затем B и B затем A — встретятся в пересечении |
| `container/list` | `l := list.New(); l.PushBack(v); for e := l.Front(); e != nil; e = e.Next() { e.Value }` |

## 02. Stack и queue

Подробно: [01_basics.md](./02_stack_queue/01_theory/01_basics.md), [02_problems.md](./02_stack_queue/01_theory/02_problems.md).

| Задача | Код |
| :--- | :--- |
| Push | `s = append(s, v)` |
| Pop | `v := s[len(s)-1]; s = s[:len(s)-1]` |
| Peek | `s[len(s)-1]` (проверить `len(s) > 0`) |
| Enqueue / Dequeue | `q = append(q, v)` / `v := q[0]; q = q[1:]` |
| Generic стек | `type Stack[T any] []T` + методы на `*Stack[T]` |
| Скобки | стек открывающих; на закрывающей сверху должна быть пара |
| Min stack за O(1) | второй стек минимумов (push при `v <= min`) |
| Очередь из двух стеков | `in` / `out`, перекладывать при пустом `out` — O(1) am |
| Sliding window max | monotonic deque — O(n) |
| BFS / DFS без рекурсии | очередь / стек + `visited` |

## 03. Set

Подробно: [01_basics.md](./03_set/01_theory/01_basics.md), [02_use_cases.md](./03_set/01_theory/02_use_cases.md).

| Задача | Код |
| :--- | :--- |
| Создать | `s := map[string]struct{}{}` (`struct{}` — 0 байт) |
| Добавить / проверить / удалить | `s[k] = struct{}{}`; `_, ok := s[k]`; `delete(s, k)` |
| Литерал | `allowed := map[string]struct{}{"admin": {}, "moderator": {}}` |
| Вариант с bool | `map[string]bool` — короче проверка `if s[k]`, +1 байт на элемент |
| Generic (1.18+) | `type Set[T comparable] map[T]struct{}` |
| Union / Intersect / Diff | обойти `a` (и `b`), проверять `Has` во втором |
| Составной ключ | массив `[2]int` или struct из comparable-полей |
| Порядок нужен | set + slice, или sorted slice + `slices.BinarySearch` |
| Конкурентно | `map` + `sync.Mutex`/`RWMutex`; `sync.Map` — read-heavy, ключи-долгожители |

Ключ map — только `comparable`: числа, строки, bool, указатели, каналы, интерфейсы, массивы и структуры из comparable. Нельзя: slice, map, func.

## 04. Tree и heap

Подробно: [01_trees.md](./04_tree_heap/01_theory/01_trees.md), [02_heap.md](./04_tree_heap/01_theory/02_heap.md).

| Обход | Порядок | Применение |
| :--- | :--- | :--- |
| Preorder | root, left, right | копирование, сериализация |
| Inorder | left, root, right | BST → отсортированный вывод |
| Postorder | left, right, root | удаление, высота |
| Level order (BFS) | по уровням, через очередь | ширина, ближайший лист |

| Операция | BST сбалансированный | BST вырожденный | `map` | Heap | Sorted slice |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Вставка | O(log n) | O(n) | O(1) am | O(log n) | O(n) |
| Поиск | O(log n) | O(n) | O(1) am | O(n) | O(log n) |
| Min | O(log n) | O(n) | O(n) | O(1) | O(1) |
| Извлечь min | O(log n) | O(n) | O(n) | O(log n) | O(n) со сдвигом |
| Построить из n | O(n log n) | — | O(n) | O(n) (`heap.Init`) | O(n log n) |

| `container/heap` | Суть |
| :--- | :--- |
| Реализовать | `Len`, `Less`, `Swap` + `Push(x any)`, `Pop() any` на указателе |
| Пользоваться | функциями пакета `heap.Init(h)`, `heap.Push(h, x)`, `heap.Pop(h)`, `heap.Fix(h, i)` |
| Max-heap | `Less` возвращает `h[i] > h[j]` |
| Tie-break | сравнить второй ключ в `Less` (например, `CreatedAt.Before`) |
| Top-K | min-heap размера K — O(n log K) |
| Медиана потока | max-heap нижней половины + min-heap верхней |

```go
package main

import (
	"container/heap"
	"fmt"
	"slices"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // > — max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old) - 1
	v := old[n]
	*h = old[:n]
	return v
}

// topK: min-heap размера k, O(n log k)
func topK(nums []int, k int) []int {
	h := &IntHeap{}
	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h) // выкидываем минимум
		}
	}
	res := []int(*h)
	slices.Sort(res)
	return res
}

func main() {
	h := &IntHeap{2, 1, 5, 0, 3}
	heap.Init(h)                                  // O(n)
	heap.Push(h, -1)                              // функция пакета, не метод!
	fmt.Println((*h)[0])                          // peek O(1): -1
	fmt.Println(heap.Pop(h))                      // -1
	fmt.Println(topK([]int{5, 1, 9, 3, 7, 2}, 3)) // [5 7 9]

	s := []int{1, 3, 5, 7}
	i, found := slices.BinarySearch(s, 5)
	fmt.Println(i, found) // 2 true
}
```

## 05. Graph

Подробно: [01_basics.md](./05_graph/01_theory/01_basics.md), [02_algorithms.md](./05_graph/01_theory/02_algorithms.md).

| Представление | Код | Память | Соседи / ребро |
| :--- | :--- | :--- | :--- |
| Adjacency list (default, разреженный) | `map[int][]int` | O(V + E) | O(degree) / O(degree) |
| Adjacency matrix (плотный) | `[][]bool` | O(V²) | O(V) / O(1) |
| Edge list | `[]Edge{From, To, Weight}` | O(E) | Kruskal: сортировка + DSU |

| Задача | Алгоритм | Сложность |
| :--- | :--- | :--- |
| Кратчайший путь без весов | BFS (`dist[n] = dist[cur] + 1`) | O(V + E) |
| Веса ≥ 0 | Dijkstra + min-heap (пропускать устаревшие записи `if it.D > dist[it.V]`) | O((V + E) log V) |
| Отрицательные веса | Bellman-Ford | O(V·E) |
| Путь в сетке с эвристикой | A* | зависит от эвристики |
| Компоненты связности | DFS/BFS из каждой непосещённой | O(V + E) |
| Цикл, неориентированный | DFS: посещённый сосед ≠ родитель | O(V + E) |
| Цикл, ориентированный | DFS трёх цветов: сосед grey → цикл | O(V + E) |
| Топосорт DAG | DFS post-order + reverse, или Kahn (in-degree 0 в очередь; обработано < V → цикл) | O(V + E) |
| Двудольность | BFS/DFS с раскраской в 2 цвета | O(V + E) |
| Динамическая связность, MST | DSU: path compression + union by rank | ≈ O(1) am |

## 06. LeetCode-паттерны

Подробно: [01_patterns.md](./06_leetcode_tasks/01_theory/01_patterns.md), [02_complexity_and_tactics.md](./06_leetcode_tasks/01_theory/02_complexity_and_tactics.md).

| Признак в условии | Паттерн |
| :--- | :--- |
| Отсортированный массив, пара/тройка | two pointers, binary search |
| Подстрока / подмассив | sliding window (`right` расширяет, `left` сжимает) |
| Частоты, «уже видели», two sum без сортировки | hash map |
| Скобки, «следующий больший» | stack / monotonic stack |
| Сетка, острова, связи | BFS / DFS / union-find |
| «Минимальное X, при котором выполнимо» | binary search по ответу |
| «Сколько способов», min/max с подзадачами | DP: сформулировать state и transition |
| Интервалы, локально лучший шаг | greedy (нужно доказательство) |
| Все перестановки / сочетания | backtracking |
| Top-K, merge K списков, медиана потока | heap |

| n | Допустимая сложность (≈ 1 с, 10⁷–10⁸ операций) |
| :--- | :--- |
| ≤ 20 | O(2ⁿ), O(n!) при n ≤ 10 |
| ≤ 10³–10⁴ | O(n²) |
| ≤ 10⁵–10⁶ | O(n log n), O(n) |
| > 10⁶ | O(n), O(log n), O(1) |

| Go на собеседовании | Код |
| :--- | :--- |
| Сортировка | `slices.Sort(s)` (1.21+), `sort.Slice(s, func(i, j int) bool {...})` |
| Бинарный поиск | `slices.BinarySearch(s, x)` (1.21+), `sort.Search(n, pred)` |
| Начальные min/max | `math.MaxInt`, `math.MinInt`; встроенные `min(a, b)`, `max(a, b)` (1.21+) |
| Символы Unicode | `s[i]` — байт; `[]rune(s)` или `for _, r := range s` |
| Счётчик | `cnt := map[rune]int{}`; `cnt[r]++` (нулевое значение 0) |

## Ловушки

- ⚠️ Linked list «для скорости» — slice почти всегда быстрее из-за cache locality.
- ⚠️ `q = q[1:]` не освобождает базовый массив → долгоживущая очередь держит память; обнулять/копировать или ring buffer.
- ⚠️ Pop из пустого слайса → panic: index out of range; сначала `len(s) > 0`.
- ⚠️ `h.Push(x)` вместо `heap.Push(h, x)` — элемент добавлен без восстановления свойства кучи.
- ⚠️ `fmt.Println((*h)[0], heap.Pop(h))` — порядок вычисления индексации и вызова не определён; peek и pop — отдельными выражениями.
- ⚠️ Методы `Push`/`Pop` для heap — на указателе (`*IntHeap`), иначе изменения потеряются.
- ⚠️ BST на отсортированных вставках вырождается в список — O(n).
- ⚠️ DFS/BFS без `visited` на графе с циклом → бесконечный цикл / рекурсия.
- ⚠️ Глубокая рекурсия DFS — стек растёт до 1 GB, затем fatal; для огромных графов — итеративно.
- ⚠️ Dijkstra с отрицательными весами даёт неверный ответ — нужен Bellman-Ford.
- ⚠️ Итерация по `map` — случайный порядок: для детерминированного вывода сортировать ключи.
- ⚠️ Обычный `map` из нескольких горутин → `fatal error: concurrent map writes`.
- ⚠️ `len(s)` — байты, а не символы; индексация строки даёт `byte`.
- ⚠️ Переполнение `int` при `(l+r)/2` на огромных индексах → `l + (r-l)/2`.

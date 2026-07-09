# Heap (priority queue)

Heap — частично-упорядоченное дерево, обычно представляется массивом. Корень — минимум (min-heap) или максимум (max-heap).

## container/heap

В stdlib Go heap реализуется через интерфейс `container/heap`. Ты пишешь сам структуру, реализуешь 5 методов, дальше пакет даёт вспомогательные функции.

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old) - 1
    v := old[n]
    *h = old[:n]
    return v
}

func main() {
    h := &IntHeap{2, 1, 5, 0, 3}
    heap.Init(h)            // O(n)
    heap.Push(h, -1)         // O(log n)
    fmt.Println(heap.Pop(h)) // -1
}
```

`heap.Push` / `heap.Pop` — функции пакета, не методы! Это путает новичков. Методы `Push`/`Pop` — внутренние, для интерфейса.

## Сложности

- `Init` — O(n).
- `Push` / `Pop` — O(log n).
- Доступ к минимуму (peek): O(1) — это `(*h)[0]`.

## Use cases

### 1. Top-K

«Найти K самых больших / маленьких из stream»:

```go
// min-heap размера K. Если новый элемент больше минимума — заменить.
```

K = размер heap'а. Вместимость и сложность — O(N log K).

### 2. Dijkstra / A*

Pathfinding в взвешенном графе. Heap хранит фронтир, на каждом шаге берём узел с минимальной дистанцией.

### 3. Scheduler / event loop

Задачи с deadline'ом — min-heap по времени. На каждом тике достаём готовые.

### 4. Median maintenance

Два heap'а: min для верхней половины, max для нижней. На каждом insert переливаем, чтобы сохранить размер. Median — корни.

## Custom comparators

Через `Less` можно реализовать любую логику:

```go
func (h ItemHeap) Less(i, j int) bool {
    if h[i].Priority != h[j].Priority {
        return h[i].Priority < h[j].Priority
    }
    return h[i].CreatedAt.Before(h[j].CreatedAt)
}
```

Тie-breaking второй колонкой — обычно по времени.

## Сравнение с sorted-slice

| Операция | Heap | Sorted slice |
|---|---|---|
| Push | O(log n) | O(n) (insert with shift) |
| Pop min | O(log n) | O(1) (pop front) или O(n) (shift) |
| Build from N | O(n) | O(n log n) |
| Iterate sorted | требует pop'ов | O(n) |

Heap — лучший выбор, когда добавления/удаления частые и порядок не нужен в каждый момент.

## Правило

- top-K, scheduling, Dijkstra → heap.
- Просто отсортировать — `slices.Sort`.
- Хочешь sorted set с диапазонными запросами — BST/B-tree, не heap.

# Graph: алгоритмы

## Connected components

Сколько изолированных подграфов:

```go
func countComponents(g map[int][]int) int {
    visited := map[int]bool{}
    count := 0
    var dfs func(int)
    dfs = func(v int) {
        visited[v] = true
        for _, n := range g[v] {
            if !visited[n] { dfs(n) }
        }
    }
    for v := range g {
        if !visited[v] {
            count++
            dfs(v)
        }
    }
    return count
}
```

## Cycle detection

### Undirected

DFS с проверкой «соседа, отличного от родителя»:

```go
var dfs func(v, parent int) bool
dfs = func(v, parent int) bool {
    visited[v] = true
    for _, n := range g[v] {
        if !visited[n] {
            if dfs(n, v) { return true }
        } else if n != parent {
            return true // нашли цикл
        }
    }
    return false
}
```

### Directed

Three-color DFS (white/grey/black). Если нашёл grey-сосед — есть цикл.

## Topological sort (DAG)

Только для направленного ациклического графа.

### DFS post-order

```go
var order []int
var visited map[int]bool

var dfs func(int)
dfs = func(v int) {
    visited[v] = true
    for _, n := range g[v] {
        if !visited[n] { dfs(n) }
    }
    order = append(order, v)
}

// потом reverse(order)
```

### Kahn's (BFS)

Считаем in-degree, кладём в очередь все с in-degree=0, на каждом шаге снимаем in-degree у соседей. Удобнее для проверки «является ли DAG»: если осталось < V — есть цикл.

## Shortest path

### BFS — невзвешенный

Уже описан в `01_basics.md`.

### Dijkstra — неотрицательные веса

Heap (min-priority queue) по дистанции:

```go
type Item struct{ V, D int }
// PQ из container/heap по D

dist := map[int]int{start: 0}
pq := &Items{{start, 0}}
heap.Init(pq)

for pq.Len() > 0 {
    it := heap.Pop(pq).(Item)
    if it.D > dist[it.V] { continue }
    for _, e := range g[it.V] {
        nd := it.D + e.W
        if d, ok := dist[e.To]; !ok || nd < d {
            dist[e.To] = nd
            heap.Push(pq, Item{e.To, nd})
        }
    }
}
```

Сложность: O((V+E) log V) с heap.

### Bellman-Ford

Для графов с отрицательными весами. O(V·E).

### A*

Dijkstra + эвристика. Хорош для путей в сетке.

## Bipartite check

Двудольный граф — вершины можно покрасить в 2 цвета так, что соседи разных цветов. BFS, чередуя цвета. Если конфликт — не bipartite.

## Union-Find (DSU)

Структура для динамической поддержки связности:

```go
type DSU struct{ parent, rank []int }

func (d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x]) // path compression
    }
    return d.parent[x]
}

func (d *DSU) Union(x, y int) {
    rx, ry := d.Find(x), d.Find(y)
    if rx == ry { return }
    if d.rank[rx] < d.rank[ry] { rx, ry = ry, rx }
    d.parent[ry] = rx
    if d.rank[rx] == d.rank[ry] { d.rank[rx]++ }
}
```

Используется в Kruskal MST, проверке связности при добавлении рёбер.

## Правило

- Кратчайший путь без весов → BFS.
- С весами ≥ 0 → Dijkstra + heap.
- DAG ordering → topo-sort.
- Динамическая связность → DSU.
- Двудольность / окраска → BFS/DFS с раскраской.

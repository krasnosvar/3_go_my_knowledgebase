# Graph: основы

Граф — множество вершин и рёбер. В Go нет встроенного типа графа — представление выбирается под задачу.

## Представления

### Adjacency list (чаще всего)

```go
type Graph map[int][]int

g := Graph{
    1: {2, 3},
    2: {4},
    3: {4},
    4: {},
}
```

Память: O(V + E). Поиск соседей: O(degree).

### Adjacency matrix

```go
matrix := [][]bool{
    {false, true,  true,  false},
    {false, false, false, true},
    // ...
}
```

Память: O(V²). Проверка ребра: O(1). Уместно для плотных графов или когда часто проверяешь «есть ли ребро».

### Edge list

```go
type Edge struct{ From, To, Weight int }
edges := []Edge{{1,2,5}, {2,3,3}}
```

Удобно для алгоритмов типа Kruskal (sort by weight + union-find).

## Виды

- **Directed / Undirected** — направленные/нет.
- **Weighted** — с весами рёбер.
- **Cyclic / DAG** — есть ли циклы. DAG поддерживает топологическую сортировку.
- **Connected** — есть путь между любыми двумя вершинами.

## BFS

Поиск в ширину. Минимальное число рёбер до каждой вершины (без весов).

```go
func bfs(g Graph, start int) map[int]int {
    dist := map[int]int{start: 0}
    q := []int{start}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for _, n := range g[cur] {
            if _, ok := dist[n]; ok { continue }
            dist[n] = dist[cur] + 1
            q = append(q, n)
        }
    }
    return dist
}
```

## DFS

Поиск в глубину. Хорош для: проверки связности, поиска компонент, детекции циклов, топосорта.

```go
func dfs(g Graph, start int) {
    visited := map[int]bool{}
    var rec func(int)
    rec = func(v int) {
        if visited[v] { return }
        visited[v] = true
        for _, n := range g[v] {
            rec(n)
        }
    }
    rec(start)
}
```

Итеративный вариант — через stack (см. stack_queue/02_problems).

## Сложность

- BFS / DFS на adjacency list: O(V + E).
- На adjacency matrix: O(V²).

## Когда какой обход

- Кратчайший путь без весов → BFS.
- Связность / компоненты / детекция циклов → DFS (или BFS).
- Топологическая сортировка → DFS post-order или Kahn's algorithm (BFS).

## Правило

- Sparse граф (мало рёбер) → adjacency list.
- Dense → adjacency matrix или edge list по контексту.
- BFS for distance, DFS for structure.

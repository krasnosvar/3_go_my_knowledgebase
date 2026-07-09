# Stack/Queue: типовые задачи

## 1. Balanced parentheses

```go
func isBalanced(s string) bool {
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
    stack := []rune{}
    for _, r := range s {
        switch r {
        case '(', '[', '{':
            stack = append(stack, r)
        case ')', ']', '}':
            if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
```

Stack хранит открывающие скобки, при закрывающей — проверяем, что сверху ровно нужная пара.

## 2. Min stack

Стек, который дополнительно отдаёт минимум за O(1):

```go
type MinStack struct {
    data, mins []int
}

func (s *MinStack) Push(v int) {
    s.data = append(s.data, v)
    if len(s.mins) == 0 || v <= s.mins[len(s.mins)-1] {
        s.mins = append(s.mins, v)
    }
}

func (s *MinStack) Pop() {
    v := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    if v == s.mins[len(s.mins)-1] {
        s.mins = s.mins[:len(s.mins)-1]
    }
}

func (s *MinStack) Min() int { return s.mins[len(s.mins)-1] }
```

Идея: вспомогательный стек минимумов хранит «текущий минимум за всю историю pop'ов».

## 3. Queue из двух стеков

Классика интервью:

```go
type Queue struct{ in, out []int }

func (q *Queue) Enqueue(v int) { q.in = append(q.in, v) }

func (q *Queue) Dequeue() (int, bool) {
    if len(q.out) == 0 {
        for len(q.in) > 0 {
            n := len(q.in) - 1
            q.out = append(q.out, q.in[n])
            q.in = q.in[:n]
        }
    }
    if len(q.out) == 0 { return 0, false }
    n := len(q.out) - 1
    v := q.out[n]
    q.out = q.out[:n]
    return v, true
}
```

Амортизированно O(1) per operation.

## 4. BFS через queue

```go
func bfs(start *Node) {
    q := []*Node{start}
    visited := map[*Node]bool{start: true}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        // обработать cur
        for _, n := range cur.Neighbours {
            if !visited[n] {
                visited[n] = true
                q = append(q, n)
            }
        }
    }
}
```

## 5. Sliding window maximum

Используется monotonic deque (двусвязная очередь, в которой только убывающие элементы). Каждый элемент входит и выходит ровно раз → O(n).

В Go нет встроенной deque — пиши на основе двух slice или fixed-size buffer.

## 6. DFS итеративно через stack

```go
func dfsIterative(start *Node) {
    stack := []*Node{start}
    visited := map[*Node]bool{}
    for len(stack) > 0 {
        n := len(stack) - 1
        cur := stack[n]
        stack = stack[:n]
        if visited[cur] { continue }
        visited[cur] = true
        // обработать cur
        for _, n := range cur.Neighbours {
            if !visited[n] {
                stack = append(stack, n)
            }
        }
    }
}
```

## Правило

- Stack — `[]T` с push/pop в конце.
- Queue до сотен элементов — `[]T` с `q[1:]`.
- Большие очереди — ring buffer или `container/list`.
- BFS — queue, DFS — stack или recursion.

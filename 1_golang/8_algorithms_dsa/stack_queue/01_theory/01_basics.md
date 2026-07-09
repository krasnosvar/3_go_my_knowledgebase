# Stack и Queue в Go

В Go нет встроенных Stack/Queue типов. Они реализуются на слайсах в одну-две строки.

## Stack (LIFO)

```go
type Stack[T any] []T

func (s *Stack[T]) Push(v T) { *s = append(*s, v) }

func (s *Stack[T]) Pop() (T, bool) {
    if len(*s) == 0 {
        var z T
        return z, false
    }
    i := len(*s) - 1
    v := (*s)[i]
    *s = (*s)[:i]
    return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(*s) == 0 {
        var z T
        return z, false
    }
    return (*s)[len(*s)-1], true
}
```

Сложности: Push, Pop, Peek — все амортизированно O(1).

## Queue (FIFO)

Самая простая на slice:

```go
type Queue[T any] []T

func (q *Queue[T]) Enqueue(v T) { *q = append(*q, v) }

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(*q) == 0 {
        var z T
        return z, false
    }
    v := (*q)[0]
    *q = (*q)[1:]
    return v, true
}
```

**Проблема**: `q[1:]` не освобождает память. Базовый массив продолжает жить, длина растёт.

Решения:

- **Ring buffer** на фиксированном слайсе с двумя индексами.
- Периодический `copy` в новый массив, когда head/tail сильно расходятся.
- `container/list` (двусвязный) — гарантирует освобождение, но cache-locality хуже.

## Generic Queue с автоочисткой

```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(v T) { q.items = append(q.items, v) }

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var z T
        return z, false
    }
    v := q.items[0]
    q.items = q.items[1:]
    if len(q.items) == 0 {
        q.items = nil // освободить базовый массив
    }
    return v, true
}
```

## Когда используется

### Stack

- Парсинг (балансировка скобок, postfix-выражения).
- DFS-обход дерева/графа без рекурсии.
- Undo/redo в редакторах.
- Function call stack — внутри runtime, недоступно напрямую.

### Queue

- BFS-обход.
- Fairness в worker pool: первое пришло — первое обработано.
- Producer/consumer без приоритетов.
- Каналы Go — фактически встроенный thread-safe queue.

## Правило

В Go: для большинства задач — slice + методы как выше. Для сетевых очередей — каналы. `container/list` — только если действительно нужны гарантии по памяти и частые удаления в середине.

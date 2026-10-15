# Linked list: типовые задачи

Список основных алгоритмических операций над связным списком. Эти паттерны встречаются на собеседованиях и решают «классику».

## 1. Reverse

Развернуть односвязный список:

```go
func reverse(head *Node) *Node {
    var prev *Node
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}
```

Идея: три указателя двигаются вместе, каждый раз перенаправляя `Next`.

## 2. Cycle detection (Floyd's tortoise and hare)

```go
func hasCycle(head *Node) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

Если есть цикл — быстрый указатель догонит медленный. Без цикла — упрётся в nil.

## 3. Find middle

Тот же fast/slow, но без цикла:

```go
func middle(head *Node) *Node {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
```

Когда `fast` доходит до конца, `slow` — посередине.

## 4. Merge two sorted lists

```go
func merge(a, b *Node) *Node {
    dummy := &Node{}
    tail := dummy
    for a != nil && b != nil {
        if a.Val < b.Val {
            tail.Next = a
            a = a.Next
        } else {
            tail.Next = b
            b = b.Next
        }
        tail = tail.Next
    }
    if a != nil { tail.Next = a }
    if b != nil { tail.Next = b }
    return dummy.Next
}
```

Dummy-узел упрощает логику без специального case для пустого результата.

## 5. Remove N-th from end

Один проход с двумя указателями: `fast` уходит на N шагов вперёд, потом оба идут до конца — `slow` будет перед удаляемым.

## 6. Detect intersection of two lists

Конкатенируем "виртуально": проходим A потом B, и B потом A. Если есть пересечение — указатели встретятся в нём.

## Общие техники

- **Two pointers** (slow/fast) — циклы, середина, k-й с конца.
- **Dummy node** — упрощает edge cases в начале списка.
- **In-place reversal** — для перестановки/обращения частей.
- **Recursion** — естественна для tree-like операций, но осторожно со стеком.

## Правило

Используй pen-and-paper / комментарии в коде, чтобы рисовать стрелки. На doubly linked list большинство задач становится тривиальным.

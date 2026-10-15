# Linked List

Linked list — структура из узлов, каждый из которых хранит данные и ссылку на следующий (и/или предыдущий) узел.

## Виды

- **Singly linked**: узел знает только о следующем.
- **Doubly linked**: знает о next и prev. Удобно удалять элементы по указателю и итерировать в обе стороны.
- **Circular**: последний указывает на первый.

## Singly linked: своя реализация

```go
type Node struct {
    Val  int
    Next *Node
}

type List struct {
    head *Node
}

func (l *List) PushFront(v int) {
    l.head = &Node{Val: v, Next: l.head}
}

func (l *List) Print() {
    for n := l.head; n != nil; n = n.Next {
        fmt.Print(n.Val, " ")
    }
    fmt.Println()
}
```

## Стандартный container/list

`container/list` — двусвязный список с интерфейсом на `interface{}`:

```go
import "container/list"

l := list.New()
l.PushBack("a")
l.PushFront("z")
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
```

## Когда уместен linked list

- Частые вставки/удаления в середину при наличии указателя на узел.
- LRU-кеш (двусвязный + map для O(1) доступа).
- Очередь задач, где удаление в середине бывает.

## Когда НЕ уместен (почти всегда)

- Когда задача решается слайсом — slice **почти всегда быстрее** на современных CPU из-за cache locality.
- Random access по индексу — у списка O(n).
- Память: каждый узел = +16 байт (указатель), плюс выравнивание.

## Сложности

| Операция | Slice | Singly list | Doubly list |
|---|---|---|---|
| Push back | O(1) am | O(n) или O(1) с tail | O(1) |
| Push front | O(n) | O(1) | O(1) |
| Pop back | O(1) | O(n) | O(1) |
| Pop front | O(n) | O(1) | O(1) |
| Insert middle (с указателем) | O(n) | O(1) | O(1) |
| Random access | O(1) | O(n) | O(n) |

## Правило

В Go выбор по умолчанию — slice. Linked list — только когда профиль или задача явно требуют.

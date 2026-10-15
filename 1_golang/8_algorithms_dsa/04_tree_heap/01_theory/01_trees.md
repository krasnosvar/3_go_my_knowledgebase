# Trees

Дерево — иерархическая структура, узел имеет родителя и потомков.

## Базовая структура

```go
type Node struct {
    Val   int
    Left  *Node
    Right *Node
}
```

Двоичное дерево — у узла до 2 детей. Двоичное дерево поиска (BST) — left.Val < root.Val < right.Val.

## Обходы

### DFS

```go
// preorder: root, left, right
func preorder(n *Node) {
    if n == nil { return }
    visit(n)
    preorder(n.Left)
    preorder(n.Right)
}

// inorder: left, root, right (для BST даёт отсортированный обход)
func inorder(n *Node) {
    if n == nil { return }
    inorder(n.Left)
    visit(n)
    inorder(n.Right)
}

// postorder: left, right, root
func postorder(n *Node) {
    if n == nil { return }
    postorder(n.Left)
    postorder(n.Right)
    visit(n)
}
```

### BFS (level order)

```go
func levelOrder(root *Node) {
    if root == nil { return }
    q := []*Node{root}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        visit(cur)
        if cur.Left != nil { q = append(q, cur.Left) }
        if cur.Right != nil { q = append(q, cur.Right) }
    }
}
```

## Типовые задачи

- **Высота / глубина**: рекурсивно `1 + max(left, right)`.
- **Diameter**: для каждого узла длина пути через него = `height(left) + height(right)`.
- **Поиск в BST**: O(log n) для сбалансированного, O(n) в худшем случае.
- **Lowest common ancestor**: рекурсия + проверка вхождения.

## BST

```go
func search(root *Node, target int) *Node {
    for root != nil {
        if target == root.Val { return root }
        if target < root.Val { root = root.Left }
        else { root = root.Right }
    }
    return nil
}
```

## Сложности

| Операция | BST balanced | BST worst | Hash map |
|---|---|---|---|
| Insert | O(log n) | O(n) | O(1) am |
| Search | O(log n) | O(n) | O(1) am |
| Inorder | O(n) | O(n) | — |

В Go для key-value обычно `map`. BST нужен, когда важен порядок ключей или диапазонные запросы.

## Балансировка

Vanilla BST вырождается в список при отсортированных вставках. В реальной разработке — самобалансирующиеся:

- AVL — строгий баланс, быстрые поиск.
- Red-black — менее строгий, быстрее вставки/удаления.

В stdlib Go их нет. Стороннее: `github.com/google/btree`, `github.com/emirpasic/gods`.

## Правило

- Иерархия данных → tree.
- Кеширование/lookup → map обычно лучше.
- Диапазонные запросы или сортированный обход → BST/B-tree.

# 04. Tree и Heap (деревья и кучи)

Деревья — естественная модель иерархий и основа индексов БД (B-tree), а куча (heap) — стандартная реализация очереди с приоритетами: планировщики, top-K, слияние отсортированных потоков, алгоритм Дейкстры. В Go кучу дают через пакет `container/heap`, деревья пишут вручную на структурах с указателями.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#04-tree-и-heap) — раздел «04. Tree и heap» общей шпаргалки раздела 8.

**Пререквизиты:** [01_linked_list](../01_linked_list/) (работа с узлами и указателями), [02_stack_queue](../02_stack_queue/) (очередь для обхода по уровням), интерфейсы из [2_oop/02_interfaces](../../2_oop/02_interfaces/) (для `heap.Interface`), рекурсия из [07_functions](../../1_go_basics/07_functions/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_trees.md](./01_theory/01_trees.md) — двоичные деревья, BST, обходы (pre/in/post-order, по уровням), высота и сбалансированность.
    *   [02_heap.md](./01_theory/02_heap.md) — двоичная куча на массиве, sift up/down, `container/heap`, priority queue, top-K.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_binary_tree_inorder](./02_examples/01_binary_tree_inorder/main.go) — двоичное дерево и in-order обход.
    *   [02_min_heap_priority_queue](./02_examples/02_min_heap_priority_queue/main.go) — очередь с приоритетами на `container/heap`.
    *   [03_top_k_shape](./02_examples/03_top_k_shape/main.go) — шаблон задачи top-K через кучу.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу), подробнее — в [03_exercises/README.md](./03_exercises/README.md):
    *   `01_tree_height` — высота двоичного дерева.
    *   `02_bst_search` — поиск в двоичном дереве поиска.
    *   `03_inorder_traversal` — симметричный (in-order) обход.
    *   `04_heap_push_pop` — min-heap на срезе вручную.
    *   `05_k_smallest` — k наименьших элементов через кучу.
    *   `06_priority_queue` — очередь с приоритетами на container/heap.
    *   `07_heapify_array` — построение кучи из массива за O(n).
    *   `08_tree_level_order` — обход дерева по уровням (BFS).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_имя_answers`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/8_algorithms_dsa/04_tree_heap`); корневой `go.mod` репозитория подхватывается автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_trees.md](./01_theory/01_trees.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_binary_tree_inorder
   ```
3. Решайте задачи по порядку, пока проверка не напечатает `PASS`:
   ```bash
   go run ./03_exercises/01_tree_height
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_tree_height_answers
   ```

## ➡️ Что дальше

[05_graph](../05_graph/) — графы, BFS и DFS.

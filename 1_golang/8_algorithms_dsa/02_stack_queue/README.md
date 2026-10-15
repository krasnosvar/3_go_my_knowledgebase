# 02. Stack и Queue (стек и очередь)

Стек (LIFO) и очередь (FIFO) — основа множества алгоритмов: проверка скобок, обход графа в ширину, монотонные очереди для скользящего окна, планировщики задач. В Go обе структуры обычно пишутся на срезе в несколько строк — важно понимать стоимость операций и ловушки (утечка памяти при `q = q[1:]`).

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#02-stack-и-queue) — раздел «02. Stack и queue» общей шпаргалки раздела 8.

**Пререквизиты:** [09_slices](../../1_go_basics/09_slices/), методы из [2_oop/01_methods](../../2_oop/01_methods/); для generic-стека — [2_oop/05_generics](../../2_oop/05_generics/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — стек и очередь на срезе, generic-варианты, ring buffer, сложность операций.
    *   [02_problems.md](./01_theory/02_problems.md) — типовые задачи: скобки, min-stack, очередь на двух стеках, монотонная очередь, BFS.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_stack_and_queue_slice](./02_examples/01_stack_and_queue_slice/main.go) — стек и очередь на срезе.
    *   [02_stack_generic](./02_examples/02_stack_generic/main.go) — generic-стек `Stack[T]`.
    *   [03_queue_with_ring_buffer_shape](./02_examples/03_queue_with_ring_buffer_shape/main.go) — очередь на кольцевом буфере.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу), подробнее — в [03_exercises/README.md](./03_exercises/README.md):
    *   `01_push_pop` — стек на срезе (Push/Pop/Peek).
    *   `02_enqueue_dequeue` — очередь на срезе (Enqueue/Dequeue).
    *   `03_balanced_brackets` — проверка сбалансированности скобок.
    *   `04_min_stack` — стек с минимумом за O(1).
    *   `05_queue_two_stacks` — очередь на двух стеках.
    *   `06_sliding_window_max` — максимум в скользящем окне (монотонная очередь).
    *   `07_bfs_queue` — обход в ширину с очередью.
    *   `08_task_scheduler` — планировщик round-robin на очереди.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_имя_answers`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/8_algorithms_dsa/02_stack_queue`); корневой `go.mod` репозитория подхватывается автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_stack_and_queue_slice
   ```
3. Решайте задачи по порядку, пока проверка не напечатает `PASS`:
   ```bash
   go run ./03_exercises/01_push_pop
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_push_pop_answers
   ```

## ➡️ Что дальше

[03_set](../03_set/) — множества на `map`.

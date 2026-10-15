# 01. Linked List (связный список)

Связный список — первая «указательная» структура данных: узлы связаны полями `Next`, а не лежат подряд в памяти. В Go-коде он встречается редко (срез почти всегда быстрее), но на собеседованиях это обязательная классика: разворот, слияние, поиск цикла, два указателя. Здесь же отрабатывается аккуратная работа с указателями и `nil`.

**Пререквизиты:** [06_pointers](../../1_go_basics/06_pointers/), [08_structs](../../1_go_basics/08_structs/), [09_slices](../../1_go_basics/09_slices/), методы из [2_oop/01_methods](../../2_oop/01_methods/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — устройство односвязного и двусвязного списка, сложность операций, `container/list`, сравнение со срезом.
    *   [02_problems.md](./01_theory/02_problems.md) — типовые задачи: dummy-узел, два указателя (slow/fast), разворот, слияние.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_manual_singly_list](./02_examples/01_manual_singly_list/main.go) — ручная реализация односвязного списка.
    *   [02_container_list_basics](./02_examples/02_container_list_basics/main.go) — базовые операции `container/list`.
    *   [03_container_list_advanced_ops](./02_examples/03_container_list_advanced_ops/main.go) — вставки, перемещения и удаление элементов `container/list`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу), подробнее — в [03_exercises/README.md](./03_exercises/README.md):
    *   `01_append` — добавление в конец списка.
    *   `02_prepend` — добавление в начало списка.
    *   `03_find_node` — поиск узла по значению.
    *   `04_delete_node` — удаление узла по значению.
    *   `05_reverse_list` — разворот списка.
    *   `06_merge_sorted_lists` — слияние двух отсортированных списков.
    *   `07_cycle_detection` — обнаружение цикла (Floyd).
    *   `08_middle_node` — поиск середины списка.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_имя_answers`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/8_algorithms_dsa/01_linked_list`); корневой `go.mod` репозитория подхватывается автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_manual_singly_list
   ```
3. Решайте задачи по порядку, пока проверка не напечатает `PASS`:
   ```bash
   go run ./03_exercises/01_append
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_append_answers
   ```

## ➡️ Что дальше

[02_stack_queue](../02_stack_queue/) — стек и очередь, которые чаще всего строятся на срезе, а не на списке.

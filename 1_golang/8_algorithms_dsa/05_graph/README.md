# 05. Graph (графы)

Графы описывают связи: зависимости сборки и миграций, маршруты, социальные связи, сетевую топологию. Минимум, который нужно уверенно писать: список смежности, BFS (кратчайший путь в невзвешенном графе), DFS (пути, циклы, компоненты), топологическая сортировка.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#05-graph) — раздел «05. Graph» общей шпаргалки раздела 8.

**Пререквизиты:** [02_stack_queue](../02_stack_queue/) (очередь для BFS), [03_set](../03_set/) (visited), рекурсия и замыкания из [07_functions](../../1_go_basics/07_functions/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — виды графов, представления (список и матрица смежности), сложность.
    *   [02_algorithms.md](./01_theory/02_algorithms.md) — BFS, DFS, компоненты связности, поиск циклов, топологическая сортировка.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_adjacency_list_build](./02_examples/01_adjacency_list_build/main.go) — построение списка смежности.
    *   [02_bfs_traversal](./02_examples/02_bfs_traversal/main.go) — обход в ширину.
    *   [03_dfs_recursive](./02_examples/03_dfs_recursive/main.go) — рекурсивный обход в глубину.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу), подробнее — в [03_exercises/README.md](./03_exercises/README.md):
    *   `01_build_graph` — построение списка смежности.
    *   `02_bfs_path` — кратчайший путь BFS в невзвешенном графе.
    *   `03_dfs_path` — поиск пути обходом в глубину.
    *   `04_connected_components` — подсчёт компонент связности.
    *   `05_cycle_detect` — цикл в ориентированном графе (три цвета).
    *   `06_topological_sort` — топологическая сортировка (Кан).
    *   `07_shortest_path_unweighted` — расстояния от источника (BFS).
    *   `08_is_bipartite` — проверка двудольности (раскраска в 2 цвета).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_имя_answers`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/8_algorithms_dsa/05_graph`); корневой `go.mod` репозитория подхватывается автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_adjacency_list_build
   ```
3. Решайте задачи по порядку, пока проверка не напечатает `PASS`:
   ```bash
   go run ./03_exercises/01_build_graph
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_build_graph_answers
   ```

## ➡️ Что дальше

[06_leetcode_tasks](../06_leetcode_tasks/) — сводная практика по паттернам.

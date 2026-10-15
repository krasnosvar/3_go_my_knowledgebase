# 06. LeetCode tasks (паттерны задач с собеседований)

Сводный блок: классические задачи с собеседований, сгруппированные по паттернам (два указателя, хеш-таблица, скользящее окно, стек, куча, BFS). Задачи опираются на структуры из предыдущих тем, поэтому идут последними.

**Пререквизиты:** темы [01](../01_linked_list/)–[05](../05_graph/) этого раздела.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_patterns.md](./01_theory/01_patterns.md) — основные паттерны решения задач и как их распознавать.
    *   [02_complexity_and_tactics.md](./01_theory/02_complexity_and_tactics.md) — оценка сложности и тактика решения на интервью.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_two_sum_template](./02_examples/01_two_sum_template/main.go) — шаблон «хеш-таблица за один проход».
    *   [02_sliding_window_template](./02_examples/02_sliding_window_template/main.go) — шаблон скользящего окна.
    *   [03_bfs_template](./02_examples/03_bfs_template/main.go) — шаблон BFS.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу), подробнее — в [03_exercises/README.md](./03_exercises/README.md):
    *   `01_arrays_merge` — слияние отсортированных массивов in-place (LeetCode 88).
    *   `02_linked_list_reverse` — разворот списка: итеративно и рекурсивно (LeetCode 206).
    *   `03_valid_parentheses` — валидные скобки (LeetCode 20) и минимум удалений (LeetCode 1249).
    *   `04_two_sum` — Two Sum через хеш-таблицу (LeetCode 1).
    *   `05_longest_substring` — самая длинная подстрока без повторов (LeetCode 3).
    *   `06_min_window` — минимальное окно, содержащее все символы (LeetCode 76).
    *   `07_top_k_heap` — k самых частых элементов (LeetCode 347).
    *   `08_graph_bfs` — число островов на сетке (LeetCode 200).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_имя_answers`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/8_algorithms_dsa/06_leetcode_tasks`); корневой `go.mod` репозитория подхватывается автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_patterns.md](./01_theory/01_patterns.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_two_sum_template
   ```
3. Решайте задачи по порядку, пока проверка не напечатает `PASS`:
   ```bash
   go run ./03_exercises/01_arrays_merge
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_arrays_merge_answers
   ```

## ➡️ Что дальше

архивные подборки [07_legacy_golang_algorithm_tasks](../07_legacy_golang_algorithm_tasks/) и [08_legacy_non_basic_data_types](../08_legacy_non_basic_data_types/); затем — [2_db_kafka_redis](../../../2_db_kafka_redis/).

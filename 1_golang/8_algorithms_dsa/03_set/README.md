# 03. Set (множество)

Встроенного множества в Go нет — его делают на `map[T]struct{}`. Это самый частый «алгоритмический» инструмент в продакшен-коде: дедупликация, проверка принадлежности за O(1), пересечения и разности списков ID, а для маленьких наборов флагов — битовые маски.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#03-set) — раздел «03. Set» общей шпаргалки раздела 8.

**Пререквизиты:** [11_maps](../../1_go_basics/11_maps/), [04_operators_bitwise](../../1_go_basics/04_operators_bitwise/) (для битовых масок), [2_oop/05_generics](../../2_oop/05_generics/) (для `Set[T comparable]`).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — множество на `map[T]struct{}`, операции и их сложность, generic-реализация.
    *   [02_use_cases.md](./01_theory/02_use_cases.md) — где множества применяются на практике и типовые паттерны.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_set_basics](./02_examples/01_set_basics/main.go) — базовые операции с множеством на `map`.
    *   [02_set_generic](./02_examples/02_set_generic/main.go) — generic-множество `Set[T comparable]`.
    *   [03_set_intersection](./02_examples/03_set_intersection/main.go) — пересечение множеств.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу), подробнее — в [03_exercises/README.md](./03_exercises/README.md):
    *   `01_deduplicate` — удаление дубликатов с сохранением порядка.
    *   `02_contains_check` — множество строк: добавление и проверка.
    *   `03_union` — объединение множеств.
    *   `04_intersection` — пересечение множеств.
    *   `05_difference` — разность множеств.
    *   `06_generic_set` — generic-множество Set[T comparable].
    *   `07_count_distinct` — подсчёт уникальных и первый повтор.
    *   `08_permission_flags` — множество флагов на битовой маске.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_имя_answers`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/8_algorithms_dsa/03_set`); корневой `go.mod` репозитория подхватывается автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_set_basics
   ```
3. Решайте задачи по порядку, пока проверка не напечатает `PASS`:
   ```bash
   go run ./03_exercises/01_deduplicate
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_deduplicate_answers
   ```

## ➡️ Что дальше

[04_tree_heap](../04_tree_heap/) — деревья и кучи.

# Тема 05. Управляющие конструкции (Control Flow)

Этот раздел посвящен изучению управляющих конструкций в Go: ветвления `if/else`, циклов `for`, переключателей `switch`, меток (labels) и смены поведения циклов в версии Go 1.22.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — `if`, `switch`, формы `for` и `range`, метки, переменная цикла в Go 1.22, ловушки.

## 📋 Пререквизиты

*   Пройдена предыдущая тема: [04_operators_bitwise](../04_operators_bitwise/).

> ⏩ **Забегание вперёд.** Некоторые примеры и упражнения используют конструкции из следующих тем:
> *   `04_range` перебирает слайсы и мапы — подробно в [09_slices](../09_slices/) и [11_maps](../11_maps/).
> *   `08_table_dispatch` использует мапу функций — см. [07_functions](../07_functions/) и [11_maps](../11_maps/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_if_else.md](./01_theory/01_if_else.md) — ветвление `if/else`, `if` с короткой инициализацией, отсутствие тернарного оператора и ранний возврат вместо вложенности.
    *   [02_switch.md](./01_theory/02_switch.md) — `switch` по значению, `switch` без выражения, `switch` с инициализацией, `fallthrough` и **Type Switch (`switch x.(type)`)**.
    *   [03_for_and_range.md](./01_theory/03_for_and_range.md) — единственный цикл `for` (классический, аналог `while`, бесконечный), `for range` по слайсам/строкам/мапам/каналам, **поведение переменной цикла до и после Go 1.22 и производительность range по большим массивам (копирование значений)**.
    *   [04_break_continue_labels.md](./01_theory/04_break_continue_labels.md) — `break`, `continue`, метки (labels) для выхода из вложенных циклов и особенность `break` внутри `select`.
    *   [05_common_pitfalls.md](./01_theory/05_common_pitfalls.md) — частые ошибки: бесконечный цикл без выхода, путаница индекса и значения в range, модификация копии через переменную цикла, злоупотребление `if/else` там, где уместен `switch`.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические паттерны: table dispatch через `switch`, конечный автомат (state machine), ранний выход, поиск в данных и event-loop через `for + select`.
*   **[02_examples/](./02_examples/)** — практические примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — обработка команд CLI с использованием switch-case и ключевого слова `fallthrough`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — системный поиск ошибок в логах с использованием меток (`OuterLoop`) для выхода и пропуска шагов.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — **обработка различных способов оплаты с помощью Type Switch (`switch x.(type)`)**.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи для самопроверки с автопроверками `PASS`/`FAIL`:
    *   `01_if_else` — проверка возраста и статуса бана с использованием guard clauses (ранний возврат).
    *   `02_switch` — маппинг кодов HTTP-статусов на их названия.
    *   `03_for_loop` — классический цикл `for` для суммирования числового диапазона.
    *   `04_range` — итерирование по слайсу и мапе через `for range`.
    *   `05_break_continue` — фильтрация четных чисел и выход при превышении лимита суммы.
    *   `06_labels` — прерывание вложенного обхода матрицы при нахождении элемента.
    *   `07_infinite_loop` — бесконечный цикл `for {}` с выходом по sentinel-значению `-1`.
    *   `08_table_dispatch` — **написание Type Switch для разграничения типов интерфейса**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Прочитайте теорию в `01_theory/` по порядку, начиная с [01_if_else.md](./01_theory/01_if_else.md) и заканчивая [06_use_cases.md](./01_theory/06_use_cases.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run`:
   ```bash
   go run ./02_examples/03_product_backend/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/06_labels/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

## ➡️ Что дальше

Когда все упражнения темы печатают `PASS`, переходите к теме [06_pointers](../06_pointers/).

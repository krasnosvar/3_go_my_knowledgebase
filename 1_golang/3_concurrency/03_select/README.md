# 03. Оператор select

`select` ждёт сразу несколько операций с каналами и выполняет ту, что готова первой. На нём строятся таймауты, неблокирующие операции, отмена и мультиплексирование.

**Зачем:** почти любой долгоживущий цикл в Go-сервисе — это `for { select { ... } }` с веткой отмены.

**Пререквизиты:** [02_channels](../02_channels/).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — семантика `select`, типовые формы, выбор таймаута, fan-in через nil-канал, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — синтаксис, случайный выбор среди готовых веток.
    *   [02_default.md](./01_theory/02_default.md) — `default` и неблокирующие операции.
    *   [03_timeout.md](./01_theory/03_timeout.md) — таймауты через `time.After`.
    *   [04_cancel.md](./01_theory/04_cancel.md) — отмена через `context` в `select`.
    *   [05_pitfalls.md](./01_theory/05_pitfalls.md) — частые ошибки: утечки таймеров, пустой `select`, nil-каналы.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — ожидание сигнала завершения.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — health-check с таймаутом.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — вычитывание двух каналов с приёмом «nil-канал».
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_basic_select](./03_exercises/01_basic_select/main.go) — первое значение из двух каналов.
    *   [02_default_non_blocking](./03_exercises/02_default_non_blocking/main.go) — `trySend`/`tryRecv` через `default`.
    *   [03_timeout](./03_exercises/03_timeout/main.go) — ожидание результата с ошибкой по таймауту.
    *   [04_two_channels](./03_exercises/04_two_channels/main.go) — вычитывание двух каналов до закрытия (nil-приём).
    *   [05_cancel_signal](./03_exercises/05_cancel_signal/main.go) — цикл обработки с выходом по `done`.
    *   [06_multiplex](./03_exercises/06_multiplex/main.go) — слияние N источников в один поток.
    *   [07_priority](./03_exercises/07_priority/main.go) — приоритет срочного канала двумя `select`.
    *   [08_retry_loop](./03_exercises/08_retry_loop/main.go) — повторы с таймаутом на каждую попытку.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/03_select`).

1. Прочитайте теорию по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_basic_select/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_basic_select_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [04_sync_package](../04_sync_package/) — примитивы синхронизации для общей памяти.

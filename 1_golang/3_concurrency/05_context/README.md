# 05. Пакет context

`context.Context` управляет временем жизни операции: отмена (`WithCancel`), таймаут (`WithTimeout`), дедлайн (`WithDeadline`) и передача служебных значений запроса (`WithValue`). Отмена распространяется от родителя ко всем потомкам.

**Зачем:** каждый входящий HTTP/gRPC-запрос несёт контекст; без него невозможно корректно прерывать запросы к БД и внешним сервисам.

**Пререквизиты:** [03_select](../03_select/), [04_sync_package](../04_sync_package/).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — API и версии, правила, таймауты и дедлайны, `Canceled` vs `DeadlineExceeded`, значения, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — интерфейс `Context`, `Background`/`TODO`.
    *   [02_cancel.md](./01_theory/02_cancel.md) — `context.WithCancel`.
    *   [03_timeout_deadline.md](./01_theory/03_timeout_deadline.md) — `WithTimeout` и `WithDeadline`.
    *   [04_values.md](./01_theory/04_values.md) — `context.WithValue` и типизированные ключи.
    *   [05_common_pitfalls.md](./01_theory/05_common_pitfalls.md) — частые ошибки: забытый `cancel`, контекст в структуре, злоупотребление `Value`.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — операция, прерываемая по таймауту.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — worker'ы, останавливаемые отменой контекста.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — request-id через `WithValue`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_cancel](./03_exercises/01_cancel/main.go) — `WithCancel` и ожидание отмены.
    *   [02_timeout](./03_exercises/02_timeout/main.go) — операция с `WithTimeout`.
    *   [03_deadline](./03_exercises/03_deadline/main.go) — абсолютный дедлайн через `WithDeadline`.
    *   [04_propagation](./03_exercises/04_propagation/main.go) — отмена родителя отменяет потомка, но не наоборот.
    *   [05_value](./03_exercises/05_value/main.go) — request-id с типизированным ключом.
    *   [06_worker_shutdown](./03_exercises/06_worker_shutdown/main.go) — worker останавливается по `ctx.Done()`.
    *   [07_http_request](./03_exercises/07_http_request/main.go) — HTTP-запрос с контекстом (`httptest`-сервер).
    *   [08_pipeline_cancel](./03_exercises/08_pipeline_cancel/main.go) — конвейер, завершающийся по отмене.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/05_context`).

1. Прочитайте теорию по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_cancel/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_cancel_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [06_atomic](../06_atomic/) — lock-free операции над простыми значениями.

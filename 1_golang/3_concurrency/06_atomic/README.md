# 06. Пакет sync/atomic

Атомарные операции (`Add`, `Load`, `Store`, `CompareAndSwap`) позволяют безопасно менять простые общие значения без мьютекса. Здесь же — когда atomic НЕ подходит и как устроена модель памяти Go.

**Зачем:** счётчики метрик, флаги состояния и «одноразовые» переключатели на atomic быстрее и проще мьютекса — если понимать их ограничения.

**Пререквизиты:** [04_sync_package](../04_sync_package/) (сравнение с `Mutex`).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — что такое атомарность, типы `atomic.Int64` и др.
    *   [02_add_load_store.md](./01_theory/02_add_load_store.md) — `Add` / `Load` / `Store`.
    *   [03_compare_swap.md](./01_theory/03_compare_swap.md) — Compare-And-Swap и CAS-циклы.
    *   [04_when_not_use.md](./01_theory/04_when_not_use.md) — когда atomic не подходит.
    *   [05_memory_model.md](./01_theory/05_memory_model.md) — модель памяти Go кратко (happens-before).
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — `atomic.AddInt64`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — `StoreInt32` / `LoadInt32` для флага здоровья.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — `CompareAndSwapInt32` для смены состояния.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_add](./03_exercises/01_add/main.go) — атомарный инкремент из многих горутин.
    *   [02_load_store](./03_exercises/02_load_store/main.go) — состояние сервиса на `atomic.Int32`.
    *   [03_compare_swap](./03_exercises/03_compare_swap/main.go) — CAS-цикл: атомарный максимум.
    *   [04_bool_flag](./03_exercises/04_bool_flag/main.go) — одноразовый флаг закрытия на `atomic.Bool`.
    *   [05_spin_wait](./03_exercises/05_spin_wait/main.go) — ожидание флага с `runtime.Gosched` и таймаутом.
    *   [06_counter](./03_exercises/06_counter/main.go) — статистика запросов на атомиках.
    *   [07_once_like](./03_exercises/07_once_like/main.go) — однократный запуск на CAS и отличие от `sync.Once`.
    *   [08_hybrid](./03_exercises/08_hybrid/main.go) — map под мьютексом + метрики на атомиках.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/06_atomic`).

1. Прочитайте теорию по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_add/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_add_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [07_patterns](../07_patterns/) — сборка изученных примитивов в типовые паттерны.

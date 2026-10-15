# 04. Пакет sync

`sync` даёт примитивы для работы с общей памятью: `WaitGroup`, `Mutex`, `RWMutex`, `Once`, `Cond`, `Pool`. Каналы удобны для передачи данных, а `sync` — для защиты общего состояния.

**Зачем:** кэши, счётчики, ленивая инициализация, пулы буферов — всё это `sync`, и ошибки в нём ведут к гонкам и deadlock'ам.

**Пререквизиты:** [01_goroutines](../01_goroutines/), [02_channels](../02_channels/), [03_select](../03_select/).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — что выбрать, `WaitGroup`, `Mutex`/`RWMutex`, `Once`/`OnceValue`, `Cond`, `Pool`, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_waitgroup.md](./01_theory/01_waitgroup.md) — `sync.WaitGroup`.
    *   [02_mutex.md](./01_theory/02_mutex.md) — `sync.Mutex`.
    *   [03_rwmutex.md](./01_theory/03_rwmutex.md) — `sync.RWMutex`.
    *   [04_once.md](./01_theory/04_once.md) — `sync.Once`.
    *   [05_cond_pool.md](./01_theory/05_cond_pool.md) — `sync.Cond` и `sync.Pool`.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — параллельный запуск команд с `WaitGroup`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — общая map статусов под `Mutex`.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — кэш на `RWMutex`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_waitgroup](./03_exercises/01_waitgroup/main.go) — параллельная сумма частей.
    *   [02_mutex](./03_exercises/02_mutex/main.go) — счётчик под `Mutex`.
    *   [03_rwmutex](./03_exercises/03_rwmutex/main.go) — конфиг: много читателей, один писатель.
    *   [04_once](./03_exercises/04_once/main.go) — ленивая инициализация ровно один раз.
    *   [05_cond](./03_exercises/05_cond/main.go) — ожидание публикации через `Cond.Broadcast`.
    *   [06_map_with_mutex](./03_exercises/06_map_with_mutex/main.go) — конкурентный подсчёт слов в map.
    *   [07_pool](./03_exercises/07_pool/main.go) — переиспользование `bytes.Buffer` через `Pool`.
    *   [08_barrier](./03_exercises/08_barrier/main.go) — барьер между двумя фазами работы.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/04_sync_package`).

1. Прочитайте теорию по порядку, начиная с [01_waitgroup.md](./01_theory/01_waitgroup.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_waitgroup/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_waitgroup_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [05_context](../05_context/) — отмена, таймауты и дедлайны.

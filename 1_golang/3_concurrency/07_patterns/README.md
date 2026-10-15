# 07. Паттерны конкурентности

Типовые конструкции из горутин и каналов: worker pool, fan-out/fan-in, pipeline, семафор, rate limit, ограниченная очередь и graceful shutdown.

**Зачем:** это готовые «кирпичи» для backend-сервисов: обработка очередей, параллельные запросы, ограничение нагрузки и корректная остановка при деплое.

**Пререквизиты:** темы [01](../01_goroutines/)–[06](../06_atomic/) этого раздела.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — задача → паттерн, worker pool, fan-out/fan-in, pipeline, semaphore, graceful shutdown, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_worker_pool.md](./01_theory/01_worker_pool.md) — worker pool.
    *   [02_fan_out_fan_in.md](./01_theory/02_fan_out_fan_in.md) — fan-out / fan-in.
    *   [03_pipeline.md](./01_theory/03_pipeline.md) — pipeline.
    *   [04_semaphore.md](./01_theory/04_semaphore.md) — семафор на буферизованном канале.
    *   [05_shutdown.md](./01_theory/05_shutdown.md) — graceful shutdown.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — простой worker над очередью задач.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — семафор, ограничивающий параллелизм.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — двухстадийный pipeline.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_worker_pool](./03_exercises/01_worker_pool/main.go) — пул фиксированного размера с сохранением порядка результатов.
    *   [02_fan_out](./03_exercises/02_fan_out/main.go) — N worker'ов читают один вход, у каждого свой выход.
    *   [03_fan_in](./03_exercises/03_fan_in/main.go) — слияние каналов с остановкой по `done`.
    *   [04_pipeline](./03_exercises/04_pipeline/main.go) — конвейер parse → square → sum.
    *   [05_semaphore](./03_exercises/05_semaphore/main.go) — ограничение параллелизма буферизованным каналом.
    *   [06_rate_limit](./03_exercises/06_rate_limit/main.go) — ограничение частоты через `time.Ticker`.
    *   [07_bounded_queue](./03_exercises/07_bounded_queue/main.go) — очередь фиксированной ёмкости с отказом при переполнении.
    *   [08_graceful_shutdown](./03_exercises/08_graceful_shutdown/main.go) — пул worker'ов с `Shutdown(ctx)`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/07_patterns`).

1. Прочитайте теорию по порядку, начиная с [01_worker_pool.md](./01_theory/01_worker_pool.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_worker_pool/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_worker_pool_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [08_race_conditions](../08_race_conditions/) — как всё это ломается и как это чинить.

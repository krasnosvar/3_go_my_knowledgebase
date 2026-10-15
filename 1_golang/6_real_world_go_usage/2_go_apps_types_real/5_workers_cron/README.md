# Workers and cron

Фоновая обработка: очереди, пул воркеров, backoff, DLQ, идемпотентность, graceful stop, защита cron от наложения, метрики.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#5-workers-and-cron) — раздел «5. Workers and cron» общей шпаргалки типов Go-приложений.

## Пререквизиты

[3_concurrency](../../../3_concurrency/) (каналы, `select`, `sync`, `atomic`, `context`).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — Workers and Cron.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_worker_pool](./02_examples/01_worker_pool/main.go) — пул воркеров на каналах.
    *   [02_retry_loop](./02_examples/02_retry_loop/main.go) — цикл повторов.
    *   [03_cron_tick](./02_examples/03_cron_tick/main.go) — периодический запуск через `time.Ticker`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_job_queue](./03_exercises/01_job_queue/main.go) — ограниченная очередь задач.
    *   [02_worker_pool](./03_exercises/02_worker_pool/main.go) — пул воркеров с сохранением порядка результатов.
    *   [03_retry_backoff](./03_exercises/03_retry_backoff/main.go) — экспоненциальный backoff с потолком.
    *   [04_dlq_move](./03_exercises/04_dlq_move/main.go) — dead letter queue.
    *   [05_idempotent_key](./03_exercises/05_idempotent_key/main.go) — идемпотентность по ключу.
    *   [06_graceful_stop](./03_exercises/06_graceful_stop/main.go) — graceful stop воркера.
    *   [07_cron_guard](./03_exercises/07_cron_guard/main.go) — защита cron-задачи от наложения запусков.
    *   [08_metrics_counter](./03_exercises/08_metrics_counter/main.go) — метрики воркера на atomic.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/5_workers_cron` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/5_workers_cron/02_examples/01_worker_pool
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/5_workers_cron/03_exercises/01_job_queue
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/5_workers_cron/04_exercises_answers/01_job_queue_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [6_k8s_operators](../6_k8s_operators/) — Kubernetes operators.

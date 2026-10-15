# Упражнения: Workers and cron

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_job_queue
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_job_queue_answers
```

## Список задач (от простого к сложному)

1. [`01_job_queue`](./01_job_queue/main.go) — ограниченная очередь задач; решение: [`01_job_queue_answers`](../04_exercises_answers/01_job_queue_answers/main.go).
2. [`02_worker_pool`](./02_worker_pool/main.go) — пул воркеров с сохранением порядка результатов; решение: [`02_worker_pool_answers`](../04_exercises_answers/02_worker_pool_answers/main.go).
3. [`03_retry_backoff`](./03_retry_backoff/main.go) — экспоненциальный backoff с потолком; решение: [`03_retry_backoff_answers`](../04_exercises_answers/03_retry_backoff_answers/main.go).
4. [`04_dlq_move`](./04_dlq_move/main.go) — dead letter queue; решение: [`04_dlq_move_answers`](../04_exercises_answers/04_dlq_move_answers/main.go).
5. [`05_idempotent_key`](./05_idempotent_key/main.go) — идемпотентность по ключу; решение: [`05_idempotent_key_answers`](../04_exercises_answers/05_idempotent_key_answers/main.go).
6. [`06_graceful_stop`](./06_graceful_stop/main.go) — graceful stop воркера; решение: [`06_graceful_stop_answers`](../04_exercises_answers/06_graceful_stop_answers/main.go).
7. [`07_cron_guard`](./07_cron_guard/main.go) — защита cron-задачи от наложения запусков; решение: [`07_cron_guard_answers`](../04_exercises_answers/07_cron_guard_answers/main.go).
8. [`08_metrics_counter`](./08_metrics_counter/main.go) — метрики воркера на atomic; решение: [`08_metrics_counter_answers`](../04_exercises_answers/08_metrics_counter_answers/main.go).

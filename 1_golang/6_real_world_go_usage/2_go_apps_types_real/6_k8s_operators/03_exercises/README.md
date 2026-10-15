# Упражнения: Kubernetes operators

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_reconcile_idempotent
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_reconcile_idempotent_answers
```

## Список задач (от простого к сложному)

1. [`01_reconcile_idempotent`](./01_reconcile_idempotent/main.go) — идемпотентный reconcile; решение: [`01_reconcile_idempotent_answers`](../04_exercises_answers/01_reconcile_idempotent_answers/main.go).
2. [`02_status_conditions`](./02_status_conditions/main.go) — status conditions; решение: [`02_status_conditions_answers`](../04_exercises_answers/02_status_conditions_answers/main.go).
3. [`03_finalizer_cleanup`](./03_finalizer_cleanup/main.go) — finalizer и очистка внешних ресурсов; решение: [`03_finalizer_cleanup_answers`](../04_exercises_answers/03_finalizer_cleanup_answers/main.go).
4. [`04_spec_validation`](./04_spec_validation/main.go) — валидация spec кастомного ресурса; решение: [`04_spec_validation_answers`](../04_exercises_answers/04_spec_validation_answers/main.go).
5. [`05_event_recording`](./05_event_recording/main.go) — запись событий с дедупликацией; решение: [`05_event_recording_answers`](../04_exercises_answers/05_event_recording_answers/main.go).
6. [`06_error_retry`](./06_error_retry/main.go) — requeue с per-item backoff; решение: [`06_error_retry_answers`](../04_exercises_answers/06_error_retry_answers/main.go).
7. [`07_owner_refs`](./07_owner_refs/main.go) — ownerReferences и каскадное удаление; решение: [`07_owner_refs_answers`](../04_exercises_answers/07_owner_refs_answers/main.go).
8. [`08_metrics_export`](./08_metrics_export/main.go) — метрики контроллера; решение: [`08_metrics_export_answers`](../04_exercises_answers/08_metrics_export_answers/main.go).

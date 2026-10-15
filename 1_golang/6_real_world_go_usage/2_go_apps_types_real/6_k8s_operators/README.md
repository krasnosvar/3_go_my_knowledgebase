# Kubernetes operators

Логика оператора на простых структурах вместо client-go: идемпотентный reconcile, conditions, finalizers, валидация spec, события, requeue, ownerReferences, метрики.

## Пререквизиты

[5_workers_cron](../5_workers_cron/); базовое знакомство с Kubernetes.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — Kubernetes Operators.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_reconcile_shape](./02_examples/01_reconcile_shape/main.go) — форма reconcile-цикла.
    *   [02_finalizer_flow](./02_examples/02_finalizer_flow/main.go) — логика финализатора.
    *   [03_status_condition](./02_examples/03_status_condition/main.go) — status condition ресурса.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_reconcile_idempotent](./03_exercises/01_reconcile_idempotent/main.go) — идемпотентный reconcile.
    *   [02_status_conditions](./03_exercises/02_status_conditions/main.go) — status conditions.
    *   [03_finalizer_cleanup](./03_exercises/03_finalizer_cleanup/main.go) — finalizer и очистка внешних ресурсов.
    *   [04_spec_validation](./03_exercises/04_spec_validation/main.go) — валидация spec кастомного ресурса.
    *   [05_event_recording](./03_exercises/05_event_recording/main.go) — запись событий с дедупликацией.
    *   [06_error_retry](./03_exercises/06_error_retry/main.go) — requeue с per-item backoff.
    *   [07_owner_refs](./03_exercises/07_owner_refs/main.go) — ownerReferences и каскадное удаление.
    *   [08_metrics_export](./03_exercises/08_metrics_export/main.go) — метрики контроллера.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/6_k8s_operators` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/6_k8s_operators/02_examples/01_reconcile_shape
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/6_k8s_operators/03_exercises/01_reconcile_idempotent
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/6_k8s_operators/04_exercises_answers/01_reconcile_idempotent_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [7_streaming_data](../7_streaming_data/) — Streaming data.

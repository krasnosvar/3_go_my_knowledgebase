# Упражнения: messaging

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`.

Запуск шаблона:

```bash
go run ./03_exercises/01_idempotent_handler/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_idempotent_handler/main.go
```

## Список задач (от простого к сложному)

1. `01_idempotent_handler` — идемпотентный обработчик: дедупликация по id, эффект применяется ровно один раз (карта `seen`).
2. `02_classify_guarantee` — классифицировать сценарий на at-most-once / at-least-once / exactly-once по порядку ack и наличию дедупликации.
3. `03_partition_by_key` — распределить сообщения по партициям `hash(key)%P` (FNV) с сохранением порядка внутри ключа.
4. `04_outbox_drain` — слив transactional outbox в брокер: повторы при потере ack, at-least-once без потери событий (хардкор).

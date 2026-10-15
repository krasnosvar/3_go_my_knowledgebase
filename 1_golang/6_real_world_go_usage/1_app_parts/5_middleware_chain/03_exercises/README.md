# Упражнения: Middleware chain

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_chain_order
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_chain_order_answers
```

## Список задач (от простого к сложному)

1. [`01_chain_order`](./01_chain_order/main.go) — сборка цепочки middleware; решение: [`01_chain_order_answers`](../04_exercises_answers/01_chain_order_answers/main.go).
2. [`02_recovery`](./02_recovery/main.go) — recovery middleware; решение: [`02_recovery_answers`](../04_exercises_answers/02_recovery_answers/main.go).
3. [`03_request_id`](./03_request_id/main.go) — request ID через заголовок и context; решение: [`03_request_id_answers`](../04_exercises_answers/03_request_id_answers/main.go).
4. [`04_status_recorder`](./04_status_recorder/main.go) — middleware, который видит статус ответа; решение: [`04_status_recorder_answers`](../04_exercises_answers/04_status_recorder_answers/main.go).

# Упражнения: Streaming data

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_decode_message
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_decode_message_answers
```

## Список задач (от простого к сложному)

1. [`01_decode_message`](./01_decode_message/main.go) — строгий разбор сообщения из топика; решение: [`01_decode_message_answers`](../04_exercises_answers/01_decode_message_answers/main.go).
2. [`02_batch_window`](./02_batch_window/main.go) — батчинг по размеру и по времени; решение: [`02_batch_window_answers`](../04_exercises_answers/02_batch_window_answers/main.go).
3. [`03_offset_commit`](./03_offset_commit/main.go) — коммит оффсетов при параллельной обработке; решение: [`03_offset_commit_answers`](../04_exercises_answers/03_offset_commit_answers/main.go).
4. [`04_retry_on_fail`](./04_retry_on_fail/main.go) — at-least-once — коммит только после успеха; решение: [`04_retry_on_fail_answers`](../04_exercises_answers/04_retry_on_fail_answers/main.go).
5. [`05_poison_message`](./05_poison_message/main.go) — poison message → DLQ, поток не стоит; решение: [`05_poison_message_answers`](../04_exercises_answers/05_poison_message_answers/main.go).
6. [`06_partition_key`](./06_partition_key/main.go) — выбор партиции по ключу; решение: [`06_partition_key_answers`](../04_exercises_answers/06_partition_key_answers/main.go).
7. [`07_consumer_rebalance`](./07_consumer_rebalance/main.go) — распределение партиций в consumer group (range assignor); решение: [`07_consumer_rebalance_answers`](../04_exercises_answers/07_consumer_rebalance_answers/main.go).
8. [`08_lag_alert`](./08_lag_alert/main.go) — consumer lag и алерт; решение: [`08_lag_alert_answers`](../04_exercises_answers/08_lag_alert_answers/main.go).

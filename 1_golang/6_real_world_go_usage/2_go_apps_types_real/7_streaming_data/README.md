# Streaming data

Консьюмеры потоков (Kafka/NATS) на модели: строгий разбор, батчинг, коммит оффсетов, at-least-once, poison messages, партиционирование, ребаланс, lag.

## Пререквизиты

[5_workers_cron](../5_workers_cron/); теория Kafka — [2_db_kafka_redis/2_kafka_rabbit_nats_redis/1_kafka](../../../../2_db_kafka_redis/2_kafka_rabbit_nats_redis/1_kafka/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — Streaming data.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_consumer_loop](./02_examples/01_consumer_loop/main.go) — цикл консьюмера.
    *   [02_batch_processing](./02_examples/02_batch_processing/main.go) — обработка батчами.
    *   [03_offset_commit](./02_examples/03_offset_commit/main.go) — коммит оффсета после обработки.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_decode_message](./03_exercises/01_decode_message/main.go) — строгий разбор сообщения из топика.
    *   [02_batch_window](./03_exercises/02_batch_window/main.go) — батчинг по размеру и по времени.
    *   [03_offset_commit](./03_exercises/03_offset_commit/main.go) — коммит оффсетов при параллельной обработке.
    *   [04_retry_on_fail](./03_exercises/04_retry_on_fail/main.go) — at-least-once — коммит только после успеха.
    *   [05_poison_message](./03_exercises/05_poison_message/main.go) — poison message → DLQ, поток не стоит.
    *   [06_partition_key](./03_exercises/06_partition_key/main.go) — выбор партиции по ключу.
    *   [07_consumer_rebalance](./03_exercises/07_consumer_rebalance/main.go) — распределение партиций в consumer group (range assignor).
    *   [08_lag_alert](./03_exercises/08_lag_alert/main.go) — consumer lag и алерт.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/7_streaming_data` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/7_streaming_data/02_examples/01_consumer_loop
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/7_streaming_data/03_exercises/01_decode_message
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/7_streaming_data/04_exercises_answers/01_decode_message_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [7_under_the_hood](../../../7_under_the_hood/) — как Go работает внутри.

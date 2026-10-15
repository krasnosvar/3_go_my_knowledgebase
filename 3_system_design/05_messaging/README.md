# 05. Messaging (Обмен сообщениями)

Асинхронная связь между сервисами через очереди и брокеры: как развязать сервисы во времени, сгладить нагрузку, гарантировать доставку и не терять сообщения. Разбираем модели брокеров (очередь vs лог), Kafka/RabbitMQ/NATS, гарантии доставки, consumer groups и партиционирование, паттерны outbox/inbox и идемпотентность, а также backpressure. Теория подкреплена runnable-примерами на стандартной библиотеке Go и упражнениями с самопроверкой.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_queues_brokers.md](./01_theory/01_queues_brokers.md) — зачем очереди и брокеры, синхронное vs асинхронное взаимодействие, ключевая разница «очередь (queue) vs журнал (log)», топики/партиции/оффсеты/ключи.
    *   [02_kafka_rabbit_nats.md](./01_theory/02_kafka_rabbit_nats.md) — Kafka vs RabbitMQ vs NATS: модели (лог / очередь с маршрутизацией / лёгкий pub-sub), сильные стороны, когда что выбирать.
    *   [03_delivery_guarantees.md](./01_theory/03_delivery_guarantees.md) — гарантии доставки at-most-once / at-least-once / exactly-once, ack/nack, retry и DLQ, дедупликация по message-id.
    *   [04_consumer_groups.md](./01_theory/04_consumer_groups.md) — consumer groups, партиционирование и параллелизм, ребаланс (stop-the-world, cooperative), порядок и коммит оффсетов.
    *   [05_outbox_inbox.md](./01_theory/05_outbox_inbox.md) — проблема двойной записи, transactional outbox, inbox/дедупликация, идемпотентные консьюмеры, effectively-once.
    *   [06_backpressure.md](./01_theory/06_backpressure.md) — backpressure, ограниченные буферы, flow control (каналы Go, prefetch RabbitMQ, pull/consumer lag Kafka), push vs pull.
    *   [07_use_cases.md](./01_theory/07_use_cases.md) — практические применения: когда брать брокер, выбор под задачу, типовые паттерны (work queue, pub/sub, outbox, saga, CQRS, DLQ), correctness-чеклист и диагностика.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib: каналы/sync; детерминированы, синхронизация через WaitGroup/close(chan), без sleep-«на удачу»):
    *   [01_inmem_queue](./02_examples/01_inmem_queue/main.go) — in-memory очередь на буферизованном канале: продюсер шлёт N сообщений, консьюмер обрабатывает; завершение через `close(chan)` + `WaitGroup`.
    *   [02_at_least_once_idempotent](./02_examples/02_at_least_once_idempotent/main.go) — доставка at-least-once с дублями + идемпотентный консьюмер, дедуплицирующий по message-id (карта `seen`); эффект применяется один раз.
    *   [03_backpressure](./02_examples/03_backpressure/main.go) — маленький буфер + медленный консьюмер: продюсер блокируется при полном буфере (backpressure), число блокировок считается детерминированно.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_idempotent_handler` — идемпотентный обработчик: дедуп по id, эффект ровно один раз.
    *   `02_classify_guarantee` — классифицировать сценарий на at-most-once / at-least-once / exactly-once.
    *   `03_partition_by_key` — распределить сообщения по партициям `hash(key)%P` с сохранением порядка внутри ключа.
    *   `04_outbox_drain` — **слив transactional outbox в брокер: повторы при потере ack, at-least-once без потери событий (Хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_queues_brokers.md](./01_theory/01_queues_brokers.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, только stdlib; работает корневой `go.mod` репозитория):
   ```bash
   go run ./02_examples/01_inmem_queue/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/01_idempotent_handler/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

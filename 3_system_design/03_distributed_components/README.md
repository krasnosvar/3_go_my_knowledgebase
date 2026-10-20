# Тема 03. Распределённые компоненты в Go (Distributed Components)

## О чём эта тема (краткий контекст)

Когда Go-сервис работает в распределённом окружении с реплицируемыми и
шардированными базами данных, брокерами сообщений и ненадёжными внешними API,
на первый план выходят управление пулами соединений, гарантии доставки событий
и защита от каскадных сбоев. В этом модуле разбираются устройство пула
`database/sql`, Consistent Hashing, паттерн Transactional Outbox, естественный
backpressure на каналах Go, Rate Limiting, Circuit Breaker, Retries с jitter и
безопасный Graceful Shutdown.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [01. Репликация, шардирование, Consistent Hashing, CAP/PACELC и транзакции](../../../5_devops_sre_knowledgebase/14_system_design/03_distributed_components_and_data/01_theory/01_data_replication_sharding_and_cap.md)
> - [02. Брокеры сообщений (Kafka, RabbitMQ, NATS), Outbox, идемпотентность и Backpressure](../../../5_devops_sre_knowledgebase/14_system_design/03_distributed_components_and_data/01_theory/02_messaging_and_async_patterns.md)
> - [03. Балансировка нагрузки, Health Checks, Rate Limiting, Circuit Breaker, Retries и Graceful Shutdown](../../../5_devops_sre_knowledgebase/14_system_design/03_distributed_components_and_data/01_theory/03_load_balancing_and_resilience_patterns.md)
> - Практические стенды с PostgreSQL, Kafka и Redis на Go:
>   [2_db_kafka_redis](../../2_db_kafka_redis/)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — тюнинг `*sql.DB`,
Consistent Hashing, Outbox, Worker Pool backpressure, `errgroup`, Token Bucket,
Circuit Breaker, Retries с jitter и порядок Graceful Shutdown.

## Пререквизиты

- [3_concurrency](../../1_golang/3_concurrency/) (каналы, `select`, `sync`,
  `context`, `atomic`).
- [01_architecture_styles](../01_architecture_styles/) и
  [02_performance_and_scaling](../02_performance_and_scaling/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория Go-специфики:
    *   [01_db_pools_and_sharding_in_go.md](./01_theory/01_db_pools_and_sharding_in_go.md)
        — внутреннее устройство пула `database/sql` (`MaxOpenConns`,
        `MaxIdleConns`), утечки соединений при незакрытом `sql.Rows`, роутинг
        на read-реплики и кольцо Consistent Hashing в Go.
    *   [02_messaging_and_backpressure_in_go.md](./01_theory/02_messaging_and_backpressure_in_go.md)
        — Transactional Outbox и идемпотентный консьюмер в Go, естественный
        backpressure и Load Shedding на буферизованных каналах, координация
        горутин через `errgroup.WithContext`.
    *   [03_resilience_and_graceful_shutdown_in_go.md](./01_theory/03_resilience_and_graceful_shutdown_in_go.md)
        — Rate Limiting (`x/time/rate` Token Bucket), конечный автомат Circuit
        Breaker, Retries с Exponential Backoff & Jitter с учётом `ctx.Done()`,
        эталонный порядок Graceful Shutdown в Kubernetes.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_consistent_hashing](./02_examples/01_consistent_hashing/main.go) —
        кольцо Consistent Hashing с виртуальными узлами (`crc32` +
        `sort.Search`).
    *   [02_shard_router](./02_examples/02_shard_router/main.go) —
        маршрутизация чтения и записи по $N$ шардам и проблема scatter-gather.
    *   [03_quorum_rw](./02_examples/03_quorum_rw/main.go) — кворумная запись и
        чтение ($R + W > N$) по репликам с версионированием.
    *   [04_inmem_queue](./02_examples/04_inmem_queue/main.go) — очередь задач
        и конкурирующие воркеры (competing consumers) на каналах.
    *   [05_at_least_once_idempotent](./02_examples/05_at_least_once_idempotent/main.go)
        — доставка `at-least-once` с ретраями и дедупликацией по `MessageID`.
    *   [06_backpressure](./02_examples/06_backpressure/main.go) — сравнение
        блокирующего backpressure и отбрасывания нагрузки (`drop` / Load
        Shedding) на буферизованном канале.
    *   [07_token_bucket](./02_examples/07_token_bucket/main.go) —
        потокобезопасный
        ограничитель скорости Token Bucket на Go.
    *   [08_circuit_breaker](./02_examples/08_circuit_breaker/main.go) —
        конечный автомат Circuit Breaker (`Closed` → `Open` → `Half-Open`).
    *   [09_retry_backoff](./02_examples/09_retry_backoff/main.go) — повторы с
        экспоненциальным backoff, jitter и отменой через `context.Context`.
    *   [10_graceful_shutdown](./02_examples/10_graceful_shutdown/main.go) —
        корректная остановка `http.Server` с доработкой активного запроса.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL`:
    *   `01_pick_shard` — детерминированный выбор шарда по ключу (`FNV-1a`).
    *   `02_reshard_moved` — оценка доли ключей, переезжающих при изменении
        числа шардов `N → M`.
    *   `03_classify_cap` — выбор `CP` или `AP` стратегии по требованиям
        домена.
    *   `04_check_quorum` — проверка строгого кворума ($R + W > N$) и
        живучести кластера при отказе узлов.
    *   `05_idempotent_handler` — идемпотентный обработчик платежей, защищённый
        от дублей `at-least-once`.
    *   `06_classify_guarantee` — определение гарантии доставки (`at-most-once`,
        `at-least-once`, `effectively-once`).
    *   `07_partition_by_key` — распределение событий по партициям с
        сохранением порядка внутри одного ключа.
    *   `08_outbox_drain` — реализация воркера выгрузки Transactional Outbox с
        защитой от повторной отправки при сбое.
    *   `09_token_bucket_allow` — метод `Allow(now)` алгоритма Token Bucket.
    *   `10_circuit_breaker_transitions` — переходы состояний Circuit Breaker
        по потоку успехов и ошибок.
    *   `11_backoff_sequence` — расчёт последовательности задержек Exponential
        Backoff с потолком `cap`.
    *   `12_call_with_timeout` — обёртка вызова с контролем дедлайна через
        `context.Context` и `select`.
    *   `13_fixed_window_limiter` — ограничитель частоты запросов Fixed Window.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Изучите фундаментальную теорию в DevOps-репозитории и прочитайте уроки в
   [01_theory/](./01_theory/).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_consistent_hashing/main.go
   go run ./02_examples/08_circuit_breaker/main.go
   go run ./02_examples/10_graceful_shutdown/main.go
   ```
3. Решите упражнения из [03_exercises/](./03_exercises/):
   ```bash
   go run ./03_exercises/01_pick_shard/main.go
   go run ./03_exercises/09_token_bucket_allow/main.go
   ```
4. Сверьтесь с ответами в [04_exercises_answers/](./04_exercises_answers/).

## Когда двигаться дальше

После изучения всех строительных блоков переходите к финальному модулю —
сборке целостных систем и расчётам для собеседований в
[04_case_studies](../04_case_studies/).

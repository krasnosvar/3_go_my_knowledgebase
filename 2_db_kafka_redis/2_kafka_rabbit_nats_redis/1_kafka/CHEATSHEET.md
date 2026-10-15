# Kafka — шпаргалка

Подробности и объяснения — в [interview_questions/theory.md](./interview_questions/theory.md)
и [basic_kafka_questions.md](./interview_questions/basic_kafka_questions.md). ⚠️ — частая ловушка.
Команды CLI не запускались — проверяйте на стенде из [docker_compose/](./docker_compose/README.md)
(Confluent 7.4 = Kafka 3.4, режим ZooKeeper).

## Модель данных

Подробно: [theory.md → Topics, Partitions, Offsets](./interview_questions/theory.md#topics-partitions-offsets).

| Понятие | Суть |
| :--- | :--- |
| Topic | именованный поток событий (≈ таблица для событий); сообщения хранятся по retention, а не удаляются после чтения |
| Partition | упорядоченный append-only лог; единица параллелизма; порядок гарантирован **только внутри** партиции |
| Offset | номер сообщения в партиции; consumer хранит свой offset → replay, чтение с начала / с конца |
| Ключ есть | `hash(key) % num_partitions` → все события одного ключа в одной партиции (порядок по сущности) |
| Ключа нет | раскидывается по партициям (round-robin; Java-клиент с 2.4 — sticky-батчами) → баланс, но без порядка |
| Broker | сервер Kafka: хранит партиции, обслуживает read/write |
| Метаданные кластера | ZooKeeper (исторически) или KRaft; KRaft production-ready с 3.3, в 4.0 ZooKeeper удалён |
| Хранение на диске | лог партиции = сегменты файлов + индексы по offset/time |

| Kafka vs классическая очередь (RabbitMQ, SQS) | Kafka | Очередь |
| :--- | :--- | :--- |
| После чтения | сообщение остаётся (retention) | удаляется |
| Читателей одного сообщения | любое число групп независимо | один consumer |
| Replay | да (сдвинуть offset) | нет |
| Масштабирование | партиции, горизонтально | сложнее |

## Когда Kafka

Подробно: [theory.md → Когда использовать Kafka?](./interview_questions/theory.md#когда-использовать-kafka).

| Да | Нет |
| :--- | :--- |
| event streaming, интеграция микросервисов | RPC / request-response → gRPC |
| централизованные логи и метрики, real-time аналитика | маленькие объёмы (< 1000 msg/день) |
| CDC, event sourcing, ETL | нужна простота → RabbitMQ |
| высокая пропускная способность | распределённые транзакции между сервисами |

## Producer

Подробно: [theory.md → Producers](./interview_questions/theory.md#producers-производители).

| `acks` | Ждём | Скорость | Риск |
| :--- | :--- | :--- | :--- |
| `0` | ничего | максимальная | потеря сообщений |
| `1` | запись на leader | баланс | потеря, если leader упал до репликации |
| `all` / `-1` | запись на все ISR | медленнее | минимальный; вместе с `min.insync.replicas` |

| Настройка (Java-имя) | sarama | Зачем |
| :--- | :--- | :--- |
| `acks=all` | `config.Producer.RequiredAcks = sarama.WaitForAll` | надёжность |
| `retries` | `config.Producer.Retry.Max = 5` | повторы при сбоях сети |
| `enable.idempotence=true` | `config.Producer.Idempotent = true` + `config.Net.MaxOpenRequests = 1` | без дублей при ретраях (producer id + sequence, в пределах партиции) |
| `compression.type` | `config.Producer.Compression = sarama.CompressionSnappy` | меньше трафика (также lz4/zstd) |
| `linger.ms` / `batch.size` | `Flush.Frequency = 100*time.Millisecond`, `Flush.Messages = 100` | батчирование: 1000 сообщений ≈ 10 round-trip вместо 1000 |
| — | `config.Producer.Return.Successes = true` | обязательно для `NewSyncProducer` |
| `message.max.bytes` (broker) | — | лимит размера; большие payload → object storage + ссылка |

Kafka 3.0+: в Java-продюсере `acks=all` и `enable.idempotence=true` по умолчанию.

## Consumer и consumer group

Подробно: [theory.md → Consumers](./interview_questions/theory.md#consumers-потребители),
[theory.md → Consumer Groups](./interview_questions/theory.md#consumer-groups).

| Правило | Следствие |
| :--- | :--- |
| Одна партиция → один consumer в группе | 3 партиции + 5 consumers = 2 простаивают |
| Один consumer может читать несколько партиций | 3 партиции + 1 consumer = работает, но медленно |
| Разные группы независимы | `analytics` и `notifications` читают весь топик каждая |
| Committed offset хранится в `__consumer_offsets` | после рестарта чтение продолжается с него |
| Начальная позиция без коммита | `sarama.OffsetNewest` (latest) / `sarama.OffsetOldest` (earliest) |

| Commit | Как | Семантика |
| :--- | :--- | :--- |
| Auto commit | периодически (по умолчанию каждые 5 c) коммитит прочитанное | возможны и потеря (упали после коммита до обработки), и дубли |
| Manual после обработки | `AutoCommit.Enable = false`; `session.MarkMessage(msg, "")` после успеха | at-least-once, нужна идемпотентная обработка |
| До обработки | commit → process | at-most-once |

| Rebalance | |
| :--- | :--- |
| Что это | перераспределение партиций между consumers группы |
| Причины | consumer добавлен/удалён/упал, не вызывает poll дольше `max.poll.interval.ms`, изменился топик |
| Чем опасен | пауза обработки, повторная обработка незакоммиченного |
| Как смягчить | cooperative-sticky assignor (incremental rebalance), правильные таймауты, graceful shutdown |

## Репликация и надёжность

Подробно: [theory.md → Brokers и Replication](./interview_questions/theory.md#brokers-и-replication).

| Понятие | Суть |
| :--- | :--- |
| Replication factor | число копий партиции на разных брокерах (RF=3 → 3× места) |
| Leader / Follower | клиенты пишут в leader партиции, followers реплицируют |
| ISR (in-sync replicas) | реплики, успевающие за leader; `acks=all` ждёт все ISR |
| Падение leader | новый leader выбирается из ISR, клиенты переключаются за секунды |
| `min.insync.replicas` | минимум ISR для записи с `acks=all`; иначе запись отклоняется (надёжность ценой доступности) |

## Семантики доставки

Подробно: [theory.md → Гарантии доставки](./interview_questions/theory.md#гарантии-доставки).

| Семантика | Producer | Consumer | Когда |
| :--- | :--- | :--- | :--- |
| At-most-once | `acks=0` | commit до обработки | метрики, некритичные логи |
| At-least-once | `acks=all` + retries | commit после обработки | большинство приложений (+ идемпотентный consumer) |
| Exactly-once | idempotence + транзакции | `isolation.level=read_committed` | платежи, read-process-write |

| Idempotent producer | Транзакции |
| :--- | :--- |
| нет дублей при ретраях в одной партиции | атомарная запись в несколько партиций + коммит offset в той же транзакции |
| `enable.idempotence=true` | `BeginTransaction` → send → `SendOffsetsToTransaction` → `CommitTransaction` |
| в sarama есть | по теории — полнее в confluent-kafka-go / franz-go |

## Retention и compaction

Подробно: [basic_kafka_questions.md → вопрос 14](./interview_questions/basic_kafka_questions.md#14-что-такое-retention-log-compaction-и-когда-что-использовать).

| Политика | Что хранит | Когда |
| :--- | :--- | :--- |
| Retention (время/размер) | удаляет старые сегменты | event log |
| Log compaction | последнее значение по ключу | состояние, справочники, CDC upsert |
| Обе | можно комбинировать | — |

## Партиции и ключ

Подробно: [theory.md → Выбор количества партиций](./interview_questions/theory.md#выбор-количества-партиций).

| Вопрос | Ответ |
| :--- | :--- |
| Сколько партиций | `max(T/P, T/C)`: T — целевой throughput, P/C — throughput producer/consumer на партицию; старт 3–6 |
| Пример | 1000 msg/s, P=500, C=100 → `max(2, 10) = 10` |
| Ключ | `user_id`, `order_id` — сущность, по которой нужен порядок |
| Глобальный порядок | только 1 партиция → нет параллелизма; обычно решают иначе |

## Паттерны и экосистема

Подробно: [theory.md → Практические паттерны](./interview_questions/theory.md#практические-паттерны).

| Задача | Решение |
| :--- | :--- |
| Ошибки обработки, poison message | ограниченный retry с backoff → DLQ-топик, затем `MarkMessage` |
| Распределённая транзакция между сервисами | Saga (choreography через события + компенсации) |
| Синхронизация БД → другие системы | CDC (Debezium читает binlog/WAL) через Kafka Connect |
| Запись в БД + публикация события | Outbox pattern |
| Хранить историю изменений | Event Sourcing; CQRS — отдельные модели записи и чтения |
| Интеграция без кода | Kafka Connect: source / sink connectors |
| Потоковая обработка в приложении | Kafka Streams (JVM); в Go — своя логика или сторонние библиотеки |
| Эволюция схем | Avro/Protobuf + Schema Registry, совместимость backward/forward/full |

## CLI (стенд)

Подробно: [docker_compose/README.md](./docker_compose/README.md).

| Задача | Команда |
| :--- | :--- |
| Поднять / статус / снести с данными | `docker compose up -d` / `docker compose ps` / `docker compose down -v` |
| Список топиков | `docker compose exec kafka kafka-topics --bootstrap-server kafka:29092 --list` |
| Партиции, leader, ISR | `docker compose exec kafka kafka-topics --bootstrap-server kafka:29092 --describe --topic test-topic` |
| Увеличить число партиций | `docker compose exec kafka kafka-topics --bootstrap-server kafka:29092 --alter --topic test-topic --partitions 3` |
| Записать (строка = сообщение, Ctrl+D) | `docker compose exec kafka kafka-console-producer --bootstrap-server kafka:29092 --topic test-topic` |
| Прочитать с начала | `docker compose exec kafka kafka-console-consumer --bootstrap-server kafka:29092 --topic test-topic --from-beginning --timeout-ms 5000` |
| Lag и offsets группы | `docker compose exec kafka kafka-consumer-groups --bootstrap-server kafka:29092 --describe --group group-1` |
| UI | `http://localhost:8080` (Kafka UI) |
| Адрес для Go-клиента с хоста | `localhost:9092`; внутри compose-сети — `kafka:29092` |

## Мониторинг

Подробно: [theory.md → Мониторинг](./interview_questions/theory.md#мониторинг).

| Где | Метрики |
| :--- | :--- |
| Consumer | lag (log end offset − committed offset), records-consumed-rate, commit-latency |
| Producer | request-latency, record-send-rate, record-error-rate, buffer-available-bytes |
| Broker | under-replicated-partitions, offline-partitions, active-controller = 1, ISR shrink/expand, диск |
| Инструменты | JMX exporter + Prometheus/Grafana, Burrow (lag) |

## Ловушки

- ⚠️ Порядок есть только внутри партиции: события одной сущности без ключа перемешаются.
- ⚠️ Consumers в группе больше, чем партиций → лишние простаивают.
- ⚠️ Добавили партиции → `hash(key) % N` меняется, новые события ключа уходят в другую партицию.
- ⚠️ `acks=all` без `min.insync.replicas` ≥ 2 → при ISR из одного leader гарантия как у `acks=1`.
- ⚠️ Auto commit + асинхронная обработка → потеря сообщений при падении.
- ⚠️ At-least-once без идемпотентного consumer → дубли в БД.
- ⚠️ Долгая обработка дольше `max.poll.interval.ms` → consumer выкинут из группы, rebalance, повторная обработка.
- ⚠️ Poison message без DLQ → consumer бесконечно падает на одном offset.
- ⚠️ Payload в мегабайты → падает throughput и репликация; храните в object storage, в Kafka — ссылку.
- ⚠️ Удалили обязательное поле из схемы → сломали потребителей; добавляйте поля с default.

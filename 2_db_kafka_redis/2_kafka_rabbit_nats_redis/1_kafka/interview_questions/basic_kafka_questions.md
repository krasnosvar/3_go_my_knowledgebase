# ТОП-20 вопросов по Apache Kafka для Go-разработчика бэкенда

## Простые вопросы (Junior уровень)

### 1. Что такое Apache Kafka и какие задачи она решает?

**Что ожидают услышать:**

- Распределенная платформа для стриминга событий (event streaming)
- Pub/Sub (producer-consumer) и хранение событий как log
- Высокая пропускная способность и горизонтальное масштабирование
- Типичные кейсы: логи, события, CDC, интеграции, очереди, аналитика

### 2. Чем Kafka отличается от классической очереди сообщений (RabbitMQ/SQS)?

**Что ожидают услышать:**

- Kafka хранит события на диске и позволяет перечитывать (retention)
- Потребители читают по offset’ам (pull), а не “удаление сообщения после ack” как основная модель
- Масштабирование через partition’ы и consumer group
- Сильные стороны: throughput, replay, event log; слабые: сложнее эксплуатация

### 3. Что такое topic, partition и offset?

**Что ожидают услышать:**

- Topic — логическая категория событий
- Partition — физическая часть топика (упорядоченность только внутри partition)
- Offset — позиция сообщения в partition (монотонно растет)
- Порядок гарантирован внутри partition, между partition — нет

### 4. Что такое producer и consumer? Какие основные настройки ты знаешь?

**Что ожидают услышать:**

- Producer публикует сообщения, consumer читает
- Producer: `acks`, `retries`, `linger.ms`, `batch.size`, `compression.type`, `enable.idempotence`
- Consumer: `group.id`, `auto.offset.reset`, `enable.auto.commit`, `max.poll.interval.ms`, `max.poll.records`

### 5. Что такое consumer group?

**Что ожидают услышать:**

- Набор consumer’ов, которые совместно читают один topic
- Каждая partition назначается максимум одному consumer’у в группе
- Масштабирование чтения = увеличение числа consumer’ов (до числа partition)
- Несколько групп могут читать один и тот же topic независимо

## Средние вопросы (Middle уровень)

### 6. Как Kafka обеспечивает отказоустойчивость? Что такое replication factor и ISR?

**Что ожидают услышать:**

- Репликация partition на несколько broker’ов
- Leader/Follower реплики, запись идет в leader
- ISR (in-sync replicas) — реплики, которые успевают за лидером
- Потеря leader → выбор нового leader из ISR

### 7. Что означает настройка `acks` у producer и как она влияет на надежность/латентность?

**Что ожидают услышать:**

- `acks=0` — не ждем подтверждения (быстро, но риск потери)
- `acks=1` — подтверждает leader после записи (частичная надежность)
- `acks=all`/`-1` — подтверждение после записи на ISR (максимальная надежность)
- Комбинация с `min.insync.replicas` влияет на доступность

### 8. Какие семантики доставки бывают (at-most-once / at-least-once / exactly-once)?

**Что ожидают услышать:**

- At-most-once: можно потерять, но без дублей (commit offset до обработки)
- At-least-once: без потерь, но возможны дубли (commit после обработки)
- Exactly-once: сложнее, чаще через транзакции Kafka + идемпотентность/детерминизм
- На практике часто строят идемпотентных consumer’ов + at-least-once

### 9. Что такое idempotent producer и чем он отличается от Kafka transactions?

**Что ожидают услышать:**

- Идемпотентный producer предотвращает дубли при ретраях в пределах одной partition (producer id + sequence)
- Требует `enable.idempotence=true`, обычно `acks=all`, ограничение `max.in.flight.requests.per.connection`
- Transactions дают “атомарность” записи в несколько partition и “read-process-write” EOS совместно с `read_committed`

### 10. Как работает commit offset’ов у consumer? В чем разница auto commit и manual commit?

**Что ожидают услышать:**

- Offset хранится в `__consumer_offsets`
- Auto commit периодически коммитит “последний прочитанный” offset → риск потерять обработку при падении
- Manual commit после успешной обработки → at-least-once
- Важны тайминги: `max.poll.interval.ms`, “rebalance” и повторная обработка

### 11. Что такое rebalance? Почему он происходит и чем опасен?

**Что ожидают услышать:**

- Перераспределение partition между consumer’ами группы
- Причины: добавили/убрали consumer, consumer “завис” и не делает poll, изменения топика
- Опасность: паузы в обработке, повторная обработка (если offset не закоммичен)
- Механизмы: cooperative-sticky assignor (incremental rebalance), правильные таймауты, graceful shutdown

### 12. Как выбрать ключ сообщения и количество partition?

**Что ожидают услышать:**

- Ключ определяет partitioning и порядок обработки по сущности (например `user_id`, `order_id`)
- Без ключа — round-robin → лучше баланс, но нет порядка по сущности
- Partition — единица параллелизма; больше partition → выше параллелизм и накладные расходы
- Учитывать рост, hotspots по ключам, ограничения на consumer group

## Продвинутые вопросы (Senior уровень)

### 13. Какие гарантии порядка дает Kafka и как проектировать обработку с учетом этого?

**Что ожидают услышать:**

- Порядок только внутри partition
- Для строгого порядка по entity нужен ключ и все события entity должны попадать в одну partition
- Если нужен глобальный порядок — обычно невозможно/дорого, решают иначе

### 14. Что такое retention, log compaction и когда что использовать?

**Что ожидают услышать:**

- Retention по времени/размеру удаляет старые сегменты
- Compaction хранит “последнее значение по ключу” (для state/таблиц, справочников)
- Можно комбинировать compaction + retention
- Примеры: CDC (upsert события) → compaction; event log → retention

### 15. Как Kafka хранит данные (segments, index) и почему важен размер сообщений?

**Что ожидают услышать:**

- Лог разбит на сегменты файлов, индексы по offset/time
- Большие сообщения ухудшают throughput/GC/replication; есть лимиты `message.max.bytes`, `fetch.max.bytes`
- Часто рекомендуют “не тащить payload 10MB”, использовать object storage + ссылку

### 16. Как устроена доставка “read-process-write” без потерь и с минимальными дублями?

**Что ожидают услышать:**

- Базово: manual commit после успешной обработки + идемпотентная обработка
- Если нужно EOS: транзакционный producer + consumer `isolation.level=read_committed`
- Outbox pattern / CDC (Debezium) для интеграции с БД
- DLQ/Retry topics для ошибок

### 17. Как обрабатывать ошибки в consumer: poison message, ретраи, DLQ?

**Что ожидают услышать:**

- Poison message — сообщение, которое всегда падает
- Ограниченный retry с backoff (внутренний или через retry topic с задержкой)
- DLQ (dead letter topic) для ручного разбора
- Важно не “убить” consumer group бесконечными падениями

### 18. Как мониторить Kafka и consumer lag? Какие метрики важны?

**Что ожидают услышать:**

- Consumer lag (разница между log end offset и committed offset)
- Producer error rate, request latency, throughput
- Broker: under-replicated partitions, ISR shrink/expand, disk usage
- Инструменты: Prometheus/JMX exporter, Grafana, Burrow (lag)

### 19. Какие подходы к schema evolution ты знаешь? Avro/Protobuf + Schema Registry?

**Что ожидают услышать:**

- Контракты сообщений и совместимость (backward/forward/full)
- Schema Registry (Confluent/Apicurio), версии схем
- Не ломать потребителей: добавление полей с default, не удалять обязательные поля
- Валидация на CI, контрактное тестирование

### 20. Опиши архитектуру высоконагруженной системы на Kafka (100k msg/s). На что обращаешь внимание?

**Что ожидают услышать:**

- Partitioning стратегия, ключи, количество partition
- Репликация, `min.insync.replicas`, `acks=all`, баланс надежности/латентности
- Producer batching + compression (lz4/zstd), connection reuse
- Consumer: параллелизм, backpressure, rate limiting, retry/DLQ
- Наблюдаемость и алерты, план на деградацию

## Дополнительные темы для обсуждения

- **Безопасность**: SASL/SCRAM, mTLS, ACL, шифрование, секреты

- **Эксплуатация**: upgrade, reassign partitions, балансировка лидеров

- **Настройки производительности**: `linger.ms`, `batch.size`, compression, fetch settings

- **Тестирование**: integration тесты с Testcontainers, contract tests

- **Интеграционные паттерны**: outbox/inbox, saga choreography, CDC (Debezium)

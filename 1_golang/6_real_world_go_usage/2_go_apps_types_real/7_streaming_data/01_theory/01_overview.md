# Streaming data

«Streaming» здесь — потребление и обработка непрерывного потока сообщений из брокера: Kafka, NATS, Pulsar, AWS Kinesis, Redis Streams.

## Базовый паттерн

```go
for {
    select {
    case <-ctx.Done():
        return
    case msg, ok := <-consumer.Messages():
        if !ok { return }
        if err := handle(ctx, msg); err != nil {
            handleFailure(msg, err)
            continue
        }
        consumer.Commit(msg.Offset)
    }
}
```

Consumer читает сообщения, обрабатывает, **коммитит offset** только после успеха.

## Delivery semantics

- **At-most-once**: сообщение может быть потеряно. Commit до обработки.
- **At-least-once**: сообщение может прийти дважды. Commit после обработки. Стандарт.
- **Exactly-once**: каждое сообщение строго один раз. Достижимо в Kafka transactions, но требует усилий и часто it's a marketing term.

В 95% случаев — at-least-once + idempotent consumer.

## Idempotency

При at-least-once consumer должен уметь обрабатывать дубликаты. Подходы:

- **Idempotency key** в сообщении, БД-таблица «обработанные ключи».
- **Upsert вместо insert** при записи в БД.
- **Внешний state** с `ON CONFLICT DO NOTHING`.

## Ordering

В Kafka сообщения упорядочены **в пределах partition**. Между partition'ами — нет.

Если порядок важен → ключ сообщения должен попадать в одну partition (например, по `userID`).

## Backpressure

Consumer медленнее producer'а → лаг растёт.

- Мониторь lag метриками.
- Параллельная обработка (worker pool на одну partition).
- Выделение «жирных» сообщений в отдельный topic.
- Если consumer хронически не справляется — добавь instances.

## Rebalance

При добавлении/удалении consumer'а в группе Kafka перераспределяет partition'ы. На время rebalance обработка стопается.

Хуки `OnRebalance` позволяют корректно зафиксировать offset и почистить state.

## Poison messages

Сообщение, которое нельзя обработать (битый формат, баг в коде). Стратегии:

- Пропустить + alert.
- Переместить в DLQ topic.
- Retry с лимитом, потом DLQ.

Никогда: бесконечно пытаться обработать одно и то же → весь pipeline стоит.

## Schema

Контракт сообщений:

- **JSON** — гибко, но без схемы.
- **Avro / Protobuf** — типизировано, schema registry, эволюция полей.

Schema registry (Confluent) — стандарт для Kafka в проде.

## Библиотеки в Go

- Kafka: `confluent-kafka-go` (через librdkafka, нужен cgo) или `segmentio/kafka-go` (чистый Go).
- NATS: `nats.go` (официальная).
- Redis Streams: `go-redis`.
- Pulsar: `apache/pulsar-client-go`.

## Правило

- At-least-once + idempotency — стандарт.
- Commit после успеха, не до.
- DLQ для poison messages.
- Партиционирование по ключу для гарантии порядка.
- Lag и rebalance — в мониторинге.

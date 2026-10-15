# Apache Kafka: Понятное объяснение основных концепций

## 📚 Содержание
1. [Что такое Kafka](#что-такое-kafka)
2. [Topics, Partitions, Offsets](#topics-partitions-offsets)
3. [Producers (Производители)](#producers-производители)
4. [Consumers (Потребители)](#consumers-потребители)
5. [Consumer Groups](#consumer-groups)
6. [Brokers и Replication](#brokers-и-replication)
7. [Гарантии доставки](#гарантии-доставки)
8. [Kafka Connect](#kafka-connect)
9. [Kafka Streams](#kafka-streams)
10. [Практические паттерны](#практические-паттерны)

---

## Что такое Kafka?

**Apache Kafka** — это распределенная платформа для потоковой обработки данных (distributed event streaming platform).

```
┌─────────────────────────────────────────────────────────┐
│              Kafka — Event Streaming Platform           │
└─────────────────────────────────────────────────────────┘

Ключевые особенности:
┌────────────────────────────────────────────────────────┐
│ • Высокая пропускная способность (миллионы msg/sec)    │
│ • Горизонтальное масштабирование                       │
│ • Fault-tolerant (отказоустойчивая)                    │
│ • Хранение сообщений на диске                          │
│ • Низкая задержка (несколько миллисекунд)              │
│ • Сообщения хранятся по времени (retention)            │
└────────────────────────────────────────────────────────┘
```

### Традиционная очередь vs Kafka

```
┌─────────────────────────────────────────────────────────┐
│        Традиционная очередь (RabbitMQ, SQS)             │
├─────────────────────────────────────────────────────────┤
│ Producer        Queue        Consumer                   │
│    │    msg1     │              │                       │
│    ├───────────► │ ────────────►│                       │
│    │    msg2     │              │                       │
│    ├───────────► │ ────────────►│                       │
│    │    msg3     │              │                       │
│    └───────────► └──────────────┘                       │
│                                                         │
│ ✅ Простая модель                                       │
│ ❌ Сообщение удаляется после чтения                     │
│ ❌ Сложно масштабировать                                │
│ ❌ Один consumer на сообщение                           │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│                      Apache Kafka                       │
├─────────────────────────────────────────────────────────┤
│ Producer        Topic         Consumers                 │
│    │                            │   │   │               │
│    │    msg1 ──────────────────►│   │   │               │
│    │    msg2 ──────────────────►│   │   │               │
│    │    msg3 ──────────────────►│   │   │               │
│    │            ▲               │   │   │               │
│    │            │ Persisted     │   │   │               │
│    │            │ (7 days)      │   │   │               │
│    │            │               │   │   │               │
│    │            └───────────────┴───┴───┘               │
│    │              Множество consumers                   │
│    │              могут читать                          │
│                                                         │
│ ✅ Сообщения хранятся                                   │
│ ✅ Множество consumers                                  │
│ ✅ Replay (перечитывание)                               │
│ ✅ Масштабируется горизонтально                         │
└─────────────────────────────────────────────────────────┘
```

### Когда использовать Kafka?

```
✅ ИСПОЛЬЗУЙТЕ Kafka для:
┌─────────────────────────────────────────────┐
│ • Event streaming (потоки событий)          │
│ • Микросервисная архитектура                │
│ • Логи и метрики (centralized logging)      │
│ • Real-time аналитика                       │
│ • CDC (Change Data Capture)                 │
│ • Event sourcing                            │
│ • Message queue между сервисами             │
│ • Data integration (ETL)                    │
│ • Высокая пропускная способность            │
└─────────────────────────────────────────────┘

❌ НЕ ИСПОЛЬЗУЙТЕ Kafka для:
┌─────────────────────────────────────────────┐
│ • RPC (request-response) — используйте gRPC │
│ • Маленькие объемы данных (<1000 msg/day)   │
│ • Нужна простота (используйте RabbitMQ)     │
│ • Транзакции между сервисами                │
└─────────────────────────────────────────────┘
```

---

## Topics, Partitions, Offsets

### Topic (Топик)

**Topic** — это категория/поток сообщений (как таблица в БД, но для событий).

```
Визуализация:
┌──────────────────────────────────────────────┐
│             Kafka Cluster                    │
│                                              │
│  Topic: "orders"                             │
│  Topic: "user-events"                        │
│  Topic: "payments"                           │
│  Topic: "notifications"                      │
│                                              │
└──────────────────────────────────────────────┘

Пример топиков в e-commerce:
• orders — новые заказы
• user-registrations — регистрации пользователей
• inventory-updates — обновления склада
• payment-confirmations — подтверждения платежей
```

### Partition (Партиция)

**Topic разделен на партиции** для параллелизма и масштабирования.

```
Topic: "orders"
┌────────────────────────────────────────────────────┐
│                                                    │
│  Partition 0:  [msg0][msg3][msg6][msg9] ...        │
│  Partition 1:  [msg1][msg4][msg7][msg10] ...       │
│  Partition 2:  [msg2][msg5][msg8][msg11] ...       │
│                                                    │
└────────────────────────────────────────────────────┘

Зачем партиции?
✅ Параллельная обработка (каждая партиция — отдельный поток)
✅ Масштабирование (больше партиций = больше throughput)
✅ Порядок сообщений внутри партиции гарантирован
⚠️ Порядок между партициями НЕ гарантирован
```

**Как сообщения распределяются по партициям?**

```
┌─────────────────────────────────────────────────┐
│          Producer отправляет сообщение          │
└────────────────┬────────────────────────────────┘
                 │
        ┌────────┴────────┐
        │ Есть key?       │
        └────────┬────────┘
                 │
    ┌────────────┴────────────┐
    ДА                       НЕТ
    │                         │
    ▼                         ▼
hash(key) % num_partitions  Round-robin
    │                         │
    ▼                         ▼
Одна и та же партиция     Случайная партиция
для одного key            каждый раз

Пример:
• Сообщение с key="user123" → всегда Partition 1
• Сообщение с key="user456" → всегда Partition 2
• Сообщение без key → Partition 0, 1, 2, 0, 1...
```

### Offset (Смещение)

**Offset** — уникальный номер сообщения внутри партиции.

```
Partition 0 топика "orders":
┌───────┬───────┬───────┬───────┬───────┬───────┐
│Offset:│   0   │   1   │   2   │   3   │   4   │
├───────┼───────┼───────┼───────┼───────┼───────┤
│ Msg:  │{...}  │{...}  │{...}  │{...}  │{...}  │
└───────┴───────┴───────┴───────┴───────┴───────┘
         ↑       ↑                       ↑
      Старые  Consumer               Новые
              читает здесь

Consumer помнит свой offset:
• Прочитал offset 0, 1 → следующий читать offset 2
• Можно вернуться назад (replay)
• Можно начать с начала (offset 0)
• Можно начать с конца (latest)
```

**Визуализация полной структуры:**

```
┌──────────────────────────────────────────────────────┐
│               Topic: "user-events"                   │
│              (3 partitions, replication=2)           │
└──────────────────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┐
         │               │               │
    Partition 0     Partition 1     Partition 2
         │               │               │
    ┌────▼────┐     ┌────▼────┐     ┌────▼────┐
    │ Offset  │     │ Offset  │     │ Offset  │
    ├─────────┤     ├─────────┤     ├─────────┤
    │ 0: {...}│     │ 0: {...}│     │ 0: {...}│
    │ 1: {...}│     │ 1: {...}│     │ 1: {...}│
    │ 2: {...}│     │ 2: {...}│     │ 2: {...}│
    │ 3: {...}│     │ 3: {...}│     │ 3: {...}│
    │   ...   │     │   ...   │     │   ...   │
    └─────────┘     └─────────┘     └─────────┘
         │               │               │
    (Broker 1)      (Broker 2)      (Broker 3)
```

---

## Producers (Производители)

### Как Producer отправляет сообщения?

```
Процесс отправки:
┌──────────────┐
│  Producer    │
└──────┬───────┘
       │
       │ 1. Создать сообщение
       │    { key: "user123", value: "clicked button" }
       │
       ▼
┌──────────────────────┐
│  Serializer          │  2. Сериализация (JSON → bytes)
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Partitioner         │  3. Выбор партиции (по key или round-robin)
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Batch (buffer)      │  4. Группировка в батчи
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Network             │  5. Отправка на Broker
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Kafka Broker        │  6. Сохранение на диск
└──────────────────────┘
```

**Пример в Go:**
```go
import (
    "github.com/IBM/sarama"
)

// Создание Producer
config := sarama.NewConfig()
config.Producer.Return.Successes = true
config.Producer.RequiredAcks = sarama.WaitForAll  // Ждать от всех реплик

producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
if err != nil {
    log.Fatal(err)
}
defer producer.Close()

// Отправка сообщения
message := &sarama.ProducerMessage{
    Topic: "user-events",
    Key:   sarama.StringEncoder("user123"),      // Ключ для партиционирования
    Value: sarama.StringEncoder(`{"event": "click", "url": "/home"}`),
}

partition, offset, err := producer.SendMessage(message)
if err != nil {
    log.Printf("Ошибка отправки: %v", err)
} else {
    log.Printf("Сообщение отправлено в partition %d offset %d", partition, offset)
}
```

### Батчирование (Batching)

```
БЕЗ батчирования:
Producer                    Kafka Broker
   │ msg1 ─────────────────►│
   │◄──────── ACK ──────────│
   │ msg2 ─────────────────►│
   │◄──────── ACK ──────────│
   │ msg3 ─────────────────►│
   │◄──────── ACK ──────────│

1000 сообщений = 1000 network round-trips ❌

С батчированием:
Producer                    Kafka Broker
   │                         │
   │ Собираем 100 сообщений  │
   │ в батч...               │
   │                         │
   │ [msg1..msg100] ────────►│
   │◄──────── ACK ──────────│
   │                         │
   │ [msg101..msg200] ──────►│
   │◄──────── ACK ──────────│

1000 сообщений = 10 network round-trips ✅
В 100 раз меньше!
```

**Настройка батчирования:**
```go
config.Producer.Compression = sarama.CompressionSnappy  // Сжатие
config.Producer.Flush.Frequency = 100 * time.Millisecond  // Отправлять каждые 100ms
config.Producer.Flush.Messages = 100                      // Или каждые 100 сообщений
config.Producer.Flush.MaxMessages = 1000                  // Макс. размер батча
```

### Acks (Подтверждения)

```
┌────────────────────────────────────────────────────┐
│              Уровни подтверждения                  │
├────────────────────────────────────────────────────┤
│                                                    │
│ acks=0 (No acknowledgment)                         │
│ Producer не ждет подтверждения                     │
│ ✅ Максимальная скорость                           │
│ ❌ Сообщения могут потеряться                      │
│                                                    │
│ acks=1 (Leader acknowledgment)                     │
│ Ждать подтверждения от leader                      │
│ ⚖️ Баланс скорость/надежность                      │
│ ⚠️ Может потеряться если leader упадет             │
│                                                    │
│ acks=all / acks=-1 (All replicas)                  │
│ Ждать подтверждения от всех реплик                 │
│ ✅ Максимальная надежность                         │
│ ❌ Медленнее                                       │
└────────────────────────────────────────────────────┘

Визуализация acks=all:
Producer                Leader      Replica 1   Replica 2
   │                      │             │           │
   │──── msg ────────────►│             │           │
   │                      │──replicate─►│           │
   │                      │──replicate─────────────►│
   │                      │◄─── ACK ────│           │
   │                      │◄─── ACK ───────────────│
   │◄───── ACK ──────────│             │           │
   │   (все подтвердили)  │             │           │
```

---

## Consumers (Потребители)

### Как Consumer читает сообщения?

```
Процесс чтения:
┌──────────────┐
│   Consumer   │
└──────┬───────┘
       │
       │ 1. Подписка на топик
       │    Subscribe("user-events")
       │
       ▼
┌──────────────────────┐
│  Kafka Broker        │  2. Fetch запрос (получить сообщения)
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Вернуть N сообщений │  3. Kafka возвращает батч
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Deserializer        │  4. Десериализация (bytes → JSON)
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Process             │  5. Обработка
└──────┬───────────────┘
       │
       ▼
┌──────────────────────┐
│  Commit offset       │  6. Сохранение прогресса
└──────────────────────┘
```

**Пример в Go:**
```go
// Создание Consumer
config := sarama.NewConfig()
config.Consumer.Return.Errors = true
config.Consumer.Offsets.Initial = sarama.OffsetNewest  // Читать с последнего

consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
if err != nil {
    log.Fatal(err)
}
defer consumer.Close()

// Подписка на партицию
partitionConsumer, err := consumer.ConsumePartition("user-events", 0, sarama.OffsetNewest)
if err != nil {
    log.Fatal(err)
}
defer partitionConsumer.Close()

// Чтение сообщений
for message := range partitionConsumer.Messages() {
    log.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s",
        message.Partition, message.Offset,
        string(message.Key), string(message.Value))
    
    // Обработка сообщения
    processEvent(message.Value)
}
```

### Offset Management (Управление смещением)

```
Consumer отслеживает свой прогресс через offset:

Partition 0:
┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐
│ 0 │ 1 │ 2 │ 3 │ 4 │ 5 │ 6 │ 7 │ 8 │ 9 │
└───┴───┴───┴───┴───┴───┴───┴───┴───┴───┘
 ✓   ✓   ✓   ✓   ▶
Прочитано      Текущая
               позиция

Committed offset = 4
(Consumer может безопасно перезапуститься с offset 4)
```

**Стратегии commit:**

```
┌────────────────────────────────────────────────────┐
│            Auto Commit (по умолчанию)              │
├────────────────────────────────────────────────────┤
│ • Автоматически каждые 5 секунд                    │
│ ✅ Просто                                          │
│ ❌ Может потерять или дублировать сообщения        │
│                                                    │
│ Пример:                                            │
│ 1. Прочитали offset 100-110                        │
│ 2. Обрабатываем...                                 │
│ 3. Auto-commit → offset 110                        │
│ 4. Crash!                                          │
│ 5. Перезапуск → начнем с offset 110                │
│ → Потеряли необработанные сообщения ❌             │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│              Manual Commit                         │
├────────────────────────────────────────────────────┤
│ • Commit только после обработки                    │
│ ✅ Гарантия at-least-once                          │
│ ⚠️ Нужно обрабатывать дубликаты                    │
│                                                    │
│ Пример:                                            │
│ 1. Прочитали сообщение offset 100                  │
│ 2. Обработали успешно                              │
│ 3. Commit offset 101                               │
│ 4. Crash!                                          │
│ 5. Перезапуск → начнем с offset 101 ✅             │
└────────────────────────────────────────────────────┘
```

**Manual commit в Go:**
```go
config.Consumer.Offsets.AutoCommit.Enable = false  // Отключить auto-commit

// Consumer Group (не partition consumer)
group, err := sarama.NewConsumerGroup(
    []string{"localhost:9092"},
    "my-consumer-group",
    config,
)

// Обработчик
func (h *Handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for message := range claim.Messages() {
        // Обработка
        err := processMessage(message)
        if err != nil {
            log.Printf("Ошибка обработки: %v", err)
            continue  // Не делаем commit, попробуем еще раз
        }
        
        // Commit только после успешной обработки
        session.MarkMessage(message, "")
    }
    return nil
}
```

---

## Consumer Groups

### Что такое Consumer Group?

**Consumer Group** — группа consumers, которые совместно читают топик.

```
БЕЗ Consumer Group (один consumer):
Topic "orders" (3 partitions):
┌──────────┐  ┌──────────┐  ┌──────────┐
│Partition0│  │Partition1│  │Partition2│
└────┬─────┘  └────┬─────┘  └────┬─────┘
     │             │             │
     └─────────────┴─────────────┘
                   │
              ┌────▼────┐
              │Consumer │
              └─────────┘

Один consumer читает все партиции
❌ Медленно
❌ Не использует параллелизм

С Consumer Group (3 consumers):
Topic "orders" (3 partitions):
┌──────────┐  ┌──────────┐  ┌──────────┐
│Partition0│  │Partition1│  │Partition2│
└────┬─────┘  └────┬─────┘  └────┬─────┘
     │             │             │
     ▼             ▼             ▼
┌─────────┐  ┌─────────┐  ┌─────────┐
│Consumer1│  │Consumer2│  │Consumer3│
└─────────┘  └─────────┘  └─────────┘
     └───────────┬───────────┘
          Consumer Group "my-group"

Каждый consumer читает свою партицию
✅ Параллельная обработка
✅ Автоматическая балансировка
```

### Rebalancing (Перебалансировка)

```
Что происходит при добавлении/удалении consumer?

Изначально (2 consumers, 4 partitions):
┌───┐ ┌───┐ ┌───┐ ┌───┐
│P0 │ │P1 │ │P2 │ │P3 │
└─┬─┘ └─┬─┘ └─┬─┘ └─┬─┘
  │     │     │     │
  └──┬──┘     └──┬──┘
     │           │
┌────▼───┐  ┌───▼────┐
│Consumer│  │Consumer│
│   1    │  │   2    │
└────────┘  └────────┘

Добавили Consumer 3:
┌───┐ ┌───┐ ┌───┐ ┌───┐
│P0 │ │P1 │ │P2 │ │P3 │
└─┬─┘ └─┬─┘ └─┬─┘ └─┬─┘
  │     │     │     │
  │     │     │     │
┌─▼──┐ ┌▼──┐ ┌▼──┐ │
│C1  │ │C2 │ │C3 │ │  ← REBALANCE!
└────┘ └───┘ └───┘ │
                   (C3 получил P3)

Consumer 2 упал:
┌───┐ ┌───┐ ┌───┐ ┌───┐
│P0 │ │P1 │ │P2 │ │P3 │
└─┬─┘ └─┬─┘ └─┬─┘ └─┬─┘
  │     │     │     │
  └──┬──┘     └──┬──┘
     │           │
┌────▼───┐  ┌───▼────┐
│Consumer│  │Consumer│  ← REBALANCE!
│   1    │  │   3    │
└────────┘  └────────┘
(C1 и C3 разделили партиции)
```

**Важные правила:**
```
┌────────────────────────────────────────────────────┐
│ 1. Одна партиция = один consumer в группе          │
│                                                    │
│    ✅ 3 партиции, 3 consumers → идеально           │
│    ⚠️ 3 партиции, 5 consumers → 2 idle             │
│    ❌ 3 партиции, 1 consumer → медленно            │
│                                                    │
│ 2. Consumer может читать несколько партиций        │
│                                                    │
│    2 партиции, 1 consumer:                         │
│    Consumer читает обе партиции                    │
│                                                    │
│ 3. Разные группы независимы                        │
│                                                    │
│    Group "analytics" → читает весь топик           │
│    Group "notifications" → читает весь топик       │
│    (оба независимо!)                               │
└────────────────────────────────────────────────────┘
```

**Пример Consumer Group в Go:**
```go
config := sarama.NewConfig()
config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
config.Consumer.Offsets.Initial = sarama.OffsetOldest

// Создать Consumer Group
group, err := sarama.NewConsumerGroup(
    []string{"localhost:9092"},
    "my-consumer-group",  // ID группы
    config,
)

// Подписаться на топики
topics := []string{"user-events", "orders"}
ctx := context.Background()

// Запустить потребление
for {
    err := group.Consume(ctx, topics, &ConsumerHandler{})
    if err != nil {
        log.Printf("Ошибка: %v", err)
    }
}

// Handler
type ConsumerHandler struct{}

func (h *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
    log.Println("Consumer group setup")
    return nil
}

func (h *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
    log.Println("Consumer group cleanup")
    return nil
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for message := range claim.Messages() {
        log.Printf("Message: %s", string(message.Value))
        session.MarkMessage(message, "")
    }
    return nil
}
```

---

## Brokers и Replication

### Kafka Cluster (Кластер брокеров)

```
┌────────────────────────────────────────────────────┐
│              Kafka Cluster                         │
│                                                    │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐         │
│  │ Broker 1 │  │ Broker 2 │  │ Broker 3 │         │
│  │ :9092    │  │ :9093    │  │ :9094    │         │
│  └──────────┘  └──────────┘  └──────────┘         │
│       │             │             │                │
│       └─────────────┴─────────────┘                │
│        KRaft controllers (метаданные)              │
│     (до Kafka 4.0 — опционально ZooKeeper)         │
└────────────────────────────────────────────────────┘

Broker — это сервер Kafka
• Хранит партиции
• Обрабатывает read/write запросы
• Координируется с другими брокерами
```

> **KRaft vs ZooKeeper.** Исторически метаданные кластера (список брокеров, топики, выбор контроллера) хранились в ZooKeeper. Начиная с Kafka 3.3 production-ready стал режим KRaft (Kafka Raft): метаданные хранят сами брокеры/контроллеры по протоколу Raft. В Kafka 4.0 поддержка ZooKeeper удалена полностью — работает только KRaft. На собеседовании про ZooKeeper стоит говорить как про устаревшую схему (и про миграцию ZooKeeper → KRaft на версиях 3.x).

### Replication (Репликация)

```
Topic "orders" с replication-factor=3:
┌────────────────────────────────────────────────────┐
│              Partition 0                           │
├────────────────────────────────────────────────────┤
│                                                    │
│  ┌─────────────┐                                   │
│  │   LEADER    │ ← Принимает все read/write        │
│  │  (Broker 1) │                                   │
│  └──────┬──────┘                                   │
│         │                                          │
│    Репликация                                      │
│         │                                          │
│   ┌─────┴──────┐                                   │
│   │            │                                   │
│ ┌─▼──────┐  ┌──▼──────┐                            │
│ │FOLLOWER│  │FOLLOWER │ ← Только реплики           │
│ │(Broker2)│  │(Broker3)│                            │
│ └────────┘  └─────────┘                            │
└────────────────────────────────────────────────────┘

✅ Отказоустойчивость: если Leader упадет, Follower станет Leader
✅ Данные не теряются
⚠️ Занимает больше места (3x для replication=3)
```

**In-Sync Replicas (ISR)**

```
┌────────────────────────────────────────────────────┐
│             In-Sync Replicas (ISR)                 │
├────────────────────────────────────────────────────┤
│                                                    │
│ ISR — реплики, которые "в синхронизации" с Leader │
│                                                    │
│ Leader (Broker 1): offset 100                      │
│ Follower (Broker 2): offset 100 ✅ In-sync         │
│ Follower (Broker 3): offset 95  ❌ Out-of-sync     │
│                                                    │
│ ISR list = [Broker 1, Broker 2]                    │
│                                                    │
│ acks=all ждет подтверждения от всех ISR            │
└────────────────────────────────────────────────────┘

Что происходит при падении Leader?
┌────────────────────────────────────┐
│ 1. Leader (Broker 1) упал          │
│ 2. Выбирается новый Leader из ISR  │
│ 3. Broker 2 становится Leader      │
│ 4. Clients автоматически           │
│    переключаются на Broker 2       │
└────────────────────────────────────┘
Время переключения: ~секунды
```

### Распределение партиций

```
Topic "orders" (6 partitions, replication=2):
┌──────────────────────────────────────────────────┐
│                   Broker 1                       │
│  P0 (Leader)   P1 (Follower)  P4 (Leader)        │
└──────────────────────────────────────────────────┘
┌──────────────────────────────────────────────────┐
│                   Broker 2                       │
│  P0 (Follower) P2 (Leader)    P5 (Follower)      │
└──────────────────────────────────────────────────┘
┌──────────────────────────────────────────────────┐
│                   Broker 3                       │
│  P1 (Leader)   P3 (Leader)    P5 (Leader)        │
└──────────────────────────────────────────────────┘

Kafka автоматически распределяет:
✅ Leaders равномерно между брокерами
✅ Replicas на разных брокерах
✅ Нагрузка балансируется
```

---

## Гарантии доставки

### Семантики доставки

```
┌────────────────────────────────────────────────────┐
│          At-most-once (максимум один раз)          │
├────────────────────────────────────────────────────┤
│ Сообщение может быть потеряно, но не дублировано   │
│                                                    │
│ Producer: acks=0                                   │
│ Consumer: auto-commit перед обработкой             │
│                                                    │
│ ✅ Максимальная скорость                           │
│ ❌ Можно потерять данные                           │
│                                                    │
│ Use case: Метрики, логи (некритичные данные)      │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│         At-least-once (минимум один раз)           │
├────────────────────────────────────────────────────┤
│ Сообщение НЕ будет потеряно, но может дублироваться│
│                                                    │
│ Producer: acks=all, retries                        │
│ Consumer: commit после обработки                   │
│                                                    │
│ ✅ Надежность                                      │
│ ⚠️ Нужно обрабатывать дубликаты                    │
│                                                    │
│ Use case: Большинство приложений                   │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│        Exactly-once (ровно один раз)               │
├────────────────────────────────────────────────────┤
│ Сообщение обработано ровно один раз                │
│                                                    │
│ Producer: idempotence + transactions               │
│ Consumer: transactions                             │
│                                                    │
│ ✅ Идеальная гарантия                              │
│ ❌ Сложнее настройка                               │
│ ❌ Медленнее                                       │
│                                                    │
│ Use case: Финансовые транзакции, платежи          │
└────────────────────────────────────────────────────┘
```

### Idempotent Producer (Идемпотентность)

```
БЕЗ idempotence:
Producer                    Kafka
   │ msg1 ─────────────────►│
   │                 [сохранено]
   │◄────X (сеть упала)     │
   │                        │
   │ msg1 retry ───────────►│
   │                 [сохранено ДВАЖДЫ] ❌
   │◄──────── ACK ──────────│

Дубликат!

С idempotence:
Producer                    Kafka
   │ msg1 (ID=100)─────────►│
   │                 [сохранено]
   │◄────X (сеть упала)     │
   │                        │
   │ msg1 retry (ID=100)───►│
   │                 [уже есть, игнорируем] ✅
   │◄──────── ACK ──────────│

Нет дубликата!
```

**Включение idempotence в Go:**
```go
config := sarama.NewConfig()
config.Producer.Idempotent = true
config.Producer.RequiredAcks = sarama.WaitForAll
config.Producer.Retry.Max = 5
config.Net.MaxOpenRequests = 1  // Важно для idempotence!

producer, err := sarama.NewSyncProducer(brokers, config)
```

### Транзакции

```
Задача: Прочитать из топика A, обработать, записать в топик B
       Атомарно (все или ничего)

БЕЗ транзакций:
1. Читаем из A       ✅
2. Обработка         ✅
3. Пишем в B         ✅
4. Commit offset в A ❌ CRASH!
→ Перезапуск → обработаем дважды ❌

С транзакциями:
BEGIN TRANSACTION
  1. Читаем из A
  2. Обработка
  3. Пишем в B
  4. Commit offset в A
COMMIT TRANSACTION

Все операции атомарны! ✅
```

**Пример транзакций (концептуально):**
```go
// Sarama пока не поддерживает транзакции полностью
// Используйте confluent-kafka-go или franz-go

// Псевдокод:
producer.BeginTransaction()

// Читаем сообщение
message := consumer.Poll()

// Обрабатываем
result := process(message)

// Пишем результат
producer.Send(topic, result)

// Commit offset
producer.SendOffsetsToTransaction(message.Offset)

// Commit транзакции
producer.CommitTransaction()
```

---

## Kafka Connect

### Что такое Kafka Connect?

**Kafka Connect** — фреймворк для интеграции Kafka с внешними системами без написания кода.

```
┌────────────────────────────────────────────────────┐
│              Kafka Connect                         │
├────────────────────────────────────────────────────┤
│                                                    │
│  Source Connectors         Sink Connectors         │
│  (читают из БД)            (пишут в БД)            │
│         │                          ▲               │
│         ▼                          │               │
│  ┌─────────────┐           ┌──────────────┐        │
│  │   MySQL     │           │ Elasticsearch│        │
│  │  Postgres   │           │   MongoDB    │        │
│  │   MongoDB   │           │     S3       │        │
│  └──────┬──────┘           └──────▲───────┘        │
│         │                          │               │
│         └──────────┐      ┌────────┘               │
│                    ▼      │                        │
│              ┌─────────────┐                        │
│              │    Kafka    │                        │
│              │   Topics    │                        │
│              └─────────────┘                        │
└────────────────────────────────────────────────────┘
```

### Примеры Connectors

**Source Connector (чтение из БД):**
```json
{
  "name": "mysql-source-connector",
  "config": {
    "connector.class": "io.debezium.connector.mysql.MySqlConnector",
    "database.hostname": "mysql.example.com",
    "database.port": "3306",
    "database.user": "kafka",
    "database.password": "secret",
    "database.server.id": "1",
    "database.server.name": "mydb",
    "table.include.list": "orders,users",
    "topic.prefix": "mysql"
  }
}

Результат:
• Изменения в таблице orders → топик mysql.orders
• Изменения в таблице users → топик mysql.users
• CDC (Change Data Capture) в реальном времени!
```

**Sink Connector (запись в БД):**
```json
{
  "name": "elasticsearch-sink-connector",
  "config": {
    "connector.class": "io.confluent.connect.elasticsearch.ElasticsearchSinkConnector",
    "topics": "user-events,orders",
    "connection.url": "http://elasticsearch:9200",
    "type.name": "_doc",
    "key.ignore": "true"
  }
}

Результат:
• Все сообщения из user-events → Elasticsearch index
• Поиск в реальном времени!
```

---

## Kafka Streams

### Что такое Kafka Streams?

**Kafka Streams** — библиотека для потоковой обработки данных прямо в вашем приложении.

```
Традиционная обработка:
┌────────┐     ┌──────────┐     ┌────────┐
│ Kafka  │────►│ Spark    │────►│ Kafka  │
│ Input  │     │ Flink    │     │ Output │
└────────┘     │(отдельный│     └────────┘
               │ кластер) │
               └──────────┘
               ❌ Сложная инфраструктура

Kafka Streams:
┌────────┐     ┌──────────────┐     ┌────────┐
│ Kafka  │────►│ Your App     │────►│ Kafka  │
│ Input  │     │ (+KStreams)  │     │ Output │
└────────┘     └──────────────┘     └────────┘
               ✅ Просто! Работает внутри приложения
```

### Пример обработки

```
Задача: Подсчитать количество кликов по URL в окне 1 минута

Input topic "clicks":
{"user": "alice", "url": "/home", "timestamp": 1000}
{"user": "bob", "url": "/home", "timestamp": 1001}
{"user": "alice", "url": "/about", "timestamp": 1002}
{"user": "charlie", "url": "/home", "timestamp": 1060}

Kafka Streams обработка:
1. Группировка по URL
2. Windowing (окно 1 минута)
3. Count

Output topic "click-counts":
{"url": "/home", "window": "00:00-00:01", "count": 2}
{"url": "/about", "window": "00:00-00:01", "count": 1}
{"url": "/home", "window": "00:01-00:02", "count": 1}
```

**Пример в Go (с использованием альтернативной библиотеки):**
```go
// Go не имеет официального Kafka Streams
// Но можно реализовать похожую логику:

func processStream(consumer *sarama.Consumer, producer sarama.SyncProducer) {
    // Читаем из input топика
    partitionConsumer, _ := consumer.ConsumePartition("clicks", 0, sarama.OffsetNewest)
    
    // Агрегация в памяти (упрощенно)
    counts := make(map[string]int)
    ticker := time.NewTicker(1 * time.Minute)
    
    for {
        select {
        case msg := <-partitionConsumer.Messages():
            var click Click
            json.Unmarshal(msg.Value, &click)
            
            // Подсчет
            counts[click.URL]++
            
        case <-ticker.C:
            // Каждую минуту отправляем результаты
            for url, count := range counts {
                result := Result{URL: url, Count: count}
                data, _ := json.Marshal(result)
                
                producer.SendMessage(&sarama.ProducerMessage{
                    Topic: "click-counts",
                    Value: sarama.ByteEncoder(data),
                })
            }
            
            // Сброс счетчиков
            counts = make(map[string]int)
        }
    }
}
```

---

## Практические паттерны

### 1. Event Sourcing

```
Вместо хранения текущего состояния, храним все события:

Традиционная БД:
┌──────────────────────────┐
│ Table: accounts          │
├──────────┬───────────────┤
│ user_id  │ balance       │
├──────────┼───────────────┤
│ alice    │ 1000          │ ← Только текущее состояние
└──────────┴───────────────┘

Event Sourcing (Kafka):
Topic "account-events":
┌────────────────────────────────────────┐
│ {"user": "alice", "event": "deposit",  │
│  "amount": 1000, "timestamp": ...}     │
│ {"user": "alice", "event": "withdraw", │
│  "amount": 200, "timestamp": ...}      │
│ {"user": "alice", "event": "deposit",  │
│  "amount": 200, "timestamp": ...}      │
└────────────────────────────────────────┘

Преимущества:
✅ Полная история изменений
✅ Можно восстановить состояние на любой момент
✅ Аудит бесплатно
✅ Можно replay и пересчитать
```

### 2. CQRS (Command Query Responsibility Segregation)

```
Разделение записи и чтения:

┌─────────────┐
│   Command   │ (Запись)
│   Service   │
└──────┬──────┘
       │
       ▼
┌────────────────┐
│ Kafka (events) │
└───┬────────┬───┘
    │        │
    │        ▼
    │   ┌──────────────┐
    │   │  Projection  │ (Создание read моделей)
    │   │   Service    │
    │   └──────┬───────┘
    │          │
    ▼          ▼
┌──────┐  ┌──────────┐
│ OLTP │  │   OLAP   │
│(Write)│  │  (Read)  │
└──────┘  └──────────┘

Write: Postgres (нормализованная)
Read: Elasticsearch (денормализованная, быстрый поиск)
```

### 3. Saga Pattern (Распределенные транзакции)

```
Задача: Оформить заказ = Резерв товара + Списание денег + Создание доставки
       Каждое действие в своем микросервисе

Choreography-based Saga (через события):

1. Order Service публикует "OrderCreated"
   ↓
2. Inventory Service слушает → резервирует товар
   → публикует "InventoryReserved"
   ↓
3. Payment Service слушает → списывает деньги
   → публикует "PaymentProcessed"
   ↓
4. Shipping Service слушает → создает доставку
   → публикует "ShippingCreated"

Если что-то пошло не так:
Payment Service → "PaymentFailed"
   ↓
Inventory Service слушает → отменяет резерв
Order Service слушает → отменяет заказ
```

### 4. CDC (Change Data Capture)

```
Синхронизация между БД через Kafka:

┌──────────────┐
│   MySQL      │ (source of truth)
│   (OLTP)     │
└──────┬───────┘
       │
       ↓ Debezium Connector
       │ (читает binlog)
       │
┌──────▼───────────────┐
│   Kafka Topics       │
│ • db.users           │
│ • db.orders          │
│ • db.products        │
└──────┬───────────────┘
       │
       ├───────────────┐
       │               │
       ▼               ▼
┌─────────────┐  ┌──────────────┐
│Elasticsearch│  │  Data Lake   │
│ (поиск)     │  │  (аналитика) │
└─────────────┘  └──────────────┘

✅ Real-time синхронизация
✅ Нет изменений в приложении
✅ Полная история изменений
```

### 5. Dead Letter Queue (DLQ)

```
Обработка ошибок:

┌────────────────┐
│  Main Topic    │
└───────┬────────┘
        │
        ▼
┌────────────────┐
│   Consumer     │
└───┬────────┬───┘
    │        │
Успех?     Ошибка?
    │        │
    ▼        ▼
 Commit   Retry 3 раза
            │
         Все еще
         ошибка?
            │
            ▼
    ┌────────────────┐
    │ DLQ Topic      │ ← Проблемные сообщения
    │ (для анализа)  │
    └────────────────┘

Логика в Go:
for message := range messages {
    var retries int
    for retries < 3 {
        err := process(message)
        if err == nil {
            session.MarkMessage(message, "")
            break
        }
        retries++
        time.Sleep(time.Second * time.Duration(retries))
    }
    
    if retries == 3 {
        // Отправить в DLQ
        sendToDLQ(message, err)
        session.MarkMessage(message, "")
    }
}
```

---

## 🎯 Резюме: Ключевые концепции Kafka

```
┌────────────────────────────────────────────────────┐
│ 1. Topics, Partitions, Offsets                     │
│    • Topic = поток событий                         │
│    • Partition = для параллелизма                  │
│    • Offset = позиция в партиции                   │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│ 2. Producers и Consumers                           │
│    • Producer → публикует события                  │
│    • Consumer → читает события                     │
│    • Батчирование для производительности           │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│ 3. Consumer Groups                                 │
│    • Параллельная обработка                        │
│    • Автоматическая балансировка                   │
│    • # consumers ≤ # partitions                    │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│ 4. Replication                                     │
│    • Отказоустойчивость                            │
│    • Leader + Followers                            │
│    • In-Sync Replicas (ISR)                        │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│ 5. Гарантии доставки                               │
│    • At-most-once (быстро, может потерять)         │
│    • At-least-once (надежно, дубликаты)            │
│    • Exactly-once (идеально, сложнее)              │
└────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────┐
│ 6. Экосистема                                      │
│    • Kafka Connect (интеграции)                    │
│    • Kafka Streams (обработка)                     │
│    • Schema Registry (схемы)                       │
└────────────────────────────────────────────────────┘
```

---

## 📚 Практические рекомендации

### Конфигурация для Production

```go
// Producer
config := sarama.NewConfig()
config.Producer.RequiredAcks = sarama.WaitForAll     // Надежность
config.Producer.Retry.Max = 5                        // Повторы
config.Producer.Return.Successes = true
config.Producer.Idempotent = true                    // Без дубликатов
config.Producer.Compression = sarama.CompressionSnappy // Сжатие
config.Producer.Flush.Messages = 100                 // Батчирование
config.Producer.Flush.Frequency = 100 * time.Millisecond

// Consumer
config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
config.Consumer.Offsets.Initial = sarama.OffsetOldest
config.Consumer.Offsets.AutoCommit.Enable = false    // Manual commit
config.Consumer.MaxProcessingTime = 10 * time.Second
```

### Выбор количества партиций

```
┌────────────────────────────────────────────────────┐
│ Формула: # partitions = max(T/P, T/C)              │
│                                                    │
│ T = target throughput (целевая пропускная способность) │
│ P = producer throughput per partition              │
│ C = consumer throughput per partition              │
│                                                    │
│ Пример:                                            │
│ • Нужно обрабатывать 1000 msg/sec                  │
│ • Producer отправляет 500 msg/sec на partition     │
│ • Consumer обрабатывает 100 msg/sec с partition    │
│                                                    │
│ # partitions = max(1000/500, 1000/100) = 10        │
│                                                    │
│ Рекомендации:                                      │
│ • Больше партиций = больше параллелизм             │
│ • Но не слишком много (overhead на rebalance)      │
│ • Начните с 3-6, масштабируйте по необходимости    │
└────────────────────────────────────────────────────┘
```

### Мониторинг

```
Ключевые метрики:
┌────────────────────────────────────────────────────┐
│ Producer:                                          │
│ • request-latency (задержка отправки)              │
│ • record-send-rate (сообщений/сек)                 │
│ • record-error-rate (ошибки)                       │
│ • buffer-available-bytes (свободная память)        │
│                                                    │
│ Consumer:                                          │
│ • records-lag (отставание от последнего offset)    │
│ • records-consumed-rate (сообщений/сек)            │
│ • commit-latency (задержка commit)                 │
│                                                    │
│ Broker:                                            │
│ • under-replicated-partitions (проблемы реплик)    │
│ • offline-partitions (недоступные партиции)        │
│ • active-controller (должен быть 1)                │
│ • request-queue-size (очередь запросов)            │
└────────────────────────────────────────────────────┘
```

---

## Вопросы для самопроверки

1. ✅ Что такое партиция и зачем они нужны?
2. ✅ Как сообщения распределяются по партициям?
3. ✅ Что такое offset и как он используется?
4. ✅ В чем разница между Consumer и Consumer Group?
5. ✅ Как работает rebalancing?
6. ✅ Что такое ISR (In-Sync Replicas)?
7. ✅ В чем разница между at-most-once и at-least-once?
8. ✅ Зачем нужна идемпотентность Producer?
9. ✅ Когда использовать Kafka Connect vs Kafka Streams?
10. ✅ Как выбрать количество партиций для топика?

Если вы понимаете ответы на эти вопросы — вы готовы работать с Kafka! 🎉

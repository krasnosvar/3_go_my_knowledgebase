# Kafka: упражнения на Go

## Зачем

Минимальный продюсер и консьюмер на библиотеке [segmentio/kafka-go](https://github.com/segmentio/kafka-go): запись сообщения в топик и чтение его в составе consumer group.

## Пререквизиты

- Go 1.22+ и Docker (с `docker compose`).
- Запущенный стенд из [../docker_compose/](../docker_compose/README.md): брокер на `localhost:9092`, топик `test-topic` создаётся автоматически.
- Теория: [../interview_questions/theory.md](../interview_questions/theory.md) (топики, партиции, consumer groups, offsets).

## Файлы

`ex1/` — отдельный Go-модуль (`go.mod` внутри):

- [ex1/producer/main.go](./ex1/producer/main.go) — пишет одно сообщение `Key-A` → `Hello Kafka from Go!` в `test-topic`.
- [ex1/consumer/main.go](./ex1/consumer/main.go) — читает `test-topic` в группе `group-1` и печатает сообщения (бесконечный цикл, выход — Ctrl+C).

## Как запустить

```bash
cd ../docker_compose && docker compose up -d && cd ../exercises/ex1

# терминал 1: консьюмер
go run ./consumer

# терминал 2: продюсер (можно запускать несколько раз)
go run ./producer
```

Ожидаемый вывод консьюмера: `received: key=Key-A value=Hello Kafka from Go!`.

## Что попробовать самостоятельно

1. Запустите два консьюмера с одним `GroupID` — при одной партиции сообщения получит только один из них. Увеличьте число партиций (`kafka-topics --alter --partitions 3`) и повторите.
2. Смените `GroupID` у второго консьюмера — теперь каждое сообщение получат оба.
3. Посмотрите lag группы: `kafka-consumer-groups --describe --group group-1` (команда — в [../docker_compose/README.md](../docker_compose/README.md)).

## Когда дальше

После этих упражнений переходите к Redis: [../../4_redis/](../../4_redis/README.md).

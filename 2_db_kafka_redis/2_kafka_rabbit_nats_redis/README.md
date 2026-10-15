# Брокеры сообщений и кэш

## Зачем

Асинхронное взаимодействие сервисов (очереди и логи событий) и быстрый in-memory кэш — две вещи, которые появляются почти в любом Go-бэкенде после базы данных.

**Уже знакомы с темой?** → шпаргалки [Kafka](./1_kafka/CHEATSHEET.md) и [Redis](./4_redis/CHEATSHEET.md).

## Пререквизиты

- Docker (с `docker compose`), Go 1.22+.
- Go: горутины, каналы, `context` ([3_concurrency](../../1_golang/3_concurrency/)).
- Полезно параллельно читать [3_system_design/05_messaging](../../3_system_design/05_messaging/) и [3_system_design/03_caching](../../3_system_design/03_caching/).

## Что внутри

- [1_kafka/](./1_kafka/README.md) — Apache Kafka: топики, партиции, offsets, consumer groups; стенд + Kafka UI, продюсер/консьюмер на Go, теория и вопросы.
- [4_redis/](./4_redis/README.md) — Redis: структуры данных, TTL, персистентность, репликация/cluster, eviction; стенд + UI, пример на Go, теория и вопросы.

**RabbitMQ и NATS пока не покрыты.** Каталог назван с расчётом на них, а номера `2_` и `3_` зарезервированы под эти подразделы — поэтому после `1_kafka` сразу идёт `4_redis`.

## Как проходить

Kafka → Redis. В каждом подразделе: теория в `interview_questions/theory.md` → стенд из `docker_compose/` → упражнения из `exercises/` → вопросы из `interview_questions/`.

## Когда дальше

Аналитическая колоночная БД — [../4_clickhouse/](../4_clickhouse/README.md).

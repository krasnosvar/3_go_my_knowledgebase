# Kafka: учебный стенд + материалы

## Зачем

Apache Kafka — распределённый лог событий: топики и партиции, offsets, consumer groups, репликация. Основа асинхронного взаимодействия сервисов и стриминга данных.

## Пререквизиты

- Docker (с `docker compose`), Go 1.22+.
- Пройденные [1_go_basics](../../../1_golang/1_go_basics/) и [3_concurrency](../../../1_golang/3_concurrency/); желательно — [1_postgres](../../1_postgres/README.md).

## Структура

- [interview_questions/](./interview_questions/README.md) — теория ([theory.md](./interview_questions/theory.md)) и ТОП-20 вопросов с ответами.
- [docker_compose/](./docker_compose/README.md) — стенд Kafka (один брокер в режиме ZooKeeper, образы Confluent 7.4) + Kafka UI; команды для работы с топиками и consumer groups.
- [exercises/](./exercises/README.md) — продюсер и консьюмер на Go (`segmentio/kafka-go`).

## Как проходить

1. Прочитайте теорию: [interview_questions/theory.md](./interview_questions/theory.md).
2. Поднимите стенд: `cd docker_compose && docker compose up -d && docker compose ps`.
3. Выполните упражнения из [exercises/](./exercises/README.md).
4. Ответьте на вопросы из [interview_questions/](./interview_questions/README.md) без подсказок.

## Когда дальше

Кэш и in-memory хранилище — [../4_redis/README.md](../4_redis/README.md).

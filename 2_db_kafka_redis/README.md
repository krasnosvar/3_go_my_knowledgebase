# Базы данных и брокеры сообщений

Этот раздел — учебный набор по хранилищам данных и брокерам сообщений в контексте Go:
теория, учебные стенды на Docker Compose, практика на Go/SQL и вопросы/ответы для самопроверки.

## 🗂 Структура

Каждый подраздел устроен одинаково:
- `docker_compose/` — учебный стенд (поднимается через `docker compose up -d`)
- `exercises/` — практика на Go/SQL от простого к сложному
- `interview_questions/` — вопросы/ответы и теория для самопроверки
- `README.md` — навигация по подразделу

Подразделы:

- `1_postgres/` — реляционная БД, SQL; стенд primary/replica + PgBouncer + pgAdmin
- `2_kafka_rabbit_nats_redis/` — брокеры сообщений и кэш; сейчас есть `1_kafka/` (Kafka на KRaft + Kafka UI) и `4_redis/` (Redis + RedisInsight)
- `4_clickhouse/` — аналитическая колоночная БД (MergeTree, партиционирование, distributed)
- `5_mongodb/` — документная NoSQL-база

## 🚀 Как проходить

1) Выбери подраздел и открой его `README.md`.
2) Подними стенд из `docker_compose/` (`docker compose up -d`, затем `docker compose ps`).
3) Пройди практику из `exercises/` по порядку.
4) Сверяйся с теорией и разбирай `interview_questions/` без подсказок.

## Быстрые ссылки

- PostgreSQL: `1_postgres/README.md`
- Kafka: `2_kafka_rabbit_nats_redis/1_kafka/README.md`
- Redis: `2_kafka_rabbit_nats_redis/4_redis/README.md`
- ClickHouse: `4_clickhouse/README.md`
- MongoDB: `5_mongodb/README.md`

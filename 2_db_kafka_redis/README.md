# Базы данных и брокеры сообщений

## Зачем

Учебный набор по хранилищам данных и брокерам сообщений в контексте Go: теория, учебные стенды на Docker Compose, практика на Go/SQL и вопросы с ответами для самопроверки.

## Пререквизиты

- Пройденный базовый Go: [1_go_basics](../1_golang/1_go_basics/), [3_concurrency](../1_golang/3_concurrency/); желательно [4_standard_library](../1_golang/4_standard_library/) (`net/http`, JSON).
- **Docker** с плагином `docker compose` — все стенды запускаются в контейнерах.
- Go 1.22+. Go-код этого раздела лежит в отдельных модулях (`go.mod` внутри каталога упражнения) и запускается из каталога модуля; зависимости скачиваются при первом `go run` (нужен интернет).

## Структура

Каждый подраздел устроен одинаково:

- `README.md` — навигация по подразделу;
- `docker_compose/` — учебный стенд (`docker compose up -d`) и README с командами;
- `exercises/` — практика на Go/SQL;
- `interview_questions/` — теория (`theory.md` или отдельный конспект) и вопросы с ответами.

Подразделы:

1. [1_postgres/](./1_postgres/README.md) — реляционная БД и SQL; стенд primary/replica + PgBouncer + pgAdmin.
2. [2_kafka_rabbit_nats_redis/](./2_kafka_rabbit_nats_redis/README.md) — брокеры сообщений и кэш: [Kafka](./2_kafka_rabbit_nats_redis/1_kafka/README.md) и [Redis](./2_kafka_rabbit_nats_redis/4_redis/README.md). RabbitMQ и NATS пока не покрыты.
3. [4_clickhouse/](./4_clickhouse/README.md) — аналитическая колоночная БД (MergeTree, партиционирование, distributed).
4. [5_mongodb/](./5_mongodb/README.md) — документная NoSQL-база.

Номер `3_` на верхнем уровне пропущен — нумерация каталогов историческая, порядок прохождения — как в списке выше.

## Как проходить

1. Откройте `README.md` подраздела и прочитайте теорию.
2. Поднимите стенд из `docker_compose/` (`docker compose up -d`, затем `docker compose ps`).
3. Пройдите практику из `exercises/` в порядке, указанном в README подраздела.
4. Ответьте на вопросы из `interview_questions/` без подсказок.

Начинать стоит с PostgreSQL: остальные подразделы опираются на понятия из него (транзакции, индексы, репликация).

## Когда дальше

Связать хранилища и брокеры с архитектурой — [3_system_design](../3_system_design/README.md).

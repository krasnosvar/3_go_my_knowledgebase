# Redis: учебный стенд + материалы

## Зачем

Redis — in-memory хранилище: кэш, счётчики, rate limiting, блокировки, очереди. Конспект покрывает структуры данных, TTL, персистентность (RDB/AOF), репликацию и cluster, политики вытеснения.

## Пререквизиты

- Docker (с `docker compose`), Go 1.22+.
- Пройденные [1_go_basics](../../../1_golang/1_go_basics/) и [3_concurrency](../../../1_golang/3_concurrency/); желательно — [1_postgres](../../1_postgres/README.md).

## Структура

- [interview_questions/](./interview_questions/README.md) — теория ([theory.md](./interview_questions/theory.md)) и ТОП-20 вопросов с ответами.
- [docker_compose/](./docker_compose/README.md) — стенд Redis (AOF) + RedisInsight + Redis Commander.
- [exercises/](./exercises/README.md) — пример на Go (`go-redis/v9`): `PING`, `SET`/`GET`, идеи для самостоятельной практики.

## Как проходить

1. Прочитайте теорию: [interview_questions/theory.md](./interview_questions/theory.md).
2. Поднимите стенд: `cd docker_compose && docker compose up -d && docker compose ps`.
3. Выполните упражнения из [exercises/](./exercises/README.md).
4. Ответьте на вопросы из [interview_questions/](./interview_questions/README.md) без подсказок.

## Когда дальше

Аналитическая колоночная БД — [../../4_clickhouse/README.md](../../4_clickhouse/README.md).

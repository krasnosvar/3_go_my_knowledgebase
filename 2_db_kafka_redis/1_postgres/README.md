# PostgreSQL: учебный стенд + материалы

## Зачем

Реляционная БД, с которой работает большинство Go-бэкендов: SQL, транзакции и уровни изоляции, индексы и `EXPLAIN`, репликация и пулинг соединений. Здесь — теория, учебный стенд на Docker Compose, практика на SQL и Go и вопросы для самопроверки.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — psql, DDL/DML/JOIN, индексы и `EXPLAIN`, уровни изоляции, блокировки, репликация и PgBouncer, ловушки.

## Пререквизиты

- Docker (с `docker compose`).
- Для Go-заданий — Go 1.22+ и пройденные [1_go_basics](../../1_golang/1_go_basics/) и [3_concurrency](../../1_golang/3_concurrency/).

## Структура

- [postgres_sql_theory.md](./postgres_sql_theory.md) — **каноничный конспект теории**: часть 1 — справочник по SQL-командам (DDL, DML, SELECT, JOIN, агрегаты, подзапросы, индексы, транзакции, представления, psql); часть 2 — концепции простым языком (нормальные формы, ACID и уровни изоляции, индексы и селективность, внешние ключи, денормализация, репликация/шардирование/партиционирование, массовые обновления, пагинация).
- [docker_compose/](./docker_compose/README.md) — стенд primary/replica + PgBouncer + pgAdmin: как подключаться, проверять репликацию и PgBouncer.
- [exercises/](./exercises/README.md) — практика:
  - [4_join_tasks/](./exercises/4_join_tasks/README.md) — задачи на JOIN с решениями;
  - [3_indexes_tasks/](./exercises/3_indexes_tasks/README.md) — индексы и оптимизация запросов;
  - [2_script_renew_task/](./exercises/2_script_renew_task/README.md) — массовое обновление 1 млн строк из Go;
  - [1_task_db/](./exercises/1_task_db/README.md) — большое итоговое задание: CRUD-сервис на Go (SQL и GORM v2).
- [interview_questions/](./interview_questions/README.md) — вопросы и разборы ответов.

## Как учиться (рекомендуемый порядок)

1. Прочитайте теорию: [postgres_sql_theory.md](./postgres_sql_theory.md) — сначала часть 1 (синтаксис), затем часть 2 (концепции).
2. Поднимите стенд:

   ```bash
   cd docker_compose
   docker compose up -d
   docker compose ps
   ```

3. Пройдите практику от простого к сложному: `4_join_tasks` → `3_indexes_tasks` → `2_script_renew_task` → `1_task_db`. Задание `1_task_db` самое объёмное (схема, API, генерация данных, аналитика, пагинация) — несмотря на номер, делайте его последним.
4. После каждой практической темы перечитайте соответствующий раздел теории и попробуйте объяснить его своими словами.
5. В конце ответьте на вопросы из [interview_questions/](./interview_questions/README.md) без подсказок.

## Когда дальше

Брокеры сообщений и кэш — [../2_kafka_rabbit_nats_redis/](../2_kafka_rabbit_nats_redis/README.md).

# ClickHouse: учебный стенд + материалы

## Зачем

ClickHouse — колоночная аналитическая (OLAP) СУБД: движки MergeTree, parts и merges, партиционирование, distributed-таблицы. Нужна, когда надо быстро агрегировать миллиарды строк.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — стенд и клиенты, движки MergeTree, ORDER BY/PARTITION BY, батчи и parts, materialized views, ловушки.

## Пререквизиты

- Docker (с `docker compose`), Go 1.22+.
- Пройденные [1_go_basics](../../1_golang/1_go_basics/) и [3_concurrency](../../1_golang/3_concurrency/); желательно — [1_postgres](../1_postgres/README.md).

## Структура

- [interview_questions/](./interview_questions/README.md) — теория ([theory.md](./interview_questions/theory.md)) и ТОП-20 вопросов с ответами.
- [docker_compose/](./docker_compose/README.md) — стенд ClickHouse Server, встроенный веб-UI `/play`, проверка через HTTP.
- [exercises/](./exercises/README.md) — подключение из Go (`clickhouse-go`) и запросы через HTTP-интерфейс.

## Как проходить

1. Прочитайте теорию: [interview_questions/theory.md](./interview_questions/theory.md).
2. Поднимите стенд: `cd docker_compose && docker compose up -d && docker compose ps`.
3. Выполните упражнения из [exercises/](./exercises/README.md).
4. Ответьте на вопросы из [interview_questions/](./interview_questions/README.md) без подсказок.

## Когда дальше

Документная NoSQL-база — [../5_mongodb/README.md](../5_mongodb/README.md).

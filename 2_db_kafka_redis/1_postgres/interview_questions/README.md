# PostgreSQL: вопросы для самопроверки

## Зачем

Вопросы и разборы ответов по SQL и PostgreSQL в формате собеседования. Читайте после теории ([../postgres_sql_theory.md](../postgres_sql_theory.md)) и хотя бы части практики: сначала отвечайте сами, потом сверяйтесь.

## Файлы (по возрастанию сложности)

- [00_databases_sql.md](./00_databases_sql.md) — список вопросов без ответов: частые, средние, продвинутые. Удобно как чек-лист.
- [0_top_questions_answers.md](./0_top_questions_answers.md) — разборы самых частых вопросов: индексы, медленные запросы, ACID, уровни изоляции, GROUP BY/HAVING.
- [1_db_basics_answers.md](./1_db_basics_answers.md) — основы: модели данных, ключи, SQL-запросы, производительность.
- [2_indexes_optimization_answers.md](./2_indexes_optimization_answers.md) — индексы и оптимизация запросов, `EXPLAIN`.
- [3_advanced_concepts_tasks.md](./3_advanced_concepts_tasks.md) — распределённые системы, проблемы параллелизма, практические задачи.

> Уровни изоляции в этих файлах описаны по стандарту SQL; где поведение PostgreSQL отличается (нет грязного чтения, нет фантомов на REPEATABLE READ), это отмечено отдельно.

## Как проходить

Возьмите [00_databases_sql.md](./00_databases_sql.md), ответьте вслух на каждый вопрос, затем найдите разбор в файлах с ответами. Запросы из ответов можно проверять на стенде из [../docker_compose/](../docker_compose/README.md).

## Когда дальше

После PostgreSQL — брокеры сообщений и кэш: [../../2_kafka_rabbit_nats_redis/](../../2_kafka_rabbit_nats_redis/README.md).

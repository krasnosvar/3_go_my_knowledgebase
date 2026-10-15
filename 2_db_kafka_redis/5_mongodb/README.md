# MongoDB: учебный стенд + материалы

## Зачем

MongoDB — документная NoSQL-база: коллекции и документы, индексы, репликация и шардирование. Полезно понимать, чем документная модель отличается от реляционной и когда она уместна.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — CRUD в mongosh и Go, `bson.M`/`bson.D`, типы индексов и ESR, aggregation, embedding vs referencing, read preference и write concern, shard key, ловушки.

## Пререквизиты

- Docker (с `docker compose`), Go 1.22+.
- Пройденные [1_go_basics](../../1_golang/1_go_basics/) и [3_concurrency](../../1_golang/3_concurrency/); желательно — [1_postgres](../1_postgres/README.md).

## Структура

- [interview_questions/](./interview_questions/README.md) — теория ([theory.md](./interview_questions/theory.md)) и ТОП-20 вопросов с ответами.
- [docker_compose/](./docker_compose/README.md) — стенд MongoDB (один контейнер) и подключение через `mongosh`.
- [exercises/](./exercises/README.md) — CRUD из Go на официальном драйвере и вручную в `mongosh`.

## Как проходить

1. Прочитайте теорию: [interview_questions/theory.md](./interview_questions/theory.md).
2. Поднимите стенд: `cd docker_compose && docker compose up -d && docker compose ps`.
3. Выполните упражнения из [exercises/](./exercises/README.md).
4. Ответьте на вопросы из [interview_questions/](./interview_questions/README.md) без подсказок.

## Когда дальше

Связать хранилища с архитектурой — [../../3_system_design/README.md](../../3_system_design/README.md).

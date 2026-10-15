# 1_task_db — мини-сервис магазина на Go + PostgreSQL

## Зачем

Самое большое задание раздела: спроектировать схему (`users`, `orders`, `items`), написать CRUD API на чистом SQL и на ORM, сгенерировать миллионы строк, написать аналитический запрос и скрипт обновления с двумя видами пагинации. Это задание объединяет почти всю теорию — поэтому его лучше делать **последним** (см. порядок в [../README.md](../README.md)).

## Пререквизиты

- Go 1.22+, Docker и стенд из [../../docker_compose/](../../docker_compose/README.md).
- Go: структуры, ошибки, `net/http`, JSON ([1_golang](../../../../1_golang/)).
- Теория: [../../postgres_sql_theory.md](../../postgres_sql_theory.md) — JOIN, агрегаты, транзакции, индексы, пагинация, массовые обновления.

## Файлы

- [task.md](./task.md) — формулировка задания.
- [postgresql_theory.md](./postgresql_theory.md) — список тем к заданию со ссылкой на каноничный конспект.
- [simple_crud_app/](./simple_crud_app/README.md) — учебная реализация CRUD: схема и данные, сервис на `database/sql` и на GORM, команды для проверки API.
  - [simple_crud_app/task-crud.md](./simple_crud_app/task-crud.md) — пошаговый гайд «как написать сервис самому».

## Как проходить

1. Прочитайте [task.md](./task.md) и попробуйте спроектировать схему сами.
2. Идите по гайду [task-crud.md](./simple_crud_app/task-crud.md), сверяясь с готовым кодом только при затруднениях.
3. Запустите готовый сервис по [simple_crud_app/README.md](./simple_crud_app/README.md) и сравните со своим.
4. Часть задания про массовое обновление разобрана в [../2_script_renew_task/](../2_script_renew_task/README.md).

## Когда дальше

Вопросы для самопроверки — [../../interview_questions/](../../interview_questions/README.md).

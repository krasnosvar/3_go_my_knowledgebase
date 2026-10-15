# Практика по индексам PostgreSQL

SQL-скрипты для изучения разных типов индексов PostgreSQL на таблице в 1 млн строк.

## Пререквизиты

- Docker и запущенный стенд из [../../docker_compose/](../../docker_compose/README.md).
- Теория: раздел «Индексы» в [../../postgres_sql_theory.md](../../postgres_sql_theory.md) (часть 1, п. 7, и часть 2, п. 5).

## Файлы

- [setup.sql](./setup.sql) — создаёт таблицу `users_index_test` и наполняет её **1 000 000** строк случайных данных (включая поля-массивы и текст).
- [practice.sql](./practice.sql) — набор команд для создания и сравнения индексов:
    - **B-tree** — базовый поиск и сравнение;
    - **Hash** — поиск по точному совпадению;
    - **GIN** — поиск в массивах (`tags`) и полнотекстовый поиск;
    - **GiST** — полнотекстовый поиск;
    - **Composite** — индекс по нескольким полям (`age`, `status`);
    - **Partial** — индексация только части таблицы.

## Как использовать

1. **Инициализация данных** (из каталога `1_postgres/docker_compose`):

   ```bash
   docker compose exec -T db-primary psql -U postgres -d postgres -v ON_ERROR_STOP=1 < ../exercises/3_indexes_tasks/setup.sql
   ```

   Проверка:

   ```bash
   docker exec postgres-primary psql -U postgres -c "SELECT * FROM users_index_test ORDER BY id LIMIT 10;"
   ```

2. **Эксперименты с индексами.** Откройте `practice.sql` и выполняйте команды по очереди в SQL-клиенте (`psql`, DBeaver, pgAdmin). Сравнивайте вывод `EXPLAIN ANALYZE` до и после создания индекса.

   Интерактивный `psql` в контейнере:

   ```bash
   docker exec -it postgres-primary psql -U postgres
   ```

> Полнотекстовый поиск в примерах настроен на английский язык (`'english'`).

## Когда дальше

После индексов переходите к [../1_task_db/](../1_task_db/README.md) — там индексы, транзакции и пагинация понадобятся в реальном сервисе.

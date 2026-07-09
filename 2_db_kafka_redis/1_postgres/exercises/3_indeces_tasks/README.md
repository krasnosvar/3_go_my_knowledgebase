docker compose exec -T db-primary psql -U postgres -d postgres -v ON_ERROR_STOP=1 < ../exercises/3_indeces_tasks/setup.sql


docker exec postgres-primary psql -U postgres -c "SELECT * FROM users_index_test ORDER BY id LIMIT 10;"


# PostgreSQL Indexing Practice

Этот проект содержит SQL скрипты для изучения работы различных типов индексов в PostgreSQL.

## Структура
- `setup.sql`: Создает таблицу `users_index_test` и наполняет ее **1,000,000** строк случайных данных. Добавлены поля для массивов и текста.
- `practice.sql`: Набор SQL команд для создания и тестирования индексов:
    - **B-tree**: Базовый поиск и сравнение.
    - **Hash**: Ускоренный поиск по точному совпадению.
    - **GIN**: Поиск в массивах (`tags`) и полнотекстовый поиск.
    - **GiST**: Полнотекстовый поиск.
    - **Composite**: Индекс по нескольким полям (`age`, `status`).
    - **Partial**: Индексация только части таблицы.

## Как использовать

1. **Инициализация БД**:
   Выполните SQL скрипт:
   ```bash
   psql -U postgres -f setup.sql
   ```

2. **Эксперименты с индексами**:
   Откройте `practice.sql` и копируйте команды по очереди в ваш SQL клиент (psql, DBeaver, pgAdmin и т.д.).
   Следите за выводом `EXPLAIN ANALYZE`, чтобы увидеть разницу в производительности.

---
*Примечание: Полнотекстовый поиск в примерах настроен на английский язык (`'english'`).*

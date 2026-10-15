# PostgreSQL — шпаргалка

Подробности и объяснения — в [postgres_sql_theory.md](./postgres_sql_theory.md) и
[interview_questions/](./interview_questions/README.md). ⚠️ — частая ловушка.
Конспект рассчитан на PostgreSQL 14+. Команды CLI/SQL не запускались — проверяйте
на стенде из [docker_compose/](./docker_compose/README.md).

## Стенд и psql

Подробно: [docker_compose/README.md](./docker_compose/README.md), [postgres_sql_theory.md](./postgres_sql_theory.md#psql).

| Задача | Команда |
| :--- | :--- |
| Поднять / статус / снести с данными | `docker compose up -d` / `docker compose ps` / `docker compose down -v` |
| psql на primary / replica | `docker compose exec -T db-primary psql -U postgres -d postgres` (`db-replica`) |
| Через PgBouncer (`:6432`, алиасы `primary`/`replica`) | `psql "postgres://postgres:admin@localhost:6432/primary"` |
| Одна команда, падать на ошибке | `psql -U postgres -d shop -v ON_ERROR_STOP=1 -c "SELECT now();"` |
| Выполнить файл | `psql ... -v ON_ERROR_STOP=1 < ./schema.sql` или `\i файл.sql` |
| Бэкап / восстановление | `pg_dump -U user -d db -f backup.sql` / `psql -U user -d db -f backup.sql`; все БД: `pg_dumpall` |

| Мета-команда | Что делает |
| :--- | :--- |
| `\l` / `\c db` | список БД / подключиться |
| `\dt` / `\d table` | таблицы / структура таблицы |
| `\di` / `\dv` / `\du` | индексы / представления / роли |
| `\x` / `\timing` | построчный вывод / замер времени |
| `\q` | выход |

## Типы и DDL

Подробно: [postgres_sql_theory.md → DDL](./postgres_sql_theory.md#ddl), [1_db_basics_answers.md](./interview_questions/1_db_basics_answers.md).

| Нужно | Тип |
| :--- | :--- |
| Автоинкремент PK | `SERIAL` (int, ~2.1 млрд) / `BIGSERIAL` (bigint) |
| Глобально уникальный id | `UUID DEFAULT gen_random_uuid()` (встроена с PG 13) |
| Деньги, точные дроби | `NUMERIC(10,2)`; `REAL`/`FLOAT` — приблизительные |
| Строка | `TEXT` / `VARCHAR(n)`; `CHAR(n)` — дополняется пробелами |
| Время | `TIMESTAMPTZ` (хранится в UTC) — выбор по умолчанию; `TIMESTAMP` — без пояса; `DATE` |
| JSON с индексами | `JSONB` (бинарный, GIN-индекс); `JSON` — просто текст |
| Перечисление / массив / бинарные | `CREATE TYPE s AS ENUM (...)` / `TEXT[]` / `BYTEA` |

| Задача | SQL |
| :--- | :--- |
| Создать таблицу | `CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY, email TEXT UNIQUE NOT NULL, created_at TIMESTAMPTZ DEFAULT now());` |
| Внешний ключ | `role_id INT REFERENCES roles(id) ON DELETE SET NULL` (или `CASCADE`) |
| Добавить / удалить столбец | `ALTER TABLE users ADD COLUMN phone VARCHAR(20);` / `DROP COLUMN phone` |
| Сменить тип / переименовать | `ALTER COLUMN age TYPE BIGINT` / `RENAME COLUMN a TO b` / `RENAME TO customers` |
| Удалить таблицу с зависимостями | `DROP TABLE IF EXISTS users CASCADE;` |
| Представление / материализованное | `CREATE VIEW v AS SELECT ...;` / `CREATE MATERIALIZED VIEW mv AS ...;` + `REFRESH MATERIALIZED VIEW CONCURRENTLY mv;` (нужен UNIQUE-индекс) |
| Партиционирование | `CREATE TABLE orders (...) PARTITION BY RANGE (created_at);`<br>`CREATE TABLE orders_2024_01 PARTITION OF orders FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');` |

## DML и SELECT

Подробно: [DML](./postgres_sql_theory.md#dml), [SELECT](./postgres_sql_theory.md#select), [агрегаты](./postgres_sql_theory.md#aggregate), [подзапросы](./postgres_sql_theory.md#subquery).

| Задача | SQL |
| :--- | :--- |
| Вставить несколько строк и вернуть id | `INSERT INTO users (name, email) VALUES ('a','a@x'), ('b','b@x') RETURNING id;` |
| Обновить / удалить по условию | `UPDATE users SET age = 31 WHERE id = 1;` / `DELETE FROM users WHERE age < 18;` |
| Очистить таблицу быстро | `TRUNCATE TABLE users;` (+ `RESTART IDENTITY` — сбросить последовательности) |
| Фильтры | `BETWEEN 20 AND 35`, `IN (...)`, `LIKE 'Ив%'`, `ILIKE '%ван%'` (без регистра), `IS NOT NULL` |
| Уникальные значения / псевдоним | `SELECT DISTINCT age` / `name AS full_name` |
| Группировка с фильтром групп | `SELECT age, COUNT(*) FROM users GROUP BY age HAVING COUNT(*) > 2;` |
| CTE / рекурсивный CTE | `WITH t AS (SELECT ...) SELECT * FROM t;` / `WITH RECURSIVE n AS (SELECT 1 AS n UNION ALL SELECT n+1 FROM n WHERE n < 10) SELECT n FROM n;` |
| Текущее время | `now()` / `CURRENT_TIMESTAMP`, `CURRENT_DATE` |
| Оценка числа строк без `COUNT(*)` | `SELECT reltuples::bigint FROM pg_class WHERE relname = 'users';` |
| Keyset-пагинация | `SELECT * FROM users WHERE id > $last ORDER BY id LIMIT 20;` |

| | WHERE | HAVING |
| :--- | :--- | :--- |
| Когда | до `GROUP BY`, фильтрует строки | после `GROUP BY`, фильтрует группы |
| Агрегаты (`COUNT`, `AVG`) | нельзя | можно |

| JOIN | Результат |
| :--- | :--- |
| `INNER JOIN` (`JOIN`) | только строки с парой в обеих таблицах |
| `LEFT JOIN` | все слева + пары справа, иначе `NULL` |
| `RIGHT JOIN` | все справа + пары слева |
| `FULL JOIN` | все строки обеих таблиц |

| Агрегат | Нюанс |
| :--- | :--- |
| `COUNT(*)` / `COUNT(col)` | все строки / без `NULL` |
| `SUM`, `AVG`, `MIN`, `MAX` | игнорируют `NULL` |

## Индексы

Подробно: [индексы (часть 2)](./postgres_sql_theory.md#5-индексы--ускорение-поиска), [2_indexes_optimization_answers.md](./interview_questions/2_indexes_optimization_answers.md).

| Тип | Когда | SQL |
| :--- | :--- | :--- |
| B-tree (по умолчанию) | `=`, `<`, `>`, `BETWEEN`, `ORDER BY`, префиксный `LIKE 'abc%'` | `CREATE INDEX idx ON users(email);` |
| Hash | только `=` | `USING hash(phone)` |
| GIN | массивы, `JSONB`, полнотекст | `USING GIN(tags)` |
| GiST | геоданные, полнотекст, ближайшие точки | `USING GIST(location)` |
| BRIN | гигантские таблицы с естественным порядком (время вставки) | `USING BRIN(created_at)` |
| SP-GiST | неравномерно распределённые данные | `USING SPGIST(...)` |
| Составной | несколько условий в `WHERE`; работает по левому префиксу | `ON users(city, age)` |
| Частичный | частые запросы к подмножеству | `ON users(email) WHERE is_active` |
| Функциональный | условие с функцией над столбцом | `ON users(LOWER(name))` |
| Уникальный частичный | уникальность среди «живых» | `CREATE UNIQUE INDEX ... ON users(email) WHERE deleted_at IS NULL;` |

| Индекс НЕ используется, если | Что делать |
| :--- | :--- |
| низкая селективность (`gender`, bool) | не индексировать / частичный индекс |
| функция над столбцом: `WHERE LOWER(name) = ...` | функциональный индекс |
| пропущен первый столбец составного: `(city, age)` и `WHERE age > 25` | переставить столбцы / отдельный индекс |
| несовпадение типов (`text = 123`) | привести тип параметра |
| `NOT IN`, отрицания; крошечная таблица | норма — Seq Scan дешевле |

Список индексов: `\di` или `SELECT indexname, indexdef FROM pg_indexes WHERE tablename = 'users';`.
Использование: `SELECT * FROM pg_stat_user_indexes;`.

## EXPLAIN и медленные запросы

Подробно: [0_top_questions_answers.md → EXPLAIN](./interview_questions/0_top_questions_answers.md#13-explain-и-explain-analyze), [2_indexes_optimization_answers.md](./interview_questions/2_indexes_optimization_answers.md#оптимизация-запросов).

| Команда | Что даёт |
| :--- | :--- |
| `EXPLAIN SELECT ...` | план и оценки, запрос **не** выполняется |
| `EXPLAIN ANALYZE SELECT ...` | реально выполняет, фактические время и строки |
| `EXPLAIN (ANALYZE, BUFFERS) ...` | + чтения из кэша/диска |
| `EXPLAIN (FORMAT JSON) ...` | план в JSON |
| `ANALYZE table;` | обновить статистику планировщика (`pg_stats`) |

| В плане | Значит |
| :--- | :--- |
| `Seq Scan` | полный проход по таблице |
| `Index Scan` / `Bitmap Heap Scan` | доступ через индекс |
| `cost=..` / `rows=` | оценка стоимости / ожидаемые строки |
| `actual time` / `loops` | реальное время шага / сколько раз повторялся |
| оценка `rows` ≪ или ≫ actual | устаревшая статистика → `ANALYZE` |

Порядок разбора: `EXPLAIN ANALYZE` → недостающие/лишние индексы → переписать запрос
(без `SELECT *`, `EXISTS` вместо `IN`, без функций над столбцом) → схема и `ANALYZE`.
Топ медленных: `SELECT * FROM pg_stat_statements ORDER BY total_exec_time DESC LIMIT 10;` (PG 13+; раньше — `total_time`).

## Транзакции и изоляция

Подробно: [транзакции](./postgres_sql_theory.md#transactions), [уровни изоляции](./postgres_sql_theory.md#4-транзакции--всё-или-ничего), [3_advanced_concepts_tasks.md](./interview_questions/3_advanced_concepts_tasks.md#проблемы-параллелизма).

| Задача | SQL |
| :--- | :--- |
| Транзакция | `BEGIN; ... COMMIT;` / `ROLLBACK;` |
| Частичный откат | `SAVEPOINT sp; ... ROLLBACK TO SAVEPOINT sp;` |
| Уровень изоляции | `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;` (первой командой в транзакции) |

| Уровень (PostgreSQL) | Dirty read | Non-repeatable | Phantom | Lost update / write skew |
| :--- | :--- | :--- | :--- | :--- |
| READ UNCOMMITTED | = READ COMMITTED | | | |
| READ COMMITTED (по умолчанию; снимок на каждый запрос) | нет | есть | есть | есть |
| REPEATABLE READ (снимок на транзакцию) | нет | нет | нет (строже стандарта) | write skew есть; конфликт UPDATE → `could not serialize access` |
| SERIALIZABLE (SSI) | нет | нет | нет | нет; цена — ошибки сериализации и повторы |

ACID: Atomicity (всё или ничего), Consistency (правила соблюдены), Isolation, Durability (после `COMMIT` переживёт сбой).

## Блокировки и конкуренция

Подробно: [0_top_questions_answers.md → SELECT FOR UPDATE](./interview_questions/0_top_questions_answers.md#14-select-for-update), [3_advanced_concepts_tasks.md](./interview_questions/3_advanced_concepts_tasks.md#проблемы-параллелизма).

| Задача | Решение |
| :--- | :--- |
| Пессимистично захватить строку | `SELECT ... WHERE id = 5 FOR UPDATE;` — держится до конца транзакции |
| Не ждать, а сразу ошибка | `FOR UPDATE NOWAIT` |
| Очередь задач: взять незанятые | `SELECT ... WHERE status = 'new' FOR UPDATE SKIP LOCKED LIMIT 10;` |
| Оптимистичная блокировка | `UPDATE p SET name = $1, version = version + 1 WHERE id = $2 AND version = $3;` → 0 строк = конфликт, повторить |
| Инкремент без гонки | `UPDATE t SET counter = counter + 1` вместо read-modify-write |
| «Проверил и вставил» | `UNIQUE`-ограничение + `INSERT` (обработать ошибку уникальности) |
| Избежать deadlock | захватывать строки в одном порядке, таймауты; PG сам находит и рвёт deadlock |
| Активные запросы / убить | `SELECT * FROM pg_stat_activity;` / `SELECT pg_terminate_backend(pid);` |

MVCC: UPDATE создаёт новую версию строки (`xmin`/`xmax`), старая становится dead tuple →
`VACUUM` (не блокирует) / `VACUUM FULL` (эксклюзивная блокировка, переписывает таблицу); autovacuum включён по умолчанию.

## Репликация, пулинг, масштабирование

Подробно: [масштабируемость](./postgres_sql_theory.md#8-масштабируемость--работа-с-большими-данными), [docker_compose/README.md](./docker_compose/README.md#health-checks).

| Задача | Команда |
| :--- | :--- |
| Реплики на primary | `SELECT application_name, state, sync_state, replay_lag FROM pg_stat_replication;` |
| Это standby? | `SELECT pg_is_in_recovery();` + `SELECT status FROM pg_stat_wal_receiver;` |
| Лаг в байтах | `SELECT pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), replay_lsn)) FROM pg_stat_replication;` |
| Физическая реплика с нуля | `pg_basebackup -h primary -U replicator -D <datadir> -P -R --slot=replica_slot` |
| PgBouncer: пулы / статистика | `psql "postgres://admin:admin@localhost:6432/pgbouncer" -c "SHOW POOLS;"` (`SHOW STATS;`, `SHOW DATABASES;`) |
| Размер таблицы с индексами | `SELECT pg_size_pretty(pg_total_relation_size('users'));` |

| Проблема | Решение |
| :--- | :--- |
| Много чтения, мало записи | реплики (чтение — с replica, запись — на primary) |
| Огромная таблица, запросы по периоду | партиционирование, удаление старых партиций |
| Одна БД не вытягивает объём / запись | шардирование (JOIN между шардами — сложно) |
| Много соединений от сервисов | PgBouncer |
| Массовый `UPDATE` 100 млн строк | батчами по 10 000 + пауза (`pg_sleep`) |
| Частые тяжёлые JOIN на чтение | денормализация / материализованное представление |
| Глубокая пагинация | keyset вместо `OFFSET` |

| Репликация | Подтверждение записи |
| :--- | :--- |
| синхронная | после записи на реплику |
| асинхронная | сразу после primary (возможен лаг и потеря хвоста при сбое) |
| логическая | переносит изменения строк, а не WAL-страницы |

## Ловушки

- ⚠️ `UPDATE`/`DELETE` без `WHERE` меняет всю таблицу.
- ⚠️ `OFFSET 100000` читает и выбрасывает все пропущенные строки — используйте keyset.
- ⚠️ В PostgreSQL `TRUNCATE` транзакционный (откатывается внутри `BEGIN`), но не запускает `ON DELETE`-триггеры и требует право `TRUNCATE`; последовательности сбрасывает только с `RESTART IDENTITY`.
- ⚠️ `CHAR(n)` в PostgreSQL не быстрее `VARCHAR`/`TEXT` — все три хранятся одинаково.
- ⚠️ `FOR UPDATE` по умолчанию ждёт; `NOWAIT` — сразу ошибка при занятой строке; `SKIP LOCKED` — пропуск занятых строк.
- ⚠️ `INSERT ... WHERE NOT EXISTS (...)` без `UNIQUE` не защищает от гонки при параллельных вставках.
- ⚠️ REPEATABLE READ/SERIALIZABLE требуют retry-цикла в приложении на ошибку сериализации (SQLSTATE `40001`).
- ⚠️ Каждый индекс замедляет `INSERT/UPDATE/DELETE` и занимает место; неиспользуемые — удалять.
- ⚠️ Случайный UUID как PK: 16 байт против 4/8, фрагментация B-tree.
- ⚠️ `VACUUM FULL` блокирует таблицу эксклюзивно — не запускать в час пик.
- ⚠️ `recovery.conf`/`standby_mode` удалены в PG 12, `wal_keep_segments` заменён на `wal_keep_size` в PG 13.
- ⚠️ `SET enable_seqscan = off` — только для диагностики, не в проде.

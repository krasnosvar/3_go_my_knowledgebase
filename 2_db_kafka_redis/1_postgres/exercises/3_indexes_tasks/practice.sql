-- practice.sql
-- Команды для тренировки индексов. Скопируйте и запустите в psql или DBeaver.

-------------------------------------------------------------------------------
-- 1. B-TREE (Стандартный индекс)
-- Лучше всего подходит для сравнений (<, <=, =, >=, >) и сортировки.
-------------------------------------------------------------------------------

-- Поиск без индекса (Seq Scan)
-- docker exec postgres-primary psql -U postgres -c "EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE email = 'user_500000@example.com';"
--                                                            QUERY PLAN                                                            
-- ---------------------------------------------------------------------------------------------------------------------------------
--  Gather  (cost=1000.00..29155.43 rows=1 width=150) (actual time=14.831..28.326 rows=1 loops=1)
--    Workers Planned: 2
--    Workers Launched: 2
--    ->  Parallel Seq Scan on users_index_test  (cost=0.00..28155.33 rows=1 width=150) (actual time=19.758..23.495 rows=0 loops=3)
--          Filter: (email = 'user_500000@example.com'::text)
--          Rows Removed by Filter: 333333
--  Planning Time: 0.557 ms
--  Execution Time: 28.430 ms
-- (8 rows)
EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE email = 'user_500000@example.com';

-- Создаем B-tree индекс
-- docker exec postgres-primary psql -U postgres -c "CREATE INDEX idx_users_email_btree ON users_index_test USING btree (email);"
-- CREATE INDEX
CREATE INDEX idx_users_email_btree ON users_index_test USING btree (email);



-- Поиск с индексом (Index Scan)
-- docker exec postgres-primary psql -U postgres -c "EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE email = 'user_500000@example.com';"
--                                                                 QUERY PLAN                                                                
-- ------------------------------------------------------------------------------------------------------------------------------------------
--  Index Scan using idx_users_email_btree on users_index_test  (cost=0.42..8.44 rows=1 width=150) (actual time=0.024..0.024 rows=1 loops=1)
--    Index Cond: (email = 'user_500000@example.com'::text)
--  Planning Time: 0.390 ms
--  Execution Time: 0.073 ms
-- (4 rows)
EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE email = 'user_500000@example.com';


-------------------------------------------------------------------------------
-- 2. HASH (Индекс для точного совпадения)
-- Работает ТУЛЬКО с оператором '=', обычно быстрее B-tree для равенства,
-- но не поддерживает диапазоны и сортировку.
-------------------------------------------------------------------------------

-- Создаем Hash индекс на username
CREATE INDEX idx_users_username_hash ON users_index_test USING hash (username);

-- Поиск по равенству (должен использовать Hash Index)
-- docker exec postgres-primary psql -U postgres -c "EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE username = 'user_700000';"
--                                                                  QUERY PLAN                                                                 
-- --------------------------------------------------------------------------------------------------------------------------------------------
--  Index Scan using idx_users_username_hash on users_index_test  (cost=0.00..8.02 rows=1 width=150) (actual time=0.024..0.024 rows=1 loops=1)
--    Index Cond: (username = 'user_700000'::text)
--  Planning Time: 0.394 ms
--  Execution Time: 0.063 ms
-- (4 rows)
EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE username = 'user_700000';


-------------------------------------------------------------------------------
-- 3. GIN (Generalized Inverted Index)
-- Идеален для массивов (Arrays), полнотекстового поиска (Full Text Search) и JSONB.
-------------------------------------------------------------------------------

-- Поиск в массиве без индекса
-- docker exec postgres-primary psql -U postgres -c "EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE tags @> ARRAY['admin'];"
--                                                          QUERY PLAN                                                         
-- ----------------------------------------------------------------------------------------------------------------------------
--  Seq Scan on users_index_test  (cost=0.00..35447.00 rows=250000 width=150) (actual time=0.044..139.575 rows=249654 loops=1)
--    Filter: (tags @> '{admin}'::text[])
--    Rows Removed by Filter: 750346
--  Planning Time: 0.956 ms
--  Execution Time: 143.845 ms
-- (5 rows)
EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE tags @> ARRAY['admin'];

-- Создаем GIN индекс на массив tags
-- docker exec postgres-primary psql -U postgres -c "CREATE INDEX idx_users_tags_gin ON users_index_test USING gin (tags);"
CREATE INDEX idx_users_tags_gin ON users_index_test USING gin (tags);

-- Поиск в массиве с индексом
-- docker exec postgres-primary psql -U postgres -c "EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE tags @> ARRAY['admin'];"
--                                                                 QUERY PLAN                                                                
-- ------------------------------------------------------------------------------------------------------------------------------------------
--  Bitmap Heap Scan on users_index_test  (cost=1696.39..27768.39 rows=250000 width=150) (actual time=15.766..82.028 rows=249654 loops=1)
--    Recheck Cond: (tags @> '{admin}'::text[])
--    Heap Blocks: exact=22946
--    ->  Bitmap Index Scan on idx_users_tags_gin  (cost=0.00..1633.89 rows=250000 width=0) (actual time=13.590..13.591 rows=249654 loops=1)
--          Index Cond: (tags @> '{admin}'::text[])
--  Planning Time: 1.310 ms
--  Execution Time: 86.351 ms
-- (7 rows)
EXPLAIN ANALYZE SELECT * FROM users_index_test WHERE tags @> ARRAY['admin'];


-------------------------------------------------------------------------------
-- 4. GiST (Generalized Search Tree)
-- Используется для гео-данных, диапазонов и полнотекстового поиска.
-- В данном случае попробуем применить для поиска по префиксу или полнотекстового поиска.
-------------------------------------------------------------------------------

-- Создаем GiST индекс для полнотекстового поиска на bio
-- Для GiST часто используют расширение pg_trgm (триграммы), если нужен поиск по LIKE
-- Но стандартно GiST хорош для tsvector

CREATE INDEX idx_users_bio_gist ON users_index_test USING gist (to_tsvector('english', bio));

-- Что такое полнотекстовый поиск:
-- Полнотекстовый поиск (Full Text Search, FTS) — это продвинутый механизм поиска текста, 
-- который, в отличие от простого LIKE '%слово%', понимает правила языка. Он умеет 
-- находить слова с учетом морфологии (разные окончания, формы слова) и игнорирует 
-- малозначимые слова (стоп-слова: предлоги, союзы).
--
-- Как это работает в PostgreSQL:
-- 1. to_tsvector('язык', text) — разбивает исходный текст на слова, удаляет стоп-слова 
--    и приводит оставшиеся к их базовой (нормальной) форме — лексемам. 
--    Например, слово 'running' (бег/бегающий) превратится в лексему 'run'.
-- 2. to_tsquery('язык', 'query') — аналогичным образом преобразует поисковый запрос 
--    в лексемы. Поддерживает логические операторы: & (И), | (ИЛИ), ! (НЕ).
-- 3. Оператор @@ — проверяет, содержит ли документ (tsvector) искомые лексемы (tsquery).
--
-- GiST индекс в данном случае работает как сигнатурное дерево: он компактно хранит 
-- хеши лексем и позволяет очень быстро отсекать строки, где нужных слов точно нет.

-- Запрос с использованием полнотекстового поиска через GiST
EXPLAIN ANALYZE 
SELECT * FROM users_index_test 
WHERE to_tsvector('english', bio) @@ to_tsquery('english', 'running');


-------------------------------------------------------------------------------
-- 5. МУЛЬТИКОЛОНОЧНЫЙ ИНДЕКС (Composite Index)
-------------------------------------------------------------------------------

-- Поиск по двум полям
EXPLAIN ANALYZE SELECT count(*) FROM users_index_test WHERE age = 25 AND status = 'active';

-- Создаем индекс
CREATE INDEX idx_users_age_status ON users_index_test (age, status);

-- Повторный поиск
EXPLAIN ANALYZE SELECT count(*) FROM users_index_test WHERE age = 25 AND status = 'active';


-------------------------------------------------------------------------------
-- 6. ЧАСТИЧНЫЙ ИНДЕКС (Partial Index)
-- Индексирует только часть строк, что экономит место.
-------------------------------------------------------------------------------

-- Индекс только для заблокированных пользователей
CREATE INDEX idx_users_blocked_only ON users_index_test (age) WHERE status = 'blocked';

-- Поиск, попадающий в условие индекса
EXPLAIN ANALYZE SELECT count(*) FROM users_index_test WHERE age > 60 AND status = 'blocked';

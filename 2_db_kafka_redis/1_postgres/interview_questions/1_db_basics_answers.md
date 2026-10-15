# Основы баз данных

## Теоретические основы
- **Что такое СУБД? Какие бывают типы?**
  СУБД (система управления базами данных) — это программный комплекс, отвечающий за создание баз, работу с ними и их администрирование.

  Как принято классифицировать СУБД:
  1. **Реляционные (RDBMS)** — таблицы с фиксированной схемой (PostgreSQL, MySQL, Oracle)
  2. **Документоориентированные** — данные хранятся документами в JSON/XML (MongoDB, CouchDB)
  3. **Колоночные** — организация хранения по столбцам (ClickHouse, Cassandra)
  4. **Ключ-значение** — простые пары «ключ → значение» (Redis, DynamoDB)
  5. **Графовые** — заточены под графовые структуры (Neo4j, ArangoDB)
  6. **In-memory** — держат данные в оперативной памяти (Redis, Memcached)
  7. **Временные ряды** — специализируются на данных с меткой времени (InfluxDB, TimescaleDB)

- **Что такое транзакция?**
  Транзакция — это набор операций над базой, который трактуется как одно неделимое действие. Она либо применяется целиком, либо не применяется вовсе — в этом суть атомарности.

  Пример в SQL:
  ```sql
  BEGIN;
  UPDATE accounts SET balance = balance - 100 WHERE id = 1;
  UPDATE accounts SET balance = balance + 100 WHERE id = 2;
  COMMIT;
  ```

- **Что такое ACID?**
  ACID — это четыре гарантии надёжности транзакций:
  - **A**tomicity (атомарность) — либо вся транзакция целиком, либо ничего
  - **C**onsistency (согласованность) — база переходит из одного корректного состояния в другое
  - **I**solation (изолированность) — параллельные транзакции не влияют друг на друга
  - **D**urability (долговечность) — успешно завершённая транзакция переживает сбой

- **Уровни изоляции транзакций**
  1. **Read Uncommitted (чтение неподтверждённых данных)**
     - видны ещё не зафиксированные изменения соседних транзакций
     - быстро, но высок риск аномалий

  2. **Read Committed (чтение подтверждённых данных)**
     - видны только уже закоммиченные изменения
     - остаются неповторяющееся и фантомное чтение

  3. **Repeatable Read (повторяемое чтение)**
     - гарантирует неизменность прочитанных в транзакции данных
     - фантомное чтение по-прежнему возможно

  4. **Serializable (сериализуемый)**
     - самый жёсткий уровень
     - транзакции выглядят так, будто выполнялись строго по очереди
     - за счёт блокировок теряется производительность

  > Это классификация по стандарту SQL. **В PostgreSQL иначе:** Read Uncommitted работает как Read Committed; Repeatable Read реализован как snapshot isolation и фантомов не допускает, остаётся write skew; Serializable (SSI) не блокирует чтения, а при конфликте откатывает транзакцию с ошибкой сериализации. Подробнее — [postgres_sql_theory.md](../postgres_sql_theory.md#concepts).

- **Что такое внешний ключ?**
  Внешний ключ (FOREIGN KEY) — ограничение, поддерживающее ссылочную целостность между таблицами.

  На что нужно обращать внимание:
  - ссылаться можно на PRIMARY KEY или на столбец с ограничением UNIQUE
  - типы связываемых полей должны быть одинаковыми

  Пример:
  ```sql
  CREATE TABLE orders (
      id SERIAL PRIMARY KEY,
      user_id INTEGER REFERENCES users(id),
      amount DECIMAL(10,2)
  );
  ```

- **Что такое constraint?**
  Ограничения (constraints) — это правила, которым обязаны подчиняться данные в таблице:
  - `PRIMARY KEY` — первичный ключ
  - `FOREIGN KEY` — внешний ключ
  - `UNIQUE` — уникальность значений
  - `CHECK` — проверка выполнения условия
  - `NOT NULL` — запрет пустых значений
  - `DEFAULT` — значение по умолчанию

- **Уникальный ключ (UNIQUE)**
  - не допускает повторов в столбце или наборе столбцов
  - в отличие от PRIMARY KEY, разрешает NULL
  - таблица может иметь сразу несколько ограничений UNIQUE

  Пример:
  ```sql
  CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      email VARCHAR(255) UNIQUE,
      username VARCHAR(50) UNIQUE
  );
  ```

- **Чем отличается Primary key от Foreign key?**
  | Характеристика | PRIMARY KEY | FOREIGN KEY |
  |----------------|-------------|-------------|
  | Назначение | Однозначно идентифицирует строку | Организует связь между таблицами |
  | NULL значения | Запрещены | Разрешены (если явно не запретить) |
  | Количество в таблице | Ровно один (допускается составной) | Можно объявить несколько |
  | Индексирование | Индекс создаётся автоматически | Индекс создаётся автоматически |
  | Ссылки | На него ссылаются другие таблицы | Сам ссылается на PRIMARY KEY или UNIQUE |

- **Какой тип данных обычно делают PK?**
  Наиболее распространённые варианты для первичного ключа:
  1. `SERIAL`/`BIGSERIAL` (PostgreSQL) либо `AUTO_INCREMENT` (MySQL) — автоинкрементные целые
  2. `INT`/`BIGINT` — когда значения задаются вручную
  3. `UUID` — для распределённых систем
  4. `IDENTITY` — стандартный SQL-механизм автоинкремента

  Примеры:
  ```sql
  -- PostgreSQL
  CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      name VARCHAR(100)
  );
  
  -- MySQL
  CREATE TABLE users (
      id INT AUTO_INCREMENT PRIMARY KEY,
      name VARCHAR(100)
  );
  ```

- **Почему UUID медленнее чем SERIAL?**
  1. **Объём** — UUID это 16 байт, тогда как SERIAL всего 4
  2. **Индексы** — ключ крупнее, значит больше страниц индекса и больше операций ввода-вывода
  3. **Кэш** — в кэш процессора помещается меньше значений UUID
  4. **Фрагментация** — случайные UUID разбросаны, из-за чего индекс фрагментируется
  5. **Сравнение** — сопоставлять 128-битные значения дольше, чем 32-битные

  При этом у UUID есть и плюсы:
  - уникальность в масштабах всего мира
  - предугадать следующий идентификатор сложно (это безопаснее)
  - значение можно сгенерировать прямо на стороне клиента

## SQL запросы
- **Что делает команда GROUP BY?**
  `GROUP BY` сворачивает строки с совпадающими значениями указанных столбцов в одну. Обычно применяется в паре с агрегатами (COUNT, SUM, AVG и т.п.).

  Пример:
  ```sql
  -- Количество заказов по каждому пользователю
  SELECT user_id, COUNT(*) as order_count
  FROM orders
  GROUP BY user_id;
  ```

- **В чем разница между HAVING и WHERE в SQL?**
  | WHERE | HAVING |
  |-------|--------|
  | Фильтрует строки до группировки | Фильтрует группы после GROUP BY |
  | Агрегатные функции запрещены | Агрегатные функции допустимы |
  | Срабатывает раньше GROUP BY | Срабатывает после GROUP BY |

  Пример:
  ```sql
  -- WHERE фильтрует строки до группировки
  SELECT department_id, AVG(salary)
  FROM employees
  WHERE hire_date > '2020-01-01'
  GROUP BY department_id
  HAVING AVG(salary) > 50000;  -- HAVING фильтрует группы
  ```

- **Что такое CTE? Какой оператор используется в CTE запросах?**
  CTE (Common Table Expression) — это именованный временный набор данных, доступный внутри основного запроса. Объявляется через ключевое слово `WITH`.

  Зачем нужен CTE:
  - делает громоздкие запросы понятнее
  - открывает возможность рекурсии
  - к одному и тому же CTE можно обращаться несколько раз

  Примеры:
  ```sql
  -- Простой CTE
  WITH regional_sales AS (
      SELECT region, SUM(amount) as total_sales
      FROM orders
      GROUP BY region
  )
  SELECT region, total_sales
  FROM regional_sales
  WHERE total_sales > 1000000;
  
  -- Рекурсивный CTE
  WITH RECURSIVE numbers AS (
      SELECT 1 AS n
      UNION ALL
      SELECT n + 1 FROM numbers WHERE n < 10
  )
  SELECT n FROM numbers;
  ```

- **Что такое подзапрос?**
  Подзапрос — это запрос, помещённый внутрь другого. Встречается в SELECT, FROM, WHERE, HAVING и других секциях.

  Какими бывают подзапросы:
  1. **Скалярный** — отдаёт ровно одно значение
     ```sql
     SELECT name, (SELECT MAX(score) FROM scores) as max_score
     FROM students;
     ```

  2. **Строковый** — возвращает несколько значений в пределах одной строки
     ```sql
     SELECT name 
     FROM products
     WHERE (price, rating) IN (SELECT MAX(price), 5 FROM products);
     ```

  3. **Столбцовый** — отдаёт один столбец с множеством строк
     ```sql
     SELECT name 
     FROM employees
     WHERE department_id IN (SELECT id FROM departments WHERE location = 'New York');
     ```

  4. **Табличный** — возвращает целую таблицу и подставляется в FROM
     ```sql
     SELECT dept_avg.*
     FROM (
         SELECT department_id, AVG(salary) as avg_salary
         FROM employees
         GROUP BY department_id
     ) as dept_avg
     WHERE avg_salary > 50000;
     ```

- **В чем разница между INNER JOIN и OUTER JOIN?**
  | Тип JOIN | Описание | Результат |
  |----------|----------|-----------|
  | INNER JOIN | Только строки, у которых нашлось соответствие в обеих таблицах | Пересечение записей |
  | LEFT OUTER JOIN | Все строки левой таблицы плюс совпадения из правой | Всё слева + совпавшее справа |
  | RIGHT OUTER JOIN | Все строки правой таблицы плюс совпадения из левой | Всё справа + совпавшее слева |
  | FULL OUTER JOIN | Все строки из обеих таблиц | Объединение LEFT и RIGHT |

  Примеры:
  ```sql
  -- INNER JOIN
  SELECT o.order_id, c.customer_name
  FROM orders o
  INNER JOIN customers c ON o.customer_id = c.id;
  
  -- LEFT JOIN (все заказы, даже без клиентов)
  SELECT o.order_id, c.customer_name
  FROM orders o
  LEFT JOIN customers c ON o.customer_id = c.id;
  
  -- RIGHT JOIN (все клиенты, даже без заказов)
  SELECT o.order_id, c.customer_name
  FROM orders o
  RIGHT JOIN customers c ON o.customer_id = c.id;
  
  -- FULL OUTER JOIN (все заказы и все клиенты)
  SELECT o.order_id, c.customer_name
  FROM orders o
  FULL OUTER JOIN customers c ON o.customer_id = c.id;
  ```
- **Как получить текущую дату и время в SQL?**
  Функции зависят от конкретной СУБД:

  ```sql
  -- PostgreSQL, SQLite
  SELECT CURRENT_TIMESTAMP;  -- Дата и время
  SELECT CURRENT_DATE;       -- Только дата
  SELECT CURRENT_TIME;       -- Только время
  
  -- MySQL, MariaDB
  SELECT NOW();             -- Дата и время
  SELECT CURDATE();         -- Только дата
  SELECT CURTIME();         -- Только время
  
  -- SQL Server
  SELECT GETDATE();         -- Дата и время
  SELECT GETUTCDATE();      -- Дата и время по UTC
  
  -- Oracle
  SELECT SYSDATE FROM DUAL;       -- Дата и время сервера
  SELECT SYSTIMESTAMP FROM DUAL;  -- Дата, время с часовым поясом
  ```

- **В чем разница между DELETE и TRUNCATE?**
  | Характеристика | DELETE | TRUNCATE |
  |----------------|--------|-----------|
  | Синтаксис | `DELETE FROM table WHERE condition` | `TRUNCATE TABLE table` |
  | Условие WHERE | Есть | Нет |
  | Триггеры | Срабатывают | Не срабатывают |
  | Счётчик автоинкремента | Сохраняется | Обнуляется |
  | Скорость | Ниже, особенно на больших таблицах | Выше, так как построчного лога нет |
  | Транзакции | Может входить в транзакцию | В большинстве СУБД — нет |
  | Права | Нужны права на DELETE | Нужны права на DROP TABLE |
  | Откат | Откатывается (внутри транзакции) | Откатить нельзя |

## Производительность
- **Как оптимизировать медленный запрос?**
  1. **Анализ запроса**
     - смотрите план через `EXPLAIN` или `EXPLAIN ANALYZE`
     - обращайте внимание на полное сканирование таблиц (FULL TABLE SCAN)
     - проверяйте, задействуются ли индексы

  2. **Оптимизация структуры**
     - добавьте недостающие индексы
     - удалите те, что не используются
     - пересмотрите схему (нормализация либо денормализация)

  3. **Оптимизация запроса**
     - берите только нужные столбцы, а не `SELECT *`
     - ограничивайте выборку через `LIMIT`
     - там, где можно, заменяйте подзапросы на JOIN
     - для больших наборов данных `EXISTS` предпочтительнее `IN`
     - не оборачивайте столбцы в функции внутри WHERE — это ломает использование индекса

  4. **Настройка СУБД**
     - подстройте буферный кэш
     - настройте параметры работы с памятью
     - обновите статистику (`ANALYZE` в PostgreSQL, `ANALYZE TABLE` в MySQL)

- **Что такое Explain и Explain ANALYZE?**
  `EXPLAIN` и `EXPLAIN ANALYZE` нужны для разбора того, как выполняется SQL-запрос.

  - **EXPLAIN** показывает предполагаемый план, не запуская сам запрос.
    ```sql
    EXPLAIN SELECT * FROM users WHERE age > 30;
    ```

  - **EXPLAIN ANALYZE** реально выполняет запрос и выдаёт план вместе с замерами.
    ```sql
    EXPLAIN ANALYZE SELECT * FROM users WHERE age > 30;
    ```

  На что смотреть в выводе:
  - **Seq Scan** против **Index Scan** — полный проход или доступ через индекс
  - **Cost** — оценочная стоимость шага
  - **Rows** — ожидаемое число обрабатываемых строк
  - **Actual Time** — реальное время
  - **Buffers** — сколько буферов задействовано

- **Как посмотреть количество текущих подключений к базе?**
  ```sql
  -- PostgreSQL
  SELECT count(*) FROM pg_stat_activity;
  
  -- MySQL
  SHOW STATUS WHERE `variable_name` = 'Threads_connected';
  
  -- SQL Server
  SELECT COUNT(*) FROM sys.dm_exec_connections;
  
  -- Oracle
  SELECT COUNT(*) FROM v$session WHERE status = 'ACTIVE';
  ```

- **Как посмотреть примерное количество записей в таблице, не используя COUNT(*)?**
  ```sql
  -- PostgreSQL
  SELECT reltuples::bigint AS estimate_count
  FROM pg_class
  WHERE relname = 'table_name';
  
  -- MySQL (InnoDB)
  SELECT TABLE_ROWS
  FROM INFORMATION_SCHEMA.TABLES
  WHERE TABLE_SCHEMA = 'database_name'
  AND TABLE_NAME = 'table_name';
  
  -- SQL Server
  SELECT p.rows
  FROM sys.tables t
  JOIN sys.partitions p ON t.object_id = p.object_id
  WHERE t.name = 'table_name' AND p.index_id IN (0, 1);
  
  -- Oracle
  SELECT num_rows
  FROM all_tables
  WHERE table_name = 'TABLE_NAME';
  ```

- **Почему иногда лучше использовать CHAR, чем VARCHAR?**
  `CHAR` бывает удачнее `VARCHAR` в таких ситуациях:

  1. **Данные фиксированной длины**
     - когда все значения столбца одинаковой ширины (коды стран, пол, статусы)
     - например, `CHAR(2)` под код страны (US, RU, GB)

  2. **Производительность**
     - фиксированная длина ускоряет сравнение и сортировку
     - не нужно отдельно хранить длину строки
     - расположение данных на диске предсказуемо

  3. **Выравнивание**
     - лучше ложится в память (для ряда СУБД это существенно)
     - меньше фрагментации при обновлениях

  4. **Индексы**
     - индекс по `CHAR` работает эффективнее
     - и занимает меньше места, чем по `VARCHAR`

  Пример использования:
  ```sql
  -- Лучше использовать CHAR(1)
  CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      gender CHAR(1),  -- 'M' или 'F'
      status CHAR(1)   -- 'A' (active), 'I' (inactive)
  );
  
  -- Лучше использовать VARCHAR
  CREATE TABLE articles (
      id SERIAL PRIMARY KEY,
      title VARCHAR(255),  -- Длина сильно варьируется
      content TEXT
  );
  ```

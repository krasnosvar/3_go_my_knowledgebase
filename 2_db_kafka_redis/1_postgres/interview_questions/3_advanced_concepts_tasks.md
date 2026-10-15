# Продвинутые концепции и практические задачи

## Распределенные системы
- **Что такое CAP теорема?**
  Согласно теореме CAP, распределённая система может гарантировать одновременно максимум два свойства из трёх:
  - **C**onsistency (согласованность) — в любой момент все узлы отдают одинаковые данные
  - **A**vailability (доступность) — на каждый запрос приходит ответ без ошибки
  - **P**artition tolerance (устойчивость к разделению) — система живёт, даже когда часть сообщений между узлами теряется

  Как это ложится на практику:
  - CA: классические реляционные БД (PostgreSQL, MySQL в стандартном режиме)
  - CP: MongoDB, HBase (жертвуют доступностью ради согласованности)
  - AP: Cassandra, DynamoDB (жертвуют согласованностью ради доступности)

- **Что такое шардирование и репликация, чем они отличаются?**
  **Репликация** — это ведение нескольких копий одних и тех же данных на разных узлах. Форматы:
  - мастер-реплика (синхронная или асинхронная)
  - мульти-мастер
  - каскадная репликация

  **Шардирование** (партиционирование) — это дробление данных на части (шарды) с раскладкой по разным узлам. Бывает:
  - горизонтальное (по строкам)
  - вертикальное (по столбцам)

  | Характеристика | Репликация | Шардирование |
  |----------------|------------|--------------|
  | Зачем нужно | Отказоустойчивость и масштабирование чтения | Масштабирование записи и чтения |
  | Что на узле | Полная копия данных | Только часть данных |
  | Запись | Уходит на все реплики | Только в нужный шард |
  | Чтение | С любой реплики | С любого шарда (иногда с последующей агрегацией) |

  Пример шардирования в PostgreSQL с помощью расширения `pg_partman`:
  ```sql
  -- Создание партиционированной таблицы
  CREATE TABLE sales (
      sale_id BIGSERIAL,
      sale_date TIMESTAMP NOT NULL,
      product_id INTEGER,
      quantity INTEGER,
      amount DECIMAL(10,2)
  ) PARTITION BY RANGE (sale_date);
  
  -- Создание партиций по месяцам
  SELECT partman.create_parent('public.sales', 'sale_date', 'native', 'monthly');
  ```

- **Как обеспечить консистентность в распределенных системах?**
  Способы держать данные согласованными:

  1. **Строгая консистентность**
     - 2PC (Two-Phase Commit)
     - консенсусные алгоритмы Paxos, Raft

  2. **Согласованность в конечном счёте**
     - асинхронная репликация
     - компенсирующие транзакции (паттерн Saga)

  3. **Модели согласованности**
     - линеаризуемость (linearizable)
     - последовательная (sequential)
     - слабая, «в итоге» (eventual)

  Пример применения паттерна Saga:
  ```go
  // Пример на Go с использованием библиотеки dtm
  saga := dtmcli.NewSaga(dtmServer, dtmcli.MustGenGid(dtmServer)).
      Add("postgres://user:pass@localhost:5432/orders",
          "update orders set status='preparing' where id=1",
          "update orders set status='created' where id=1").
      Add("postgres://user:pass@localhost:5432/inventory",
          "update inventory set quantity=quantity-1 where product_id=100",
          "update inventory set quantity=quantity+1 where product_id=100")
  
  err := saga.Submit()
  ```

- **Что такое MVCC (Multiversion Concurrency Control)?**
  MVCC — это подход к управлению одновременным доступом, при котором каждая транзакция работает со снимком данных на момент своего старта.

  Как это устроено:
  1. У каждой строки есть два служебных поля: `xmin` (транзакция-создатель) и `xmax` (транзакция-удалитель)
  2. При изменении строки заводится её новая версия
  3. Старые версии живут, пока их видит хотя бы одна транзакция

  Что это даёт:
  - читающие не блокируют пишущих и наоборот
  - при чтении блокировки не нужны
  - высокая производительность на чтении

  Пример в PostgreSQL:
  ```sql
  -- Просмотр версий строк
  SELECT xmin, xmax, * FROM table_name;
  
  -- Настройка уровня изоляции
  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
  ```

- **Как работают распределенные транзакции?**
  Распределённая транзакция затрагивает сразу несколько баз или сервисов. Основные протоколы:

  1. **2PC (Two-Phase Commit)**
     - фаза подготовки: координатор спрашивает участников, готовы ли они
     - фаза фиксации/отката: координатор объявляет итоговое решение

  2. **3PC (Three-Phase Commit)**
     - добавлена промежуточная фаза предварительной фиксации
     - реже блокируется, если координатор упал

  3. **Saga**
     - цепочка локальных транзакций
     - для отката используются компенсирующие операции

  Пример 2PC в PostgreSQL:
  ```sql
  -- Начало распределенной транзакции
  BEGIN;
  
  -- Выполнение операций в разных БД
  INSERT INTO db1.public.orders VALUES (...);
  INSERT INTO db2.public.inventory VALUES (...);
  
  -- Подготовка к фиксации
  PREPARE TRANSACTION 'transaction_id';
  
  -- Фиксация на всех узлах
  COMMIT PREPARED 'transaction_id';
  
  -- Или откат
  -- ROLLBACK PREPARED 'transaction_id';
  ```

## Проблемы параллелизма
- **Что такое deadlock и как его избежать?**
  Deadlock (взаимоблокировка) возникает, когда две (или больше) транзакции держат ресурсы, каждый из которых нужен другой стороне, — и все встают в бесконечное ожидание.

  Четыре условия, при которых возможен deadlock:
  1. взаимное исключение (ресурс занят только одной транзакцией)
  2. удержание с ожиданием (транзакция держит одно и просит ещё)
  3. отсутствие вытеснения (ресурс нельзя отнять принудительно)
  4. циклическое ожидание

  Как снизить риск:
  - выставлять таймауты
  - всегда захватывать блокировки в одинаковом порядке
  - полагаться на автодетект (PostgreSQL сам находит и разрывает deadlock)

  Пример deadlock в PostgreSQL:
  ```sql
  -- Сессия 1
  BEGIN;
  UPDATE accounts SET balance = balance - 100 WHERE id = 1;
  
  -- Сессия 2
  BEGIN;
  UPDATE accounts SET balance = balance - 200 WHERE id = 2;
  
  -- Deadlock при обратном порядке блокировок
  -- Сессия 1
  UPDATE accounts SET balance = balance + 100 WHERE id = 2;  -- Блокируется
  
  -- Сессия 2
  UPDATE accounts SET balance = balance + 200 WHERE id = 1;  -- Deadlock!
  ```

- **Проблемы транзакций (Transaction problems)**
  Типичные аномалии при параллельном доступе:

  1. **Грязное чтение (Dirty Read)**
     - видны чужие ещё не зафиксированные изменения
     - лечится уровнем READ COMMITTED и выше

  2. **Неповторяющееся чтение (Non-repeatable Read)**
     - при повторном чтении тех же строк транзакция получает уже другие значения
     - лечится REPEATABLE READ или SERIALIZABLE

  3. **Фантомное чтение (Phantom Read)**
     - при повторном чтении диапазона появляются новые строки
     - по стандарту SQL лечится SERIALIZABLE или блокировкой диапазона
     - в PostgreSQL фантомов нет уже на REPEATABLE READ (snapshot isolation); там остаётся только write skew, от которого защищает SERIALIZABLE

  4. **Потерянное обновление (Lost Update)**
     - две транзакции читают и правят одну строку, второй апдейт затирает первый
     - лечится SELECT FOR UPDATE либо оптимистичной блокировкой

- **Что такое race condition в контексте БД?**
  Race condition (состояние гонки) — это когда результат зависит от порядка выполнения операций, а сам порядок непредсказуем.

  Как это выглядит в БД:

  1. **Проблема инкремента**
     ```sql
     -- Транзакция 1         -- Транзакция 2
     SELECT counter FROM t;   SELECT counter FROM t;
     -- counter = 5           -- counter = 5
     UPDATE t SET counter=6;  UPDATE t SET counter=6;
     -- Ожидалось 7, но будет 6
     ```

  2. **Проблема «проверил-и-вставил»**
     ```sql
     -- Небезопасно (race condition)
     IF NOT EXISTS (SELECT 1 FROM users WHERE email='test@example.com')
     THEN
         INSERT INTO users (email) VALUES ('test@example.com');
     END IF;
     
     -- Безопасно (использование уникального ограничения)
     INSERT INTO users (email) 
     SELECT 'test@example.com'
     WHERE NOT EXISTS (SELECT 1 FROM users WHERE email='test@example.com');
     ```

- **Как работает SELECT FOR UPDATE?**
  `SELECT FOR UPDATE` удерживает выбранные строки заблокированными до конца транзакции, не давая другим их менять.

  Способы применения:
  ```sql
  -- Блокировка строки
  BEGIN;
  SELECT * FROM accounts WHERE id = 1 FOR UPDATE;
  -- Другие транзакции не могут изменить эту строку до COMMIT
  UPDATE accounts SET balance = balance - 100 WHERE id = 1;
  COMMIT;
  
  -- Блокировка с пропуском заблокированных строк
  SELECT * FROM orders WHERE status = 'new' FOR UPDATE SKIP LOCKED LIMIT 10;
  
  -- Блокировка с ожиданием (по умолчанию)
  SELECT * FROM products WHERE id = 100 FOR UPDATE NOWAIT; -- Вызовет ошибку при блокировке
  ```

- **Как обрабатывать конфликты при параллельном доступе?**
  Основные стратегии:

  1. **Пессимистичная блокировка**
     - блокируем на уровне строки или таблицы
     - `SELECT FOR UPDATE`, `LOCK TABLE`

  2. **Оптимистичная блокировка**
     - работаем через версию строки (номер версии или timestamp)
     - пример:
       ```sql
       -- Обновление с проверкой версии
       UPDATE products 
       SET name = 'New Name', version = version + 1 
       WHERE id = 1 AND version = 5;
       
       -- Проверка количества обновленных строк
       IF row_count = 0 THEN
           -- Обработка конфликта
       END IF;
       ```

  3. **Повторные попытки (Retry)**
     ```python
     def update_with_retry():
         max_retries = 3
         for attempt in range(max_retries):
             try:
                 # Попытка обновления
                 return update_record()
             except ConcurrentModificationError:
                 if attempt == max_retries - 1:
                     raise
                 time.sleep(0.1 * (2 ** attempt))  # Экспоненциальная задержка
     ```

## Практические задачи

### Задача 1: Книжный магазин

**Решение запросов:**

1. **Список книг, проданных сегодня со скидкой**
   ```sql
   SELECT 
       b.book_id,
       b.title,
       a.name AS author,
       b.price,
       b.discount,
       b.price * (1 - b.discount/100) AS final_price,
       SUM(s.count) AS total_sold
   FROM books b
   JOIN authors a ON b.author_id = a.author_id
   JOIN solds s ON b.book_id = s.book_id
   WHERE DATE(s.date) = CURRENT_DATE
     AND b.discount > 0
   GROUP BY b.book_id, a.name
   ORDER BY total_sold DESC;
   ```

2. **Топ 5 книг по популярности (по количеству продаж)**
   ```sql
   SELECT 
       b.book_id,
       b.title,
       a.name AS author,
       SUM(s.count) AS total_sold,
       SUM(s.count * b.price * (1 - b.discount/100)) AS total_revenue
   FROM books b
   JOIN authors a ON b.author_id = a.author_id
   JOIN solds s ON b.book_id = s.book_id
   GROUP BY b.book_id, a.name
   ORDER BY total_sold DESC
   LIMIT 5;
   ```

3. **Рекомендуемые индексы**
   ```sql
   -- Для запросов по дате продажи
   CREATE INDEX idx_solds_date ON solds(date);
   
   -- Для соединения таблиц
   CREATE INDEX idx_books_author_id ON books(author_id);
   CREATE INDEX idx_solds_book_id ON solds(book_id);
   CREATE INDEX idx_ratings_book_id ON ratings(book_id);
   
   -- Для часто используемых фильтров
   CREATE INDEX idx_books_discount ON books(discount) WHERE discount > 0;
   CREATE INDEX idx_books_genre ON books(genre);
   ```

### Задача 2: Оптимизация запроса

**Исходный запрос (пример):**
```sql
SELECT u.user_id, u.username, COUNT(o.order_id) as order_count
FROM users u
LEFT JOIN orders o ON u.user_id = o.user_id
WHERE u.registration_date > '2023-01-01'
GROUP BY u.user_id, u.username
HAVING COUNT(o.order_id) > 5
ORDER BY order_count DESC;
```

**Что можно улучшить:**

1. **Заменить LEFT JOIN на INNER JOIN**
   ```sql
   -- Если нам нужны только пользователи с заказами
   SELECT u.user_id, u.username, COUNT(*) as order_count
   FROM users u
   INNER JOIN orders o ON u.user_id = o.user_id
   WHERE u.registration_date > '2023-01-01'
   GROUP BY u.user_id, u.username
   HAVING COUNT(*) > 5
   ORDER BY order_count DESC;
   ```

2. **Добавить индексы**
   ```sql
   CREATE INDEX idx_users_registration_date ON users(registration_date);
   CREATE INDEX idx_orders_user_id ON orders(user_id);
   ```

3. **Вынести подсчёт заказов в отдельный CTE**
   ```sql
   WITH user_orders AS (
       SELECT user_id, COUNT(*) as order_count
       FROM orders
       GROUP BY user_id
       HAVING COUNT(*) > 5
   )
   SELECT u.user_id, u.username, uo.order_count
   FROM users u
   JOIN user_orders uo ON u.user_id = uo.user_id
   WHERE u.registration_date > '2023-01-01'
   ORDER BY uo.order_count DESC;
   ```

### Задача 3: Анализ продаж

**Запрос для вывода магазинов с количеством заказов > 100 за сентябрь 2023:**

```sql
WITH monthly_orders AS (
    SELECT 
        s.store_id,
        s.store_name,
        COUNT(o.order_id) as order_count,
        SUM(o.amount) as total_sales
    FROM stores s
    JOIN orders o ON s.store_id = o.store_id
    WHERE o.order_date >= '2023-09-01' 
      AND o.order_date < '2023-10-01'
    GROUP BY s.store_id, s.store_name
)
SELECT 
    store_id,
    store_name,
    order_count,
    total_sales
FROM monthly_orders
WHERE order_count > 100
ORDER BY total_sales DESC;
```

**Как ускорить:**
1. Завести индекс по ID магазина и дате заказа:
   ```sql
   CREATE INDEX idx_orders_store_date ON orders(store_id, order_date);
   ```

2. При больших объёмах — использовать материализованное представление:
   ```sql
   CREATE MATERIALIZED VIEW mv_monthly_store_sales AS
   SELECT 
       s.store_id,
       s.store_name,
       DATE_TRUNC('month', o.order_date) as month,
       COUNT(o.order_id) as order_count,
       SUM(o.amount) as total_sales
   FROM stores s
   JOIN orders o ON s.store_id = o.store_id
   GROUP BY s.store_id, s.store_name, DATE_TRUNC('month', o.order_date);
   
   CREATE UNIQUE INDEX idx_mv_monthly_store_sales ON mv_monthly_store_sales(store_id, month);
   ```

## Администрирование

- **Как посмотреть текущие подключения к базе?**
  ```sql
  -- PostgreSQL
  SELECT * FROM pg_stat_activity;
  
  -- MySQL
  SHOW PROCESSLIST;
  
  -- SQL Server
  EXEC sp_who2;
  
  -- Oracle
  SELECT * FROM v$session;
  ```

- **Как мониторить производительность запросов?**
  ```sql
  -- Просмотр медленных запросов (PostgreSQL)
  SELECT * FROM pg_stat_statements ORDER BY total_time DESC LIMIT 10;
  
  -- Просмотр использования индексов
  SELECT * FROM pg_stat_user_indexes;
  
  -- Мониторинг блокировок
  SELECT blocked_locks.pid AS blocked_pid,
         blocked_activity.usename AS blocked_user,
         blocking_locks.pid AS blocking_pid,
         blocking_activity.usename AS blocking_user
  FROM pg_catalog.pg_locks blocked_locks
  JOIN pg_catalog.pg_stat_activity blocked_activity 
      ON blocked_activity.pid = blocked_locks.pid
  JOIN pg_catalog.pg_locks blocking_locks 
      ON blocking_locks.locktype = blocked_locks.locktype
      AND blocking_locks.database IS NOT DISTINCT FROM blocked_locks.database
      AND blocking_locks.relation IS NOT DISTINCT FROM blocked_locks.relation
      AND blocking_locks.page IS NOT DISTINCT FROM blocked_locks.page
      AND blocking_locks.tuple IS NOT DISTINCT FROM blocked_locks.tuple
      AND blocking_locks.virtualxid IS NOT DISTINCT FROM blocked_locks.virtualxid
      AND blocking_locks.transactionid IS NOT DISTINCT FROM blocked_locks.transactionid
      AND blocking_locks.classid IS NOT DISTINCT FROM blocked_locks.classid
      AND blocking_locks.objid IS NOT DISTINCT FROM blocked_locks.objid
      AND blocking_locks.objsubid IS NOT DISTINCT FROM blocked_locks.objsubid
      AND blocking_locks.pid != blocked_locks.pid
  JOIN pg_catalog.pg_stat_activity blocking_activity 
      ON blocking_activity.pid = blocking_locks.pid
  WHERE NOT blocked_locks.granted;
  ```

- **Как настроить репликацию?**
  **PostgreSQL (логическая репликация):**
  1. На основном сервере (postgresql.conf):
     ```
     wal_level = logical
     max_wal_senders = 10
     max_replication_slots = 10
     ```
  2. Создать пользователя для репликации:
     ```sql
     CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'password';
     ```
  3. Настроить аутентификацию (pg_hba.conf):
     ```
     host    replication     replicator      replica_ip/32       md5
     ```
  4. На реплике создать слот репликации:
     ```sql
     SELECT * FROM pg_create_physical_replication_slot('replica_slot');
     ```
  5. Запустить репликацию:
     ```bash
     pg_basebackup -h master_host -U replicator -D /var/lib/postgresql/14/main -P -R --slot=replica_slot
     ```

- **Как сделать бэкап базы данных?**
  ```bash
  # PostgreSQL
  pg_dump -U username -d dbname -f backup.sql
  pg_dumpall -U username -f all_databases.sql
  
  # MySQL
  mysqldump -u username -p dbname > backup.sql
  mysqldump --all-databases -u username -p > all_databases.sql
  
  # SQL Server
  BACKUP DATABASE dbname TO DISK = '/path/to/backup.bak';
  
  # MongoDB
  mongodump --host localhost --port 27017 --db dbname --out /backup/path
  ```

- **Как восстановить базу из бэкапа?**
  ```bash
  # PostgreSQL
  psql -U username -d dbname -f backup.sql
  
  # MySQL
  mysql -u username -p dbname < backup.sql
  
  # SQL Server
  RESTORE DATABASE dbname FROM DISK = '/path/to/backup.bak';
  
  # MongoDB
  mongorestore --host localhost --port 27017 --db dbname /backup/path/dbname
  ```

  **Что важно помнить про бэкапы:**
  - периодически проверять, что восстановление реально работает
  - держать несколько копий в разных местах
  - автоматизировать снятие бэкапов
  - следить за тем, что бэкапы успешно создаются

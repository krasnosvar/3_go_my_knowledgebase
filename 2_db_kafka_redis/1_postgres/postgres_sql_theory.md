# 📘 PostgreSQL: справочник по SQL и основные концепции

Каноничный конспект теории по PostgreSQL в этой базе. Часть 1 — справочник по SQL-командам, часть 2 — концепции «простым языком» (транзакции, изоляция, индексы, масштабирование).

# Часть 1. Справочник по базовым SQL-командам

---

## 📋 Содержание

1. [DDL — Создание и изменение структуры](#ddl)
2. [DML — Работа с данными](#dml)
3. [SELECT — Выборка данных](#select)
4. [JOIN — Объединение таблиц](#join)
5. [Агрегатные функции](#aggregate)
6. [Подзапросы](#subquery)
7. [Индексы](#indexes)
8. [Транзакции](#transactions)
9. [Представления (Views)](#views)
10. [Полезные команды psql](#psql)

Часть 2 — [Концепции простым языком](#concepts): JOIN изнутри, агрегаты, нормальные формы, транзакции и уровни изоляции, индексы и селективность, внешние ключи, денормализация, масштабирование, массовые обновления, пагинация.

---

## 1. DDL — Создание и изменение структуры <a name="ddl"></a>

### Создание таблицы

```sql
CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    age        INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Разбор типов данных, которые можно назначить полям таблицы
CREATE TABLE IF NOT EXISTS users (

    -- 🔑 Идентификатор строки, значение растёт само.
    -- Варианты: SERIAL (int, до ~2.1 млрд), BIGSERIAL (bigint, до ~9.2 квинтиллиона), UUID
    id              SERIAL PRIMARY KEY,

    -- 📝 Текст с ограничением по длине.
    -- Варианты: VARCHAR(n) — не более n символов, CHAR(n) — ровно n символов, TEXT — без предела
    name            VARCHAR(100) NOT NULL,

    -- 📧 Текст, значения которого не должны повторяться.
    -- Варианты: VARCHAR(n), TEXT
    -- Под UNIQUE PostgreSQL сам заводит индекс
    email           VARCHAR(255) UNIQUE NOT NULL,

    -- 🔢 Целое число.
    -- Варианты: SMALLINT (до ~32к), INT (до ~2.1 млрд), BIGINT (до ~9.2 квинтиллиона)
    age             INT,

    -- 💰 Дробь / денежные суммы.
    -- Варианты: NUMERIC(10,2) — точно (для финансов), FLOAT/REAL — приблизительно (для расчётов)
    salary          NUMERIC(10, 2),

    -- ✅ Логическое значение: true / false / NULL
    -- Вариант: BOOLEAN
    is_active       BOOLEAN DEFAULT true,

    -- 📅 Дата без времени. Формат: YYYY-MM-DD
    -- Вариант: DATE
    birth_date      DATE,

    -- 🕐 Дата и время без часового пояса.
    -- Варианты: TIMESTAMP — без TZ, TIMESTAMPTZ — с часовым поясом (предпочтительно)
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- 🌍 Дата и время с учётом часового пояса (оптимальный выбор для прода).
    -- Хранится в UTC, показывается в часовом поясе сессии
    updated_at      TIMESTAMPTZ DEFAULT NOW(),

    -- 🔗 Внешний ключ — ссылка на строку другой таблицы.
    -- Варианты: INT, BIGINT, UUID — тип обязан совпадать с PK родителя
    -- ON DELETE CASCADE — вслед за родителем удаляется и эта строка
    -- ON DELETE SET NULL — при удалении родителя поле обнуляется в NULL
    role_id         INT REFERENCES roles(id) ON DELETE SET NULL,

    -- 🗂️ Данные в JSON. JSONB — бинарное представление с поддержкой индексов (обычно берут его)
    -- Варианты: JSON — обычный текст, JSONB — бинарный, индексируемый
    metadata        JSONB,

    -- 🖼️ Двоичные данные (файлы, картинки).
    -- Вариант: BYTEA — до 1 ГБ. Крупные файлы лучше держать на диске, а в БД класть путь в VARCHAR
    avatar          BYTEA,

    -- 🆔 Идентификатор, уникальный в глобальном масштабе.
    -- Вариант: UUID. Как сгенерировать: gen_random_uuid() (pg 13+) или uuid_generate_v4()
    external_id     UUID DEFAULT gen_random_uuid(),

    -- 📋 Перечисление — допустимы только заранее заданные значения.
    -- Тип нужно объявить заранее: CREATE TYPE status_enum AS ENUM ('active', 'banned')
    status          status_enum DEFAULT 'active',

    -- 🔢 Массив однотипных значений.
    -- Варианты: INT[], TEXT[], VARCHAR[] и прочие
    tags            TEXT[]

);

```

### Изменение таблицы

```sql
-- Добавить столбец
ALTER TABLE users ADD COLUMN phone VARCHAR(20);

-- Изменить тип столбца
ALTER TABLE users ALTER COLUMN age TYPE BIGINT;

-- Переименовать столбец
ALTER TABLE users RENAME COLUMN phone TO phone_number;

-- Удалить столбец
ALTER TABLE users DROP COLUMN phone_number;

-- Переименовать таблицу
ALTER TABLE users RENAME TO customers;
```

### Удаление таблицы

```sql
DROP TABLE IF EXISTS users;

-- Удалить таблицу вместе с зависимостями
DROP TABLE IF EXISTS users CASCADE;
```

---

## 2. DML — Работа с данными <a name="dml"></a>

### INSERT — Добавление данных

```sql
-- Одна строка
INSERT INTO users (name, email, age)
VALUES ('Иван', 'ivan@mail.ru', 30);

-- Несколько строк
INSERT INTO users (name, email, age) VALUES
    ('Анна',   'anna@mail.ru',   25),
    ('Пётр',   'petr@mail.ru',   35),
    ('Мария',  'maria@mail.ru',  28);

-- Вернуть сгенерированный ID
INSERT INTO users (name, email) VALUES ('Олег', 'oleg@mail.ru')
RETURNING id;
```

### UPDATE — Обновление данных

```sql
-- Обновить одно поле
UPDATE users SET age = 31 WHERE id = 1;

-- Обновить несколько полей
UPDATE users SET name = 'Иван Иванов', age = 32 WHERE id = 1;

-- Обновить все строки
UPDATE users SET age = age + 1;
```

### DELETE — Удаление данных

```sql
-- Удалить конкретную строку
DELETE FROM users WHERE id = 1;

-- Удалить по условию
DELETE FROM users WHERE age < 18;

-- Удалить все строки (оставить таблицу)
DELETE FROM users;

-- Быстрая очистка таблицы
TRUNCATE TABLE users;
```

---

## 3. SELECT — Выборка данных <a name="select"></a>

### Базовые запросы

```sql
-- Все столбцы
SELECT * FROM users;

-- Определённые столбцы
SELECT id, name, email FROM users;

-- С псевдонимами
SELECT id, name AS full_name, email AS mail FROM users;

-- Уникальные значения
SELECT DISTINCT age FROM users;
```

### WHERE — Фильтрация

```sql
SELECT * FROM users WHERE age > 25;
SELECT * FROM users WHERE age BETWEEN 20 AND 35;
SELECT * FROM users WHERE name IN ('Иван', 'Анна');
SELECT * FROM users WHERE name LIKE 'Ив%';      -- начинается на "Ив"
SELECT * FROM users WHERE name ILIKE '%ван%';   -- без учёта регистра
SELECT * FROM users WHERE email IS NOT NULL;
```

### ORDER BY, LIMIT, OFFSET

```sql
-- Сортировка
SELECT * FROM users ORDER BY age DESC;
SELECT * FROM users ORDER BY name ASC, age DESC;

-- Ограничение количества строк
SELECT * FROM users LIMIT 10;

-- Пагинация
SELECT * FROM users ORDER BY id LIMIT 10 OFFSET 20;
```

---

## 4. JOIN — Объединение таблиц <a name="join"></a>

```sql
-- INNER JOIN — только совпадающие строки
SELECT u.id, u.name, o.id AS order_id
FROM users u
JOIN orders o ON u.id = o.user_id;

-- LEFT JOIN — все из левой + совпадения из правой
SELECT u.name, o.id AS order_id
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;

-- RIGHT JOIN — все из правой + совпадения из левой
SELECT u.name, o.id AS order_id
FROM users u
RIGHT JOIN orders o ON u.id = o.user_id;

-- FULL JOIN — все строки из обеих таблиц
SELECT u.name, o.id AS order_id
FROM users u
FULL JOIN orders o ON u.id = o.user_id;
```

| Тип JOIN    | Описание                                            |
|-------------|-----------------------------------------------------|
| `INNER JOIN`| Строки, для которых есть пара в обеих таблицах       |
| `LEFT JOIN` | Все строки слева + совпавшие справа (иначе NULL)     |
| `RIGHT JOIN`| Все строки справа + совпавшие слева                 |
| `FULL JOIN` | Все строки из обеих таблиц                          |

---

## 5. Агрегатные функции <a name="aggregate"></a>

```sql
SELECT COUNT(*)           FROM users;                   -- количество строк
SELECT COUNT(age)         FROM users;                   -- без NULL
SELECT AVG(age)           FROM users;                   -- среднее значение
SELECT SUM(age)           FROM users;                   -- сумма
SELECT MIN(age)           FROM users;                   -- минимум
SELECT MAX(age)           FROM users;                   -- максимум

-- GROUP BY — группировка
SELECT age, COUNT(*) AS total
FROM users
GROUP BY age
ORDER BY total DESC;

-- HAVING — фильтр после группировки
SELECT age, COUNT(*) AS total
FROM users
GROUP BY age
HAVING COUNT(*) > 2;
```

---

## 6. Подзапросы <a name="subquery"></a>

```sql
-- В WHERE
SELECT name FROM users
WHERE id IN (SELECT user_id FROM orders);

-- В FROM (как таблица)
SELECT sub.name, sub.total
FROM (
    SELECT u.name, COUNT(o.id) AS total
    FROM users u
    LEFT JOIN orders o ON u.id = o.user_id
    GROUP BY u.name
) AS sub
WHERE sub.total > 0;

-- CTE (Common Table Expression)
WITH active_users AS (
    SELECT * FROM users WHERE age BETWEEN 18 AND 60
)
SELECT * FROM active_users WHERE name LIKE 'А%';
```

---

## 7. Индексы <a name="indexes"></a>

```sql
-- Создать индекс
CREATE INDEX idx_users_email ON users(email);

-- Уникальный индекс
CREATE UNIQUE INDEX idx_users_email_unique ON users(email);

-- Составной индекс
CREATE INDEX idx_users_name_age ON users(name, age);

-- Удалить индекс
DROP INDEX IF EXISTS idx_users_email;

-- Посмотреть список индексов таблицы
SELECT indexname, indexdef
FROM pg_indexes
WHERE tablename = 'users';
```

> 💡 Индексы делают `SELECT` быстрее, но замедляют `INSERT`, `UPDATE` и `DELETE`. Заводите их на столбцах, по которым чаще всего идёт фильтрация.

---

## 8. Транзакции <a name="transactions"></a>

```sql
-- Начало транзакции
BEGIN;

INSERT INTO users (name, email) VALUES ('Тест', 'test@mail.ru');
UPDATE users SET age = 99 WHERE name = 'Тест';

-- Подтвердить изменения
COMMIT;

-- Или отменить все изменения
ROLLBACK;
```

### Точки сохранения (SAVEPOINT)

```sql
BEGIN;

INSERT INTO users (name, email) VALUES ('Тест', 'test@mail.ru');
SAVEPOINT my_savepoint;

DELETE FROM users WHERE name = 'Тест';

-- Откатиться только до точки сохранения
ROLLBACK TO SAVEPOINT my_savepoint;

COMMIT;
```

---

## 9. Представления (Views) <a name="views"></a>

```sql
-- Создать представление
CREATE VIEW active_users AS
SELECT id, name, email
FROM users
WHERE age BETWEEN 18 AND 60;

-- Использовать представление как таблицу
SELECT * FROM active_users;

-- Обновляемое представление
CREATE OR REPLACE VIEW active_users AS
SELECT id, name, email
FROM users
WHERE age >= 18;

-- Удалить представление
DROP VIEW IF EXISTS active_users;
```

---

## 10. Полезные команды psql <a name="psql"></a>

| Команда              | Описание                              |
|----------------------|---------------------------------------|
| `\l`                 | Список баз данных                     |
| `\c имя_бд`          | Подключиться к базе данных            |
| `\dt`                | Список таблиц                         |
| `\d имя_таблицы`     | Структура таблицы                     |
| `\di`                | Список индексов                       |
| `\dv`                | Список представлений                  |
| `\du`                | Список пользователей                  |
| `\timing`            | Включить/выключить замер времени      |
| `\x`                 | Расширенный вывод (построчно)         |
| `\q`                 | Выйти из psql                         |
| `\i файл.sql`        | Выполнить SQL-файл                    |

---

## 🔑 Шпаргалка по типам данных

| Тип              | Описание                        |
|------------------|---------------------------------|
| `SERIAL`         | Автоинкремент (int)             |
| `BIGSERIAL`      | Автоинкремент (bigint)          |
| `INT / BIGINT`   | Целое число                     |
| `NUMERIC(p, s)`  | Точное дробное число            |
| `FLOAT / REAL`   | Число с плавающей точкой        |
| `VARCHAR(n)`     | Строка до n символов            |
| `TEXT`           | Строка без ограничений          |
| `BOOLEAN`        | true / false                    |
| `DATE`           | Дата (YYYY-MM-DD)               |
| `TIMESTAMP`      | Дата и время                    |
| `UUID`           | Уникальный идентификатор        |
| `JSONB`          | JSON (бинарный, с индексами)    |

---

# Часть 2. Концепции простым языком <a name="concepts"></a>

Шпаргалка по основам PostgreSQL с понятными примерами из жизни.

Шпаргалка по основам PostgreSQL с понятными примерами из жизни.

---

## 1. JOIN — Объединение таблиц

### Что это?
Способ получить данные из нескольких таблиц одновременно.

### Пример из жизни
У вас есть две таблицы:
- `users` (пользователи): id, имя
- `orders` (заказы): id, user_id, сумма

Вы хотите узнать: **кто и сколько потратил**
```sql
-- INNER JOIN — только те пользователи, у которых ЕСТЬ заказы
SELECT users.name, orders.amount
FROM users
INNER JOIN orders ON users.id = orders.user_id;

-- LEFT JOIN — ВСЕ пользователи (даже без заказов)
SELECT users.name, orders.amount
FROM users
LEFT JOIN orders ON users.id = orders.user_id;
-- Если у пользователя нет заказов → amount будет NULL
```

### Как PostgreSQL это делает?

**Nested Loop** (вложенный цикл) — для маленьких таблиц:
```
Для каждого пользователя:
  Проверить все заказы
```

**Hash Join** — для больших таблиц:
```
1. Создать хеш-таблицу из users
2. Пройтись по orders и быстро найти совпадения
```

**Merge Join** — если данные уже отсортированы:
```
Идти по обеим таблицам синхронно (как слияние двух списков)
```

---

## 2. Агрегатные функции — Подсчёт статистики

### Что это?
Функции, которые считают что-то для группы строк.

### Примеры
```sql
-- Сколько всего заказов?
SELECT COUNT(*) FROM orders;

-- Какая общая сумма продаж?
SELECT SUM(amount) FROM orders;

-- Средний чек
SELECT AVG(amount) FROM orders;

-- Сколько заказов у каждого пользователя?
SELECT user_id, COUNT(*), SUM(amount)
FROM orders
GROUP BY user_id;
```

### Производительность

❌ **Медленно:**
```sql
-- Группировка по неиндексированной колонке
SELECT category, COUNT(*) 
FROM products 
GROUP BY category;  -- Полное сканирование таблицы
```

✅ **Быстро:**
```sql
-- Есть индекс на category
CREATE INDEX idx_category ON products(category);
-- Теперь запрос летает
```

**Лайфхак:** `COUNT(*)` быстрее `COUNT(column)` (не проверяет NULL)

---

## 3. Нормальные формы — Как правильно хранить данные

### Зачем?
Чтобы не дублировать данные и не запутаться в обновлениях.

### 1NF — Атомарность

❌ **Плохо:**
```sql
CREATE TABLE orders (
  id INT,
  items TEXT  -- "яблоки, груши, бананы" ← список в одной ячейке!
);
```

✅ **Хорошо:**
```sql
CREATE TABLE orders (id INT);
CREATE TABLE order_items (
  order_id INT,
  item_name TEXT
);
-- Каждый товар — отдельная строка
```

### 2NF — Нет частичных зависимостей

❌ **Плохо:**
```sql
CREATE TABLE order_items (
  order_id INT,
  product_id INT,
  product_name TEXT,     -- ← зависит только от product_id
  product_price DECIMAL, -- ← зависит только от product_id
  PRIMARY KEY (order_id, product_id)
);
-- product_name повторяется в каждом заказе!
```

✅ **Хорошо:**
```sql
CREATE TABLE products (
  id INT PRIMARY KEY,
  name TEXT,
  price DECIMAL
);

CREATE TABLE order_items (
  order_id INT,
  product_id INT,  -- ← ссылка на products
  quantity INT
);
```

### 3NF — Нет транзитивных зависимостей

❌ **Плохо:**
```sql
CREATE TABLE employees (
  id INT,
  name TEXT,
  department_id INT,
  department_name TEXT  -- ← зависит от department_id, а не от id!
);
```

✅ **Хорошо:**
```sql
CREATE TABLE departments (
  id INT,
  name TEXT
);

CREATE TABLE employees (
  id INT,
  name TEXT,
  department_id INT  -- ← ссылка
);
```

---

## 4. Транзакции — Всё или ничего

### Что это?
Группа операций, которые выполняются вместе или не выполняются вообще.

### Пример: Перевод денег
```sql
BEGIN;  -- Начало транзакции

  -- Снять у Васи
  UPDATE accounts SET balance = balance - 100 WHERE user_id = 1;
  
  -- Добавить Пете
  UPDATE accounts SET balance = balance + 100 WHERE user_id = 2;

COMMIT;  -- Применить изменения
-- Если что-то пойдёт не так → ROLLBACK (откатить)
```

### ACID — 4 гарантии

- **A (Atomicity)** — всё или ничего (либо оба UPDATE выполнятся, либо ни один)
- **C (Consistency)** — правила БД соблюдаются (balance не уйдёт в минус)
- **I (Isolation)** — транзакции не мешают друг другу
- **D (Durability)** — после COMMIT данные не потеряются (даже если сервер упадёт)

### Уровни изоляции

| Уровень | Что видно | Какие аномалии остаются (в PostgreSQL) |
|---------|-----------|----------|
| **READ UNCOMMITTED** | В PostgreSQL ведёт себя как READ COMMITTED (грязного чтения нет) | Как у READ COMMITTED |
| **READ COMMITTED** (по умолчанию) | Только сохранённые данные; снимок берётся на каждый запрос | Неповторяемое чтение, фантомы, потерянные обновления |
| **REPEATABLE READ** | Один снимок на всю транзакцию (snapshot isolation) | Фантомов нет; остаётся write skew (аномалия сериализации). Конкурирующий UPDATE той же строки → ошибка `could not serialize access`, транзакцию надо повторить |
| **SERIALIZABLE** | Полная изоляция (SSI) | Аномалий нет; цена — ошибки сериализации и повторы транзакций |

> В стандарте SQL уровень REPEATABLE READ допускает фантомные чтения. PostgreSQL строже стандарта: его REPEATABLE READ реализован как snapshot isolation, поэтому фантомов нет, но возможен write skew (две транзакции читают пересекающиеся данные и пишут в разные строки, нарушая общий инвариант). От write skew защищает только SERIALIZABLE (или явные блокировки `SELECT ... FOR UPDATE`).

### SELECT FOR UPDATE — Блокировка строк
```sql
BEGIN;

  -- Забронировать строку (другие не смогут её изменить)
  SELECT * FROM products WHERE id = 5 FOR UPDATE;
  
  -- Теперь можно безопасно обновить
  UPDATE products SET stock = stock - 1 WHERE id = 5;

COMMIT;
```

**Когда использовать:** При редактировании товаров в корзине, резервировании мест.

---

## 5. Индексы — Ускорение поиска

### Что это?
Как закладка в книге — помогает быстро найти нужную страницу.

### Без индекса:
```sql
SELECT * FROM users WHERE email = 'ivan@mail.ru';
-- PostgreSQL прочитает ВСЕ 1 млн строк → медленно
```

### С индексом:
```sql
CREATE INDEX idx_email ON users(email);
-- Теперь поиск мгновенный (как поиск в словаре)
```

### Типы индексов

**B-tree** (по умолчанию) — для любых сравнений:
```sql
CREATE INDEX idx_age ON users(age);
-- Работает для: =, <, >, BETWEEN, ORDER BY
```

**Hash** — только для точного совпадения:
```sql
CREATE INDEX idx_phone ON users USING hash(phone);
-- Работает только для: phone = '...'
-- Быстрее B-tree, но ограничен
```

### Селективность — насколько индекс полезен

✅ **Высокая селективность** (хороший индекс):
```sql
email: ivan@mail.ru, maria@gmail.com, petr@yandex.ru
-- Все значения разные → индекс эффективен
```

❌ **Низкая селективность** (плохой индекс):
```sql
gender: М, Ж, М, М, Ж, М, Ж
-- Всего 2 значения → индекс бесполезен
```

### Составные индексы
```sql
CREATE INDEX idx_city_age ON users(city, age);

-- ✅ Работает:
WHERE city = 'Москва' AND age > 25
WHERE city = 'Москва'

-- ❌ НЕ работает:
WHERE age > 25  -- city не указан
```

**Правило:** Порядок колонок важен! Первая колонка должна быть в WHERE.

### EXPLAIN — Как проверить, используется ли индекс?
```sql
EXPLAIN SELECT * FROM users WHERE email = 'test@mail.ru';

-- Вывод:
-- Index Scan using idx_email  ← хорошо, индекс используется
-- Seq Scan on users           ← плохо, полное сканирование таблицы
```

---

## 6. Внешние ключи — Защита от ошибок

### Что это?
Правило: "В таблице заказов user_id должен существовать в таблице users"
```sql
CREATE TABLE users (
  id INT PRIMARY KEY,
  name TEXT
);

CREATE TABLE orders (
  id INT PRIMARY KEY,
  user_id INT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- ✅ Можно:
INSERT INTO orders (user_id) VALUES (5);  -- если user 5 существует

-- ❌ Ошибка:
INSERT INTO orders (user_id) VALUES (999);  -- такого user нет
```

### Влияние на производительность

**Минусы:**
- Каждый INSERT/UPDATE проверяет существование связанной записи → медленнее
- При удалении родителя нужно обработать детей

**Каскадное удаление:**
```sql
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE

-- Удалили пользователя → автоматически удалятся все его заказы
```

---

## 7. Денормализация — Когда нарушать правила

### Зачем?
Нормализация → много JOIN → медленные запросы.

### Пример

**Нормализованная схема:**
```sql
SELECT users.name, orders.amount
FROM users
JOIN orders ON users.id = orders.user_id;
-- JOIN на каждый запрос → медленно
```

**Денормализованная схема:**
```sql
CREATE TABLE orders (
  id INT,
  user_id INT,
  user_name TEXT,  -- ← дублируем имя!
  amount DECIMAL
);

SELECT user_name, amount FROM orders;
-- Без JOIN → быстро
```

### Когда денормализовать?

✅ **Да:**
- Очень частые читающие запросы
- JOIN тормозит работу
- Данные редко меняются (имя пользователя)

❌ **Нет:**
- Часто обновляемые данные (баланс счёта)
- Небольшие таблицы
- Критична консистентность

### Пример из реальной жизни
```sql
-- Вместо COUNT(*) при каждом запросе:
CREATE TABLE users (
  id INT,
  name TEXT,
  order_count INT DEFAULT 0  -- ← кешируем счётчик
);

-- При создании заказа обновляем:
UPDATE users SET order_count = order_count + 1 WHERE id = ...;
```

---

## 8. Масштабируемость — Работа с большими данными

### Репликация — Копии БД
```
Master (главная)
  ├─ Slave 1 (копия для чтения)
  ├─ Slave 2 (копия для чтения)
  └─ Slave 3 (копия для чтения)
```

**Когда:** Много читающих запросов, мало пишущих.

**Пример:** Интернет-магазин
- Master → оформление заказов (запись)
- Slave → просмотр каталога (чтение)

---

### Шардирование — Разделение данных
```
Сервер 1: пользователи с ID 1-1000000
Сервер 2: пользователи с ID 1000001-2000000
Сервер 3: пользователи с ID 2000001-3000000
```

**Когда:** Одна БД не справляется, данные огромные.

**Проблема:** Сложно делать JOIN между серверами.

---

### Партиционирование — Разделение таблицы
```sql
-- Вместо одной большой таблицы:
CREATE TABLE orders (
  id INT,
  created_at DATE,
  amount DECIMAL
) PARTITION BY RANGE (created_at);

-- Создаём партиции:
CREATE TABLE orders_2024_01 PARTITION OF orders
  FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

CREATE TABLE orders_2024_02 PARTITION OF orders
  FOR VALUES FROM ('2024-02-01') TO ('2024-03-01');
```

**Плюсы:**
- Запрос за январь → читает только orders_2024_01
- Можно удалить старые партиции (архивация)

---

### Холодное/Горячее хранение

**Горячее** (SSD, быстро):
- Заказы за последние 30 дней
- Активные пользователи

**Холодное** (HDD, медленно, дёшево):
- Заказы за 2020 год
- Удалённые аккаунты

---

## 9. Обновление 100 млн записей — Как не сломать БД?

### Проблема
```sql
UPDATE users SET status = 'active';  -- 100 млн строк
-- БД зависнет на несколько часов!
```

### Решение — Батчами (порциями)
```sql
DO $$
BEGIN
  LOOP
    -- Обновляем по 10,000 записей
    UPDATE users SET status = 'active'
    WHERE id IN (
      SELECT id FROM users 
      WHERE status != 'active' 
      LIMIT 10000
    );
    
    -- Нет больше записей → выходим
    EXIT WHEN NOT FOUND;
    
    -- Пауза 100 мс (даём БД отдохнуть)
    PERFORM pg_sleep(0.1);
  END LOOP;
END $$;
```

### Альтернатива — Временная таблица
```sql
-- 1. Создать новую таблицу
CREATE TABLE users_new AS 
SELECT id, name, 'active' as status FROM users;

-- 2. Переключить атомарно
BEGIN;
  DROP TABLE users;
  ALTER TABLE users_new RENAME TO users;
COMMIT;
```

---

## 10. Пагинация — Страницы в результатах

### Классический способ (LIMIT/OFFSET)
```sql
-- Страница 1 (строки 1-20)
SELECT * FROM users ORDER BY id LIMIT 20 OFFSET 0;

-- Страница 2 (строки 21-40)
SELECT * FROM users ORDER BY id LIMIT 20 OFFSET 20;

-- Страница 100 (строки 1981-2000)
SELECT * FROM users ORDER BY id LIMIT 20 OFFSET 1980;
```

### Проблема

На странице 100 PostgreSQL:
1. Читает первые 1980 строк
2. Пропускает их
3. Возвращает следующие 20

**Это медленно!**

---

### Keyset Pagination (правильный способ)
```sql
-- Страница 1
SELECT * FROM users ORDER BY id LIMIT 20;
-- Последний id = 20

-- Страница 2
SELECT * FROM users WHERE id > 20 ORDER BY id LIMIT 20;
-- Последний id = 40

-- Страница 3
SELECT * FROM users WHERE id > 40 ORDER BY id LIMIT 20;
```

**Плюсы:**
- Использует индекс → всегда быстро
- Работает на любой странице

**Минусы:**
- Нельзя перейти сразу на страницу 100 (только вперёд/назад)

---

## Шпаргалка — Когда что использовать

| Задача | Решение |
|--------|---------|
| Медленный поиск | Создать индекс |
| Много JOIN | Денормализация или оптимизация запросов |
| Много читающих запросов | Репликация |
| Огромная таблица | Партиционирование |
| Массовое обновление | Батчи + паузы |
| Пагинация больших данных | Keyset pagination |
| Конкуренция за данные | SELECT FOR UPDATE |
| Защита от битых ссылок | FOREIGN KEY |

---

## Полезные команды
```sql
-- Посмотреть план запроса
EXPLAIN ANALYZE SELECT ...;

-- Список индексов
\di

-- Размер таблицы
SELECT pg_size_pretty(pg_total_relation_size('users'));

-- Активные запросы
SELECT * FROM pg_stat_activity;

-- Убить долгий запрос
SELECT pg_terminate_backend(pid);
```

---

*Конспект рассчитан на PostgreSQL 14 и новее*

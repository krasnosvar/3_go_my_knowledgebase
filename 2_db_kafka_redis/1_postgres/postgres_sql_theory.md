# 📘 Справочник по базовым SQL-командам PostgreSQL

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

*Справочник рассчитан на PostgreSQL 14 и новее*

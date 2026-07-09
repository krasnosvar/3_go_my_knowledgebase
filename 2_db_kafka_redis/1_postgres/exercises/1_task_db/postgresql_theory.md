Теория PostgreSQL (PgSQL)
1. Join: Виды и механизмы работы.
2. Агрегатные функции: Использование и производительность.
3. Нормальные формы: Понимание первых трех (1NF, 2NF, 3NF).
4. Транзакции: 
- * Принципы ACID.
- 4 уровня изоляции транзакций.
- SELECT FOR UPDATE (блокировки).
5. Индексы:
- Типы: Hash, B-tree.
- Понятие селективности.
- Составные и уникальные индексы.
- Использование команды EXPLAIN для анализа запросов.
6. Внешние ключи (Foreign Keys): Назначение и влияние на производительность.
7. Денормализация: Когда и зачем отходить от нормальных форм.
8. Масштабируемость:
- Репликация.
- Шардирование.
- Партиционирование.
- Концепция «Холодное/Горячее» хранение данных.
9. Кейс: Как выполнить запрос на обновление 100 млн записей, не «положив» базу?
10. Пагинация: Использование LIMIT / OFFSET.


# PostgreSQL — Простым языком

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

| Уровень | Что видно | Проблемы |
|---------|-----------|----------|
| **READ COMMITTED** (по умолчанию) | Только сохранённые данные | Данные могут измениться между запросами |
| **REPEATABLE READ** | Снимок на момент начала транзакции | Фантомные чтения |
| **SERIALIZABLE** | Полная изоляция | Медленнее, конфликты |

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

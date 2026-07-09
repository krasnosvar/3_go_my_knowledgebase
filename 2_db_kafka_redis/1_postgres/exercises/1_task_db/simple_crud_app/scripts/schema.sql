-- 1. Создаём базу данных "shop", если её ещё нет
CREATE DATABASE shop;

-- 2. Подключаемся к базе "shop"
\c shop;

-- 3. Создаём таблицу пользователей
--    - id: уникальный идентификатор
--    - order_count: количество заказов (по умолчанию 0)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    order_count INT DEFAULT 0
);

-- 4. Создаём таблицу заказов
--    - id: уникальный идентификатор заказа
--    - user_id: ссылка на пользователя
--    - created_at: время создания заказа (по умолчанию текущее)
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now()
);

-- 5. Создаём таблицу товаров
--    - id: уникальный идентификатор товара
--    - order_id: ссылка на заказ
--    - price: цена товара
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    price NUMERIC(10,2) NOT NULL
);

-- 6. Генерация 1 млн пользователей
--    Каждый пользователь получает order_count = 0
INSERT INTO users(order_count)
SELECT 0
FROM generate_series(1, 1000000);

-- 7. Генерация 1 млн заказов
--    - user_id выбирается случайным образом из диапазона 1..999999
--    - created_at случайно распределён в пределах последних 365 дней
--      (т.е. за последний год)
INSERT INTO orders(user_id, created_at)
SELECT (random()*999999)::int + 1, now() - (random() * interval '365 days')
FROM generate_series(1, 1000000);

-- 8. Генерация 1 млн товаров
--    - order_id выбирается случайным образом из диапазона 1..999999
--    - price случайное число от 0 до 100 с двумя знаками после запятой
INSERT INTO items(order_id, price)
SELECT (random()*999999)::int + 1, (random()*100)::numeric(10,2)
FROM generate_series(1, 1000000);

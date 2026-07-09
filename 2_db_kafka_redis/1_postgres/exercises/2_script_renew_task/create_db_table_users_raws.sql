/*
Разбор конструкции:
generate_series(1, 1000000): Эта функция создает набор строк от 1 до 1го миллиона. Это позволяет нам выполнить INSERT 1 миллион раз за один запрос, без использования циклов в приложении.

ARRAY[...]: Мы создаем массив строк с нашими именами прямо в запросе. В PostgreSQL массивы индексируются с 1.

Выбор случайного имени:
  random(): Возвращает случайное число от 0.0 до 1.0.
  * 10: Масштабируем до 0.0 - 10.0.
  + 1: Сдвигаем диапазон до 1.0 - 11.0 (чтобы индексы были от 1).
  floor(...): Округляем вниз до целого числа (1, 2, ..., 10).
  ARRAY[...][ индекс ]: Берем элемент массива по полученному случайному индексу.
*/
-- Drop the users table if it exists to ensure a clean slate
DROP TABLE IF EXISTS users;

-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    order_count INT
);

-- Insert 1 million rows with random names and order_count = 0
INSERT INTO users (name, order_count)
SELECT 
    -- Here we select a random name from the array of names:
    -- order_count = 0
    -- random() returns 0.0 to 1.0, so * 10 gives 0.0 to 10.0, + 1 gives 1.0 to 11.0. floor() rounds down to int 1-10.
    (ARRAY[
        'Den', 'Max', 'John', 'Bob', 'Alice', 
        'Ann', 'Vasya', 'Petya', 'Masha', 'Olya'
    ])[floor(random() * 10 + 1)], 
    0 
-- generate_series creates 1 million rows, effectively looping 1 million times
FROM generate_series(1, 1000000); 

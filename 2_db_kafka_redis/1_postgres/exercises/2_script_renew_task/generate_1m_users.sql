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
    (ARRAY['Den', 'Max', 'John', 'Bob', 'Alice', 'Ann', 'Vasya', 'Petya', 'Masha', 'Olya'])[floor(random() * 10 + 1)],
    0
FROM generate_series(1, 1000000);

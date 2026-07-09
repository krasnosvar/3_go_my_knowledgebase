-- setup.sql
-- This script creates a table and populates it with 1,000,000 rows.
-- Added columns for GIN and GiST practice.

DROP TABLE IF EXISTS users_index_test;

CREATE TABLE users_index_test (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    age INT NOT NULL,
    status TEXT NOT NULL,
    tags TEXT[] NOT NULL, -- For GIN
    bio TEXT NOT NULL,     -- For GIN/GiST (Full Text Search)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Generate random data
INSERT INTO users_index_test (username, email, age, status, tags, bio, created_at)
SELECT 
    'user_' || i AS username,
    'user_' || i || '@example.com' AS email,
    (random() * 80 + 18)::INT AS age,
    (ARRAY['active', 'blocked', 'pending'])[floor(random() * 3 + 1)] AS status,
    CASE floor(random() * 4)::INT
        WHEN 0 THEN ARRAY['dev', 'admin']
        WHEN 1 THEN ARRAY['qa', 'user']
        WHEN 2 THEN ARRAY['manager', 'dev']
        ELSE ARRAY['user']
    END AS tags,
    'This is a sample bio for user ' || i || '. They like ' || (ARRAY['coding', 'reading', 'running', 'swimming'])[floor(random() * 4 + 1)] || '.' AS bio,
    NOW() - (random() * (interval '365 days')) AS created_at
FROM generate_series(1, 1000000) AS i;

ANALYZE users_index_test;

SELECT count(*) FROM users_index_test;

-- ============================================================
-- Инициализация БД для практики SQL JOIN
-- Автоматически выполняется при первом запуске docker compose
-- ============================================================

-- ============================================================
-- ТАБЛИЦЫ
-- ============================================================

-- Отделы компании
CREATE TABLE departments (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    budget      NUMERIC(12, 2) NOT NULL DEFAULT 0,
    location    VARCHAR(100)
);

-- Сотрудники
CREATE TABLE employees (
    id              SERIAL PRIMARY KEY,
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    email           VARCHAR(100) UNIQUE NOT NULL,
    department_id   INT REFERENCES departments(id),  -- может быть NULL (сотрудник без отдела)
    hire_date       DATE NOT NULL,
    salary          NUMERIC(10, 2) NOT NULL
);

-- Проекты
CREATE TABLE projects (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    start_date  DATE NOT NULL,
    end_date    DATE,          -- NULL = проект ещё идёт
    budget      NUMERIC(12, 2) NOT NULL DEFAULT 0
);

-- Назначения сотрудников на проекты (many-to-many)
CREATE TABLE project_assignments (
    id          SERIAL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES employees(id),
    project_id  INT NOT NULL REFERENCES projects(id),
    role        VARCHAR(50) NOT NULL,  -- 'lead', 'developer', 'tester', 'analyst'
    assigned_at DATE NOT NULL DEFAULT CURRENT_DATE,
    UNIQUE(employee_id, project_id)
);

-- Клиенты
CREATE TABLE clients (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    industry    VARCHAR(50),
    country     VARCHAR(50) NOT NULL DEFAULT 'Russia'
);

-- Заказы клиентов
CREATE TABLE orders (
    id          SERIAL PRIMARY KEY,
    client_id   INT NOT NULL REFERENCES clients(id),
    project_id  INT REFERENCES projects(id),  -- может быть NULL
    amount      NUMERIC(12, 2) NOT NULL,
    status      VARCHAR(20) NOT NULL DEFAULT 'new',  -- 'new', 'in_progress', 'completed', 'cancelled'
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP
);

-- ============================================================
-- ДАННЫЕ
-- ============================================================

-- Отделы (5 штук, один без сотрудников — для LEFT JOIN проверок)
INSERT INTO departments (name, budget, location) VALUES
    ('Backend',    5000000, 'Москва'),
    ('Frontend',   3000000, 'Москва'),
    ('QA',         2000000, 'Санкт-Петербург'),
    ('DevOps',     2500000, 'Удалённо'),
    ('Analytics',  1500000, 'Казань');   -- в этом отделе не будет сотрудников

-- Сотрудники (10 штук, один без отдела)
INSERT INTO employees (first_name, last_name, email, department_id, hire_date, salary) VALUES
    ('Иван',    'Петров',     'ivan.petrov@company.ru',     1, '2020-03-15', 250000),
    ('Анна',    'Сидорова',   'anna.sidorova@company.ru',   1, '2021-07-01', 200000),
    ('Максим',  'Козлов',     'maxim.kozlov@company.ru',    1, '2022-01-10', 180000),
    ('Елена',   'Новикова',   'elena.novikova@company.ru',  2, '2020-09-20', 190000),
    ('Дмитрий', 'Волков',     'dmitry.volkov@company.ru',   2, '2023-02-14', 160000),
    ('Ольга',   'Морозова',   'olga.morozova@company.ru',   3, '2019-11-01', 170000),
    ('Сергей',  'Лебедев',    'sergey.lebedev@company.ru',  3, '2021-05-20', 150000),
    ('Мария',   'Кузнецова',  'maria.kuznetsova@company.ru',4, '2020-06-15', 220000),
    ('Алексей', 'Соколов',    'alexey.sokolov@company.ru',  4, '2022-08-01', 190000),
    ('Наталья', 'Попова',     'natalya.popova@company.ru', NULL, '2023-06-01', 140000); -- без отдела!

-- Проекты (6 штук, один без назначений)
INSERT INTO projects (name, start_date, end_date, budget) VALUES
    ('Маркетплейс v2',    '2023-01-15', '2023-12-31', 15000000),
    ('Мобильное приложение', '2023-06-01', NULL,       8000000),   -- ещё идёт
    ('Внутренний портал',  '2022-03-01', '2022-12-15', 3000000),
    ('Аналитическая платформа', '2024-01-10', NULL,    12000000),  -- ещё идёт
    ('CI/CD пайплайн',    '2023-04-01', '2023-09-30',  2000000),
    ('Миграция на K8s',   '2024-06-01', NULL,          5000000);   -- без назначений — для проверки

-- Назначения на проекты
INSERT INTO project_assignments (employee_id, project_id, role, assigned_at) VALUES
    (1, 1, 'lead',      '2023-01-15'),
    (2, 1, 'developer', '2023-02-01'),
    (3, 1, 'developer', '2023-03-01'),
    (6, 1, 'tester',    '2023-04-01'),
    (4, 2, 'lead',      '2023-06-01'),
    (5, 2, 'developer', '2023-06-15'),
    (3, 2, 'developer', '2023-07-01'),
    (1, 3, 'lead',      '2022-03-01'),
    (6, 3, 'tester',    '2022-04-01'),
    (8, 4, 'lead',      '2024-01-10'),
    (9, 4, 'developer', '2024-01-15'),
    (2, 4, 'developer', '2024-02-01'),
    (8, 5, 'lead',      '2023-04-01'),
    (9, 5, 'developer', '2023-05-01'),
    (7, 5, 'tester',    '2023-05-15');

-- Клиенты (5 штук, один без заказов)
INSERT INTO clients (name, industry, country) VALUES
    ('ООО «Ромашка»',       'Ритейл',       'Russia'),
    ('ИП Сергеев',           'Логистика',    'Russia'),
    ('TechCorp Inc.',        'IT',           'USA'),
    ('Global Trade Ltd.',    'Финансы',      'UK'),
    ('ООО «Пустышка»',      'Консалтинг',   'Russia');  -- без заказов

-- Заказы (8 штук)
INSERT INTO orders (client_id, project_id, amount, status, created_at, completed_at) VALUES
    (1, 1, 5000000,  'completed',   '2023-02-01', '2023-11-30'),
    (1, 2, 3000000,  'in_progress', '2023-07-01', NULL),
    (2, 1, 2000000,  'completed',   '2023-03-15', '2023-12-20'),
    (3, 4, 8000000,  'in_progress', '2024-01-20', NULL),
    (3, NULL, 1500000, 'new',       '2024-03-01', NULL),         -- заказ без проекта
    (4, 3, 3000000,  'completed',   '2022-05-10', '2022-12-10'),
    (2, 5, 1000000,  'completed',   '2023-05-01', '2023-09-25'),
    (1, 4, 4000000,  'cancelled',   '2024-02-15', NULL);

-- ============================================================
-- ПРОВЕРКА: можно выполнить после старта
-- ============================================================
-- SELECT 'departments' AS tbl, COUNT(*) FROM departments
-- UNION ALL SELECT 'employees', COUNT(*) FROM employees
-- UNION ALL SELECT 'projects', COUNT(*) FROM projects
-- UNION ALL SELECT 'project_assignments', COUNT(*) FROM project_assignments
-- UNION ALL SELECT 'clients', COUNT(*) FROM clients
-- UNION ALL SELECT 'orders', COUNT(*) FROM orders;

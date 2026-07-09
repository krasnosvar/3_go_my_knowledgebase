-- ============================================================
-- Решение задачи 3: Множественные JOIN — сотрудники на проектах
-- ============================================================

-- РЕШЕНИЕ:
SELECT
    e.first_name,           -- имя сотрудника
    e.last_name,            -- фамилия
    p.name AS project_name, -- название проекта
    pa.role                 -- роль на проекте
FROM employees e
    -- 1) Соединяем сотрудников с их назначениями на проекты
    JOIN project_assignments pa ON e.id = pa.employee_id
    -- 2) Соединяем назначения с информацией о проекте
    JOIN projects p ON pa.project_id = p.id
WHERE
    p.end_date IS NULL      -- только активные проекты (без даты окончания)
ORDER BY
    p.name,                 -- сначала по проекту
    pa.role;                -- потом по роли

-- РЕЗУЛЬТАТ:
-- +----------+-----------+---------------------------+-----------+
-- | Мария    | Кузнецова | Аналитическая платформа   | developer |  ← ой, lead!
-- | ...      | ...       | ...                       | ...       |
-- +----------+-----------+---------------------------+-----------+
-- (порядок зависит от сортировки строк — 'developer' < 'lead' < 'tester')

-- ============================================================
-- ПОЯСНЕНИЯ:
-- ============================================================
--
-- Здесь используются 3 таблицы, соединённые через 2 JOIN'а:
--
--   employees ──JOIN──> project_assignments ──JOIN──> projects
--     (кто)        (кто на каком проекте)        (информация о проекте)
--
-- Это типичный паттерн для связи many-to-many:
-- employees <─── project_assignments ───> projects
--
-- project_assignments — это СВЯЗУЮЩАЯ таблица (junction/bridge table).
-- Она содержит FK на обе стороны + дополнительные атрибуты (role, assigned_at).
--
-- ============================================================
-- АЛЬТЕРНАТИВА: с указанием отдела сотрудника
-- ============================================================
-- Можно добавить ещё один JOIN для отдела:
--
-- SELECT e.first_name, e.last_name, d.name AS dept, p.name AS project, pa.role
-- FROM employees e
--     JOIN project_assignments pa ON e.id = pa.employee_id
--     JOIN projects p ON pa.project_id = p.id
--     LEFT JOIN departments d ON e.department_id = d.id
--     -- LEFT JOIN, чтобы не потерять сотрудников без отдела
-- WHERE p.end_date IS NULL
-- ORDER BY p.name, pa.role;

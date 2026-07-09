-- ============================================================
-- Решение задачи 5: Комплексный отчёт по проектам
-- ============================================================

-- РЕШЕНИЕ:
SELECT
    p.name AS project_name,

    -- Статус проекта через CASE
    CASE
        WHEN p.end_date IS NULL THEN 'active'
        ELSE 'completed'
    END AS status,

    -- Количество участников (COUNT по НЕ-NULL значениям)
    COUNT(DISTINCT pa.employee_id) AS team_size,
    --     ^^^^^^^^
    -- DISTINCT нужен, потому что если не использовать подзапросы для заказов,
    -- JOIN с orders может создать дубли строк из project_assignments.

    -- Лид проекта (из подзапроса)
    lead.first_name AS lead_first_name,
    lead.last_name  AS lead_last_name,

    -- Общая сумма заказов (COALESCE для проектов без заказов)
    COALESCE(order_totals.total, 0) AS total_orders

FROM projects p

    -- LEFT JOIN назначений (для подсчёта участников)
    LEFT JOIN project_assignments pa ON p.id = pa.project_id
    -- LEFT JOIN, чтобы проект "Миграция на K8s" (без участников) не пропал

    -- LEFT JOIN для получения лида проекта
    -- Используем подзапрос, чтобы получить ровно 1 строку на проект
    LEFT JOIN LATERAL (
        SELECT e.first_name, e.last_name
        FROM project_assignments pa2
            JOIN employees e ON pa2.employee_id = e.id
        WHERE pa2.project_id = p.id
          AND pa2.role = 'lead'
        LIMIT 1  -- на случай если несколько лидов (защита от дублей)
    ) lead ON TRUE
    -- LATERAL позволяет подзапросу ссылаться на p.id из внешнего запроса.
    -- Это PostgreSQL-специфичная фича, аналог CROSS APPLY в MSSQL.

    -- LEFT JOIN для суммы заказов (предагрегированной)
    LEFT JOIN (
        SELECT project_id, SUM(amount) AS total
        FROM orders
        GROUP BY project_id
    ) order_totals ON p.id = order_totals.project_id
    -- Аграгируем заказы в подзапросе ЗАРАНЕЕ, чтобы избежать
    -- дублирования строк при JOIN с другими таблицами.
    -- Если джоинить orders напрямую, SUM будет некорректной
    -- из-за дублей, созданных JOIN с project_assignments.

GROUP BY
    p.id, p.name, p.end_date,
    lead.first_name, lead.last_name,
    order_totals.total

ORDER BY total_orders DESC;

-- РЕЗУЛЬТАТ:
-- +--------------------------+-----------+------+---------+-----------+--------------+
-- | project_name             | status    | team | lead_fn | lead_ln   | total_orders |
-- +--------------------------+-----------+------+---------+-----------+--------------+
-- | Аналитическая платформа  | active    |    3 | Мария   | Кузнецова |     12000000 |
-- | Маркетплейс v2           | completed |    4 | Иван    | Петров    |      7000000 |
-- | Внутренний портал        | completed |    2 | Иван    | Петров    |      3000000 |
-- | Мобильное приложение     | active    |    3 | Елена   | Новикова  |      3000000 |
-- | CI/CD пайплайн           | completed |    3 | Мария   | Кузнецова |      1000000 |
-- | Миграция на K8s          | active    |    0 | NULL    | NULL      |            0 |
-- +--------------------------+-----------+------+---------+-----------+--------------+

-- ============================================================
-- АЛЬТЕРНАТИВНОЕ РЕШЕНИЕ (без LATERAL — более портативное):
-- ============================================================

-- SELECT
--     p.name AS project_name,
--     CASE WHEN p.end_date IS NULL THEN 'active' ELSE 'completed' END AS status,
--     COALESCE(team.cnt, 0) AS team_size,
--     lead_e.first_name AS lead_first_name,
--     lead_e.last_name AS lead_last_name,
--     COALESCE(ot.total, 0) AS total_orders
-- FROM projects p
--     -- Подсчёт участников в подзапросе
--     LEFT JOIN (
--         SELECT project_id, COUNT(*) AS cnt
--         FROM project_assignments
--         GROUP BY project_id
--     ) team ON p.id = team.project_id
--     -- Поиск лида
--     LEFT JOIN project_assignments lead_pa
--         ON p.id = lead_pa.project_id AND lead_pa.role = 'lead'
--     LEFT JOIN employees lead_e
--         ON lead_pa.employee_id = lead_e.id
--     -- Сумма заказов
--     LEFT JOIN (
--         SELECT project_id, SUM(amount) AS total
--         FROM orders
--         GROUP BY project_id
--     ) ot ON p.id = ot.project_id
-- ORDER BY total_orders DESC;

-- ============================================================
-- КЛЮЧЕВЫЕ МОМЕНТЫ И ЛОВУШКИ:
-- ============================================================
--
-- 1. ПРОБЛЕМА ДУБЛИРОВАНИЯ (fan-out problem):
--    Если JOIN orders напрямую (без подзапроса), а у проекта 4 назначения
--    и 2 заказа → получим 4×2=8 строк → SUM будет в 4 раза больше!
--    Решение: агрегировать в подзапросе ДО JOIN.
--
-- 2. LATERAL JOIN — мощная PostgreSQL-фича:
--    Позволяет подзапросу ссылаться на колонки из предыдущих FROM/JOIN.
--    Без LATERAL подзапрос в JOIN "не видит" внешнюю таблицу.
--
-- 3. COUNT(DISTINCT ...) vs подзапрос:
--    COUNT(DISTINCT pa.employee_id) — простое решение, но работает
--    только если нет fan-out от других JOIN'ов.
--    Подзапрос с COUNT — надёжнее.
--
-- 4. Порядок JOIN'ов не влияет на результат (оптимизатор решает сам),
--    но влияет на читаемость. Ставьте LEFT JOIN'ы в логическом порядке.

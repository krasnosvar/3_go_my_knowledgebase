# Практика SQL JOIN — 5 задач с нарастающей сложностью

## Быстрый старт

```bash
# Запуск PostgreSQL в Docker
docker compose up -d

# Подключение (пароль: devpass)
psql -h localhost -p 5433 -U dev -d join_practice

# Выполнение запроса прямо в контейнере (без интерактивной сессии)
docker exec -it join_practice_db psql -U dev -d join_practice -c "SELECT * FROM departments;"

# Зайти в консоль БД внутри контейнера
docker exec -it join_practice_db psql -U dev -d join_practice

# Или через pgcli (если установлен)
pgcli -h localhost -p 5433 -U dev -d join_practice

# Остановка и удаление данных
docker compose down -v
```

## Схема данных

```
departments ──< employees ──< project_assignments >── projects
                                                         │
clients ──< orders ─────────────────────────────────────┘
```

**6 таблиц:** departments, employees, projects, project_assignments, clients, orders

Таблицы содержат "ловушки" для JOIN'ов:
- Сотрудник без отдела (`department_id IS NULL`)
- Отдел без сотрудников (Analytics)
- Проект без назначений (Миграция на K8s)
- Клиент без заказов (ООО «Пустышка»)
- Заказ без привязки к проекту (`project_id IS NULL`)

## Задачи

| # | Файл | Тема | Сложность |
|---|------|------|-----------|
| 1 | `task_01_inner_join_employees.sql` | INNER JOIN | ★☆☆☆☆ |
| 2 | `task_02_left_join_departments.sql` | LEFT JOIN + COUNT | ★★☆☆☆ |
| 3 | `task_03_multi_join_projects.sql` | JOIN 3 таблиц (many-to-many) | ★★★☆☆ |
| 4 | `task_04_clients_orders.sql` | LEFT JOIN + ON vs WHERE, NOT EXISTS | ★★★★☆ |
| 5 | `task_05_project_report.sql` | Комплексный отчёт, LATERAL, подзапросы | ★★★★★ |

## Решения

Лежат в папке `solutions/`. Открывайте только после попытки решения!

## Ключевые темы для собеседования

- **INNER vs LEFT JOIN** — когда строки пропадают
- **COUNT(\*) vs COUNT(column)** — поведение с NULL
- **SUM + COALESCE** — SUM возвращает NULL, не 0
- **ON vs WHERE в LEFT JOIN** — критичная разница
- **Fan-out problem** — дублирование при множественных JOIN'ах
- **3 способа найти "без связей"** — LEFT JOIN+IS NULL / NOT EXISTS / NOT IN

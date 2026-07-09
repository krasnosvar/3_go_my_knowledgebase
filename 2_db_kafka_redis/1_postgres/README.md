# Базы данных (PostgreSQL): учебный стенд + материалы

Этот каталог — учебный набор по SQL/PostgreSQL:
- **теория** (понятно “как студенту”)
- **практика на Docker Compose** (упражнения от простого к сложному)
- **исходные вопросы/ответы** (в `interview_questions/`)

## Структура

- `docker_compose/`
  - `docker-compose.yaml` — поднимает primary/replica + PgBouncer + pgAdmin
  - `README.md` — как подключаться, команды, статусы, PgBouncer admin
- `exercises/`
  - пошаговые **упражнения** от простого к сложному на твоём стенде:
    - `1_task_db/` — базовый CRUD-сервис на Go (в т.ч. вариант на GORM v2)
    - `2_script_renew_task/` — конкурентное обновление строк, worker pool, генерация 1M пользователей
    - `3_indeces_tasks/` — индексы и оптимизация запросов (`setup.sql`, `practice.sql`)
    - `4_join_tasks/` — задачи на JOIN с решениями (`solutions/`)
- `interview_questions/`
  - исходные материалы (вопросы/ответы/задачи), на которых основаны упражнения
- `postgres_sql_theory.md`
  - конспект **теории**: SQL, ACID, изоляции, MVCC, блокировки, индексы, EXPLAIN, репликация, PgBouncer, CAP

## Как учиться (рекомендуемый порядок)

1) Подними стенд:

```bash
cd docker_compose
docker compose up -d
docker compose ps
```

2) Пройди практику:
- открой `exercises/` и делай упражнения по порядку (`1_task_db` → `2_script_renew_task` → `3_indeces_tasks` → `4_join_tasks`)
- в каждой подпапке есть команды, SQL и ожидаемые наблюдения

3) После каждой практической темы:
- прочитай соответствующий раздел в `postgres_sql_theory.md`
- вернись и попробуй объяснить это своими словами

4) В конце:
- пробегись по `interview_questions/` и попробуй отвечать без подсказок

## Быстрые ссылки

- Практика: `exercises/`
- Теория: `postgres_sql_theory.md`
- Как подключаться/проверять PgBouncer/репликацию: `docker_compose/README.md`
- Вопросы/ответы: `interview_questions/`

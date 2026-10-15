# PostgreSQL: практика

## Зачем

Закрепить теорию из [../postgres_sql_theory.md](../postgres_sql_theory.md) на живой базе: JOIN, индексы и `EXPLAIN`, массовые обновления из Go, CRUD-сервис.

## Пререквизиты

- Docker (с `docker compose`); для Go-заданий — Go 1.22+.
- Стенд из [../docker_compose/](../docker_compose/README.md) (кроме `4_join_tasks` — у него свой compose).
- Прочитанная теория по соответствующей теме.

## Задания (рекомендуемый порядок — от простого к сложному)

1. [4_join_tasks/](./4_join_tasks/README.md) — 5 SQL-задач на JOIN с нарастающей сложностью и решениями. Только SQL, отдельный маленький стенд. Лучший старт.
2. [3_indexes_tasks/](./3_indexes_tasks/README.md) — индексы B-tree/Hash/GIN/GiST/составные/частичные на 1 млн строк, сравнение через `EXPLAIN ANALYZE`.
3. [2_script_renew_task/](./2_script_renew_task/README.md) — обновить 1 млн строк из Go: последовательно, конкурентно, батчами, worker pool. Нужны горутины и `sync`.
4. [1_task_db/](./1_task_db/README.md) — **большое** итоговое задание: схема, CRUD API на SQL и GORM, генерация данных, аналитика, пагинация.

Номера каталогов исторические: `1_task_db` идёт первым по номеру, но по объёму это финальное задание — начинайте с `4_join_tasks`.

## Как запускать

Каждое задание описывает команды в своём README. Go-код лежит в отдельных модулях (`go.mod` внутри каталога), запуск — `go run .` или `go run ./<подкаталог>` из каталога модуля.

## Когда дальше

После практики — [../interview_questions/](../interview_questions/README.md).

# simple_crud_app — CRUD-сервис на Go + PostgreSQL

Учебная реализация первой части задания [../task.md](../task.md): HTTP API поверх таблицы `users` на чистом SQL и на GORM. Пошаговый гайд «как написать самому» — [task-crud.md](./task-crud.md).

## Файлы

- [scripts/schema.sql](./scripts/schema.sql) — создаёт базу `shop`, таблицы `users`/`orders`/`items` и наполняет их 1 млн строк каждая.
- [crud-service/](./crud-service/) — отдельный Go-модуль на `database/sql` + `github.com/lib/pq`, порт `:8080`:
  - `db.go` — `InitDB()`: открытие соединения и `Ping`;
  - `models.go` — структуры, описывающие строки таблиц;
  - `db_functions.go` — функции с SQL-запросами;
  - `handlers.go` — HTTP-хендлеры `/users`, вызывающие функции из `db_functions.go`;
  - `handlers_tx.go` — те же операции в транзакциях (`/users_working_with_transactions`).
- [crud-service-gorm-v2/](./crud-service-gorm-v2/) — то же на GORM v2, порт `:8888`.

## Шаг 1. База

Поднимите стенд из [../../../docker_compose/](../../../docker_compose/README.md) и выполните схему (команды — из каталога `2_db_kafka_redis`):

```bash
cd 1_postgres/docker_compose
docker compose up -d
docker compose exec -T db-primary psql -U postgres -d postgres -v ON_ERROR_STOP=1 < ../exercises/1_task_db/simple_crud_app/scripts/schema.sql
```

## Шаг 2. Запуск сервиса

```bash
cd crud-service
go run .            # строка подключения берётся из DATABASE_URL, по умолчанию postgres://postgres:admin@localhost:5432/shop

# или вариант на GORM
cd ../crud-service-gorm-v2
go run .
```

Как создавался модуль с нуля: `go mod init crud-service`, затем файлы по порядку `db.go` → `models.go` → `db_functions.go` → `handlers.go` → `handlers_tx.go` → `main.go`.

## Шаг 3. Проверка API

Таблица `users` (`crud-service`, порт 8080):

```bash
# Create
curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"order_count": 5}'

# Read
curl -X GET "http://localhost:8080/users?id=1"

# Update
curl -X PUT http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"ID":1,"OrderCount":10}'

# Delete
curl -X DELETE "http://localhost:8080/users?id=1"
```

Те же операции в транзакциях:

```bash
curl -X POST http://localhost:8080/users_working_with_transactions \
     -H "Content-Type: application/json" \
     -d '{"order_count": 5}'

curl -X PUT http://localhost:8080/users_working_with_transactions \
     -H "Content-Type: application/json" \
     -d '{"ID":1,"OrderCount":15}'

curl -X DELETE "http://localhost:8080/users_working_with_transactions?id=1"
```

Вариант на GORM (`crud-service-gorm-v2`, порт 8888):

```bash
curl -X POST http://localhost:8888/users \
     -H "Content-Type: application/json" \
     -d '{"order_count": 5}'

curl -X GET "http://localhost:8888/users?id=1"

curl -X PUT http://localhost:8888/users \
     -H "Content-Type: application/json" \
     -d '{"ID":1,"OrderCount":10}'

curl -X DELETE "http://localhost:8888/users?id=1"
```

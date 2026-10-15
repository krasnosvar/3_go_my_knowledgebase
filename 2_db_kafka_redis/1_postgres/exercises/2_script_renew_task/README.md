# Массовое обновление 1 млн строк из Go

## Задача

Есть таблица `users` с полями `id`, `name`, `order_count` и 1 млн строк. Имена пользователей вида "Den", "Max", "John", "Bob", "Alice", "Ann", "Vasya", "Petya", "Masha", "Olya"; `order_count` у всех 0. Необходимо всем увеличить `order_count` на 1 — и сравнить разные подходы по скорости и нагрузке на БД.

## Пререквизиты

- Go 1.22+, Docker и запущенный стенд из [../../docker_compose/](../../docker_compose/README.md) (PostgreSQL на `localhost:5432`, пароль `admin`).
- Темы: горутины, `sync.WaitGroup`, worker pool ([3_concurrency](../../../../1_golang/3_concurrency/)); батчевые обновления — часть 2, п. 9 в [../../postgres_sql_theory.md](../../postgres_sql_theory.md).

## Файлы

- [create_db_table_users_raws.sql](./create_db_table_users_raws.sql) и [generate_1m_users.sql](./generate_1m_users.sql) — создание таблицы и генерация 1 млн пользователей.
- `go.mod` — отдельный модуль (драйвер `pgx`). Каждый вариант решения — отдельная программа:
  - [1_without_concurrency/main.go](./1_without_concurrency/main.go) — последовательно, без конкурентности;
  - [2_simple_concurrency/main.go](./2_simple_concurrency/main.go) — горутина на батч;
  - [3_proper_query/main.go](./3_proper_query/main.go) — батчевый `UPDATE ... WHERE id BETWEEN` (правильный путь);
  - [4_worker_pool/main.go](./4_worker_pool/main.go) — пул воркеров для сравнения.

## Как запустить

```bash
# из каталога 1_postgres/docker_compose: подготовить данные
docker compose exec -T db-primary psql -U postgres -d postgres -v ON_ERROR_STOP=1 < ../exercises/2_script_renew_task/create_db_table_users_raws.sql

# из каталога 2_script_renew_task: запустить нужный вариант
go run ./1_without_concurrency
go run ./3_proper_query
```

Строка подключения берётся из `DATABASE_URL`, по умолчанию `postgres://postgres:admin@localhost:5432/postgres?sslmode=disable`.

Проверка результата:

```
docker exec postgres-primary psql -U postgres -c "SELECT * FROM users ORDER BY id LIMIT 10;"
 id | name  | order_count 
----+-------+-------------
  1 | Den   |           1
  2 | Alice |           1
  3 | Max   |           1
  4 | Max   |           1
  5 | Olya  |           1
  6 | Olya  |           1
  7 | Max   |           1
  8 | Olya  |           1
  9 | Max   |           1
 10 | Bob   |           1
(10 rows)

```


## Когда использовать что?

✅ **Batch Update (один SQL на батч)**, когда:

- массовое обновление по условию (`SET x = y WHERE ...`);
- вся логика — это SQL (вычисления, проверки);
- нет внешних зависимостей (API, файлы).

```sql
-- Пометить старые заказы
UPDATE orders SET status = 'archived' WHERE created_at < NOW() - INTERVAL '1 year';

-- Пересчитать счётчики
UPDATE users SET order_count = (SELECT COUNT(*) FROM orders WHERE user_id = users.id);

-- Обновить цены с учётом инфляции
UPDATE products SET price = price * 1.05 WHERE category = 'electronics';
```

✅ **Worker Pool в Go**, когда для каждой записи нужно:

- 🌐 вызвать внешний API;
- 📁 прочитать/записать файл;
- 🧮 выполнить сложные вычисления в Go (не SQL);
- 📧 отправить email/уведомление.

```go
// Загрузить аватары пользователей из внешнего API
for userID := range jobs {
    avatar := fetchAvatarFromAPI(userID) // HTTP-запрос
    db.Exec("UPDATE users SET avatar = $1 WHERE id = $2", avatar, userID)
}

// Обработать изображения
for productID := range jobs {
    image := loadImage(productID)
    thumbnail := resizeImage(image) // CPU-intensive операция
    saveThumbnail(thumbnail)
}
```

## 🎓 Золотое правило

> **«Отдавай работу туда, где она выполняется эффективнее»**

```
PostgreSQL ❤️:                    Go ❤️:
- Работа с множествами данных     - Внешние API
- JOIN'ы, GROUP BY, агрегации     - Сложная бизнес-логика
- Обновления по условию           - Работа с файлами
- Транзакционность                - Параллельные вычисления
```

💡 В этой задаче:

```sql
UPDATE users SET order_count = 1 WHERE id BETWEEN 1 AND 100000;
```

- ✅ простое условие (`BETWEEN`);
- ✅ простое значение (константа 1);
- ✅ никаких внешних зависимостей;
- ✅ никаких вычислений.

→ PostgreSQL справится за доли секунды на батч, Go + Worker Pool — заметно медленнее (каждая строка — отдельный запрос по сети).

## Когда дальше

Следующая практика — [../4_join_tasks/](../4_join_tasks/README.md) или [../1_task_db/](../1_task_db/README.md) (см. порядок в [../README.md](../README.md)).

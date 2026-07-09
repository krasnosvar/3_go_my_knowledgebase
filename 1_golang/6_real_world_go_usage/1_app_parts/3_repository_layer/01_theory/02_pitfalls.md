# Repository: pitfalls

## 1. Бизнес-логика в SQL

```sql
SELECT u.*, CASE WHEN u.status = 'active' AND u.last_login > NOW() - INTERVAL '7 days'
    THEN 'engaged' ELSE 'inactive' END AS segment
```

Бизнес-правило «engaged = active + visited in 7 days» теперь живёт в SQL. Service не контролирует.

Лучше: достать данные, посчитать в service. Если проблема perf — оптимизируй явно.

## 2. Не возвращать ErrNotFound

```go
func (r *Repo) Get(id int) (User, error) {
    var u User
    err := db.QueryRow(...).Scan(...)
    return u, err // sql.ErrNoRows утекает наружу
}
```

Service вынужден импортировать `database/sql` и проверять `sql.ErrNoRows`. Маппи на границе:

```go
if errors.Is(err, sql.ErrNoRows) { return User{}, ErrNotFound }
```

## 3. SELECT *

```go
db.Query("SELECT * FROM users")
```

При добавлении колонки `Scan` ломается. Всегда явно перечисляй поля:

```sql
SELECT id, email, created_at FROM users
```

## 4. Не использовать context

```go
db.Query("SELECT ...") // не отменяется
```

Только `*Context` варианты: `QueryContext`, `ExecContext`, `QueryRowContext`. Иначе при отмене сервис продолжает грузить БД.

## 5. SQL injection

```go
db.Query("SELECT * FROM users WHERE name = '" + name + "'") // дыра
```

Параметризованные запросы:

```go
db.QueryContext(ctx, "SELECT ... WHERE name = $1", name)
```

Драйвер escape'ит автоматически, инъекция невозможна.

## 6. N+1

```go
for _, u := range users {
    o, _ := orderRepo.GetByUser(ctx, u.ID) // N запросов
}
```

100 пользователей → 100 запросов. Для каждого — round trip. Решения:

- JOIN в одном запросе.
- `WHERE user_id IN (...)` + group в Go.
- DataLoader pattern для batched fetch.

## 7. Соединения не возвращаются в pool

```go
rows, _ := db.Query(...)
// забыли rows.Close()
```

Connection остаётся занятым. Под нагрузкой пул исчерпывается. Всегда:

```go
rows, err := db.QueryContext(ctx, ...)
if err != nil { return err }
defer rows.Close()
```

## 8. Транзакция без таймаута

```go
tx, _ := db.BeginTx(ctx, nil) // ctx без timeout
// долгая работа
```

Зависший клиент держит row-level locks, блокирует других. Всегда context.WithTimeout перед BeginTx, обязательно `defer tx.Rollback()`.

## 9. Migrations отдельно от кода

Если миграции лежат «где-то на CI», легко отстать от схемы. Embed их в бинарник через `//go:embed` и применяй на старте — детерминировано.

## 10. ORM везде

Тяжёлый ORM (gorm, ent) удобен на старте, но скрывает SQL. Когда нужен fine control над запросами — пишешь сырой SQL рядом, и получаешь два mental model.

Для большинства Go-проектов: `pgx` или `database/sql` + `sqlx` или `sqlc` (codegen из SQL → типизированный код). Меньше магии, больше контроля.

## Правило

Repo — границы между Go-типами и SQL. Каждая утечка SQL-семантики в service — повод задуматься.

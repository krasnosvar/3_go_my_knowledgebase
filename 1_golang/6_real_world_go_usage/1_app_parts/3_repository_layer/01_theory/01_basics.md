# Repository layer

Repository — слой между service и хранилищем. Скрывает детали БД (SQL, NoSQL, API), отдаёт типизированные методы.

## Базовый контракт

```go
type UserRepo interface {
    Get(ctx context.Context, id int) (User, error)
    GetByEmail(ctx context.Context, email string) (User, error)
    Save(ctx context.Context, u *User) error
    Delete(ctx context.Context, id int) error
    List(ctx context.Context, filter Filter, page Pagination) ([]User, error)
}
```

Интерфейс описывает потребности **service** (см. 2_service_layer). Реализация — отдельно (postgres, in-memory для тестов).

## Реализация для PostgreSQL

```go
type postgresUserRepo struct {
    db *sql.DB
}

func (r *postgresUserRepo) Get(ctx context.Context, id int) (User, error) {
    var u User
    row := r.db.QueryRowContext(ctx,
        `SELECT id, email, created_at FROM users WHERE id = $1`, id)
    if err := row.Scan(&u.ID, &u.Email, &u.CreatedAt); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return User{}, ErrNotFound
        }
        return User{}, fmt.Errorf("get user %d: %w", id, err)
    }
    return u, nil
}
```

## Маппинг ошибок

Низкоуровневые ошибки (sql.ErrNoRows, pgErr с конкретным кодом) **на границе репозитория** превращаются в доменные:

- `sql.ErrNoRows` → `ErrNotFound`.
- pgError 23505 (unique violation) → `ErrAlreadyExists`.
- pgError 23503 (FK violation) → `ErrConflict`.

Service знает только про доменные. Это снимает зависимость от драйвера.

## Транзакции

Repo предоставляет методы, которые можно использовать как **внутри**, так и **вне** транзакции:

```go
type DBTX interface {
    QueryRowContext(...) *sql.Row
    ExecContext(...) (sql.Result, error)
    // ...
}

func (r *Repo) Get(ctx context.Context, db DBTX, id int) (User, error) { ... }
```

Service сам решает, открывать транзакцию или нет.

Альтернатива — `TxRunner`:

```go
type TxRunner interface {
    RunTx(ctx context.Context, fn func(ctx context.Context) error) error
}
```

Внутри `fn` контекст содержит транзакцию, repo автоматически её использует.

## Pagination

Универсального решения нет, но типичные подходы:

- **Offset-based**: `LIMIT 20 OFFSET 40`. Просто, но плохо масштабируется (offset медленно растёт).
- **Cursor-based**: `WHERE id > last_seen ORDER BY id LIMIT 20`. Стабильно при добавлениях, требует упорядочивания.

```go
type Pagination struct {
    Limit  int
    Cursor string // id или encoded multi-key
}
```

## Не делать в repo

- Бизнес-логику ("если статус active, то ...").
- HTTP / транспорт.
- Side effects кроме записи (sending emails, kicking jobs).
- Логику между несколькими сущностями (это сервис).

## Тестирование

- Unit-тесты сервиса используют fake repo (in-memory map).
- Integration-тесты repo — на реальной БД (testcontainers, отдельная схема).

## Правило

Repo — **тонкая прослойка над хранилищем**. Знает SQL, не знает бизнес.

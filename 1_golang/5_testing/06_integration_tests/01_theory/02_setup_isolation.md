# Setup и изоляция

## Проблема общего state

Если два теста бегут на одной БД и один заполняет таблицу, второй увидит чужие данные. Отсюда — flaky тесты, зависимость от порядка, ад при добавлении новых.

## Стратегии изоляции

### 1. Один контейнер на тест

```go
func TestX(t *testing.T) {
    pg := startPostgres(t)
    defer pg.Stop()
    // ...
}
```

Самый простой, но медленный. Подходит для редких особо чувствительных тестов.

### 2. Один контейнер на пакет, схема на тест

```go
func TestMain(m *testing.M) {
    pg = startPostgres()
    defer pg.Stop()
    os.Exit(m.Run())
}

func TestX(t *testing.T) {
    schema := "test_" + randomID()
    db := pg.connectSchema(schema)
    t.Cleanup(func() { db.Exec("DROP SCHEMA " + schema + " CASCADE") })
    // ...
}
```

Быстро (один процесс БД), изолировано (отдельная схема).

### 3. Truncate в Cleanup

```go
func TestX(t *testing.T) {
    db := globalDB()
    t.Cleanup(func() { db.Exec("TRUNCATE users, orders RESTART IDENTITY CASCADE") })
    // ...
}
```

Быстро, но не работает при `t.Parallel()` — два параллельных теста очистят друг другу данные.

### 4. Транзакция-обёртка

```go
func TestX(t *testing.T) {
    tx, _ := db.Begin()
    t.Cleanup(func() { tx.Rollback() })
    repo := NewRepo(tx)
    // ...
}
```

Каждый тест в своей транзакции, всегда rollback. Очень быстро, идеально для CRUD-тестов.

Не работает, если код использует свои транзакции — `BEGIN` внутри уже не сработает.

## Миграции

Перед тестами обязательно — реальные миграции (`golang-migrate`, `goose`):

```go
func runMigrations(t *testing.T, db *sql.DB) {
    t.Helper()
    m, _ := migrate.NewWithDatabaseInstance(...)
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        t.Fatal(err)
    }
}
```

Тесты должны работать на той же схеме, что прод.

## Параллельность

`t.Parallel()` + integration-тесты с реальной БД работает только при правильной изоляции.

Если два параллельных теста используют общую таблицу без транзакции — будет race и flakes.

## Skip без условий

```go
if testing.Short() {
    t.Skip("skipping integration test in -short mode")
}
```

`go test -short` — пропускает интеграционные тесты. Используется в pre-commit, чтобы быстро проверить unit.

## Build tags для отделения

```go
//go:build integration
```

Без тега — не компилируется. Запуск: `go test -tags=integration ./...`. CI прогоняет с тегом, локально — без.

## Правило

- Изоляция через схему/транзакцию — стандарт для БД.
- Каждый тест должен быть запускаем в любом порядке с любым набором других тестов.
- Длинные интеграционные — за тегом или флагом, чтобы не тормозить обычный workflow.

# Integration tests: основы

Unit-тест проверяет одну функцию изолированно. Integration-тест проверяет, что **несколько компонентов работают вместе**: сервис + БД, handler + middleware, продюсер + Kafka.

## Когда нужны

- Реальная SQL-логика (`pgx`, `database/sql`): запросы, транзакции, индексы.
- HTTP-роутинг + middleware-цепочка.
- Контракты с очередями (Kafka, NATS, Redis).
- End-to-end сценарии: «зарегистрировался → получил welcome email».

## Когда unit достаточно

- Чистая бизнес-логика без внешних зависимостей.
- Алгоритмы, валидаторы, парсеры.
- Service + mocked repo.

## Базовый пример: HTTP

```go
func TestHandler(t *testing.T) {
    rr := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/health", nil)

    router().ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("status = %d; want 200", rr.Code)
    }
}
```

`httptest.NewRecorder` и `httptest.NewServer` — стандартные инструменты для HTTP-интеграционных тестов без поднятия сети.

## С реальной БД

Через `testcontainers-go`:

```go
func TestRepoIntegration(t *testing.T) {
    ctx := context.Background()
    pg, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:16-alpine"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("user"),
        postgres.WithPassword("pass"),
    )
    if err != nil { t.Fatal(err) }
    t.Cleanup(func() { _ = pg.Terminate(ctx) })

    dsn, _ := pg.ConnectionString(ctx)
    db, _ := sql.Open("postgres", dsn)
    runMigrations(t, db)

    repo := New(db)
    // ... вызываем методы repo, проверяем
}
```

Контейнер поднимается, миграции запускаются, тесты идут на реальной БД.

## Изоляция тестов

Каждый integration test — на своей БД (или на схеме). Иначе тесты друг другу мешают.

Варианты:

- Один контейнер на пакет, отдельная схема в каждом тесте.
- Один контейнер на тест (медленнее, но проще).
- Truncate таблиц в `t.Cleanup`.

## Скорость

Integration-тесты в 10–100 раз медленнее unit. Поэтому:

- Юнит покрывает основную логику, integration — критические сценарии.
- Тег `integration` для отделения:

```go
//go:build integration

package repo

func TestRealDB(t *testing.T) { ... }
```

Запуск: `go test -tags=integration ./...`. Без тега тесты пропускаются — локально и в CI разный набор.

## Правило

- Unit покрывает 80% логики.
- Integration ловит то, что в unit невозможно увидеть: SQL-ошибки, race на реальном клиенте, баги middleware-цепочек.
- Integration без mock — иначе это всё ещё unit.

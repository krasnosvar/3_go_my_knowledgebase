# Service layer

Service (use-case) — слой, где живёт бизнес-логика. Между handler'ом и репозиторием.

## Зачем нужен

- **Изоляция бизнес-правил**: транспорт (HTTP/gRPC) и хранилище (PG/Mongo) могут меняться, бизнес-логика — нет.
- **Тестируемость**: handler через mock service, service через mock repo. Чистые unit-тесты на ядро без внешних зависимостей.
- **Композиция**: один service может использовать несколько repo, других service'ов, внешних клиентов.

## Базовая форма

```go
type UserService interface {
    Register(ctx context.Context, email string) (User, error)
    Block(ctx context.Context, id int) error
}

type userService struct {
    repo   UserRepo
    mailer Mailer
    clock  Clock
}

func NewUserService(repo UserRepo, mailer Mailer, clock Clock) UserService {
    return &userService{repo: repo, mailer: mailer, clock: clock}
}

func (s *userService) Register(ctx context.Context, email string) (User, error) {
    if !validEmail(email) {
        return User{}, ErrInvalidEmail
    }
    u := User{Email: email, CreatedAt: s.clock.Now()}
    if err := s.repo.Save(ctx, &u); err != nil {
        return User{}, fmt.Errorf("save: %w", err)
    }
    _ = s.mailer.SendWelcome(ctx, u.Email) // не критично
    return u, nil
}
```

## Что внутри service

- Валидация бизнес-правил (не форматов — это handler).
- Координация: «сначала проверь, потом списать, потом отправить event».
- Транзакции (через repo).
- Обогащение данных, вычисления.

## Что НЕ внутри service

- HTTP-специфичный код: статусы, body, query.
- SQL / ORM / драйверы БД.
- Логи в stdout (используй `slog` через зависимость).
- Глобальные синглтоны.

## Зависимости через интерфейсы

Service принимает интерфейсы, не конкретные типы:

```go
type UserRepo interface {
    Save(ctx context.Context, u *User) error
    Get(ctx context.Context, id int) (User, error)
}
```

В тестах подсовывается fake. В проде — `postgresUserRepo`.

## Возврат

- Конкретный тип (`User`, не интерфейс): «принимай интерфейсы, возвращай структуры».
- `error` всегда последний.
- Sentinel-ошибки или typed errors на бизнес-условия (`ErrAlreadyExists`).

## Транзакции

Если несколько операций должны быть атомарны — открывай транзакцию в service:

```go
err := s.tx.RunTx(ctx, func(ctx context.Context) error {
    if err := s.repo.Lock(ctx, id); err != nil { return err }
    if err := s.repo.Update(ctx, ...); err != nil { return err }
    return s.repo.Insert(ctx, ...)
})
```

Конкретная реализация — через интерфейс `TxRunner` или подобное, чтобы оставить service чистым.

## Правило

Service — про бизнес. Если в нём приходится знать про SQL, HTTP или конкретный движок очереди — это утечка слоя.

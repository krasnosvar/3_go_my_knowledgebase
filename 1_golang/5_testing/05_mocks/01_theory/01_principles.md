# Mocks: принципы

В Go моки — это просто реализация интерфейса. Никаких frameworks обязательно не нужно.

## Базовая идея

Сервис принимает интерфейс — в проде подсовываем настоящую реализацию, в тесте — поддельную.

```go
type UserRepo interface {
    Get(ctx context.Context, id int) (User, error)
}

type Service struct{ repo UserRepo }

func (s *Service) Greet(ctx context.Context, id int) (string, error) {
    u, err := s.repo.Get(ctx, id)
    if err != nil { return "", err }
    return "hello " + u.Name, nil
}
```

В тесте:

```go
type fakeRepo struct {
    user User
    err  error
}

func (r fakeRepo) Get(_ context.Context, _ int) (User, error) {
    return r.user, r.err
}

func TestGreet(t *testing.T) {
    svc := &Service{repo: fakeRepo{user: User{Name: "den"}}}
    got, _ := svc.Greet(context.Background(), 1)
    if got != "hello den" { t.Fatal("wrong greeting") }
}
```

Полминуты кода — и тест без БД, без сети, мгновенный.

## Когда моки нужны

- Зависимости с внешним state (БД, HTTP, файлы, время).
- Дорогие или флаки операции в unit-тестах.
- Сценарии ошибок, которые сложно воспроизвести (timeout, network error).

## Когда моки НЕ нужны

- Маленькие чистые функции — тестируй напрямую.
- Логика, которая не пересекает интерфейс — мок ничего не даст.
- Когда мок переписывает половину тестируемого кода — это сигнал, что зависимости плохо разделены.

## Где определять интерфейс

В Go идиома: **интерфейс у потребителя**. Сервис описывает «мне нужен такой UserRepo», конкретный `postgresRepo` живёт в другом пакете и про этот интерфейс не знает.

```go
// service/service.go
type UserRepo interface { Get(ctx, id) (User, error) }

// repo/postgres.go
type PostgresRepo struct { ... }
func (r *PostgresRepo) Get(...) // по утиной типизации удовлетворяет интерфейсу
```

Это снимает циклы импортов и упрощает тесты.

## Размер интерфейса

Маленький. 1–3 метода. Если для теста нужно «реализовать 12 методов» — интерфейс слишком большой.

## Hand-made vs generated

- Маленькие моки — пиши руками. 5 строк, всё видно.
- Сложные с проверкой вызовов / порядка — `gomock`, `mockery`. Но это сторонняя зависимость и шаблонный код.

В большинстве случаев hand-made + пара флажков достаточно.

## Правило

Маленькие интерфейсы → простые fake-структуры. Если без mock-фреймворка не справляешься — посмотри, не велик ли интерфейс или не пытаешься ли ты тестировать слишком много за один раз.

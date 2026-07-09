# Handmade fakes

## Простой fake

```go
type fakeRepo struct {
    user User
    err  error
}

func (r fakeRepo) Get(_ context.Context, _ int) (User, error) {
    return r.user, r.err
}
```

Покрывает 80% случаев. Для каждого теста — свой экземпляр с нужными значениями.

## Fake с памятью

Часто нужен «чуть умнее» fake, который записывает поведение или возвращает разные ответы:

```go
type fakeRepo struct {
    items map[int]User
    saveCalls int
}

func (r *fakeRepo) Get(_ context.Context, id int) (User, error) {
    u, ok := r.items[id]
    if !ok { return User{}, ErrNotFound }
    return u, nil
}

func (r *fakeRepo) Save(_ context.Context, u User) error {
    r.items[u.ID] = u
    r.saveCalls++
    return nil
}
```

Тест проверяет:

```go
if r.saveCalls != 1 { t.Errorf("Save called %d times", r.saveCalls) }
```

## Fake с управляемой ошибкой

```go
type fakeRepo struct {
    nextErr error
}

func (r *fakeRepo) Get(...) (User, error) {
    if r.nextErr != nil {
        err := r.nextErr
        r.nextErr = nil
        return User{}, err
    }
    return User{...}, nil
}
```

Удобно симулировать «первый вызов падает, второй ок».

## Spy

Spy — fake, который **записывает аргументы** для последующей проверки:

```go
type spyMailer struct {
    sent []string
}

func (s *spyMailer) Send(to string) error {
    s.sent = append(s.sent, to)
    return nil
}
```

В тесте `len(s.sent)` и `s.sent[0]` показывают, что и сколько раз было вызвано.

## Stub

Stub — fake, возвращающий **фиксированный ответ**, без логики. Самый простой тип:

```go
type stubClock struct{}

func (stubClock) Now() time.Time {
    return time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
}
```

Полезно для замораживания времени, рандома, конфигов.

## Расположение

Маленькие fake'и — в том же `_test.go` файле. Большие переиспользуемые — в `testutil` пакете внутри проекта.

```
internal/
  user/
    service.go
    service_test.go
  testutil/
    fakes.go
```

## Не превращай fake в реальную реализацию

Если fake усложняется до 200 строк бизнес-логики — что-то не так. Возможно, тестируется слишком многое за раз, или интерфейс слишком жирный.

## Правило

- 5-15 строк fake — идеал.
- Поведение fake предсказуемо, состояние видно.
- Для каждого теста — fresh экземпляр.

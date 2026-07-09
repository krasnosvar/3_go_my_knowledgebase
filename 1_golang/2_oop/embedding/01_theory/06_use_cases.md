# Use-cases embedding

## 1. Декораторы для сервисов

Встраивание — лучший способ «обернуть» тип и поменять часть поведения без дубликатов.

```go
type loggingRepo struct {
    UserRepo
    log *slog.Logger
}

func (r loggingRepo) Save(ctx context.Context, u User) error {
    r.log.Info("save user", "id", u.ID)
    return r.UserRepo.Save(ctx, u)
}
```

Все методы `UserRepo` автоматически доступны, а `Save` мы заменили своим — с логами.

## 2. Расширения типов из stdlib

Часто нужно добавить поля к `http.Request`, `sql.Tx`, `os.File`:

```go
type FancyFile struct {
    *os.File
    bytesRead int64
}

func (f *FancyFile) Read(p []byte) (int, error) {
    n, err := f.File.Read(p)
    f.bytesRead += int64(n)
    return n, err
}
```

Внешний тип автоматически `io.Reader`, потому что встроенный — `io.Reader`.

## 3. Композиция маленьких «миксинов»

Логгер, метрики, clock как отдельные встраиваемые типы:

```go
type Service struct {
    Logger
    Metrics
    Clock
    // ...
}
```

Каждый миксин — простая структура с методами, а сервис получает их все.

## 4. DSL и fluent API

В builder-паттернах бывает удобно встраивать «набор опций»:

```go
type HTTPOptions struct {
    Timeout time.Duration
    Retry   int
}

type Client struct {
    HTTPOptions
    baseURL string
}
```

## 5. Интерфейсная композиция

Большой интерфейс — сумма маленьких:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

То же embedding, но для интерфейсов. Стандартный способ наращивать контракт.

## 6. Поддельная «наследственность» (осторожно)

Иногда хочется «Base» с общими методами, «наследующийся» в потомках:

```go
type Base struct{ id int }
func (b Base) ID() int { return b.id }

type User struct { Base }
type Order struct { Base }
```

Формально работает, но это плохой путь. В Go лучше делать общий интерфейс `Identifiable` или явные функции. Пожалей читателя кода.

## Когда НЕ использовать

- Когда в голове всплывает слово «наследование» — скорее всего, лучше интерфейс.
- Если у встраиваемого типа слишком много экспортируемых методов, а тебе нужна пара — проще поля.
- Конфликты имён — сигнал, что конструкция избыточна.

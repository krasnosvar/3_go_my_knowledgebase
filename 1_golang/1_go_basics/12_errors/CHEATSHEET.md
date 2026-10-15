# Ошибки — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии указаны там, где поведение зависит от них.

## Основы

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Задача | Код |
| :--- | :--- |
| Интерфейс | `type error interface { Error() string }` |
| Простая ошибка | `errors.New("division by zero")` |
| С подстановкой | `fmt.Errorf("invalid age: %d", age)` |
| Сигнатура | `func f() (T, error)` — `error` последним; успех → `return v, nil` |
| Проверка (early return) | `if err != nil { return err }` |
| Ресурс | `f, err := os.Open(p); if err != nil { return err }; defer f.Close()` |
| Стиль текста | с маленькой буквы, без точки: `"open config: no such file"` |
| Нулевое значение | `var err error` → `nil` — «ошибки нет» |

## Оборачивание и разбор цепочки

Подробно: [02_wrapping_is_as.md](./01_theory/02_wrapping_is_as.md).

| Задача | Код |
| :--- | :--- |
| Обернуть с контекстом (1.13+) | `fmt.Errorf("save user %d: %w", id, err)` |
| Несколько `%w` в одном вызове (1.20+) | `fmt.Errorf("a: %w, b: %w", e1, e2)` |
| Только текст, без цепочки | `fmt.Errorf("...: %v", err)` — `Is`/`As` уже не найдут |
| Сравнить со значением по цепочке | `errors.Is(err, ErrNotFound)` |
| Достать тип из цепочки | `var ve *ValidationError; if errors.As(err, &ve) { ve.Field }` |
| То же, generic (1.26+) | `ve, ok := errors.AsType[*ValidationError](err)` |
| Снять один слой | `errors.Unwrap(err)` (не раскрывает `Join` и несколько `%w`) |
| Объединить (1.20+) | `errors.Join(e1, e2)` — `nil` пропускаются; `Join(nil, nil) == nil` |

| | `errors.Is` | `errors.As` |
| :--- | :--- | :--- |
| Ищет | конкретное **значение** (sentinel) | значение конкретного **типа** |
| Второй аргумент | `target error` | указатель на переменную: `&ve` |
| Результат | `bool` | `bool` + заполняет переменную |
| Типичный пример | `errors.Is(err, sql.ErrNoRows)`, `context.Canceled` | `*pgconn.PgError` → `pgErr.Code == "23505"` |

```go
var ErrNotFound = errors.New("not found")

err := fmt.Errorf("get user 42: %w", ErrNotFound)
fmt.Println(err)                                             // get user 42: not found
fmt.Println(errors.Is(err, ErrNotFound), err == ErrNotFound) // true false

j := errors.Join(ErrNotFound, nil, errors.New("age: neg"))
fmt.Printf("%q\n", j.Error())          // "not found\nage: neg"
fmt.Println(errors.Is(j, ErrNotFound)) // true
```

## Свои ошибки

Подробно: [03_custom_errors.md](./01_theory/03_custom_errors.md).

| Вид | Объявление | Проверка |
| :--- | :--- | :--- |
| Sentinel | `var ErrNotFound = errors.New("not found")` | `errors.Is` |
| Типизированная | `type ValidationError struct{ Field, Message string }` + `func (e *ValidationError) Error() string` | `errors.As` |
| С причиной | поле `Err error` + `func (e *DBError) Unwrap() error { return e.Err }` | `Is`/`As` идут сквозь неё |

```go
type NotFoundError struct{}

func (e *NotFoundError) Error() string { return "not found" }

func check() error {
	var e *NotFoundError // nil-указатель
	return e             // ⚠️ интерфейс (type=*NotFoundError, value=nil) != nil
}

fmt.Println(check() != nil) // true — возвращайте явный nil
```

## panic vs error

Подробно: [04_panic_recover_boundary.md](./01_theory/04_panic_recover_boundary.md).

| Ситуация | Чем |
| :--- | :--- |
| Неверный ввод, сеть, БД, файлы | `error` |
| Нарушен инвариант («не может случиться») | `panic` |
| Не удалась обязательная инициализация (regexp-константа, порт) | `panic` / `log.Fatal` при старте |
| Handler/воркер упал — не ронять процесс | `defer` + `recover` на границе → `error` / 500 |
| Panic в чужой горутине | ловить только в ней самой: свой `defer recover` в каждой долгоживущей |

## Паттерны и use-cases

Подробно: [05_patterns.md](./01_theory/05_patterns.md), [06_use_cases.md](./01_theory/06_use_cases.md).

| Слой | Что делать с ошибкой |
| :--- | :--- |
| Библиотека / репозиторий | sentinel или тип; мапить `sql.ErrNoRows` → свой `ErrNotFound` |
| Сервис | оборачивать `%w` контекстом операции и ID |
| Транспорт (HTTP/gRPC) | `errors.Is`/`As` → статус-код и текст для пользователя |
| Ретраи | явно: `errors.Is(err, ErrTransient)` или интерфейс `Temporary() bool` |
| Отмена | `context.Canceled`, `context.DeadlineExceeded` — нормальное завершение, не шуметь в логах |
| Валидация нескольких полей | `errors.Join(validateName(n), validateAge(a))` |

Цепочка читается как трасса: `create user: save user 42: insert user: duplicate key`.

## Ловушки

- ⚠️ `err == ErrX` ломается после оборачивания — только `errors.Is`.
- ⚠️ `%v` вместо `%w` рвёт цепочку: `Is`/`As` перестают находить причину.
- ⚠️ `errors.As` со значением вместо указателя (`errors.As(err, ve)`) → panic; нужен `&ve`.
- ⚠️ Typed nil: вернули `*MyErr(nil)` как `error` → `err != nil` истинно.
- ⚠️ Сравнение текста ошибок (`err.Error() == "..."`) — хрупко.
- ⚠️ `_ = doIt()` и пустой `if err != nil {}` — проглоченная ошибка.
- ⚠️ Лог + возврат одной и той же ошибки → дубли в логах; либо лог, либо возврат.
- ⚠️ `recover` без `defer` или в другой горутине не сработает.

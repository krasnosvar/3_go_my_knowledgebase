# context.WithValue

`WithValue` прикрепляет к контексту пару ключ-значение. Используется, чтобы передать через цепочку функций данные, которые по смыслу — не бизнес-параметры: request ID, user ID из middleware, trace span, tenant и т. п.

## Синтаксис

```go
type ctxKey string

const userKey ctxKey = "user"

ctx := context.WithValue(parent, userKey, user)

// где-то ниже
if u, ok := ctx.Value(userKey).(User); ok {
    // ...
}
```

## Типизированный ключ

Используй **свой тип** для ключа, не `string`. Это защищает от коллизий между пакетами.

```go
type ctxKey struct{ name string }

var userKey = ctxKey{"user"}
```

`ctxKey{...}` из разных пакетов никогда не совпадут.

## Что класть в контекст

Хорошие кандидаты:

- request ID / trace ID;
- authenticated user / role;
- tenant / organization;
- logger с уже привязанными полями;
- deadline-related помощники.

## Чего НЕ класть

- Бизнес-параметры функции. Они идут обычными аргументами.
- Опции, меняющие поведение: логика должна быть видима в сигнатуре.
- Большие объекты или mutable-state.
- Дескрипторы ресурсов (БД, HTTP-клиенты) — это зависимости, им место в структуре сервиса.

## Производительность

`ctx.Value(key)` ищет ключ, идя по цепочке контекстов. Сложность линейная от длины цепочки. В hot path это заметно — не клади туда часто читаемые мелочи.

## Типовой паттерн middleware

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        u := authenticate(r)
        ctx := context.WithValue(r.Context(), userKey, u)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    if u, ok := r.Context().Value(userKey).(User); ok {
        // используем
    }
}
```

## Типовой паттерн хелперов

Оборачивай `Value` в типизированный аксессор, чтобы не повторять assertion:

```go
func UserFrom(ctx context.Context) (User, bool) {
    u, ok := ctx.Value(userKey).(User)
    return u, ok
}
```

Снаружи остаётся простой API, внутри — один assert.

## Рекомендация

- `WithValue` — для cross-cutting concerns.
- Делай типизированные ключи и аксессоры.
- Документируй, что передаётся через контекст, и не расширяй этот список без нужды.

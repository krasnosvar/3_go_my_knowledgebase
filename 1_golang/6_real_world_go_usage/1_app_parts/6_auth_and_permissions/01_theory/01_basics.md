# Auth and Permissions

Аутентификация — кто ты. Авторизация — что тебе можно. Это разные вещи и обычно разные слои.

## Authentication (authn)

Способы:

- **Session cookies** — сервер хранит сессию, клиент шлёт session ID в cookie. Подходит для web-фронтов.
- **JWT (Bearer token)** — токен подписан, сервер не хранит. Удобно для API и микросервисов.
- **OAuth 2.0 / OIDC** — делегированная аутентификация (Google, GitHub).
- **API keys** — длинная случайная строка для service-to-service.
- **mTLS** — клиентские сертификаты, для строгой межсервисной аутентификации.

## JWT-pattern

```go
import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    UserID int    `json:"uid"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func parseToken(s string) (Claims, error) {
    token, err := jwt.ParseWithClaims(s, &Claims{}, func(t *jwt.Token) (any, error) {
        return secret, nil
    })
    if err != nil { return Claims{}, err }
    c, ok := token.Claims.(*Claims)
    if !ok || !token.Valid { return Claims{}, ErrInvalidToken }
    return *c, nil
}
```

JWT хранит данные в payload — не клади туда секреты.

## Authorization (authz)

Стратегии:

- **Roles**: «admin», «editor», «viewer». Просто, но негибко.
- **Permissions**: каждое действие = строковый код, у роли набор. Компромисс.
- **ABAC**: атрибут-ориентированно, политики типа «можно если owner == user.id и status != deleted». Гибко, сложно.
- **RBAC + scopes**: классика для API — токен содержит `scope: read:users write:orders`.

## В коде

Auth middleware кладёт user в context. Authz обычно — на handler / service-уровне:

```go
func canEditOrder(u User, o Order) bool {
    return u.Role == "admin" || (u.Role == "user" && o.UserID == u.ID)
}

func handler(w, r) {
    u := userFrom(r.Context())
    o := ...
    if !canEditOrder(u, o) {
        http.Error(w, "forbidden", 403)
        return
    }
    // ...
}
```

Или объявление политик в декларативной форме (через casbin, opa, и т. п.) — но для большинства проектов проще руками.

## Хранение паролей

- НЕ plain text. НЕ MD5/SHA-1. Только медленные функции с солью.
- `bcrypt` (default), `argon2id` (современнее, рекомендуется OWASP).
- Соль и параметры (cost, time/memory) хранятся в hash-строке самого алгоритма.

```go
import "golang.org/x/crypto/bcrypt"
hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
err := bcrypt.CompareHashAndPassword(hash, []byte(input))
```

## Защита от brute-force

- Rate limit на login endpoint (per-IP и per-user).
- CAPTCHA после N неудачных попыток.
- Аудит-лог всех неуспешных логинов.
- Уведомление пользователю о входе с нового устройства/IP.

## Refresh tokens

Access token короткий (15 мин), refresh — длинный (дни). Refresh хранится в HttpOnly cookie или secure storage. Этим решается trade-off «безопасность vs UX».

## Правило

- Не выдумывай свою криптографию.
- Используй стандартные libraries (`golang.org/x/crypto`, `github.com/golang-jwt/jwt`).
- Authn и authz — разные слои.
- Все секреты — через Vault / Secret Manager, никогда в коде.

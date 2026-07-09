# Auth/Authz: pitfalls

## 1. Хранение паролей в plain или быстрым хешем

```go
hash := sha256.Sum256([]byte(password))
```

SHA — слишком быстрый. GPU подбирает миллиарды/сек. Только bcrypt / argon2 / scrypt.

## 2. Своя реализация токенов

«Я просто base64 подпишу HMAC'ом, JWT слишком тяжёлый» → 99% делают это с уязвимостями. Используй проверенную библиотеку.

## 3. JWT с алгоритмом none

В JWT можно указать `alg: none` — без подписи. Если парсер не проверяет алгоритм явно — токен с `none` проходит. **Проверяй expected algorithm**:

```go
jwt.ParseWithClaims(s, &Claims{}, func(t *jwt.Token) (any, error) {
    if t.Method.Alg() != "HS256" { return nil, ErrUnexpectedAlg }
    return secret, nil
})
```

## 4. Секреты в JWT payload

JWT — **подписан**, не зашифрован. Payload читается любым, кто видит токен. Не клади туда email, телефоны, secrets.

## 5. Длинные access токены

Access живёт сутки → украденный токен годен сутки. Refresh-pattern: access короткий (15 мин), refresh длинный с возможностью отозвать.

## 6. Нет rate limit на login

Брутфорс пароля без ограничений. Минимум: per-IP rate limit на login endpoint.

## 7. Сравнение токенов через `==`

```go
if token == expected { ... } // timing attack
```

Атакующий по времени ответа угадывает токен по символам. Используй `subtle.ConstantTimeCompare`:

```go
if subtle.ConstantTimeCompare([]byte(token), []byte(expected)) == 1 { ... }
```

## 8. Authorization in handler без интерфейса

```go
func handler(w, r) {
    u := userFromCtx(r.Context())
    if u.Role != "admin" && u.Role != "moderator" {
        http.Error(w, "forbidden", 403); return
    }
    // ...
}
```

В трёх местах — три разных проверки. Вынеси в функцию `canEdit(u, resource) bool` и используй везде.

## 9. CSRF-незащищённые мутации

Если используешь cookie-based auth → CSRF атаки. Решения:

- SameSite=Strict cookies.
- CSRF-токен в каждом мутирующем запросе (POST/PUT/DELETE).
- Для API с Bearer token CSRF не актуален.

## 10. Логирование секретов

```go
slog.Info("login", "user", email, "password", password) // в логах prod
```

Никогда. Маскируй или просто не логируй чувствительные поля.

## 11. Permission check на front, не на back

«Кнопка скрыта в UI» — не security. Backend обязан проверять каждый endpoint.

## 12. RBAC превратился в IF-цепочку

```go
if u.Role == "admin" { /* ... */ }
else if u.Role == "user" && o.UserID == u.ID { /* ... */ }
else if u.Role == "moderator" && o.Status == "pending" { /* ... */ }
```

Растёт с каждым новым use-case. Вынеси в политики (decision functions), или возьми OPA / casbin.

## Правило

- bcrypt / argon2 для паролей.
- Подпись JWT обязательна, алгоритм проверяется явно.
- Authz в одном месте, не размазана по handler'ам.
- Любые секреты — не в логах, не в JWT payload, не в коде.

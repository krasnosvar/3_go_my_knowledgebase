# Упражнения: Auth and permissions

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_bearer_token
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_bearer_token_answers
```

## Список задач (от простого к сложному)

1. [`01_bearer_token`](./01_bearer_token/main.go) — authn-middleware по Bearer-токену; решение: [`01_bearer_token_answers`](../04_exercises_answers/01_bearer_token_answers/main.go).
2. [`02_permission_bits`](./02_permission_bits/main.go) — права как битовые флаги и роли; решение: [`02_permission_bits_answers`](../04_exercises_answers/02_permission_bits_answers/main.go).
3. [`03_hmac_token`](./03_hmac_token/main.go) — подписанный токен (идея JWT на HMAC-SHA256); решение: [`03_hmac_token_answers`](../04_exercises_answers/03_hmac_token_answers/main.go).
4. [`04_ownership`](./04_ownership/main.go) — authz на уровне ресурса (ownership); решение: [`04_ownership_answers`](../04_exercises_answers/04_ownership_answers/main.go).

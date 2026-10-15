# Auth and permissions

Аутентификация (кто ты) и авторизация (что тебе можно): Bearer-токены, подпись HMAC, роли и права на уровне ресурса.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#6-auth-and-permissions) — раздел «6. Auth and permissions» общей шпаргалки частей приложения.

## Пререквизиты

[5_middleware_chain](../5_middleware_chain/), битовые операции из [1_go_basics/04_operators_bitwise](../../../1_go_basics/04_operators_bitwise/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — Auth and Permissions.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Auth/Authz: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_role_check](./02_examples/01_role_check/main.go) — проверка роли пользователя.
    *   [02_token_claim_shape](./02_examples/02_token_claim_shape/main.go) — claims токена как структура.
    *   [03_permission_set](./02_examples/03_permission_set/main.go) — набор разрешений как множество.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_bearer_token](./03_exercises/01_bearer_token/main.go) — authn-middleware по Bearer-токену.
    *   [02_permission_bits](./03_exercises/02_permission_bits/main.go) — права как битовые флаги и роли.
    *   [03_hmac_token](./03_exercises/03_hmac_token/main.go) — подписанный токен (идея JWT на HMAC-SHA256).
    *   [04_ownership](./03_exercises/04_ownership/main.go) — authz на уровне ресурса (ownership).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/6_auth_and_permissions` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/6_auth_and_permissions/02_examples/01_role_check
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/6_auth_and_permissions/03_exercises/01_bearer_token
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/6_auth_and_permissions/04_exercises_answers/01_bearer_token_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [7_cache_layer](../7_cache_layer/) — Cache layer.

# Упражнения: Backend API

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_routes
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_routes_answers
```

## Список задач (от простого к сложному)

1. [`01_routes`](./01_routes/main.go) — маршруты REST-ресурса на stdlib ServeMux (Go 1.22+); решение: [`01_routes_answers`](../04_exercises_answers/01_routes_answers/main.go).
2. [`02_json_response`](./02_json_response/main.go) — единые хелперы JSON-ответов; решение: [`02_json_response_answers`](../04_exercises_answers/02_json_response_answers/main.go).
3. [`03_error_handling`](./03_error_handling/main.go) — handler, возвращающий error; решение: [`03_error_handling_answers`](../04_exercises_answers/03_error_handling_answers/main.go).
4. [`04_middleware_chain`](./04_middleware_chain/main.go) — цепочка logging + auth с прерыванием; решение: [`04_middleware_chain_answers`](../04_exercises_answers/04_middleware_chain_answers/main.go).
5. [`05_repo_abstraction`](./05_repo_abstraction/main.go) — handler зависит от интерфейса хранилища; решение: [`05_repo_abstraction_answers`](../04_exercises_answers/05_repo_abstraction_answers/main.go).
6. [`06_context_timeout`](./06_context_timeout/main.go) — таймаут на вызов зависимости; решение: [`06_context_timeout_answers`](../04_exercises_answers/06_context_timeout_answers/main.go).
7. [`07_pagination`](./07_pagination/main.go) — пагинация списка в API; решение: [`07_pagination_answers`](../04_exercises_answers/07_pagination_answers/main.go).
8. [`08_healthcheck`](./08_healthcheck/main.go) — liveness и readiness; решение: [`08_healthcheck_answers`](../04_exercises_answers/08_healthcheck_answers/main.go).

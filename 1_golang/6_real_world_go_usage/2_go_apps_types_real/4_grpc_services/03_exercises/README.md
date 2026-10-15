# Упражнения: gRPC services

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_status_codes
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_status_codes_answers
```

## Список задач (от простого к сложному)

1. [`01_status_codes`](./01_status_codes/main.go) — gRPC status codes; решение: [`01_status_codes_answers`](../04_exercises_answers/01_status_codes_answers/main.go).
2. [`02_deadline_handling`](./02_deadline_handling/main.go) — бюджет времени из deadline клиента; решение: [`02_deadline_handling_answers`](../04_exercises_answers/02_deadline_handling_answers/main.go).
3. [`03_request_validation`](./03_request_validation/main.go) — валидация запроса с перечнем нарушений; решение: [`03_request_validation_answers`](../04_exercises_answers/03_request_validation_answers/main.go).
4. [`04_interceptor_chain`](./04_interceptor_chain/main.go) — цепочка unary-интерсепторов; решение: [`04_interceptor_chain_answers`](../04_exercises_answers/04_interceptor_chain_answers/main.go).
5. [`05_auth_metadata`](./05_auth_metadata/main.go) — аутентификация по metadata; решение: [`05_auth_metadata_answers`](../04_exercises_answers/05_auth_metadata_answers/main.go).
6. [`06_stream_loop`](./06_stream_loop/main.go) — server-streaming цикл; решение: [`06_stream_loop_answers`](../04_exercises_answers/06_stream_loop_answers/main.go).
7. [`07_retry_policy`](./07_retry_policy/main.go) — политика ретраев клиента; решение: [`07_retry_policy_answers`](../04_exercises_answers/07_retry_policy_answers/main.go).
8. [`08_health_rpc`](./08_health_rpc/main.go) — сервис здоровья (grpc.health.v1); решение: [`08_health_rpc_answers`](../04_exercises_answers/08_health_rpc_answers/main.go).

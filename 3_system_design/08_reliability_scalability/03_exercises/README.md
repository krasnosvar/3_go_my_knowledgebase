# Упражнения: reliability & scalability

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`.

Все проверки детерминированы: время и случайность управляемы (логические тики, переданный `now`, фиксированный seed), поэтому тесты не флакают.

Запуск шаблона:

```bash
go run ./03_exercises/01_token_bucket_allow/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_token_bucket_allow_answers/main.go
```

## Список задач (от простого к сложному)

1. `01_token_bucket_allow` — реализовать `Allow()` для token bucket (списать токен или отказать).
2. `02_circuit_breaker_transitions` — реализовать переходы состояний circuit breaker в `OnResult` (closed/open/half-open).
3. `03_backoff_sequence` — вычислить последовательность задержек exponential backoff с потолком.
4. `04_call_with_timeout` — обернуть вызов таймаутом через `context` (результат либо `ctx.Err()`).
5. `05_fixed_window_limiter` — реализовать fixed-window rate limiter на управляемом времени (**Хардкор**).

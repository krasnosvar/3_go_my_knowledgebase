# Упражнения: context

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_cancel/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_cancel_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_cancel](./01_cancel/main.go) — `WithCancel` и ожидание отмены. Решение: [01_cancel_answers](../04_exercises_answers/01_cancel_answers/main.go).
2. [02_timeout](./02_timeout/main.go) — операция с `WithTimeout`. Решение: [02_timeout_answers](../04_exercises_answers/02_timeout_answers/main.go).
3. [03_deadline](./03_deadline/main.go) — абсолютный дедлайн через `WithDeadline`. Решение: [03_deadline_answers](../04_exercises_answers/03_deadline_answers/main.go).
4. [04_propagation](./04_propagation/main.go) — отмена родителя отменяет потомка, но не наоборот. Решение: [04_propagation_answers](../04_exercises_answers/04_propagation_answers/main.go).
5. [05_value](./05_value/main.go) — request-id с типизированным ключом. Решение: [05_value_answers](../04_exercises_answers/05_value_answers/main.go).
6. [06_worker_shutdown](./06_worker_shutdown/main.go) — worker останавливается по `ctx.Done()`. Решение: [06_worker_shutdown_answers](../04_exercises_answers/06_worker_shutdown_answers/main.go).
7. [07_http_request](./07_http_request/main.go) — HTTP-запрос с контекстом (`httptest`-сервер). Решение: [07_http_request_answers](../04_exercises_answers/07_http_request_answers/main.go).
8. [08_pipeline_cancel](./08_pipeline_cancel/main.go) — конвейер, завершающийся по отмене. Решение: [08_pipeline_cancel_answers](../04_exercises_answers/08_pipeline_cancel_answers/main.go).

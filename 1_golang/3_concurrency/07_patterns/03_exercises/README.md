# Упражнения: patterns

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_worker_pool/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_worker_pool_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_worker_pool](./01_worker_pool/main.go) — пул фиксированного размера с сохранением порядка результатов. Решение: [01_worker_pool_answers](../04_exercises_answers/01_worker_pool_answers/main.go).
2. [02_fan_out](./02_fan_out/main.go) — N worker'ов читают один вход, у каждого свой выход. Решение: [02_fan_out_answers](../04_exercises_answers/02_fan_out_answers/main.go).
3. [03_fan_in](./03_fan_in/main.go) — слияние каналов с остановкой по `done`. Решение: [03_fan_in_answers](../04_exercises_answers/03_fan_in_answers/main.go).
4. [04_pipeline](./04_pipeline/main.go) — конвейер parse → square → sum. Решение: [04_pipeline_answers](../04_exercises_answers/04_pipeline_answers/main.go).
5. [05_semaphore](./05_semaphore/main.go) — ограничение параллелизма буферизованным каналом. Решение: [05_semaphore_answers](../04_exercises_answers/05_semaphore_answers/main.go).
6. [06_rate_limit](./06_rate_limit/main.go) — ограничение частоты через `time.Ticker`. Решение: [06_rate_limit_answers](../04_exercises_answers/06_rate_limit_answers/main.go).
7. [07_bounded_queue](./07_bounded_queue/main.go) — очередь фиксированной ёмкости с отказом при переполнении. Решение: [07_bounded_queue_answers](../04_exercises_answers/07_bounded_queue_answers/main.go).
8. [08_graceful_shutdown](./08_graceful_shutdown/main.go) — пул worker'ов с `Shutdown(ctx)`. Решение: [08_graceful_shutdown_answers](../04_exercises_answers/08_graceful_shutdown_answers/main.go).

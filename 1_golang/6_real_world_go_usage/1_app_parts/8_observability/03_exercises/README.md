# Упражнения: Observability

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_structured_log
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_structured_log_answers
```

## Список задач (от простого к сложному)

1. [`01_structured_log`](./01_structured_log/main.go) — структурные логи через log/slog; решение: [`01_structured_log_answers`](../04_exercises_answers/01_structured_log_answers/main.go).
2. [`02_counter_exposition`](./02_counter_exposition/main.go) — счётчик с лейблами и вывод в формате Prometheus; решение: [`02_counter_exposition_answers`](../04_exercises_answers/02_counter_exposition_answers/main.go).
3. [`03_latency_histogram`](./03_latency_histogram/main.go) — гистограмма латентности (как у Prometheus); решение: [`03_latency_histogram_answers`](../04_exercises_answers/03_latency_histogram_answers/main.go).
4. [`04_trace_propagation`](./04_trace_propagation/main.go) — проброс trace ID между сервисами; решение: [`04_trace_propagation_answers`](../04_exercises_answers/04_trace_propagation_answers/main.go).

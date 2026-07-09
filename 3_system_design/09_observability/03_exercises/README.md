# Упражнения: observability

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`.

Запуск шаблона:

```bash
go run ./03_exercises/01_red_metrics/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_red_metrics/main.go
```

## Список задач (от простого к сложному)

1. `01_red_metrics` — по списку запросов посчитать RED: rate (rps), долю ошибок и среднюю длительность.
2. `02_prometheus_format` — отформатировать метрику в одну строку Prometheus-экспозиции `name{k="v",...} value`.
3. `03_readiness_handler` — readiness-хендлер, отвечающий 200/503 в зависимости от флага готовности (`atomic.Bool`).
4. `04_parse_log_line` — разобрать строку структурного лога (logfmt `key=value`) в `map[string]string`.

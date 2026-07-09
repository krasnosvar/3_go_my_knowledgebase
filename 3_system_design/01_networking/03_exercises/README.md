# Упражнения: networking

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`.

Запуск шаблона:

```bash
go run ./03_exercises/01_status_category/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_status_category/main.go
```

## Список задач (от простого к сложному)

1. `01_status_category` — по коду HTTP-статуса вернуть категорию (1xx..5xx / unknown).
2. `02_parse_url` — разобрать URL на части через `net/url` (scheme, host, port, path, query).
3. `03_tcp_line_protocol` — строчный протокол поверх TCP: прочитать строку, вернуть в верхнем регистре.
4. `04_http_handler` — HTTP-хэндлер с корректным статусом (200/400) и заголовком Content-Type.
5. `05_udp_sum_server` — UDP-сервер, суммирующий числа из одной датаграммы (хардкор).

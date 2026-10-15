# Упражнения: net/http, net

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, без внешней сети.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_status_codes/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_status_codes_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_status_codes`](./01_status_codes/main.go) — маппинг доменных ошибок в HTTP-коды через `errors.Is`.
2. [`02_json_handler`](./02_json_handler/main.go) — JSON-хендлер: `Content-Type`, 405 + `Allow` для не-GET.
3. [`03_query_params`](./03_query_params/main.go) — разбор `page`/`limit` с дефолтами и границами.
4. [`04_client_timeout`](./04_client_timeout/main.go) — `http.Client{Timeout}`, проверка статуса, закрытие тела.
5. [`05_retry_client`](./05_retry_client/main.go) — повторы на 5xx и сетевых ошибках, 4xx не повторяем.
6. [`06_tcp_echo_shape`](./06_tcp_echo_shape/main.go) — TCP echo-сервер на `net` + `bufio.Scanner`.
7. [`07_graceful_stop`](./07_graceful_stop/main.go) — `Shutdown` дожидается запросов на лету, `ErrServerClosed` — не ошибка.
8. [`08_health_endpoint`](./08_health_endpoint/main.go) — `/healthz` (liveness) и `/readyz` (readiness, 503).

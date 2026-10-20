# 05. net/http, net

`net/http` — основа большинства Go-сервисов: хендлеры, роутинг, клиенты с
таймаутами, graceful shutdown, health-эндпоинты. Пакет `net` — уровень ниже:
TCP-слушатели и соединения. Все упражнения работают локально (`httptest`,
`127.0.0.1:0`) без внешней сети.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — сервер
(`ServeMux` 1.22+, таймауты, shutdown), клиент, `net` TCP/UDP, ловушки.

## 🧭 Пререквизиты

- [03_encoding_json_csv_base64](../03_encoding_json_csv_base64/) — JSON
- [04_io_bufio_os_filepath](../04_io_bufio_os_filepath/) —
  `io.Reader`/`Writer`, `bufio`
- [3_concurrency/01_goroutines](../../3_concurrency/01_goroutines/) и
  [3_concurrency/05_context](../../3_concurrency/05_context/) — горутины и
  `context` для таймаутов и остановки

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_http_server.md](./01_theory/01_http_server.md) — сервер:
        `Handler`/`HandlerFunc`, ответ, middleware, таймауты, graceful
        shutdown.
    *   [02_http_client.md](./01_theory/02_http_client.md) — клиент: почему не
        `DefaultClient`, `Request`, проверка статуса, `Transport`, ретраи.
    *   [03_net_basics.md](./01_theory/03_net_basics.md) — `net`:
        `Dial`/`Listen`, таймауты и контексты, UDP, unix-сокеты, частые
        ошибки.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib,
    завершаются сами):
    *   [01_http_handler_shape](./02_examples/01_http_handler_shape/main.go) —
        хендлер, проверенный через `httptest`.
    *   [02_http_client_shape](./02_examples/02_http_client_shape/main.go) —
        GET к локальному `httptest`-серверу с проверкой ошибки до `resp`.
    *   [03_net_listener_shape](./02_examples/03_net_listener_shape/main.go) —
        TCP-слушатель на свободном порту.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками
    `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_status_codes](./03_exercises/01_status_codes/main.go) — маппинг
        доменных ошибок в HTTP-коды через `errors.Is`.
    *   [02_json_handler](./03_exercises/02_json_handler/main.go) —
        JSON-хендлер: `Content-Type`, 405 + `Allow` для не-GET.
    *   [03_query_params](./03_exercises/03_query_params/main.go) — разбор
        `page`/`limit` с дефолтами и границами.
    *   [04_client_timeout](./03_exercises/04_client_timeout/main.go) —
        `http.Client{Timeout}`, проверка статуса, закрытие тела.
    *   [05_retry_client](./03_exercises/05_retry_client/main.go) — повторы на
        5xx и сетевых ошибках, 4xx не повторяем.
    *   [06_tcp_echo_shape](./03_exercises/06_tcp_echo_shape/main.go) — TCP
        echo-сервер на `net` + `bufio.Scanner`.
    *   [07_graceful_stop](./03_exercises/07_graceful_stop/main.go) —
        `Shutdown` дожидается запросов на лету, `ErrServerClosed` — не ошибка.
    *   [08_health_endpoint](./03_exercises/08_health_endpoint/main.go) —
        `/healthz` (liveness) и `/readyz` (readiness, 503).
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения
    тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Команды выполняются из папки темы
(`1_golang/4_standard_library/05_net_http_net`).

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с
   [01_http_server.md](./01_theory/01_http_server.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_http_handler_shape/main.go
   ```
3. Решите задачи в `03_exercises/` и запустите самопроверку (в стартовом
   состоянии печатает `FAIL`):
   ```bash
   go run ./03_exercises/01_status_codes/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_status_codes_answers/main.go
   ```

## ➡️ Дальше

[5_testing](../../5_testing/) — тесты для такого кода (`httptest`,
table-driven); сетевые основы — в
[3_system_design/01_architecture_styles](../../../3_system_design/01_architecture_styles/).

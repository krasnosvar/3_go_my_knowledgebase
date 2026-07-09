# 02. API Styles (Стили API)

Как сервисы общаются между собой и с клиентами: основные стили API (REST, gRPC/Protobuf, GraphQL), их сильные и слабые стороны и осознанный выбор под задачу. Теория подкреплена runnable-примерами на стандартной библиотеке Go (только REST — gRPC/GraphQL требуют внешних зависимостей и разобраны теоретически) и упражнениями с самопроверкой.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_rest.md](./01_theory/01_rest.md) — REST: ресурсы и URL, HTTP-методы и идемпотентность (safe/idempotent), коды статусов по операциям, версионирование (path/header/query), пагинация (offset/limit vs cursor).
    *   [02_grpc.md](./01_theory/02_grpc.md) — gRPC: RPC поверх HTTP/2, четыре типа вызовов (unary + server/client/bidi streaming), генерация кода из `.proto`, где выигрывает и в чём слаб.
    *   [03_protobuf.md](./01_theory/03_protobuf.md) — Protobuf: схема (`.proto`), бинарное кодирование (TLV/varint), эволюция сообщений и обратная совместимость, неприкосновенность номеров полей и `reserved`.
    *   [04_graphql.md](./01_theory/04_graphql.md) — GraphQL: схема (SDL), query/mutation, решение over/under-fetching, проблема N+1 на сервере и DataLoader.
    *   [05_choosing.md](./01_theory/05_choosing.md) — как выбирать: REST vs gRPC vs GraphQL (внешние API / межсервис / гибкие клиенты), дерево решений и гибридные архитектуры.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические применения: типичные сценарии, correctness-паттерны (идемпотентность/retry, таймауты, пагинация, статус-коды, эволюция контракта) и частые ошибки.
*   **[02_examples/](./02_examples/)** — runnable-примеры (каждый поднимает сервер через `httptest` на `127.0.0.1:0`, обменивается данными и завершается; только stdlib):
    *   [01_rest_crud](./02_examples/01_rest_crud/main.go) — маленький REST-ресурс in-memory: GET/POST/PUT/DELETE над `/users`, коды 200/201/204/404/405, заголовок `Location`.
    *   [02_pagination](./02_examples/02_pagination/main.go) — пагинация `limit`/`offset` через query-параметры: клиент проходит коллекцию страницами по ссылке `next`.
    *   [03_idempotency](./02_examples/03_idempotency/main.go) — идемпотентность: повтор PUT даёт то же состояние, повтор POST создаёт дубликат, а POST с `Idempotency-Key` дедуплицируется.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_method_semantics` — по имени HTTP-метода вернуть флаги safe/idempotent.
    *   `02_paginate` — собрать пагинированный ответ (окно `[offset, offset+limit)` + метаданные и `HasNext`).
    *   `03_rest_status` — REST-хэндлер `/users/{id}` с корректными кодами 200/204/404/405 и заголовком `Allow`.
    *   `04_parse_filter` — разбор query-строки фильтра (`q`, `limit`, `active`) с дефолтами через `net/url` и `strconv`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_rest.md](./01_theory/01_rest.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, без внешних зависимостей и без `go.mod`):
   ```bash
   go run ./02_examples/01_rest_crud/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/01_method_semantics/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

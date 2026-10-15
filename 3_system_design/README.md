# System Design Knowledge Base

Этот раздел — личная база для системного изучения System Design: сеть, API, кэширование, масштабирование данных, обмен сообщениями, архитектура, принципы проектирования, надёжность, наблюдаемость и подготовка к system design интервью.

Акцент — на практике, применимой в Go-бэкенде: не абстрактная теория, а то, что реально помогает проектировать сервисы и проходить собеседования.

## Формат подтем

Каждая подтема устроена одинаково:

- `01_theory/` — краткая теория по подтемам (без воды);
- `02_examples/` — runnable-примеры (только stdlib);
- `03_exercises/` — упражнения от простого к сложному с самопроверкой `PASS`/`FAIL`;
- `04_exercises_answers/` — эталонные решения (`NN_name_answers/`).

## Пререквизиты

- Пройдены [1_go_basics](../1_golang/1_go_basics/) и [2_oop](../1_golang/2_oop/); для 05/08 полезны [3_concurrency](../1_golang/3_concurrency/).
- Примеры и упражнения запускаются из папки подтемы: `go run ./03_exercises/01_.../main.go`.

## Подразделы

1. [01_networking](01_networking/) — сеть: TCP/UDP, модель OSI/TCP-IP, HTTP/HTTPS, HTTP/2 и HTTP/3, DNS, TLS.
2. [02_api_styles](02_api_styles/) — стили API: REST, gRPC, Protobuf, GraphQL и выбор между ними.
3. [03_caching](03_caching/) — кэширование: паттерны, TTL/eviction, типичные проблемы, Redis как кэш.
4. [04_data_scaling](04_data_scaling/) — масштабирование данных: репликация, шардирование, CAP, согласованность, SQL vs NoSQL.
5. [05_messaging](05_messaging/) — обмен сообщениями: очереди и брокеры, гарантии доставки, outbox/inbox, backpressure.
6. [06_architecture](06_architecture/) — архитектура: монолит vs микросервисы, Clean Architecture, DDD, DI, API Gateway.
7. [07_design_principles](07_design_principles/) — принципы проектирования: SOLID, DRY/KISS/YAGNI, паттерны, 12-factor.
8. [08_reliability_scalability](08_reliability_scalability/) — надёжность и масштабируемость: балансировка, rate limiting, circuit breaker, retries, graceful shutdown.
9. [09_observability](09_observability/) — наблюдаемость: логи, метрики (RED/USE), трейсинг, пробы, алертинг.
10. [10_design_interview](10_design_interview/) — подход к system design интервью и список типовых задач.

## 🚀 Как проходить

- Идти по порядку от `01_networking` к `10_design_interview`: фундамент (сеть, API) → данные и сообщения → архитектура и принципы → эксплуатация (надёжность, наблюдаемость) → интервью.
- Разделы 01–02 — база, на которую опирается всё остальное; не пропускать.
- Разделы 03–05 — как масштабировать состояние и связывать сервисы.
- Разделы 06–07 — как раскладывать код и сервисы по границам.
- Разделы 08–09 — что нужно, чтобы система жила в проде.
- Раздел 10 — сборка всего вместе: тренировка на типовых задачах.


# Типы реальных Go-приложений

Как устроены типичные приложения на Go и какие задачи в них решаются. В каждой папке: обзорная теория, runnable-примеры, 8 упражнений с самопроверкой `PASS`/`FAIL` и решения. Внешние библиотеки (cobra, grpc-go, client-go, kafka-клиенты) не используются: ключевая логика смоделирована на стандартной библиотеке, чтобы всё запускалось офлайн.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — что выбрать, CLI, backend и graceful shutdown, демоны, gRPC-коды, retry/cron, операторы, streaming, ловушки.

## Пререквизиты

[1_app_parts](../1_app_parts/) и [3_concurrency](../../3_concurrency/).

## Темы (по порядку)

1. [1_cli_tools/](./1_cli_tools/) — CLI: `flag.FlagSet`, подкоманды, env fallback, коды выхода, stdout/stderr, dry-run, help.
2. [2_backend/](./2_backend/) — HTTP backend (REST): маршруты, JSON-ответы, ошибки, middleware, хранилище, таймауты, пагинация, health.
3. [3_system_apps/](./3_system_apps/) — системные утилиты и демоны: сигналы, остановка, PID-файл, пути, логи, ФС, внешние процессы.
4. [4_grpc_services/](./4_grpc_services/) — gRPC: status codes, deadlines, валидация, интерсепторы, metadata, стриминг, ретраи, health.
5. [5_workers_cron/](./5_workers_cron/) — воркеры и cron: очереди, пул, backoff, DLQ, идемпотентность, graceful stop.
6. [6_k8s_operators/](./6_k8s_operators/) — Kubernetes operator: reconcile, conditions, finalizers, requeue, ownerReferences.
7. [7_streaming_data/](./7_streaming_data/) — стриминг (Kafka/NATS): оффсеты, at-least-once, poison messages, партиции, ребаланс, lag.

## Как запускать

```bash
cd 1_cli_tools
go run ./03_exercises/01_arg_parse                  # заготовка → FAIL
go run ./04_exercises_answers/01_arg_parse_answers  # решение → PASS
```

## Когда дальше

[7_under_the_hood](../../7_under_the_hood/) или практика с реальными сервисами в [2_db_kafka_redis](../../../2_db_kafka_redis/).

# App Parts

Строительные блоки реального Go-приложения. Каждый блок — отдельная тема с теорией, примерами, упражнениями с самопроверкой `PASS`/`FAIL` и решениями. Всё на стандартной библиотеке, без сети и внешних сервисов.

## Пререквизиты

[1_go_basics](../../1_go_basics/), интерфейсы из [2_oop](../../2_oop/), `context` и `sync` из [3_concurrency](../../3_concurrency/), `net/http` из [4_standard_library](../../4_standard_library/).

## Темы (по порядку)

1. [1_http_handlers/](./1_http_handlers/) — тонкий handler: маппинг ошибок в статусы, query-параметры, JSON, маршрутизация `ServeMux` (4 упражнения).
2. [2_service_layer/](./2_service_layer/) — бизнес-правила, перевод ошибок репозитория, инъекция часов (4 упражнения).
3. [3_repository_layer/](./3_repository_layer/) — контракт репозитория, маппинг ошибок драйвера, offset и keyset пагинация (4 упражнения).
4. [4_config_and_env/](./4_config_and_env/) — окружение и дефолты, типизированный разбор, валидация, приоритет источников (4 упражнения).
5. [5_middleware_chain/](./5_middleware_chain/) — порядок цепочки, recovery, request ID, перехват статуса ответа (4 упражнения).
6. [6_auth_and_permissions/](./6_auth_and_permissions/) — Bearer-токен, права битовыми флагами, HMAC-токен, ownership (4 упражнения).
7. [7_cache_layer/](./7_cache_layer/) — кеш-декоратор, инвалидация на записи, negative caching, TTL с jitter (4 упражнения).
8. [8_observability/](./8_observability/) — `slog`, счётчики и гистограммы в формате Prometheus, проброс trace ID (4 упражнения).

## Как запускать

Из каталога темы (модуль — в корне репозитория):

```bash
cd 1_http_handlers
go run ./02_examples/01_health_handler
go run ./03_exercises/01_status_mapping            # заготовка → FAIL
go run ./04_exercises_answers/01_status_mapping_answers  # решение → PASS
```

## Когда дальше

[2_go_apps_types_real/](../2_go_apps_types_real/) — как эти части собираются в приложения разных типов.

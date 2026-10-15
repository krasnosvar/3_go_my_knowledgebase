# Real world Go usage

Этот раздел про Go **в контексте реального рабочего приложения**. Материалы не разложены по пакетам стандартной библиотеки (testing/http/db…), потому что в проекте они всегда идут вместе. Вместо этого:

- в [1_app_parts/](./1_app_parts/) собраны **части приложения** (handler, service, repository, config, middleware, auth, cache, observability) — кубики, которые проектируют через интерфейсы и тестируют по отдельности;
- в [2_go_apps_types_real/](./2_go_apps_types_real/) — **типы Go-приложений** (CLI, backend API, демоны, gRPC, воркеры, k8s-операторы, стриминг) и их типичные задачи.

## Пререквизиты

[1_go_basics](../1_go_basics/), [2_oop](../2_oop/) (интерфейсы), [3_concurrency](../3_concurrency/), [4_standard_library](../4_standard_library/), желательно [5_testing](../5_testing/).

## Порядок

1. [1_app_parts/](./1_app_parts/) — 8 тем по порядку, в каждой: теория, примеры, 4 упражнения с самопроверкой и решения.
2. [2_go_apps_types_real/](./2_go_apps_types_real/) — 7 типов приложений, в каждом 8 упражнений с самопроверкой и решения.
3. [2_go_apps_types_real/tasks_from_interview/](./2_go_apps_types_real/tasks_from_interview/) — задачи с собеседований (нужна сеть).

## Как запускать

Всё — пакеты корневого модуля репозитория, внешних зависимостей нет:

```bash
# из корня репозитория
go run ./1_golang/6_real_world_go_usage/1_app_parts/1_http_handlers/03_exercises/01_status_mapping
go vet ./1_golang/6_real_world_go_usage/...
```

Заготовка упражнения печатает `FAIL`, решённая — `PASS`. Решения лежат в `04_exercises_answers/<имя>_answers/`.

## Когда дальше

После раздела — [7_under_the_hood](../7_under_the_hood/) (runtime, память, профилирование) или практика с реальными БД и брокерами в [2_db_kafka_redis](../../2_db_kafka_redis/).

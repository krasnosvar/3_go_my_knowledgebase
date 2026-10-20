# Тема 02. Производительность и масштабирование в Go (Performance & Scaling)

## О чём эта тема (краткий контекст)

Высокая пропускная способность и предсказуемо низкая латентность (p99) в Go
достигаются пониманием того, как рантайм управляет горутинами, памятью и
сборщиком мусора в контейнерах. В этом модуле мы разбираем масштабирование
планировщика `G-M-P`, настройку `GOMAXPROCS` и `GOMEMLIMIT`, боевую диагностику
узких мест через `pprof`, `runtime/trace` и `log/slog`, а также устройство
высококонкурентных in-memory кэшей и защиту от cache stampede через
`singleflight`.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [01. Узкие места, закон Литтла и методология профилирования](../../../5_devops_sre_knowledgebase/14_system_design/02_performance_caching_and_scaling/01_theory/01_bottlenecks_and_profiling_methodology.md)
> - [02. Многоуровневое кэширование, паттерны и аномалии кэша](../../../5_devops_sre_knowledgebase/14_system_design/02_performance_caching_and_scaling/01_theory/02_caching_levels_and_patterns.md)
> - [03. Масштабирование High-Load систем (C10k, автоскейлинг, Kubernetes)](../../../5_devops_sre_knowledgebase/14_system_design/02_performance_caching_and_scaling/01_theory/03_scaling_high_load_systems.md)
> - Наблюдаемость и мониторинг:
>   [8_observability](../../../5_devops_sre_knowledgebase/8_observability/README.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — `GOMAXPROCS` и
`GOMEMLIMIT`, команды `pprof` и `trace`, `log/slog`, шардированный L1-кэш,
`singleflight`, ловушки GC.

## Пререквизиты

- [3_concurrency](../../1_golang/3_concurrency/) (`sync.Mutex`, `sync.RWMutex`,
  каналы, `context`).
- [7_under_the_hood](../../1_golang/7_under_the_hood/) (устройство рантайма и
  GC).
- [01_architecture_styles](../01_architecture_styles/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория Go-специфики:
    *   [01_go_scheduler_memory_and_scaling.md](./01_theory/01_go_scheduler_memory_and_scaling.md)
        — планировщик `G-M-P`, троттлинг CPU в контейнерах (`GOMAXPROCS`),
        защита от OOMKill (`GOMEMLIMIT`), escape analysis и влияние аллокаций
        на p99.
    *   [02_observability_and_profiling_in_go.md](./01_theory/02_observability_and_profiling_in_go.md)
        — профилирование в продакшене через `net/http/pprof` (CPU, heap,
        allocs, goroutine, mutex), `runtime/trace`, `runtime/metrics` и
        zero-alloc логирование `log/slog`.
    *   [03_caching_in_go.md](./01_theory/03_caching_in_go.md) —
        потокобезопасный
        L1-кэш (`RWMutex`, шардированная мапа, `sync.Map`), подавление cache
        stampede через `singleflight` и обход сканирования указателей GC на
        больших кэшах.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_structured_logging](./02_examples/01_structured_logging/main.go) —
        структурное JSON-логирование через `log/slog` с пробросом `request_id`
        и `trace_id`.
    *   [02_inprocess_metrics](./02_examples/02_inprocess_metrics/main.go) —
        потокобезопасный сбор RED-метрик (Rate, Errors, Duration) и расчёт p50,
        p95, p99.
    *   [03_health_readiness](./02_examples/03_health_readiness/main.go) —
        разделение проб `/healthz` (liveness) и `/readyz` (readiness).
    *   [04_cache_aside](./02_examples/04_cache_aside/main.go) — паттерн
        Cache-Aside с TTL и подсчётом hit/miss ratio.
    *   [05_lru_eviction](./02_examples/05_lru_eviction/main.go) — реализация
        LRU-кэша за $O(1)$ на `map` + двусвязном списке.
    *   [06_stampede_singleflight](./02_examples/06_stampede_singleflight/main.go)
        — защита от cache stampede (thundering herd) при конкурентных промахах
        кэша (аналог `singleflight`).
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL`:
    *   `01_ttl_cache` — потокобезопасный кэш с истечением срока жизни (TTL).
    *   `02_lru_cache` — кэш фиксированной ёмкости с вытеснением давно не
        использовавшихся ключей (LRU).
    *   `03_hit_ratio` — расчёт эффективности кэша (`HitRatio`) по потоку
        обращений.
    *   `04_invalidation` — инвалидация кэша при записи и удалении
        (Cache-Aside).
    *   `05_red_metrics` — агрегация RED-метрик (`Requests`, `Errors`,
        `ErrorRate`,
        перцентиль латентности).
    *   `06_prometheus_format` — сериализация метрик в текстовый exposition
        формат Prometheus.
    *   `07_readiness_handler` — HTTP-хендлер готовности `/readyz`, проверяющий
        зависимости.
    *   `08_parse_log_line` — парсинг структурной `key=value` строки лога.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Ознакомьтесь с базовой теорией в DevOps-репозитории и прочитайте уроки в
   [01_theory/](./01_theory/).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_structured_logging/main.go
   go run ./02_examples/06_stampede_singleflight/main.go
   ```
3. Выполните упражнения из [03_exercises/](./03_exercises/):
   ```bash
   go run ./03_exercises/01_ttl_cache/main.go
   ```
4. Сверьтесь с эталонными решениями в
   [04_exercises_answers/](./04_exercises_answers/).

## Когда двигаться дальше

После освоения настройки рантайма Go в контейнерах, профилирования и
конкурентного кэширования переходите к распределённым компонентам:
[03_distributed_components](../03_distributed_components/).

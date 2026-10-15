# 08. Reliability & Scalability (Надёжность и масштабируемость)

Что нужно, чтобы система выдерживала растущую нагрузку и отказы: как распределять трафик, как масштабироваться, как защищаться от перегрузки и как корректно работать в условиях сбоев. Теория подкреплена runnable-примерами на стандартной библиотеке Go (детерминированными и завершающимися) и упражнениями с самопроверкой.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — алгоритмы балансировки, L4/L7, шардирование, rate limiting, circuit breaker, ретраи и таймауты, graceful shutdown, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_load_balancing.md](./01_theory/01_load_balancing.md) — балансировка нагрузки: алгоритмы (round-robin, least-connections, hash / consistent hashing), L4 vs L7, health checks и разделение liveness/readiness.
    *   [02_scaling.md](./01_theory/02_scaling.md) — горизонтальное vs вертикальное масштабирование, stateless-сервисы как ключ к scale-out, шардирование нагрузки и данных.
    *   [03_rate_limiting.md](./01_theory/03_rate_limiting.md) — ограничение частоты: token bucket, leaky bucket, fixed window и sliding window; локальный и распределённый лимит.
    *   [04_circuit_breaker.md](./01_theory/04_circuit_breaker.md) — предохранитель: состояния closed/open/half-open, зачем нужен, защита от каскадного отказа.
    *   [05_retries_timeouts.md](./01_theory/05_retries_timeouts.md) — таймауты, deadline propagation, повторы с exponential backoff + jitter, идемпотентность как условие безопасного retry.
    *   [06_graceful_shutdown.md](./01_theory/06_graceful_shutdown.md) — корректное завершение: SIGTERM + context, дренаж соединений, укладка в grace-период.
    *   [07_use_cases.md](./01_theory/07_use_cases.md) — как приёмы складываются в цельную стратегию устойчивого масштабируемого сервиса; типичные антипаттерны.
*   **[02_examples/](./02_examples/)** — runnable-примеры (детерминированные, без гонок по реальному времени, гарантированно завершаются):
    *   [01_token_bucket](./02_examples/01_token_bucket/main.go) — рейт-лимитер token bucket на логических тиках; демонстрация burst и отказов (429).
    *   [02_circuit_breaker](./02_examples/02_circuit_breaker/main.go) — предохранитель с переходами closed → open → half-open → closed; печать всех переходов.
    *   [03_retry_backoff](./02_examples/03_retry_backoff/main.go) — retry с exponential backoff + jitter вокруг нестабильной функции; jitter детерминирован (фиксированный seed).
    *   [04_graceful_shutdown](./02_examples/04_graceful_shutdown/main.go) — HTTP-сервер на `127.0.0.1:0`, запрос из горутины, затем `Server.Shutdown` с дренажом; без висящего сервера.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_token_bucket_allow` — реализовать `Allow()` для token bucket.
    *   `02_circuit_breaker_transitions` — реализовать переходы состояний в `OnResult`.
    *   `03_backoff_sequence` — вычислить последовательность задержек exponential backoff.
    *   `04_call_with_timeout` — обернуть вызов таймаутом через `context`.
    *   `05_fixed_window_limiter` — **fixed-window rate limiter на управляемом времени (Хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_load_balancing.md](./01_theory/01_load_balancing.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (всё офлайн, только stdlib; работает корневой `go.mod` репозитория):
   ```bash
   go run ./02_examples/01_token_bucket/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/01_token_bucket_allow/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

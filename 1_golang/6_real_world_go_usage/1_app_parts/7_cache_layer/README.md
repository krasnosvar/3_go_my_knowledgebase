# Cache layer

Кеш как декоратор репозитория: cache-aside, инвалидация на записи, negative caching и TTL с jitter.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#7-cache-layer) — раздел «7. Cache layer» общей шпаргалки частей приложения.

## Пререквизиты

[3_repository_layer](../3_repository_layer/); подробнее о кешировании — [3_system_design/03_caching](../../../../3_system_design/03_caching/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — Cache layer.
    *   [02_pitfalls.md](./01_theory/02_pitfalls.md) — Cache: pitfalls.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_in_memory_cache_shape](./02_examples/01_in_memory_cache_shape/main.go) — простейший in-memory кеш на map.
    *   [02_ttl_shape](./02_examples/02_ttl_shape/main.go) — проверка истечения TTL.
    *   [03_cache_aside_shape](./02_examples/03_cache_aside_shape/main.go) — последовательность шагов cache-aside.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_cached_repo_decorator](./03_exercises/01_cached_repo_decorator/main.go) — кеш как декоратор репозитория (cache-aside).
    *   [02_invalidate_on_write](./03_exercises/02_invalidate_on_write/main.go) — инвалидация кеша на записи.
    *   [03_negative_cache](./03_exercises/03_negative_cache/main.go) — negative caching.
    *   [04_ttl_jitter](./03_exercises/04_ttl_jitter/main.go) — TTL с jitter против одновременного протухания.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/1_app_parts/7_cache_layer` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_basics.md](./01_theory/01_basics.md), затем [02_pitfalls.md](./01_theory/02_pitfalls.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/7_cache_layer/02_examples/01_in_memory_cache_shape
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/7_cache_layer/03_exercises/01_cached_repo_decorator
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/1_app_parts/7_cache_layer/04_exercises_answers/01_cached_repo_decorator_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [8_observability](../8_observability/) — Observability.

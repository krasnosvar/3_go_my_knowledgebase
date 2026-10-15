# Упражнения: Cache layer

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_cached_repo_decorator
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_cached_repo_decorator_answers
```

## Список задач (от простого к сложному)

1. [`01_cached_repo_decorator`](./01_cached_repo_decorator/main.go) — кеш как декоратор репозитория (cache-aside); решение: [`01_cached_repo_decorator_answers`](../04_exercises_answers/01_cached_repo_decorator_answers/main.go).
2. [`02_invalidate_on_write`](./02_invalidate_on_write/main.go) — инвалидация кеша на записи; решение: [`02_invalidate_on_write_answers`](../04_exercises_answers/02_invalidate_on_write_answers/main.go).
3. [`03_negative_cache`](./03_negative_cache/main.go) — negative caching; решение: [`03_negative_cache_answers`](../04_exercises_answers/03_negative_cache_answers/main.go).
4. [`04_ttl_jitter`](./04_ttl_jitter/main.go) — TTL с jitter против одновременного протухания; решение: [`04_ttl_jitter_answers`](../04_exercises_answers/04_ttl_jitter_answers/main.go).

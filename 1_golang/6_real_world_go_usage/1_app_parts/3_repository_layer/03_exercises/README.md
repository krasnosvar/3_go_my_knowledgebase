# Упражнения: Repository layer

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_inmem_repo
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_inmem_repo_answers
```

## Список задач (от простого к сложному)

1. [`01_inmem_repo`](./01_inmem_repo/main.go) — in-memory репозиторий по контракту; решение: [`01_inmem_repo_answers`](../04_exercises_answers/01_inmem_repo_answers/main.go).
2. [`02_map_db_errors`](./02_map_db_errors/main.go) — маппинг ошибок драйвера в доменные; решение: [`02_map_db_errors_answers`](../04_exercises_answers/02_map_db_errors_answers/main.go).
3. [`03_offset_pagination`](./03_offset_pagination/main.go) — offset-пагинация в репозитории; решение: [`03_offset_pagination_answers`](../04_exercises_answers/03_offset_pagination_answers/main.go).
4. [`04_keyset_pagination`](./04_keyset_pagination/main.go) — keyset (cursor) пагинация; решение: [`04_keyset_pagination_answers`](../04_exercises_answers/04_keyset_pagination_answers/main.go).

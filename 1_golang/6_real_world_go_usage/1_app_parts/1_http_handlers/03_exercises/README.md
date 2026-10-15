# Упражнения: HTTP handlers

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_status_mapping
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_status_mapping_answers
```

## Список задач (от простого к сложному)

1. [`01_status_mapping`](./01_status_mapping/main.go) — маппинг доменных ошибок в HTTP-статусы; решение: [`01_status_mapping_answers`](../04_exercises_answers/01_status_mapping_answers/main.go).
2. [`02_query_params`](./02_query_params/main.go) — разбор query-параметров пагинации; решение: [`02_query_params_answers`](../04_exercises_answers/02_query_params_answers/main.go).
3. [`03_json_create`](./03_json_create/main.go) — handler создания ресурса с JSON-телом; решение: [`03_json_create_answers`](../04_exercises_answers/03_json_create_answers/main.go).
4. [`04_method_routing`](./04_method_routing/main.go) — маршрутизация по методу и path-параметру (Go 1.22+ ServeMux); решение: [`04_method_routing_answers`](../04_exercises_answers/04_method_routing_answers/main.go).

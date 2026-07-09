# Упражнения: api_styles

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`.

Запуск шаблона:

```bash
go run ./03_exercises/01_method_semantics/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_method_semantics/main.go
```

## Список задач (от простого к сложному)

1. `01_method_semantics` — по имени HTTP-метода вернуть флаги safe/idempotent.
2. `02_paginate` — собрать пагинированный ответ (окно `[offset, offset+limit)` + метаданные, включая `HasNext`).
3. `03_rest_status` — REST-хэндлер `/users/{id}` с корректными кодами 200/204/404/405 и заголовком `Allow`.
4. `04_parse_filter` — разобрать query-строку фильтра (`q`, `limit`, `active`) с дефолтами через `net/url` и `strconv`.

# 5.6 Integration tests

## Зачем

Интеграционный тест проверяет код вместе с реальными частями окружения: HTTP-стеком, файловой системой, БД. Здесь всё работает без Docker: `httptest`, `t.TempDir` и режим `-short` для медленных тестов. Подход с настоящей БД (testcontainers) описан в теории.

## Пререквизиты

- [05_mocks](../05_mocks/);
- `net/http` и `os` из [4_standard_library](../../4_standard_library/).

## Структура

- [01_theory/](./01_theory/):
  - [01_basics.md](./01_theory/01_basics.md) — что считать интеграционным тестом, `httptest`, testcontainers;
  - [02_setup_isolation.md](./01_theory/02_setup_isolation.md) — изоляция состояния между тестами, миграции, skip и build tags;
  - [03_flakiness.md](./01_theory/03_flakiness.md) — нестабильные тесты: причины и лечение;
- [02_examples/](./02_examples/) — примеры кода с тестами:
  - [01_httptest_server](./02_examples/01_httptest_server/httptest_server_test.go) — handler через `httptest.NewRecorder`;
  - [02_tempdir_roundtrip](./02_examples/02_tempdir_roundtrip/store_test.go) — сохранение и чтение JSON-файла в `t.TempDir`;
  - [03_short_mode](./02_examples/03_short_mode/fetch_test.go) — реальный HTTP-сервер `httptest.NewServer` и пропуск в `-short`;
- [03_exercises/](./03_exercises/) — 4 задачи с самопроверкой (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как запускать

Из папки темы (`1_golang/5_testing/06_integration_tests`):

```bash
go test -v ./02_examples/...                 # примеры
go test ./03_exercises/01_recorder_handler/          # заготовка: FAIL, пока не решена
go test ./04_exercises_answers/01_recorder_handler_answers/   # эталон: ok
```

## Дальше

раздел пройден — переходите к [6_real_world_go_usage](../../6_real_world_go_usage/).

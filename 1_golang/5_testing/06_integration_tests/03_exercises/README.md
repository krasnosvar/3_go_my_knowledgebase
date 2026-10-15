# Упражнения: Integration tests

Самопроверка в этом разделе — сам `go test`:

- **«напишите тест»** (большинство задач): рядом с кодом лежит `*_test.go`
  с заглушкой `t.Fatal("TODO ...")`, поэтому в стартовом состоянии
  `go test` печатает `FAIL`. Напишите проверки, удалите `t.Fatal` — тест
  должен стать `ok`. Чтобы убедиться, что тест действительно что-то ловит,
  временно сломайте код под тестом: тест обязан упасть.
- **«почините код»** (задачи 02–04 в `04_benchmarks`): тест уже написан и
  падает, править нужно только код под тестом.

Запуск (из папки `06_integration_tests`):

```bash
go test -v ./03_exercises/01_recorder_handler/
```

Эталонное решение:

```bash
go test -v ./04_exercises_answers/01_recorder_handler_answers/
```

## Список задач

1. `01_recorder_handler` — handler: код, заголовки и JSON-тело через `httptest.NewRecorder`.
2. `02_httptest_server` — HTTP-клиент против `httptest.NewServer`: 200, 404, 500.
3. `03_tempdir_files` — round-trip через настоящие файлы в `t.TempDir`.
4. `04_short_skip` — медленный тест с `testing.Short()` и запуск с `-short`.

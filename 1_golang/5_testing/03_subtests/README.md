# 5.3 Subtests

## Зачем

`t.Run` превращает каждый случай в именованный подтест: его можно запустить отдельно (`-run 'Test/имя'`), распараллелить (`t.Parallel`) и снабдить своим setup/cleanup.

## Пререквизиты

- [02_table_driven_tests](../02_table_driven_tests/);
- замыкания из [1_go_basics/07_functions](../../1_go_basics/07_functions/).

## Структура

- [01_theory/](./01_theory/):
  - [01_basics.md](./01_theory/01_basics.md) — `t.Run`, имена подтестов, запуск по `-run`;
  - [02_parallel_and_setup.md](./01_theory/02_parallel_and_setup.md) — `t.Parallel`, setup/teardown, `t.Cleanup`, `t.TempDir`;
- [02_examples/](./02_examples/) — примеры кода с тестами:
  - [01_status_classifier](./02_examples/01_status_classifier/status_classifier_test.go) — подтесты `t.Run` для классификатора кодов;
  - [02_parallel_subtests](./02_examples/02_parallel_subtests/square_test.go) — параллельные подтесты;
  - [03_cleanup_helper](./02_examples/03_cleanup_helper/counter_test.go) — хелпер setup + `t.Cleanup` + `t.TempDir`;
- [03_exercises/](./03_exercises/) — 4 задачи с самопроверкой (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как запускать

Из папки темы (`1_golang/5_testing/03_subtests`):

```bash
go test -v ./02_examples/...                 # примеры
go test ./03_exercises/01_t_run/          # заготовка: FAIL, пока не решена
go test ./04_exercises_answers/01_t_run_answers/   # эталон: ok
```

## Дальше

[04_benchmarks](../04_benchmarks/).

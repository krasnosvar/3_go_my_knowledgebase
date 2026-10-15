# 5.4 Benchmarks

## Зачем

Бенчмарки измеряют время и аллокации (`go test -bench`, `-benchmem`). `testing.AllocsPerRun` превращает требование «без лишних аллокаций» в обычный тест, который ломается при регрессии.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#04-бенчмарки) — раздел «04. Бенчмарки» общей шпаргалки раздела Testing.

## Пререквизиты

- [03_subtests](../03_subtests/) — `b.Run` устроен так же, как `t.Run`;
- `strings.Builder`, `strconv` из [1_go_basics/10_strings](../../1_go_basics/10_strings/).

## Структура

- [01_theory/](./01_theory/):
  - [01_basics.md](./01_theory/01_basics.md) — `BenchmarkXxx`, `b.N`, запуск и флаги;
  - [02_interpretation.md](./01_theory/02_interpretation.md) — как читать `ns/op`, `B/op`, `allocs/op`, сравнение через `benchstat`;
  - [03_alloc_optimization.md](./01_theory/03_alloc_optimization.md) — приёмы уменьшения аллокаций;
- [02_examples/](./02_examples/) — примеры кода с тестами:
  - [01_concat](./02_examples/01_concat/concat_test.go) — простейший бенчмарк;
  - [02_sub_benchmarks](./02_examples/02_sub_benchmarks/sum_test.go) — под-бенчмарки по размерам, `b.ResetTimer`, `b.ReportAllocs`;
  - [03_allocs_per_run](./02_examples/03_allocs_per_run/greet_test.go) — тест на число аллокаций через `testing.AllocsPerRun`;
- [03_exercises/](./03_exercises/) — 4 задачи с самопроверкой (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как запускать

Из папки темы (`1_golang/5_testing/04_benchmarks`):

```bash
go test -v ./02_examples/...                 # примеры
go test ./03_exercises/01_write_benchmark/          # заготовка: FAIL, пока не решена
go test ./04_exercises_answers/01_write_benchmark_answers/   # эталон: ok
```

## Дальше

[05_mocks](../05_mocks/).

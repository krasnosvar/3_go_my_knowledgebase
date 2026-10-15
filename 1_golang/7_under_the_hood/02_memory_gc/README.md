# 7.2 Memory and GC

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#02-память-и-gc) — раздел «02. Память и GC» общей шпаргалки раздела 7.

## Зачем

Аллокации в куче — главный источник нагрузки на GC и частая причина
просадок latency. Здесь — где живут переменные, как работает цикл GC,
что значит `GOGC`, и приёмы, которые убирают лишние аллокации.

## Пререквизиты

- [01_runtime](../01_runtime/) — особенно escape analysis;
- [5_testing/04_benchmarks](../../5_testing/04_benchmarks/) — `b.ReportAllocs`
  и `testing.AllocsPerRun`.

## Структура

- [01_theory/](./01_theory/):
  - [01_stack_vs_heap.md](./01_theory/01_stack_vs_heap.md) — стек и куча, escape analysis;
  - [02_gc_cycle.md](./01_theory/02_gc_cycle.md) — фазы GC, write barrier, `GOGC`, `GOMEMLIMIT`;
  - [03_optimization.md](./01_theory/03_optimization.md) — как уменьшать аллокации.
- [02_examples/](./02_examples/):
  - [01_escape_analysis_preview](./02_examples/01_escape_analysis_preview/main.go) — указатель на локальную переменную уходит в кучу;
  - [02_gc_stats_read](./02_examples/02_gc_stats_read/main.go) — чтение `MemStats`;
  - [03_object_pool_shape](./02_examples/03_object_pool_shape/main.go) — `sync.Pool` для переиспользования буферов.
- [03_exercises/](./03_exercises/) — 8 задач с самопроверкой `PASS`/`FAIL` (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как проходить

1. Прочитайте теорию по порядку.
2. Запустите примеры (из этой папки):
   ```bash
   go run ./02_examples/01_escape_analysis_preview
   go build -gcflags=-m ./02_examples/01_escape_analysis_preview
   ```
3. Решайте упражнения до `PASS`:
   ```bash
   go run ./03_exercises/01_stack_vs_heap
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_stack_vs_heap_answers
   ```

## Дальше

[04_profiling](../04_profiling/) — как находить такие проблемы в реальном
сервисе через `pprof`.

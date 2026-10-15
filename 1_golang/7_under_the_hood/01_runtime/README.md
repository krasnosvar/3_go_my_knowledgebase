# 7.1 Runtime Internals

## Зачем

Runtime встроен в каждый Go-бинарник: он планирует горутины, растит их
стеки, управляет памятью и GC. Понимание модели GMP и escape analysis
объясняет, почему горутины дешёвые, когда код «тормозит» без видимых
причин и откуда берутся аллокации.

## Пререквизиты

- [3_concurrency/01_goroutines](../../3_concurrency/01_goroutines/) и
  [3_concurrency/04_sync_package](../../3_concurrency/04_sync_package/);
- указатели и слайсы из [1_go_basics](../../1_go_basics/).

## Структура

- [01_theory/](./01_theory/):
  - [01_runtime_overview.md](./01_theory/01_runtime_overview.md) — что входит в runtime, диагностический API, `GODEBUG`;
  - [02_scheduler_gmp.md](./01_theory/02_scheduler_gmp.md) — планировщик: G, M, P, очереди, work stealing;
  - [03_stack_and_preemption.md](./01_theory/03_stack_and_preemption.md) — рост стека горутины и вытеснение;
  - [04_escape_analysis.md](./01_theory/04_escape_analysis.md) — когда переменная уходит в кучу;
  - [05_runtime_deep_dive.md](./01_theory/05_runtime_deep_dive.md) — подробный разбор runtime.
- [02_examples/](./02_examples/):
  - [01_scheduler_shape](./02_examples/01_scheduler_shape/main.go) — `GOMAXPROCS` и `NumCPU`;
  - [02_memstats_preview](./02_examples/02_memstats_preview/main.go) — `runtime.ReadMemStats` и число горутин;
  - [03_gomaxprocs_shape](./02_examples/03_gomaxprocs_shape/main.go) — изменение `GOMAXPROCS`.
- [03_exercises/](./03_exercises/) — 8 задач с самопроверкой `PASS`/`FAIL` (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как проходить

1. Прочитайте теорию по порядку.
2. Запустите примеры (из этой папки):
   ```bash
   go run ./02_examples/01_scheduler_shape
   ```
3. Решайте упражнения: заготовка компилируется и печатает `FAIL`, цель — `PASS`:
   ```bash
   go run ./03_exercises/01_goroutine_states
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./04_exercises_answers/01_goroutine_states_answers
   ```

## Дальше

[02_memory_gc](../02_memory_gc/) — куча, GC и оптимизация аллокаций.

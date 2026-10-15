# 7. Go Under The Hood

Как Go устроен внутри: runtime и планировщик, память и сборщик мусора,
установка и модули, профилирование, reflection и `unsafe`.

## Зачем

Без этих знаний можно писать на Go, но нельзя объяснить, почему сервис
тормозит, течёт по памяти или «залипает». Здесь — модель того, что делает
runtime с вашим кодом, и инструменты, чтобы это увидеть (`-gcflags=-m`,
`runtime/metrics`, `pprof`).

## Пререквизиты

- [1_go_basics](../1_go_basics/) — указатели, слайсы, map;
- [3_concurrency](../3_concurrency/) — горутины, каналы, `sync`;
- [5_testing/04_benchmarks](../5_testing/04_benchmarks/) — бенчмарки и
  `testing.AllocsPerRun` (используются в упражнениях по памяти).

## Порядок прохождения

1. [01_runtime](./01_runtime/) — runtime, планировщик GMP, стек, escape
   analysis. Теория, примеры, 8 упражнений с самопроверкой.
2. [02_memory_gc](./02_memory_gc/) — стек и куча, цикл GC, `GOGC`,
   оптимизация аллокаций. Теория, примеры, 8 упражнений с самопроверкой.
3. [03_compiler](./03_compiler/) — структура установки Go, модули,
   layout проекта. Только теория.
4. [04_profiling](./04_profiling/) — `pprof`, типы профилей, учебный
   «полигон» с четырьмя типичными проблемами.
5. [05_reflection_unsafe](./05_reflection_unsafe/) — `reflect` и `unsafe`:
   когда нужны и почему их избегают. Только теория.

## Как запускать

Все команды — из корня репозитория (там лежит `go.mod`):

```bash
go run ./1_golang/7_under_the_hood/01_runtime/02_examples/01_scheduler_shape
go run ./1_golang/7_under_the_hood/01_runtime/03_exercises/01_goroutine_states
```

Или из папки темы, например `cd 1_golang/7_under_the_hood/01_runtime`:

```bash
go run ./03_exercises/01_goroutine_states
```

## Дальше

[8_algorithms_dsa](../8_algorithms_dsa/) — алгоритмы и структуры данных;
либо вернитесь к [6_real_world_go_usage](../6_real_world_go_usage/) и
примените профилирование к реальному сервису.

# Go Knowledge Base

Раздел по Go: от базового синтаксиса до конкурентности, стандартной библиотеки,
testing, runtime, real-world backend-практик и алгоритмов.

## Формат тем

Большинство тем устроены по единому шаблону:

- `01_theory/` — краткая теория;
- `02_examples/` — runnable-примеры;
- `03_exercises/` — упражнения;
- `04_exercises_answers/` — решения для самопроверки (`NN_name_answers/`);
- `README.md` — навигация по теме.

Упражнения содержат `TODO` и встроенную самопроверку: в стартовом состоянии
они компилируются и печатают `FAIL`, после решения — `PASS`. Готовые варианты
лежат рядом в `04_exercises_answers/`. Запуск — из папки темы:
`go run ./03_exercises/01_xxx` (как установить Go — см. раздел
«0. Подготовка» в [корневом README](../README.md)).

## Разделы

1. [1_go_basics](./1_go_basics/) — переменные, типы, `fmt`, операторы,
   control flow, pointers, functions, structs, slices, strings, maps, errors.
2. [2_oop](./2_oop/) — methods, interfaces, embedding, type assertions,
   type switches, generics.
3. [3_concurrency](./3_concurrency/) — goroutines, channels, `select`,
   `sync`, `context`, `atomic`, concurrency patterns, race conditions.
4. [4_standard_library](./4_standard_library/) — `fmt`, `strconv`, `strings`,
   `regexp`, `sort`, `time`, `encoding/*`, `io`, `os`, `filepath`,
   `net/http`, `net`.
5. [5_testing](./5_testing/) — unit tests, table-driven tests, subtests,
   benchmarks, mocks, integration tests.
6. [6_real_world_go_usage](./6_real_world_go_usage/) — части backend-приложений
   и типы Go-приложений: CLI, backend API, system apps, gRPC, workers/cron,
   Kubernetes operators.
7. [7_under_the_hood](./7_under_the_hood/) — runtime, scheduler, memory,
   GC, compiler, modules, profiling, reflection, `unsafe`.
8. [8_algorithms_dsa](./8_algorithms_dsa/) — linked list, stack/queue, set,
   trees/heaps, graphs, leetcode-style patterns и архив дополнительных задач.

## Как проходить

1. Начать с `1_go_basics`, чтобы закрыть синтаксис и базовые типы.
2. Перейти к `2_oop` и `3_concurrency`: это основа Go-собеседований и
   практической backend-разработки.
3. Параллельно разбирать `4_standard_library` и `5_testing`, запуская примеры
   локально.
4. После базы идти в `6_real_world_go_usage` и `7_under_the_hood`.
5. `8_algorithms_dsa` использовать как отдельный блок для задач и интервью.

## Проверка

Весь раздел входит в корневой Go-модуль репозитория:

```bash
go vet ./1_golang/...
go test ./1_golang/5_testing/.../02_examples/... ./1_golang/5_testing/.../04_exercises_answers/...
```

Заготовки `5_testing/*/03_exercises` падают в `go test` намеренно — до решения.

## Критерии качества темы

Готовая тема должна:

- иметь понятный `README.md`;
- содержать короткую теорию без лишней воды;
- иметь запускаемые примеры;
- давать упражнения от простого к сложному;
- хранить решения отдельно от условий.

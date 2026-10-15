# 3. Конкурентность в Go (Concurrency)

Раздел о том, как писать конкурентный код на Go: горутины и каналы, `select`, пакеты `sync`, `context` и `sync/atomic`, типовые паттерны и разбор ошибок — гонок, deadlock'ов, livelock'ов и голодания.

**Зачем:** конкурентность — главная сильная сторона Go и обязательная тема backend-собеседований; без неё не понять, как работают HTTP-серверы, worker'ы и клиенты к БД.

**Пререквизиты:** [1_go_basics](../1_go_basics/) и [2_oop](../2_oop/) (методы, интерфейсы, замыкания).

## 🗂 Темы (проходить по порядку)

1. [01_goroutines](./01_goroutines/) — Горутины (goroutines).
2. [02_channels](./02_channels/) — Каналы (channels).
3. [03_select](./03_select/) — Оператор select.
4. [04_sync_package](./04_sync_package/) — Пакет sync.
5. [05_context](./05_context/) — Пакет context.
6. [06_atomic](./06_atomic/) — Пакет sync/atomic.
7. [07_patterns](./07_patterns/) — Паттерны конкурентности.
8. [08_race_conditions](./08_race_conditions/) — Гонки, deadlock, livelock, starvation.

## Формат темы

Каждая тема устроена одинаково:

- `01_theory/` — теория по подтемам;
- `02_examples/` — короткие runnable-примеры;
- `03_exercises/` — 8 задач с автопроверкой `PASS`/`FAIL`;
- `04_exercises_answers/` — эталонные решения (`NN_name_answers`).

## 🚀 Как запускать

Из папки темы, например `1_golang/3_concurrency/02_channels`:

```bash
go run ./02_examples/01_cli_utilities/main.go
go run -race ./03_exercises/01_unbuffered/main.go
go run -race ./04_exercises_answers/01_unbuffered_answers/main.go
```

Флаг `-race` включает детектор гонок: для конкурентного кода запускайте с ним всегда.

## ➡️ Что дальше

После раздела — [4_standard_library](../4_standard_library/) и [5_testing](../5_testing/); устройство планировщика и модели памяти — в [7_under_the_hood](../7_under_the_hood/).

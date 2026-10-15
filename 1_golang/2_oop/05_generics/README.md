# 05. Generics (обобщённое программирование)

Generics (Go 1.18+) позволяют писать функции и типы, работающие с разными типами без `any` и без копипасты: `Map`/`Filter`, контейнеры, репозитории. Тема разбирает type parameters, constraints (`any`, `comparable`, `cmp.Ordered`, свои union с `~`), обобщённые структуры и интерфейсы и — важно — когда generics НЕ нужны.

## Пререквизиты

[02_interfaces](../02_interfaces/) (constraint — это интерфейс) и [04_type_assertions_switches](../04_type_assertions_switches/); слайсы и map — [1_go_basics/09_slices](../../1_go_basics/09_slices/), [11_maps](../../1_go_basics/11_maps/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_type_parameters.md](./01_theory/01_type_parameters.md) — синтаксис `[T any]`, вывод типов, явная инстанциация.
    *   [02_constraints.md](./01_theory/02_constraints.md) — constraints: `any`, `comparable`, union, `~`, `cmp.Ordered`.
    *   [03_generic_structs.md](./01_theory/03_generic_structs.md) — обобщённые структуры и их методы.
    *   [04_generic_interfaces.md](./01_theory/04_generic_interfaces.md) — обобщённые интерфейсы.
    *   [05_pitfalls.md](./01_theory/05_pitfalls.md) — ограничения и ошибки: нет generic-методов, premature genericization.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — где generics реально полезны.
*   **[02_examples/](./02_examples/)** — runnable-примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — обобщённая `First`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — `Sum` с union-constraint `Number`.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — обобщённый `Repo[T]`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_generic_func](./03_exercises/01_generic_func/main.go) — `Map` и `Filter`.
    *   [02_constraints](./03_exercises/02_constraints/main.go) — `Sum` со своим `Number` и `Max` с `cmp.Ordered`.
    *   [03_generic_struct](./03_exercises/03_generic_struct/main.go) — `Optional[T]`: `Some`, `Get`, `OrElse`.
    *   [04_generic_interface](./03_exercises/04_generic_interface/main.go) — `memRepo[T]` под интерфейс `Repository[T]`.
    *   [05_comparable](./03_exercises/05_comparable/main.go) — `Index` и `Unique` с `comparable`.
    *   [06_type_set](./03_exercises/06_type_set/main.go) — type set `Integer` с `~`, `Abs`, `Clamp`.
    *   [07_any_alias](./03_exercises/07_any_alias/main.go) — `any` как алиас `interface{}`: `Join` и `Count`.
    *   [08_stack](./03_exercises/08_stack/main.go) — обобщённый стек `Stack[T]`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`NN_name_answers/`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/2_oop/05_generics`); корневой `go.mod` репозитория подхватится автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/) по порядку, начиная с [01_type_parameters.md](./01_theory/01_type_parameters.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи по порядку. Заготовка печатает `FAIL`, пока не решена:
   ```bash
   go run ./03_exercises/01_generic_func/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_generic_func_answers/main.go
   ```

## Что дальше

Блок ООП пройден. Дальше — [3_concurrency](../../3_concurrency/): горутины, каналы и синхронизация. Обзорная статья по всему блоку — [00_overview.md](../00_overview.md).

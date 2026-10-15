# 01. Методы (Methods)

Метод — функция, привязанная к типу через receiver. На методах держится всё «ООП» в Go: инкапсуляция поведения рядом с данными, реализация интерфейсов, конструкторы `NewXxx`. Главное в теме — понять разницу value/pointer receiver и method set: от неё зависит, меняется ли состояние и реализует ли тип интерфейс.

## Пререквизиты

[1_go_basics](../../1_go_basics/): особенно [06_pointers](../../1_go_basics/06_pointers/), [07_functions](../../1_go_basics/07_functions/) и [08_structs](../../1_go_basics/08_structs/). Для упражнения 08 — ошибки из [12_errors](../../1_go_basics/12_errors/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — синтаксис метода, receiver, методы у любых именованных типов.
    *   [02_value_pointer.md](./01_theory/02_value_pointer.md) — value vs pointer receiver: копирование, мутация, когда что выбирать.
    *   [03_method_set.md](./01_theory/03_method_set.md) — method set значения и указателя, адресуемость, связь с интерфейсами.
    *   [04_best_practices.md](./01_theory/04_best_practices.md) — соглашения: единообразные receiver'ы, имена, конструкторы.
    *   [05_pitfalls.md](./01_theory/05_pitfalls.md) — частые ошибки: потерянная мутация, nil receiver, копирование мьютекса.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — где методы применяются в реальном коде.
*   **[02_examples/](./02_examples/)** — runnable-примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — `String()` с value receiver для версии CLI.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — счётчик с pointer receiver.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — метод валидации сущности.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_value_receiver](./03_exercises/01_value_receiver/main.go) — `Area`/`Perimeter` у `Rect` с value receiver.
    *   [02_pointer_receiver](./03_exercises/02_pointer_receiver/main.go) — `Inc`/`Add`/`Reset` у счётчика с pointer receiver.
    *   [03_method_set](./03_exercises/03_method_set/main.go) — исправить receiver'ы так, чтобы `T{}` реализовывал `Namer`, а `Setter` — только `*T`.
    *   [04_constructor_usage](./03_exercises/04_constructor_usage/main.go) — конструктор `NewUser` с ролью по умолчанию и метод `Greeting`.
    *   [05_mutation](./03_exercises/05_mutation/main.go) — баг потерянной мутации в `Deposit` + `Withdraw` с проверкой баланса.
    *   [06_copy_cost](./03_exercises/06_copy_cost/main.go) — перевести методы «тяжёлой» структуры на pointer receiver.
    *   [07_fluent_style](./03_exercises/07_fluent_style/main.go) — построитель SQL-запроса с цепочкой методов.
    *   [08_validation_method](./03_exercises/08_validation_method/main.go) — `Validate() error` с порядком проверок.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`NN_name_answers/`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/2_oop/01_methods`); корневой `go.mod` репозитория подхватится автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/) по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи по порядку. Заготовка печатает `FAIL`, пока не решена:
   ```bash
   go run ./03_exercises/01_value_receiver/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_value_receiver_answers/main.go
   ```

## Что дальше

Следующая подтема — [02_interfaces](../02_interfaces/) (интерфейсы).

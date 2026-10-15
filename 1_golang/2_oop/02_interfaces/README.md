# 02. Интерфейсы (Interfaces)

Интерфейсы — центральный инструмент проектирования в Go: полиморфизм, слабая связность, подмена зависимостей в тестах. Реализация неявная (нет `implements`), поэтому важно понимать, как receiver влияет на реализацию, как устроен интерфейс внутри (тип + значение) и откуда берётся ловушка typed nil.

## Пререквизиты

[01_methods](../01_methods/) — method set и receiver'ы обязательны. Ошибки (`error`) — [1_go_basics/12_errors](../../1_go_basics/12_errors/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — объявление интерфейса, неявная реализация, пустой интерфейс.
    *   [02_pointer_vs_value.md](./01_theory/02_pointer_vs_value.md) — как value/pointer receiver влияет на реализацию интерфейса.
    *   [03_assertions_and_switch.md](./01_theory/03_assertions_and_switch.md) — type assertion и type switch (подробнее — в [04_type_assertions_switches](../04_type_assertions_switches/)).
    *   [04_nil_interface.md](./01_theory/04_nil_interface.md) — nil interface vs интерфейс с nil-указателем внутри.
    *   [05_design_principles.md](./01_theory/05_design_principles.md) — маленькие интерфейсы, «accept interfaces, return structs», интерфейс на стороне потребителя.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — реальные сценарии: репозитории, плагины, моки.
*   **[02_examples/](./02_examples/)** — runnable-примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — интерфейс `Command` для подкоманд CLI.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — `HealthChecker` для проверки узлов.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — `UserRepo` с in-memory реализацией.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_basic_interface](./03_exercises/01_basic_interface/main.go) — `Cat`/`Dog` реализуют `Speaker`, функция `chorus`.
    *   [02_pointer_receiver](./03_exercises/02_pointer_receiver/main.go) — `*counter` реализует `Counter`, а значение — нет.
    *   [03_type_assertions](./03_exercises/03_type_assertions/main.go) — безопасный `v.(int)` в форме comma-ok.
    *   [04_type_switch](./03_exercises/04_type_switch/main.go) — `describe` через type switch, включая `nil` и `error`.
    *   [05_nil_interface](./03_exercises/05_nil_interface/main.go) — исправить возврат typed nil как `error`.
    *   [06_small_interfaces](./03_exercises/06_small_interfaces/main.go) — `copyAll` поверх маленьких `Reader`/`Writer`.
    *   [07_dependency_injection](./03_exercises/07_dependency_injection/main.go) — `Service.Notify` через внедрённый `Sender`.
    *   [08_mocking](./03_exercises/08_mocking/main.go) — `greet` с подменяемыми часами `Clock`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`NN_name_answers/`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/2_oop/02_interfaces`); корневой `go.mod` репозитория подхватится автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/) по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи по порядку. Заготовка печатает `FAIL`, пока не решена:
   ```bash
   go run ./03_exercises/01_basic_interface/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_basic_interface_answers/main.go
   ```

## Что дальше

Следующая подтема — [03_embedding](../03_embedding/) (встраивание).

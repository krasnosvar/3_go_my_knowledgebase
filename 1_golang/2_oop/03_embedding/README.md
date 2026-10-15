# 03. Встраивание (Embedding)

В Go нет наследования — вместо него композиция через встраивание (embedding). Встроенный тип «продвигает» свои поля и методы во внешний, что позволяет переиспользовать поведение и собирать реализации интерфейсов из частей. Тема учит, как работает promotion, как «перекрыть» метод, что делать при конфликте имён и чем опасен встроенный nil-указатель.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — синтаксис, правила method promotion для `T` и `*T`, перекрытие и конфликты имён, применение, ловушки.

## Пререквизиты

[01_methods](../01_methods/) и [02_interfaces](../02_interfaces/); структуры — [1_go_basics/08_structs](../../1_go_basics/08_structs/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — синтаксис встраивания, имя встроенного поля.
    *   [02_method_promotion.md](./01_theory/02_method_promotion.md) — продвижение методов и полей, влияние на method set.
    *   [03_override.md](./01_theory/03_override.md) — «переопределение» метода и явный вызов встроенного.
    *   [04_name_conflicts.md](./01_theory/04_name_conflicts.md) — конфликты имён на одной глубине и правила разрешения.
    *   [05_pointer_embedding.md](./01_theory/05_pointer_embedding.md) — встраивание `*T`: разделяемое состояние и nil-паника.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — композиция в реальном коде: логгеры, метрики, общие поля.
*   **[02_examples/](./02_examples/)** — runnable-примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — общий `BaseCmd` для команд CLI.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — встроенные метрики в воркере.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — общие поля `Timestamps` у сущности.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_basic_embedding](./03_exercises/01_basic_embedding/main.go) — встроить `A` в `B`, чтобы появился `Hello`.
    *   [02_method_promotion](./03_exercises/02_method_promotion/main.go) — встроить `Logger` и собрать сервис конструктором.
    *   [03_override](./03_exercises/03_override/main.go) — свой `Print` у `Child`, вызывающий `Parent.Print`.
    *   [04_conflict](./03_exercises/04_conflict/main.go) — разрешить конфликт `Name()` двух встроенных типов.
    *   [05_pointer_embedding](./03_exercises/05_pointer_embedding/main.go) — инициализировать встроенный `*Counter` в конструкторе.
    *   [06_composition](./03_exercises/06_composition/main.go) — `API` из `Auth` и журнала аудита.
    *   [07_interface](./03_exercises/07_interface/main.go) — реализовать `Runner` через встроенный `Base`.
    *   [08_constructor](./03_exercises/08_constructor/main.go) — `NewChild` с заполнением встроенного `Base`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`NN_name_answers/`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/2_oop/03_embedding`); корневой `go.mod` репозитория подхватится автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/) по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи по порядку. Заготовка печатает `FAIL`, пока не решена:
   ```bash
   go run ./03_exercises/01_basic_embedding/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_basic_embedding_answers/main.go
   ```

## Что дальше

Следующая подтема — [04_type_assertions_switches](../04_type_assertions_switches/) (type assertions и type switches).

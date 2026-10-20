# 06. SOLID, DRY/KISS/YAGNI и паттерны проектирования в Go

Принципы и паттерны объектно-ориентированного дизайна в идиоматичном Go:
SOLID (SRP, OCP, LSP, ISP, DIP), DRY/KISS/YAGNI и классические паттерны GoF
(Functional Options, Strategy, Decorator/Middleware, Adapter, Factory). Акцент
— на том, как в Go без классов и наследования эти задачи решаются через узкие
интерфейсы, композицию и функции первого класса.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — SOLID в Go,
DRY/KISS/YAGNI, паттерны GoF в идиоматичном Go, запахи кода на review, ловушки.

## Пререквизиты

[01_methods](../01_methods/), [02_interfaces](../02_interfaces/) и
[03_embedding](../03_embedding/); функции как значения —
[1_go_basics/07_functions](../../1_go_basics/07_functions/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_solid.md](./01_theory/01_solid.md) — пять принципов SOLID (SRP,
        OCP, LSP, ISP, DIP) с примерами на Go; выражение через интерфейсы и
        композицию.
    *   [02_dry_kiss_yagni.md](./01_theory/02_dry_kiss_yagni.md) — DRY, KISS и
        YAGNI: борьба с лишней сложностью; дублирование знания против ложного
        DRY («a little copying is better than a little dependency»).
    *   [03_design_patterns.md](./01_theory/03_design_patterns.md) — паттерны
        GoF по группам (порождающие, структурные, поведенческие) и их
        упрощение в Go до интерфейса, func-типа или обёртки.
    *   [04_use_cases.md](./01_theory/04_use_cases.md) — практические
        применения: тестируемость через DIP, middleware как декораторы,
        strategy для сменных алгоритмов, functional options, запахи кода на
        code review.
*   **[02_examples/](./02_examples/)** — runnable-примеры:
    *   [01_solid_dip](./02_examples/01_solid_dip/main.go) — Dependency
        Inversion «до/после»: высокоуровневый модуль зависит от абстракции
        `Notifier`, обе реализации работают через один интерфейс.
    *   [02_strategy_pattern](./02_examples/02_strategy_pattern/main.go) —
        Strategy идиоматично через func-тип: подстановка разных стратегий
        скидки и сортировки (как `sort.Slice`).
    *   [03_decorator_pattern](./02_examples/03_decorator_pattern/main.go) —
        Decorator: обёртка вокруг интерфейса `Service` (логирование, счётчик)
        и middleware над `http.Handler`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL`
    (заготовки компилируются сразу):
    *   [01_srp_refactor](./03_exercises/01_srp_refactor/main.go) — SRP:
        разнести подсчёт суммы и её форматирование по разным
        ответственностям.
    *   [02_strategy_func](./03_exercises/02_strategy_func/main.go) — Strategy
        через func-тип: стратегии цены как значения-функции.
    *   [03_decorator](./03_exercises/03_decorator/main.go) — Decorator:
        добавить поведение обёрткой, не меняя исходный тип.
    *   [04_ocp_open_closed](./03_exercises/04_ocp_open_closed/main.go) —
        OCP через интерфейс: сумма площадей фигур без `switch` по типу.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения
    (`NN_name_answers/`).

## 🚀 Как проходить тему

Команды запускаются из папки темы
(`1_golang/2_oop/06_solid_and_design_patterns`); корневой `go.mod` репозитория
подхватится автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/) по порядку, начиная с
   [01_solid.md](./01_theory/01_solid.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_solid_dip/main.go
   ```
3. Решайте задачи по порядку. Заготовка печатает `FAIL`, пока не решена:
   ```bash
   go run ./03_exercises/01_srp_refactor/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_srp_refactor_answers/main.go
   ```

## Что дальше

Блок ООП пройден. Дальше — [3_concurrency](../../3_concurrency/): горутины,
каналы и синхронизация. Архитектура сервисов (Clean Architecture, DDD,
12-Factor App) — в
[3_system_design/01_architecture_styles](../../../3_system_design/01_architecture_styles/).
Обзорная статья по всему блоку ООП — [00_overview.md](../00_overview.md).

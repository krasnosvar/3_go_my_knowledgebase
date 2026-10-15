# 07. Design Principles (Принципы проектирования)

Принципы и паттерны, помогающие писать поддерживаемый, расширяемый код и разворачивать приложения предсказуемо: SOLID, DRY/KISS/YAGNI, паттерны проектирования и 12-factor app. Акцент — на **идиоматичном Go**: интерфейсы, композиция и функции-значения вместо слепого переноса ООП-паттернов и иерархий наследования. Теория подкреплена runnable-примерами на стандартной библиотеке и упражнениями с самопроверкой.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — SOLID в Go, DRY/KISS/YAGNI, паттерны в идиоматичном Go, 12-factor, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_solid.md](./01_theory/01_solid.md) — SOLID: пять принципов (SRP, OCP, LSP, ISP, DIP) с короткими примерами на Go; как они выражаются через интерфейсы и композицию, а не наследование.
    *   [02_dry_kiss_yagni.md](./01_theory/02_dry_kiss_yagni.md) — DRY / KISS / YAGNI: борьба с лишней сложностью; дублирование знания против ложного DRY («a little copying is better than a little dependency»).
    *   [03_design_patterns.md](./01_theory/03_design_patterns.md) — паттерны по группам (порождающие / структурные / поведенческие) кратко; как в Go они упрощаются до интерфейса, func-типа или обёртки.
    *   [04_twelve_factor.md](./01_theory/04_twelve_factor.md) — 12-factor app: двенадцать факторов cloud-native приложения и их отражение в Go (env-config, port binding, graceful shutdown, stateless, логи в stdout).
    *   [05_use_cases.md](./01_theory/05_use_cases.md) — практические применения: тестируемость через DIP, middleware как декораторы, strategy для сменных алгоритмов, functional options, 12-factor на деплое, запахи кода на review.
*   **[02_examples/](./02_examples/)** — runnable-примеры (детерминированные, только stdlib, завершаются):
    *   [01_solid_dip](./02_examples/01_solid_dip/main.go) — Dependency Inversion «до/после»: высокоуровневый модуль зависит от абстракции `Notifier`, обе реализации работают через один интерфейс.
    *   [02_strategy_pattern](./02_examples/02_strategy_pattern/main.go) — Strategy идиоматично через func-тип: подстановка разных стратегий скидки и сортировки (как `sort.Slice`).
    *   [03_decorator_pattern](./02_examples/03_decorator_pattern/main.go) — Decorator: обёртка вокруг интерфейса `Service` (логирование, счётчик) и middleware над `http.Handler`.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_srp_refactor` — SRP: разнести подсчёт суммы и её форматирование по разным ответственностям.
    *   `02_strategy_func` — Strategy через func-тип: стратегии цены как значения-функции.
    *   `03_decorator` — Decorator: добавить поведение обёрткой, не меняя исходный тип.
    *   `04_ocp_open_closed` — **OCP через интерфейс: сумма площадей фигур без `switch` по типу (Хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_solid.md](./01_theory/01_solid.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, только stdlib; работает корневой `go.mod` репозитория):
   ```bash
   go run ./02_examples/01_solid_dip/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/04_ocp_open_closed/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

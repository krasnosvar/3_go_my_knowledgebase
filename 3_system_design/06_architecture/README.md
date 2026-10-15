# 06. Architecture (Архитектура)

Как раскладывать код и сервисы по границам: стили архитектуры (монолит против микросервисов), слои и направление зависимостей (Clean Architecture), доменное моделирование (DDD), внедрение зависимостей (DI) и инфраструктурные компоненты, связывающие сервисы (API Gateway, service discovery, BFF). Теория подкреплена runnable-примерами на стандартной библиотеке Go и упражнениями с самопроверкой.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — монолит vs микросервисы, слои Clean Architecture, DDD, DI, gateway/discovery/BFF, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_monolith_vs_microservices.md](./01_theory/01_monolith_vs_microservices.md) — монолит против микросервисов: плюсы/минусы каждого, когда что оправдано, распределённый монолит как антипаттерн.
    *   [02_clean_architecture.md](./01_theory/02_clean_architecture.md) — Clean Architecture: слои (entity/usecase/adapters/frameworks), правило зависимостей, порты и адаптеры, инверсия зависимостей.
    *   [03_ddd_basics.md](./01_theory/03_ddd_basics.md) — основы DDD: ubiquitous language, bounded context, entity / value object / aggregate и aggregate root.
    *   [04_dependency_injection.md](./01_theory/04_dependency_injection.md) — DI в Go через интерфейсы и конструкторы, composition root, ручной DI против wire.
    *   [05_gateway_discovery.md](./01_theory/05_gateway_discovery.md) — API Gateway, service discovery (client-/server-side, реестр), BFF — кратко.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические применения: выбор стиля, нарезка сервисов по bounded contexts, структура кода внутри сервиса, типичные ошибки.
*   **[02_examples/](./02_examples/)** — runnable-примеры (детерминированные, только stdlib, завершаются):
    *   [01_dependency_injection](./02_examples/01_dependency_injection/main.go) — DI через интерфейс: usecase зависит от порта `Repository`, в main подставляется in-memory реализация; показана подмена зависимости.
    *   [02_clean_layers](./02_examples/02_clean_layers/main.go) — срез Clean Architecture: entity + порт `UserRepository` + usecase + in-memory адаптер; сценарий создать/получить пользователя.
    *   [03_ports_adapters](./02_examples/03_ports_adapters/main.go) — один порт `Notifier` и два адаптера (console и «email»-заглушка), выбор адаптера в main.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_implement_repository` — реализовать порт `Repository` (in-memory) и usecase поверх него.
    *   `02_constructor_injection` — внедрить зависимость `Clock` через конструктор (детерминированный `FixedClock`).
    *   `03_fix_dependency_direction` — исправить направление зависимостей: ввести порт `Logger` в слое usecase.
    *   `04_swap_adapter` — **подмена адаптера через порт `PaymentGateway` (хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_monolith_vs_microservices.md](./01_theory/01_monolith_vs_microservices.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, только stdlib; работает корневой `go.mod` репозитория):
   ```bash
   go run ./02_examples/01_dependency_injection/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/01_implement_repository/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

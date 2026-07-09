# Упражнения: architecture

Каждая задача — это `main.go` с функцией/типом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL`) — ваша цель довести их до `PASS`.

Запуск шаблона:

```bash
go run ./03_exercises/01_implement_repository/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_implement_repository/main.go
```

## Список задач (от простого к сложному)

1. `01_implement_repository` — реализовать порт `Repository` (in-memory) и usecase поверх него (`StoreAndLoad`).
2. `02_constructor_injection` — внедрить зависимость `Clock` через конструктор; результат детерминирован за счёт `FixedClock`.
3. `03_fix_dependency_direction` — исправить направление зависимостей: ввести порт `Logger` в слое usecase, чтобы адаптер зависел от порта, а не наоборот.
4. `04_swap_adapter` — **подмена адаптера через порт (хардкор)**: два адаптера `PaymentGateway` (успех/отказ), один usecase работает поверх обоих.

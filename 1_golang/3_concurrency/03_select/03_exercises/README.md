# Упражнения: select

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_basic_select/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_basic_select_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_basic_select](./01_basic_select/main.go) — первое значение из двух каналов. Решение: [01_basic_select_answers](../04_exercises_answers/01_basic_select_answers/main.go).
2. [02_default_non_blocking](./02_default_non_blocking/main.go) — `trySend`/`tryRecv` через `default`. Решение: [02_default_non_blocking_answers](../04_exercises_answers/02_default_non_blocking_answers/main.go).
3. [03_timeout](./03_timeout/main.go) — ожидание результата с ошибкой по таймауту. Решение: [03_timeout_answers](../04_exercises_answers/03_timeout_answers/main.go).
4. [04_two_channels](./04_two_channels/main.go) — вычитывание двух каналов до закрытия (nil-приём). Решение: [04_two_channels_answers](../04_exercises_answers/04_two_channels_answers/main.go).
5. [05_cancel_signal](./05_cancel_signal/main.go) — цикл обработки с выходом по `done`. Решение: [05_cancel_signal_answers](../04_exercises_answers/05_cancel_signal_answers/main.go).
6. [06_multiplex](./06_multiplex/main.go) — слияние N источников в один поток. Решение: [06_multiplex_answers](../04_exercises_answers/06_multiplex_answers/main.go).
7. [07_priority](./07_priority/main.go) — приоритет срочного канала двумя `select`. Решение: [07_priority_answers](../04_exercises_answers/07_priority_answers/main.go).
8. [08_retry_loop](./08_retry_loop/main.go) — повторы с таймаутом на каждую попытку. Решение: [08_retry_loop_answers](../04_exercises_answers/08_retry_loop_answers/main.go).

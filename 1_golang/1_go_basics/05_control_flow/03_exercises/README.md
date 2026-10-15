# Упражнения: Тема 05. Управляющие конструкции (Control Flow)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`05_control_flow/`):

```bash
go run ./03_exercises/01_if_else                 # заготовка
go run ./04_exercises_answers/01_if_else_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_if_else`](./01_if_else/main.go) | проверка возраста и статуса бана с использованием guard clauses (ранний возврат). | [`01_if_else_answers`](../04_exercises_answers/01_if_else_answers/main.go) |
| 2 | [`02_switch`](./02_switch/main.go) | маппинг кодов HTTP-статусов на их названия. | [`02_switch_answers`](../04_exercises_answers/02_switch_answers/main.go) |
| 3 | [`03_for_loop`](./03_for_loop/main.go) | классический цикл `for` для суммирования числового диапазона. | [`03_for_loop_answers`](../04_exercises_answers/03_for_loop_answers/main.go) |
| 4 | [`04_range`](./04_range/main.go) | итерирование по слайсу и мапе через `for range`. | [`04_range_answers`](../04_exercises_answers/04_range_answers/main.go) |
| 5 | [`05_break_continue`](./05_break_continue/main.go) | фильтрация четных чисел и выход при превышении лимита суммы. | [`05_break_continue_answers`](../04_exercises_answers/05_break_continue_answers/main.go) |
| 6 | [`06_labels`](./06_labels/main.go) | прерывание вложенного обхода матрицы при нахождении элемента. | [`06_labels_answers`](../04_exercises_answers/06_labels_answers/main.go) |
| 7 | [`07_infinite_loop`](./07_infinite_loop/main.go) | бесконечный цикл `for {}` с выходом по sentinel-значению `-1`. | [`07_infinite_loop_answers`](../04_exercises_answers/07_infinite_loop_answers/main.go) |
| 8 | [`08_table_dispatch`](./08_table_dispatch/main.go) | **написание Type Switch для разграничения типов интерфейса**. | [`08_table_dispatch_answers`](../04_exercises_answers/08_table_dispatch_answers/main.go) |

Решения открывайте только после собственной попытки.

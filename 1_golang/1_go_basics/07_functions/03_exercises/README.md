# Упражнения: Тема 07. Функции (Functions)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`07_functions/`):

```bash
go run ./03_exercises/01_basic_declaration                 # заготовка
go run ./04_exercises_answers/01_basic_declaration_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_basic_declaration`](./01_basic_declaration/main.go) | объявление функции форматирования строки приветствия. | [`01_basic_declaration_answers`](../04_exercises_answers/01_basic_declaration_answers/main.go) |
| 2 | [`02_multiple_return`](./02_multiple_return/main.go) | возвращение частного и остатка при делении с проверкой на ноль. | [`02_multiple_return_answers`](../04_exercises_answers/02_multiple_return_answers/main.go) |
| 3 | [`03_variadic`](./03_variadic/main.go) | суммирование переменного числа аргументов. | [`03_variadic_answers`](../04_exercises_answers/03_variadic_answers/main.go) |
| 4 | [`04_first_class`](./04_first_class/main.go) | применение функции-операции к каждому элементу слайса. | [`04_first_class_answers`](../04_exercises_answers/04_first_class_answers/main.go) |
| 5 | [`05_closure_counter`](./05_closure_counter/main.go) | создание счетчика с инкрементируемым шагом через замыкание. | [`05_closure_counter_answers`](../04_exercises_answers/05_closure_counter_answers/main.go) |
| 6 | [`06_defer_order`](./06_defer_order/main.go) | модификация именованного возвращаемого значения внутри `defer`. | [`06_defer_order_answers`](../04_exercises_answers/06_defer_order_answers/main.go) |
| 7 | [`07_error_wrapper`](./07_error_wrapper/main.go) | безопасный запуск функции с перехватом паник и конвертацией их в `error`. | [`07_error_wrapper_answers`](../04_exercises_answers/07_error_wrapper_answers/main.go) |
| 8 | [`08_generic_min`](./08_generic_min/main.go) | **написание обобщенной функции `Min` с ограничением `cmp.Ordered`**. | [`08_generic_min_answers`](../04_exercises_answers/08_generic_min_answers/main.go) |

Решения открывайте только после собственной попытки.

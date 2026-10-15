# Упражнения: Тема 12. Обработка ошибок (Errors)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`12_errors/`):

```bash
go run ./03_exercises/01_basic_error                 # заготовка
go run ./04_exercises_answers/01_basic_error_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_basic_error`](./01_basic_error/main.go) | создание и возврат базовой ошибки при некорректных аргументах. | [`01_basic_error_answers`](../04_exercises_answers/01_basic_error_answers/main.go) |
| 2 | [`02_wrap`](./02_wrap/main.go) | обертывание ошибки с добавлением контекста выполнения. | [`02_wrap_answers`](../04_exercises_answers/02_wrap_answers/main.go) |
| 3 | [`03_is`](./03_is/main.go) | рекурсивная проверка наличия sentinel error в цепочке ошибок. | [`03_is_answers`](../04_exercises_answers/03_is_answers/main.go) |
| 4 | [`04_as`](./04_as/main.go) | поиск и извлечение полей кастомной ошибки `ValidationError` из цепочки. | [`04_as_answers`](../04_exercises_answers/04_as_answers/main.go) |
| 5 | [`05_custom_type`](./05_custom_type/main.go) | описание кастомного типа ошибки, реализующего интерфейс `error` с pointer receiver. | [`05_custom_type_answers`](../04_exercises_answers/05_custom_type_answers/main.go) |
| 6 | [`06_defer_close`](./06_defer_close/main.go) | **идиоматическая обработка ошибок закрытия файлов/ресурсов в defer с обновлением именованного возврата**. | [`06_defer_close_answers`](../04_exercises_answers/06_defer_close_answers/main.go) |
| 7 | [`07_multi_step`](./07_multi_step/main.go) | многоэтапное обертывание ошибок по стеку вызовов. | [`07_multi_step_answers`](../04_exercises_answers/07_multi_step_answers/main.go) |
| 8 | [`08_retry`](./08_retry/main.go) | **обнаружение и исправление ловушки типизированного nil (typed nil error) в интерфейсах**. | [`08_retry_answers`](../04_exercises_answers/08_retry_answers/main.go) |

Решения открывайте только после собственной попытки.

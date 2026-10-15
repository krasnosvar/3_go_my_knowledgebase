# Упражнения: Тема 01. Переменные и константы (Variables & Constants)

Каждая задача — это `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии проверки печатают `FAIL`) — ваша цель довести их до `PASS`.

Команды запускаются из папки темы (`01_variables_constants/`):

```bash
go run ./03_exercises/01_var_decl                 # заготовка
go run ./04_exercises_answers/01_var_decl_answers # эталонное решение
```

## Список задач (от простого к сложному)

| # | Задача | Что тренирует | Решение |
|---|---|---|---|
| 1 | [`01_var_decl`](./01_var_decl/main.go) | базовое объявление переменных. | [`01_var_decl_answers`](../04_exercises_answers/01_var_decl_answers/main.go) |
| 2 | [`02_short_decl`](./02_short_decl/main.go) | краткая форма `:=`. | [`02_short_decl_answers`](../04_exercises_answers/02_short_decl_answers/main.go) |
| 3 | [`03_const`](./03_const/main.go) | объявление констант. | [`03_const_answers`](../04_exercises_answers/03_const_answers/main.go) |
| 4 | [`04_iota`](./04_iota/main.go) | генерация последовательности. | [`04_iota_answers`](../04_exercises_answers/04_iota_answers/main.go) |
| 5 | [`05_zero_values`](./05_zero_values/main.go) | проверка дефолтных нулевых значений. | [`05_zero_values_answers`](../04_exercises_answers/05_zero_values_answers/main.go) |
| 6 | [`06_scope`](./06_scope/main.go) | область видимости. | [`06_scope_answers`](../04_exercises_answers/06_scope_answers/main.go) |
| 7 | [`07_shadowing`](./07_shadowing/main.go) | **исправление ловушки затенения переменной (баг при инициализации базы данных)**. | [`07_shadowing_answers`](../04_exercises_answers/07_shadowing_answers/main.go) |
| 8 | [`08_group_decl`](./08_group_decl/main.go) | групповое объявление `var (...)`. | [`08_group_decl_answers`](../04_exercises_answers/08_group_decl_answers/main.go) |

Решения открывайте только после собственной попытки.

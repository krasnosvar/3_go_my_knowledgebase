# Изучение context в Go

`context.Context` нужен для управления временем жизни операций:

- отмена (`WithCancel`),
- таймаут (`WithTimeout`),
- дедлайн (`WithDeadline`),
- проброс служебных значений (`WithValue`).

## Структура

- `01_theory/` — база и практики.
- `02_examples/` — 3 runnable примера.
- `03_exercises/` — 8 задач.
- `04_exercises_answers/` — решения.

## Порядок изучения

1. Теория `01_theory/01_...` -> `01_theory/06_use_cases.md`.
2. Примеры в `02_examples`.
3. Задачи в `03_exercises`.
4. Проверка через `04_exercises_answers`.

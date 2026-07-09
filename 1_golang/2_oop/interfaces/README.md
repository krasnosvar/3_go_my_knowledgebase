# Изучение interfaces в Go

Интерфейсы — центральная концепция Go-проектирования.

В этой теме важно понять:

- неявную реализацию интерфейсов;
- влияние pointer/value receiver на реализацию интерфейса;
- type assertion и type switch;
- ловушку `nil interface`;
- принцип "accept interfaces, return structs".

## Структура

- `01_theory/` — теория по интерфейсам и практикам дизайна.
- `02_examples/` — 3 примера из CLI, инфраструктуры и backend.
- `03_exercises/` — 8 практических задач.
- `04_exercises_answers/` — ответы для самопроверки.
- `_legacy/` — прежние материалы до пересборки.

## Порядок изучения

1. Пройти `01_theory/01_basics.md` -> `01_theory/05_design_principles.md`.
2. Прочитать `01_theory/06_use_cases.md`.
3. Запустить `02_examples/*`.
4. Решить `03_exercises/*` по порядку.
5. Свериться с `04_exercises_answers/*`.

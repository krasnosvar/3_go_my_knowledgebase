# 2. ООП в Go: методы, интерфейсы, встраивание, generics, SOLID

## Зачем

В Go нет классов и наследования, но есть всё, что нужно для объектного дизайна:
методы у типов, неявные интерфейсы, композиция через встраивание, generics,
принципы SOLID и идиоматичные паттерны проектирования. Этот блок учит
пользоваться ими идиоматично — так, как это делают в production-коде и
спрашивают на собеседованиях.

## Пререквизиты

Пройден [1_go_basics](../1_go_basics/) — особенно указатели, функции,
структуры, слайсы, map и ошибки.

## Обзор

[00_overview.md](./00_overview.md) — обзорная статья «Все об ООП в Go»:
структуры, методы, конструкторы, инкапсуляция, встраивание, интерфейсы, DI,
generics и сравнение с классическим ООП. Удобно прочитать в начале для общей
картины и вернуться в конце как к шпаргалке.

## Порядок прохождения

1. [01_methods](./01_methods/) — value/pointer receiver, method set,
   конструкторы `NewXxx`.
2. [02_interfaces](./02_interfaces/) — неявная реализация, typed nil,
   маленькие интерфейсы, DI и моки.
3. [03_embedding](./03_embedding/) — композиция вместо наследования:
   promotion, «переопределение», конфликты имён.
4. [04_type_assertions_switches](./04_type_assertions_switches/) — `v.(T)` без
   паник, type switch, ветвление по интерфейсам.
5. [05_generics](./05_generics/) — type parameters, constraints (`comparable`,
   `cmp.Ordered`, `~`), обобщённые типы.
6. [06_solid_and_design_patterns](./06_solid_and_design_patterns/) — SOLID,
   DRY/KISS/YAGNI и паттерны проектирования (Functional Options, Strategy,
   Decorator/Middleware) в идиоматичном Go.

Каждая подтема устроена одинаково: `01_theory/` → `02_examples/` →
`03_exercises/` (самопроверка `PASS`/`FAIL`) → `04_exercises_answers/`.

## Как запускать

Из папки подтемы, например:

```bash
cd 1_golang/2_oop/01_methods
go run ./03_exercises/01_value_receiver/main.go
```

Проверить, что весь блок компилируется (из корня репозитория):

```bash
go vet ./1_golang/2_oop/...
```

## Что дальше

[3_concurrency](../3_concurrency/) — горутины, каналы, `select`, `sync`,
`context`. Интерфейсы и методы из этого блока там используются постоянно.

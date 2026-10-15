# 5.5 Mocks, fakes, spies

## Зачем

Чтобы тестировать бизнес-логику без БД, сети и часов, зависимости описывают маленькими интерфейсами и в тестах подставляют рукописные fakes/stubs/spies. Библиотеки моков не нужны.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#05-моки-и-fakes) — раздел «05. Моки и fakes» общей шпаргалки раздела Testing.

## Пререквизиты

- интерфейсы: [2_oop/02_interfaces](../../2_oop/02_interfaces/);
- [03_subtests](../03_subtests/).

## Структура

- [01_theory/](./01_theory/):
  - [01_principles.md](./01_theory/01_principles.md) — базовая идея, когда моки нужны и когда нет, где определять интерфейс;
  - [02_handmade_fakes.md](./01_theory/02_handmade_fakes.md) — рукописные fakes на Go;
  - [03_anti_patterns.md](./01_theory/03_anti_patterns.md) — анти-паттерны: мок всего подряд, проверка реализации вместо поведения;
- [02_examples/](./02_examples/) — примеры кода с тестами:
  - [01_service_repo](./02_examples/01_service_repo/service_repo_test.go) — fake-репозиторий для сервиса;
  - [02_fake_clock](./02_examples/02_fake_clock/token_test.go) — подмена времени через функцию `now`;
  - [03_spy](./02_examples/03_spy/audit_test.go) — spy записывает порядок вызовов;
- [03_exercises/](./03_exercises/) — 4 задачи с самопроверкой (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как запускать

Из папки темы (`1_golang/5_testing/05_mocks`):

```bash
go test -v ./02_examples/...                 # примеры
go test ./03_exercises/01_fake_repo/          # заготовка: FAIL, пока не решена
go test ./04_exercises_answers/01_fake_repo_answers/   # эталон: ok
```

## Дальше

[06_integration_tests](../06_integration_tests/).

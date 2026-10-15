# 5.1 Unit tests

## Зачем

Unit-тест проверяет одну функцию или тип в изоляции: быстро, детерминированно, без сети и диска. Это основа всего остального раздела: команда `go test`, файлы `*_test.go`, `t.Errorf` / `t.Fatalf`, покрытие.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#01-unit-тесты) — раздел «01. Unit-тесты» общей шпаргалки раздела Testing.

## Пререквизиты

- [1_go_basics](../../1_go_basics/) — функции, структуры, ошибки (`errors.Is`);
- [2_oop/01_methods](../../2_oop/01_methods/) — методы (упражнение 03).

## Структура

- [01_theory/](./01_theory/):
  - [01_basics.md](./01_theory/01_basics.md) — файлы `_test.go`, сигнатура `TestXxx`, `go test`, `t.Error` vs `t.Fatal`;
  - [02_arrange_act_assert.md](./01_theory/02_arrange_act_assert.md) — структура теста Arrange–Act–Assert;
  - [03_pitfalls_and_coverage.md](./01_theory/03_pitfalls_and_coverage.md) — частые ошибки и покрытие (`-cover`, `-coverprofile`);
- [02_examples/](./02_examples/) — примеры кода с тестами:
  - [01_math](./02_examples/01_math/math_test.go) — минимальный тест функции `Add`;
  - [02_validation](./02_examples/02_validation/validation_test.go) — тест булевой функции `IsAdult`;
  - [03_error_case](./02_examples/03_error_case/parse_test.go) — проверка ветки с ошибкой;
- [03_exercises/](./03_exercises/) — 7 задач с самопроверкой (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как запускать

Из папки темы (`1_golang/5_testing/01_unit_tests`):

```bash
go test -v ./02_examples/...                 # примеры
go test ./03_exercises/01_sum/          # заготовка: FAIL, пока не решена
go test ./04_exercises_answers/01_sum_answers/   # эталон: ok
```

## Дальше

[02_table_driven_tests](../02_table_driven_tests/).

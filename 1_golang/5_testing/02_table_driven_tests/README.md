# 5.2 Table-driven tests

## Зачем

Табличные тесты — идиома Go: случаи описываются данными (слайс или map структур), а проверка пишется один раз. Новый случай — одна строка в таблице.

## Пререквизиты

- [01_unit_tests](../01_unit_tests/) — `go test`, `t.Errorf`;
- анонимные структуры и слайсы из [1_go_basics](../../1_go_basics/).

## Структура

- [01_theory/](./01_theory/):
  - [01_pattern.md](./01_theory/01_pattern.md) — паттерн: таблица, цикл, сообщения об ошибках;
  - [02_pitfalls.md](./01_theory/02_pitfalls.md) — типичные ошибки: захват переменной цикла, `t.Fatal` в цикле, сравнение слайсов;
- [02_examples/](./02_examples/) — примеры кода с тестами:
  - [01_split_words](./02_examples/01_split_words/split_words_test.go) — таблица со слайсом в ожидании (`reflect.DeepEqual`);
  - [02_want_err](./02_examples/02_want_err/port_test.go) — таблица с полем `wantErr`;
  - [03_map_cases](./02_examples/03_map_cases/grade_test.go) — таблица в виде map + `t.Run`;
- [03_exercises/](./03_exercises/) — 4 задачи с самопроверкой (список — в [README упражнений](./03_exercises/README.md)).
- [04_exercises_answers/](./04_exercises_answers/) — эталонные решения (`NN_name_answers`).

## Как запускать

Из папки темы (`1_golang/5_testing/02_table_driven_tests`):

```bash
go test -v ./02_examples/...                 # примеры
go test ./03_exercises/01_basic_table/          # заготовка: FAIL, пока не решена
go test ./04_exercises_answers/01_basic_table_answers/   # эталон: ok
```

## Дальше

[03_subtests](../03_subtests/).

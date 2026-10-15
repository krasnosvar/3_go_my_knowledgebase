# 04. Type assertions и type switches

Когда значение лежит в интерфейсе (`any`, `error`, `io.Reader`), иногда нужно узнать конкретный тип или проверить, реализует ли оно другой интерфейс. Для этого есть type assertion `v.(T)` и type switch. Тема учит делать это без паник и не злоупотреблять `any` там, где хватит обычного интерфейса.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — формы `v.(T)`, type switch и тип переменной в ветках, `any` vs интерфейс vs generics, ловушки.

## Пререквизиты

[02_interfaces](../02_interfaces/) (устройство интерфейса) и [03_embedding](../03_embedding/). Для упражнений 02 и 06 — `fmt.Errorf` и `strconv` из [1_go_basics](../../1_go_basics/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_assertions.md](./01_theory/01_assertions.md) — `v.(T)` в одно- и двухзначной форме, паника при несовпадении.
    *   [02_type_switch.md](./01_theory/02_type_switch.md) — `switch x := v.(type)`, несколько типов в case, `nil`.
    *   [03_pitfalls.md](./01_theory/03_pitfalls.md) — частые ошибки: паники, порядок case, typed nil.
    *   [04_use_any.md](./01_theory/04_use_any.md) — `any` / `interface{}`: когда уместен и когда нет.
    *   [05_design.md](./01_theory/05_design.md) — как проектировать без лишних assertion'ов.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — реальные сценарии: JSON, ошибки, опциональные интерфейсы.
*   **[02_examples/](./02_examples/)** — runnable-примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — assertion аргумента CLI к строке.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — type switch по коду статуса в `any`.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — извлечение `id` из `map[string]any`-payload через assertion.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_assertion_ok](./03_exercises/01_assertion_ok/main.go) — `asString` через comma-ok.
    *   [02_assertion_fail](./03_exercises/02_assertion_fail/main.go) — `mustString` без паники, с ошибкой и `%T`.
    *   [03_type_switch_basic](./03_exercises/03_type_switch_basic/main.go) — `kind` по базовым типам.
    *   [04_type_switch_struct](./03_exercises/04_type_switch_struct/main.go) — площадь фигуры по типу структуры, включая `*Rect` и nil.
    *   [05_any_parser](./03_exercises/05_any_parser/main.go) — `format` для int/string/[]int/map/nil.
    *   [06_safe_cast](./03_exercises/06_safe_cast/main.go) — `toIntOr` — приведение «чего угодно» к int.
    *   [07_interface_branching](./03_exercises/07_interface_branching/main.go) — ветвление по интерфейсам `error` и `fmt.Stringer`.
    *   [08_json_like](./03_exercises/08_json_like/main.go) — `getPath`/`getInt` по JSON-подобной `map[string]any`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`NN_name_answers/`).

## 🚀 Как проходить тему

Команды запускаются из папки темы (`1_golang/2_oop/04_type_assertions_switches`); корневой `go.mod` репозитория подхватится автоматически.

1. Прочитайте теорию в [01_theory/](./01_theory/) по порядку, начиная с [01_assertions.md](./01_theory/01_assertions.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи по порядку. Заготовка печатает `FAIL`, пока не решена:
   ```bash
   go run ./03_exercises/01_assertion_ok/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_assertion_ok_answers/main.go
   ```

## Что дальше

Следующая подтема — [05_generics](../05_generics/) (generics).

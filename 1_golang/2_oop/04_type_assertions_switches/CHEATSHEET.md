# Type assertions и type switches — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Type assertion

Подробно: [01_assertions.md](./01_theory/01_assertions.md).

| Задача | Код | Если тип не тот |
| :--- | :--- | :--- |
| Извлечь конкретный тип | `s := i.(string)` | **panic** |
| Безопасная форма | `s, ok := i.(string)` | `ok == false`, `s` — нулевое значение |
| Проверить интерфейс | `rw, ok := w.(io.ReadWriter)` | `ok == false` |
| Достать из `map[string]any` | `if v, ok := m["port"].(int); ok { ... }` | ветка не выполнится |
| Ошибка определённого типа (по цепочке) | `var ne *net.OpError; if errors.As(err, &ne) { ne.Op }` | `false` |

| Assertion `y.(int)` | Conversion `int(y)` |
| :--- | :--- |
| только из интерфейсного значения | между конкретными совместимыми типами |
| не меняет значение, только «распаковывает» | преобразует значение |
| `x.(int)` для `x int` — ошибка компиляции | `i.(string)` при `i = 42` — всегда промах |

## Type switch

Подробно: [02_type_switch.md](./01_theory/02_type_switch.md).

| Конструкция | Тип `x` в ветке |
| :--- | :--- |
| `case int:` | `int` |
| `case int, int64:` (несколько типов) | тип исходного интерфейса (`any`) |
| `case fmt.Stringer:` | `fmt.Stringer` — любой реализующий тип |
| `case nil:` | nil-интерфейс (нет ни типа, ни значения) |
| `default:` | тип исходного интерфейса |

Ветки проверяются сверху вниз — узкие типы выше общих интерфейсов. `fallthrough` запрещён.

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

type temp float64

func (t temp) String() string { return fmt.Sprintf("%.1f°C", float64(t)) }

func describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int, int64:
		return fmt.Sprintf("целое %v (%T)", x, x) // x — any
	case string:
		return "строка " + x // x — string
	case fmt.Stringer:
		return "stringer " + x.String()
	default:
		return fmt.Sprintf("unknown %T", x)
	}
}

func main() {
	for _, v := range []any{nil, 42, int64(7), "go", temp(21.5), 3.14} {
		fmt.Println(describe(v))
	}
	// nil / целое 42 (int) / целое 7 (int64) / строка go / stringer 21.5°C / unknown float64

	var r io.Reader = strings.NewReader("abc")
	if wt, ok := r.(io.WriterTo); ok { // опциональный интерфейс
		n, _ := wt.WriteTo(io.Discard)
		fmt.Println("WriterTo:", n) // WriterTo: 3
	}
	var a any = "x"
	n, ok := a.(int)
	fmt.Println(n, ok) // 0 false — без ok была бы паника
}
```

## any и когда что выбрать

Подробно: [04_use_any.md](./01_theory/04_use_any.md), [05_design.md](./01_theory/05_design.md), [06_use_cases.md](./01_theory/06_use_cases.md).

| Задача | Инструмент |
| :--- | :--- |
| JSON с неизвестной схемой | `json.Unmarshal(data, &raw)` → switch: `map[string]any`, `[]any`, `float64`, `string`, `bool`, `nil` |
| «Если реализация умеет — используем» | assertion к опциональному интерфейсу (`Flusher`, `io.WriterTo`, `io.Closer`) |
| Классификация ошибок | `errors.Is` / `errors.As`, а не assertion |
| Ошибка → HTTP-статус | обычный `switch { case errors.Is(err, ErrNotFound): return 404 }` |
| Событие → обработчик (мало типов) | type switch по `*UserCreated`, `*OrderPlaced` + `default` с логом `%T` |
| Одинаковый метод во всех ветках | интерфейс с этим методом вместо switch |
| `func F(xs []any) any` + `.(int)` внутри | generic `func F[T Number](xs []T) T` |
| Внешние данные | assert/switch один раз на границе (handler/decoder), дальше — типизированные значения |

`any` — alias `interface{}` (Go 1.18+), полностью идентичен. Упаковка значения в интерфейс может стоить аллокации.

## Ловушки

Подробно: [03_pitfalls.md](./01_theory/03_pitfalls.md).

- ⚠️ `i.(T)` без `ok` на пользовательском вводе → runtime panic.
- ⚠️ В интерфейсе лежит `T`, а ассертят `*T` (или наоборот) → `ok == false`.
- ⚠️ `v.(I)` при `I` с методом на `*S`, а в `v` лежит `S{}` → `false`: value не реализует `I`.
- ⚠️ Type switch без `default` на внешних данных → неожиданный тип (например, `float64` из JSON) молча игнорируется.
- ⚠️ `case int, int64:` → `x` остаётся `any`, арифметика не компилируется.
- ⚠️ Typed nil: `var p *MyType; var v any = p` → `v != nil`, `v.(*MyType)` даёт `ok == true` и nil-указатель — проверяйте nil после assertion.
- ⚠️ Assertion/switch в hot path на миллионах значений заметно дороже вызова метода интерфейса.
- ⚠️ `if _, ok := v.(*mytype)` внутри своего же пакета — признак, что не хватает интерфейса.
- ⚠️ Функция принимает `any` и молча поддерживает только 3 типа — документируйте набор.

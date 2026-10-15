# Управление потоком — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Семантика циклов зависит от строки `go` в `go.mod` (в этом репозитории — `go 1.22`).

## if

Подробно: [01_if_else.md](./01_theory/01_if_else.md).

| Задача | Код |
| :--- | :--- |
| Ветвление | `if x > 0 { ... } else if x == 0 { ... } else { ... }` |
| С инициализацией (scope — только `if/else`) | `if v, err := parse(s); err == nil { use(v) } else { log.Println(err) }` |
| Проверка ключа в map | `if v, ok := m[k]; ok { ... }` |
| Тернарный оператор | нет — полноценный `if` |
| Ранний выход вместо вложенности | `if !ok { return }` |

## switch

Подробно: [02_switch.md](./01_theory/02_switch.md).

| Форма | Код |
| :--- | :--- |
| По значению, несколько значений в case | `switch day { case "sat", "sun": ...; default: ... }` |
| Без выражения (= цепочка `if`) | `switch { case age < 18: ...; case age < 65: ... }` |
| С инициализацией | `switch x := compute(); { case x < 0: ... }` |
| Провалиться в следующий case | `fallthrough` (последней строкой case, без проверки условия) |
| По типу | `switch v := x.(type) { case int: ...; case string: ... }` |
| Диспетчер команд | `switch cmd { case "build": return runBuild(); default: return fmt.Errorf("unknown command: %s", cmd) }` |

Каждый `case` завершается сам — `break` писать не нужно.

## for

Подробно: [03_for_and_range.md](./01_theory/03_for_and_range.md).

| Форма | Код |
| :--- | :--- |
| Классический | `for i := 0; i < 10; i++ { }` |
| Два указателя | `for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 { }` |
| Как `while` | `for n < 100 { n *= 2 }` |
| Бесконечный | `for { ... break ... }` |
| По целому (1.22+) | `for i := range 10 { }` → 0..9 |
| По функции-итератору (1.23+) | `for k := range maps.Keys(m) { }`, свой: `func(yield func(T) bool)` |
| Event loop | `for { select { case <-ctx.Done(): return; case v := <-ch: handle(v) } }` |

| `range` по | 1-е значение | 2-е значение |
| :--- | :--- | :--- |
| Слайс / массив | индекс | копия элемента |
| Строка | индекс **байта** | `rune` |
| Map | ключ (порядок случайный) | значение |
| Канал | значение (до `close`) | — |
| Целое `n` (1.22+) | 0..n−1 | — |

| Нужно | Код |
| :--- | :--- |
| Только значения | `for _, v := range s` |
| Только индексы | `for i := range s` |
| Изменить элементы на месте | `for i := range items { items[i].Done = true }` |
| Не копировать большие элементы массива | `for i := range arr { use(arr[i]) }` или `range &arr` |

## break, continue, метки

Подробно: [04_break_continue_labels.md](./01_theory/04_break_continue_labels.md).

| Оператор | Действует на |
| :--- | :--- |
| `break` | ближайший `for`, `switch` или `select` |
| `continue` | следующую итерацию ближайшего `for` |
| `break Label` | цикл/`switch`/`select` с этой меткой |
| `continue Label` | следующую итерацию внешнего цикла с меткой |
| `goto Label` | есть, но почти не используется |

## Переменная цикла (1.22)

Подробно: [03_for_and_range.md](./01_theory/03_for_and_range.md#переменные-цикла-и-горутины).

| Версия `go` в go.mod | `for _, v := range ...` с `&v` или `go func(){ use(v) }()` |
| :--- | :--- |
| < 1.22 | одна `v` на весь цикл → все указатели/горутины видят последнее значение; лечили `v := v` |
| ≥ 1.22 | новая `v` на каждой итерации — `v := v` больше не нужен |

```go
package main

import (
	"fmt"
	"maps"
	"slices"
)

func classify(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprint("int ", v)
	case string:
		return "string " + v
	default:
		return fmt.Sprintf("other %T", v)
	}
}

func main() {
	for i := range 3 { // Go 1.22+: range по целому
		fmt.Print(i, " ")
	}
	fmt.Println() // 0 1 2

	var ptrs []*int
	for _, v := range []int{1, 2, 3} {
		ptrs = append(ptrs, &v) // Go 1.22+: новая v на каждой итерации
	}
	fmt.Println(*ptrs[0], *ptrs[1], *ptrs[2]) // 1 2 3 (до 1.22: 3 3 3)

	for i, r := range "Привет" {
		if i > 2 {
			break
		}
		fmt.Printf("%d:%c ", i, r)
	}
	fmt.Println() // 0:П 2:р

Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j == 3 {
				fmt.Println("break Outer at", i, j) // break Outer at 1 2
				break Outer
			}
		}
	}

	switch n := 1; n {
	case 1:
		fmt.Print("one ")
		fallthrough
	case 2:
		fmt.Println("<= 2") // one <= 2
	}

	m := map[string]int{"b": 2, "a": 1}
	for _, k := range slices.Sorted(maps.Keys(m)) { // Go 1.23+: итераторы
		fmt.Print(k, "=", m[k], " ")
	}
	fmt.Println() // a=1 b=2

	fmt.Println(classify(42), "|", classify("go"), "|", classify(1.5))
	// int 42 | string go | other float64
}
```

## Ловушки

Подробно: [05_common_pitfalls.md](./01_theory/05_common_pitfalls.md), [06_use_cases.md](./01_theory/06_use_cases.md).

- ⚠️ `break` внутри `select` или `switch` в цикле выходит только из них — для выхода из `for` нужна метка.
- ⚠️ `for i := range nums` даёт индекс, а не значение.
- ⚠️ `v.Done = true` в `for _, v := range items` меняет копию — писать через `items[i]`.
- ⚠️ `for { }` без `return`/`break`/`ctx.Done()` — вечный цикл.
- ⚠️ Индексы при `range` по строке — байтовые: у кириллицы шаг 2.
- ⚠️ Порядок `range` по map случайный — для стабильного вывода сортировать ключи.
- ⚠️ Условие обязано быть `bool`: `if x` для `int` и `if x = 1` не компилируются.
- ⚠️ `defer` в теле цикла срабатывает при выходе из функции, а не из итерации — закрывать ресурс в отдельной функции.
- ⚠️ Пустой `default: // ничего` — шум; длинная лестница `if/else` по одной переменной — повод для `switch`.

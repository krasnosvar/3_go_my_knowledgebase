# Переменные и константы — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Объявление

Подробно: [01_vars_and_zero_values.md](./01_theory/01_vars_and_zero_values.md), [03_var_short_decl_scope.md](./01_theory/03_var_short_decl_scope.md).

| Задача | Код | Где работает |
| :--- | :--- | :--- |
| Нулевое значение типа | `var a int` | везде |
| Тип + значение | `var b int = 10` | везде |
| Тип выводится | `var c = 10` (→ `int`) | везде |
| Несколько сразу | `var x, y int = 1, 2` | везде |
| Блок связанных переменных | `var ( mu sync.Mutex; cache = make(map[string]int) )` | везде |
| Короткое объявление | `name := "den"`, `name, age := "den", 30` | только в функциях |
| Переприсвоить часть | `a, err := f()` → `b, err := g()` (новая `b`, та же `err`) | только в функциях |
| Отбросить значение | `_, err := f()` | везде |
| Подмена при сборке | `var Version = "dev"` + `go build -ldflags "-X main.Version=1.2.3"` | уровень пакета |

| Контекст | Предпочтительно |
| :--- | :--- |
| Уровень пакета | `var` (`:=` там запрещён) |
| Локально, тип очевиден из значения | `:=` |
| Нужен явный тип / zero-init (`var x *T`, интерфейс) | `var` |

## Нулевые значения

Подробно: [01_vars_and_zero_values.md](./01_theory/01_vars_and_zero_values.md).

| Тип | Zero value | Что можно с zero value |
| :--- | :--- | :--- |
| `int`, `float64`, … | `0` | всё |
| `bool` | `false` | всё |
| `string` | `""` | всё |
| `*T`, `func`, `interface` | `nil` | сравнить с `nil`; вызов/разыменование → panic |
| `[]T` | `nil` | `len`, `range`, `append` |
| `map[K]V` | `nil` | чтение → zero value; ⚠️ запись → panic, сначала `make` |
| `chan T` | `nil` | ⚠️ отправка/приём блокируют навсегда |
| `struct` | все поля — свои zero values | всё |

## Область видимости

Подробно: [03_var_short_decl_scope.md](./01_theory/03_var_short_decl_scope.md).

| Где объявлено | Видно |
| :--- | :--- |
| Уровень пакета, имя с заглавной (`Name`) | во всём пакете + экспортируется |
| Уровень пакета, имя со строчной (`name`) | во всём пакете |
| Блок `{ ... }` | до конца блока |
| `if v := f(); cond { }` / `for` / `switch` с инициализацией | только внутри конструкции (включая `else`) |
| Неиспользованная локальная переменная | ошибка компиляции (`declared and not used`) |

## Константы и iota

Подробно: [04_constants_iota.md](./01_theory/04_constants_iota.md).

| Задача | Код |
| :--- | :--- |
| Untyped константа (подстраивается под `int`, `int64`, …) | `const N = 10` |
| Типизированная | `const N int32 = 10` |
| Группа | `const ( MaxRetries = 3; DefaultPort = 8080 )` |
| Последовательность с 0 | `const ( A = iota; B; C )` → 0, 1, 2 |
| Enum на именованном типе | `type Status int` + `const ( StatusPending Status = iota; StatusActive )` |
| Пропустить 0 | `const ( _ = iota; KB = 1 << (10 * iota); MB; GB )` |
| Битовые флаги | `const ( PermRead Perm = 1 << iota; PermWrite; PermExec )` → 1, 2, 4 |
| Читаемый вывод enum | метод `func (s Status) String() string` |

| Правило iota | |
| :--- | :--- |
| Старт | 0 в каждом новом блоке `const ( ... )` |
| Инкремент | на каждую строку блока, включая `_` |
| Пропущенное выражение | повторяется выражение предыдущей строки с новым `iota` |
| Значение константы | только constant expression (литералы, арифметика, другие константы) |

```go
package main

import "fmt"

type Perm uint8

const (
	PermRead  Perm = 1 << iota // 1
	PermWrite                  // 2
	PermExec                   // 4
)

const (
	_  = iota             // 0 пропускаем
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
)

func main() {
	var n int    // 0
	var s string // ""
	x, err := 1, error(nil)
	y, err := 2, fmt.Errorf("boom") // y новая, err переприсвоена
	if x := 10; x > y {             // x затеняет внешнюю
		fmt.Println("inner x:", x)
	}
	p := PermRead | PermExec
	fmt.Println(n, s == "", x, y, err, p&PermWrite != 0, KB, MB)
	// inner x: 10
	// 0 true 1 2 boom false 1024 1048576
}
```

## Shadowing

Подробно: [05_shadowing.md](./01_theory/05_shadowing.md).

| Ситуация | Что происходит | Как правильно |
| :--- | :--- | :--- |
| `x := 2` во вложенном блоке | новая `x`, внешняя не меняется | `x = 2` |
| `err := doB()` внутри `if` | новая `err`, внешняя остаётся от `doA()` | `err = doB()` |
| `result := compute()` в замыкании | внешняя `result` не меняется | `result = compute()` |
| Поиск | `staticcheck`, `golangci-lint` | — |

## Ключевые слова (25)

Подробно: [02_reserved_keywords.md](./01_theory/02_reserved_keywords.md).

| Группа | Слова |
| :--- | :--- |
| Объявления | `var` `const` `func` `type` `struct` `interface` `map` `chan` `import` `package` |
| Управление потоком | `if` `else` `switch` `case` `default` `fallthrough` `for` `range` `break` `continue` `goto` `return` |
| Конкурентность, отложенный вызов | `go` `defer` `select` |
| Предопределённые идентификаторы (не ключевые слова) | `true` `false` `nil` `int` `len` `make` … — переопределить можно, но не нужно |

## Ловушки

- ⚠️ `a, b := f()` без новых имён слева в том же блоке → ошибка `no new variables on left side of :=`.
- ⚠️ `:=` во вложенном блоке создаёт новую переменную — частый источник «потерянной» `err`.
- ⚠️ Запись в nil-мапу (`var m map[string]int; m["x"] = 1`) → panic.
- ⚠️ `:=` на уровне пакета не компилируется — только `var`.
- ⚠️ Константа, которая на деле зависит от окружения, должна быть конфигом (env/flags), а не `const`.
- ⚠️ Глобальная `var`, изменяемая из разных мест, — скрытая связность; допустимы `os.Stdout`, `http.DefaultClient`, `Version`.
- ⚠️ `_ = x` вместо удаления ненужной переменной — маскирует мёртвый код.

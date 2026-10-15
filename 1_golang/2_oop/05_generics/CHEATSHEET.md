# Generics — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Generics — с Go 1.18; `cmp.Ordered`, пакеты `slices`/`maps` — с 1.21; generic type aliases — с 1.24.

## Type parameters

Подробно: [01_type_parameters.md](./01_theory/01_type_parameters.md).

| Задача | Код |
| :--- | :--- |
| Generic-функция | `func Identity[T any](v T) T { return v }` |
| Несколько параметров | `func Get[K comparable, V any](m map[K]V, k K) (V, bool)` |
| Вывод типа из аргументов | `Identity(10)` → `T = int` |
| Явная инстанциация | `Identity[int](10)`, `Get[string, User](users, "alice")` |
| Функция как значение | `f := Identity[int]` (без инстанциации нельзя) |
| Параметры зависят друг от друга | `func First[S ~[]E, E any](s S) E` — подходит и `type IDs []int` |
| Нулевое значение `T` | `var zero T; return zero, false` |

Реализация — stenciling по GC shape: не reflection (как Java) и не полная мономорфизация (как C++).

## Constraints

Подробно: [02_constraints.md](./01_theory/02_constraints.md).

| Constraint | Что разрешает | Пример |
| :--- | :--- | :--- |
| `any` | любой тип; операций почти нет | `func First[T any](s []T) T` |
| `comparable` | `==`, `!=`; ключ map | `func Index[T comparable](s []T, v T) int` |
| `cmp.Ordered` (1.21+) | `<`, `<=`, `>`, `>=`: целые, float, строки и типы на их основе | `func Max[T cmp.Ordered](a, b T) T` |
| Интерфейс с методами | вызов методов | `func Join[T fmt.Stringer](xs []T, sep string) string` |
| Union | перечисленные типы ровно | `interface{ int \| int64 \| float64 }` |
| `~T` (approximation) | все типы с underlying `T` | `interface{ ~int \| ~float64 }` — подходит `type Celsius float64` |
| `golang.org/x/exp/constraints` | `Integer`, `Float`, `Signed`, `Unsigned`, `Complex` | внешний модуль; часто проще свой union |

Минимально достаточный constraint: `any` → `comparable` → `cmp.Ordered` / свой union. `~` с интерфейсом (`~io.Reader`) запрещён.

## Generic типы и интерфейсы

Подробно: [03_generic_structs.md](./01_theory/03_generic_structs.md), [04_generic_interfaces.md](./01_theory/04_generic_interfaces.md).

| Задача | Код |
| :--- | :--- |
| Generic struct | `type Stack[T any] struct { items []T }` |
| Метод (параметры типа повторяются у receiver) | `func (s *Stack[T]) Push(v T)` |
| Инстанциация | `s := Stack[int]{}` / `var s Stack[string]` |
| Несколько параметров | `type Cache[K comparable, V any] struct { m map[K]V }` |
| Generic на map | `type Set[T comparable] map[T]struct{}` |
| Generic alias (1.24+) | `type Set[T comparable] = map[T]struct{}` |
| Generic интерфейс (тип значения) | `type Repo[T any] interface { Save(ctx context.Context, v T) error }` |
| Использование | `var c Container[int] = &IntBox{}` |
| Композиция | `type ReadWriter[T any] interface { Reader[T]; Writer[T] }` |
| Метод с собственным `[U any]` | **нельзя** (и в 1.26): `method must have no type parameters` → свободная функция `StackMap[T, U any](s *Stack[T], f func(T) U) []U` |

Constraint-интерфейс (`interface{ ~int \| ~float64 }`) — только в `[...]`; generic-интерфейс — обычный тип значения.

## Пример

```go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](xs []T) T {
	var s T
	for _, x := range xs {
		s += x
	}
	return s
}

func Map[T, U any](xs []T, f func(T) U) []U {
	res := make([]U, 0, len(xs))
	for _, x := range xs {
		res = append(res, f(x))
	}
	return res
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

type Celsius float64

func main() {
	fmt.Println(Sum([]Celsius{1, 2.5})) // 3.5 — благодаря ~float64
	toStr := func(i int) string { return fmt.Sprint(i * 10) }
	fmt.Println(Map([]int{1, 2}, toStr))  // [10 20]
	fmt.Println(Max("a", "b"), Max(3, 2)) // b 3
	fmt.Println(Max(1, 2.5))              // 2.5 — T = float64 (Go 1.21+)

	var s Stack[string]
	s.Push("x")
	v, ok := s.Pop()
	_, ok2 := s.Pop()
	fmt.Println(v, ok, ok2) // x true false

	fmt.Println(slices.Index([]int{5, 7}, 7), slices.Max([]int{3, 9, 1})) // 1 9
}
```

## Когда generics

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

| Да | Нет |
| :--- | :--- |
| Контейнеры: `Set`, `Stack`, `Queue`, `RingBuffer`, `LRUCache` | используется ровно с одним типом |
| `Map`/`Filter`/`Reduce`, если stdlib не покрывает | есть готовое: `slices.Index`, `slices.Contains`, `maps.Clone`, `maps.Equal` |
| `Result[T]` / `Option[T]`, `Box[T]` вместо `any` + assertion | сигнатура читается хуже двух простых функций |
| Числовые алгоритмы: одна `Sum[T Number]` вместо пяти | разное поведение по типам → интерфейс с методом |
| Типовой CRUD `Repo[T]` | бизнес-логика — всё равно в конкретных сервисах |

## Ловушки

Подробно: [05_pitfalls.md](./01_theory/05_pitfalls.md).

- ⚠️ Нет generic-методов и методов с type parameter в интерфейсе → свободные функции.
- ⚠️ Нет `T | error` / sum types → `type Result[T any] struct { V T; Err error }`.
- ⚠️ `[T any]` и потом ничего нельзя сделать с `T` → constraint слишком слабый.
- ⚠️ `==` на generic struct с полем `T` работает только при `T comparable`.
- ⚠️ Нет параметров-значений (`func F[N int]()`) и variadic type parameters.
- ⚠️ Длинные union в сигнатурах → вынести в именованный constraint (`Number`).
- ⚠️ Каждая GC shape — своя копия кода: сотни инстанциаций раздувают бинарник.
- ⚠️ `x/exp/constraints.Ordered` и `cmp.Ordered` вперемешку → на Go 1.21+ только `cmp.Ordered`.
- ⚠️ `Max(1, 2.5)` до Go 1.21 — ошибка вывода типов; с 1.21 нетипизированные константы разных видов дают `float64`. `Max(a, b)` с `a int`, `b float64` — ошибка всегда.

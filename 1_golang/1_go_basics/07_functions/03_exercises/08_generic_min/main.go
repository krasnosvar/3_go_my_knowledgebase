// ⏩ Забегание вперёд: дженерики и ограничения (constraints) подробно — в 2_oop/05_generics
// (после интерфейсов 2_oop/02_interfaces). Здесь достаточно синтаксиса func F[T cmp.Ordered](...).

package main

import (
	"cmp"
	"fmt"
)

// Min — обобщённая функция, возвращающая минимальное из двух значений типа T.
// В качестве ограничения (constraint) используется cmp.Ordered из стандартной библиотеки Go (1.21+).
func Min[T cmp.Ordered](a, b T) T {
	// TODO: Верните меньшее из a и b (сравнение через <).
	var zero T
	return zero
}

func main() {
	fmt.Println("--- Проверка 08_generic_min ---")

	t1 := Min(10, 20) == 10
	t2 := Min("apple", "banana") == "apple"
	t3 := Min(3.14, 2.71) == 2.71

	if t1 && t2 && t3 {
		fmt.Println("Тест дженерика Min: PASS")
	} else {
		fmt.Println("Тест дженерика Min: FAIL")
	}
}

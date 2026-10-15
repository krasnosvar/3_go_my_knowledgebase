package main

import (
	"cmp"
	"fmt"
)

// Min возвращает наименьшее из двух значений с помощью дженериков
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("--- Проверка 08_generic_min (Решение) ---")

	t1 := Min(10, 20) == 10
	t2 := Min("apple", "banana") == "apple"
	t3 := Min(3.14, 2.71) == 2.71

	if t1 && t2 && t3 {
		fmt.Println("Тест дженерика Min: PASS")
	} else {
		fmt.Println("Тест дженерика Min: FAIL")
	}
}

package main

import (
	"cmp"
	"fmt"
)

// TODO: Напишите обобщенную функцию Min, принимающую два параметра типа T и возвращающую минимальный из них.
// В качестве ограничения (constraint) используйте cmp.Ordered из стандартной библиотеки Go.
// Сигнатура: func Min[T cmp.Ordered](a, b T) T { ... }

var _ = cmp.Compare[int] // Держим импорт активным для компиляции

func main() {
	fmt.Println("--- Проверка 08_generic_min ---")

	// ВНИМАНИЕ: Раскомментируйте блок ниже после объявления функции Min.

	/*
		t1 := Min(10, 20) == 10
		t2 := Min("apple", "banana") == "apple"
		t3 := Min(3.14, 2.71) == 2.71

		if t1 && t2 && t3 {
			fmt.Println("Тест дженерика Min: PASS")
		} else {
			fmt.Println("Тест дженерика Min: FAIL")
		}
	*/

	fmt.Println("Раскомментируйте проверку в main")
}

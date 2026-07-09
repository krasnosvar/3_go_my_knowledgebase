package main

import (
	"fmt"
)

// runSubsliceDemo принимает слайс s ( гарантированно len=4, cap=4)
// 1. Создает подлайс sub = s[1:3]
// 2. Меняет первый элемент sub[0] на 99
// 3. Добавляет в sub элемент 100 с помощью append (внимание: превысит ли это емкость sub?)
// 4. Возвращает измененный оригинальный слайс s и измененный подлайс sub.
func runSubsliceDemo(s []int) ([]int, []int) {
	// TODO: Реализуйте шаги 1-3.
	return nil, nil
}

func main() {
	fmt.Println("--- Проверка 04_subslice_pitfalls ---")

	s := []int{10, 20, 30, 40} // len=4, cap=4
	orig, sub := runSubsliceDemo(s)

	// Разбор логики:
	// s[1:3] дает sub со ссылкой на [20, 30], len=2, cap=3.
	// sub[0] = 99 меняет оригинальный массив, теперь orig[1] == 99.
	// append(sub, 100) дописывает 100 в свободную ячейку cap (индекс 3 в оригинале).
	// Таким образом, orig[3] становится равным 100.
	// Результаты:
	// orig должен быть [10, 99, 30, 100]
	// sub должен быть [99, 30, 100]

	t1 := false
	if len(orig) == 4 {
		t1 = orig[0] == 10 && orig[1] == 99 && orig[2] == 30 && orig[3] == 100
	}
	t2 := false
	if len(sub) == 3 {
		t2 = sub[0] == 99 && sub[1] == 30 && sub[2] == 100
	}

	if t1 && t2 {
		fmt.Println("Тест подлайсов: PASS")
	} else {
		fmt.Printf("Тест подлайсов: FAIL (orig=%v, sub=%v)\n", orig, sub)
	}
}

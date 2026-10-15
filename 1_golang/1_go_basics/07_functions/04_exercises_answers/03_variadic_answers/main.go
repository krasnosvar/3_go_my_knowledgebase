package main

import (
	"fmt"
)

// sum возвращает сумму вариативных аргументов
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("--- Проверка 03_variadic (Решение) ---")

	t1 := sum(1, 2, 3, 4) == 10
	t2 := sum() == 0
	t3 := sum(5) == 5

	if t1 && t2 && t3 {
		fmt.Println("Тест variadic sum: PASS")
	} else {
		fmt.Printf("Тест variadic sum: FAIL (t1=%t, t2=%t, t3=%t)\n", t1, t2, t3)
	}
}

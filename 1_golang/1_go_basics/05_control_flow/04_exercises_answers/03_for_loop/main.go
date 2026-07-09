package main

import (
	"fmt"
)

// sumN возвращает сумму всех чисел от 1 до n
func sumN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println("--- Проверка 03_for_loop (Решение) ---")

	t1 := sumN(5) == 15
	t2 := sumN(10) == 55
	t3 := sumN(0) == 0

	if t1 && t2 && t3 {
		fmt.Println("Тест sumN: PASS")
	} else {
		fmt.Printf("Тест sumN: FAIL (t1=%t, t2=%t, t3=%t)\n", t1, t2, t3)
	}
}

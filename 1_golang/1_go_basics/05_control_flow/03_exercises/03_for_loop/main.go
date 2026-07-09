package main

import (
	"fmt"
)

// sumN возвращает сумму всех чисел от 1 до n включительно с помощью классического цикла for.
// Если n <= 0, возвращает 0.
func sumN(n int) int {
	// TODO: Реализуйте цикл for.
	return 0
}

func main() {
	fmt.Println("--- Проверка 03_for_loop ---")

	t1 := sumN(5) == 15  // 1+2+3+4+5 = 15
	t2 := sumN(10) == 55 // 55
	t3 := sumN(0) == 0

	if t1 && t2 && t3 {
		fmt.Println("Тест sumN: PASS")
	} else {
		fmt.Printf("Тест sumN: FAIL (t1=%t, t2=%t, t3=%t)\n", t1, t2, t3)
	}
}

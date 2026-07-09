package main

import (
	"fmt"
)

// sumEvenUntilLimit суммирует только четные числа из слайса nums до превышения лимита
func sumEvenUntilLimit(nums []int, limit int) int {
	sum := 0
	for _, val := range nums {
		if val%2 != 0 {
			continue
		}
		sum += val
		if sum > limit {
			break
		}
	}
	return sum
}

func main() {
	fmt.Println("--- Проверка 05_break_continue (Решение) ---")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	r1 := sumEvenUntilLimit(s, 10)
	t1 := r1 == 12

	r2 := sumEvenUntilLimit(s, 100)
	t2 := r2 == 20

	if t1 && t2 {
		fmt.Println("Тест break/continue: PASS")
	} else {
		fmt.Printf("Тест break/continue: FAIL (r1=%d, r2=%d)\n", r1, r2)
	}
}

package main

import (
	"fmt"
)

// filterPositive оставляет в nums только положительные числа (> 0).
// Операция должна быть произведена in-place (без выделения нового backing array через make).
// Используйте паттерн nums[:0] и append.
func filterPositive(nums []int) []int {
	// TODO: Реализуйте in-place фильтрацию.
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_in_place_ops ---")

	s := []int{-1, 5, 0, -3, 10}

	// Запоминаем указатель на исходный backing array
	origPtr := fmt.Sprintf("%p", s)

	res := filterPositive(s)
	resPtr := fmt.Sprintf("%p", res)

	t1 := len(res) == 2 && res[0] == 5 && res[1] == 10
	t2 := origPtr == resPtr // Backing array должен остаться тем же!

	if t1 && t2 {
		fmt.Println("Тест filterPositive: PASS")
	} else {
		fmt.Printf("Тест filterPositive: FAIL (len=%d, val=%v, samePtr=%t)\n", len(res), res, t1 && t2)
	}
}

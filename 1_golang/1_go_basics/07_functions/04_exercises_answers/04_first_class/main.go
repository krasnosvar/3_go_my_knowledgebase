package main

import (
	"fmt"
)

// applyOperation применяет функцию-операцию к слайсу чисел
func applyOperation(nums []int, op func(int) int) []int {
	res := make([]int, len(nums))
	for i, val := range nums {
		res[i] = op(val)
	}
	return res
}

func main() {
	fmt.Println("--- Проверка 04_first_class (Решение) ---")

	s := []int{1, 2, 3}
	doubleOp := func(x int) int { return x * 2 }
	res := applyOperation(s, doubleOp)

	if len(res) == 3 && res[0] == 2 && res[1] == 4 && res[2] == 6 {
		fmt.Println("Тест applyOperation: PASS")
	} else {
		fmt.Printf("Тест applyOperation: FAIL (получено %v)\n", res)
	}
}

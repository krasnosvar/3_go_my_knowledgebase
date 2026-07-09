package main

import (
	"fmt"
)

// applyOperation принимает слайс чисел и функцию-операцию,
// применяет операцию к каждому элементу и возвращает новый слайс.
func applyOperation(nums []int, op func(int) int) []int {
	// TODO: Реализуйте применение функции op.
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_first_class ---")

	s := []int{1, 2, 3}

	// Операция: умножение на 2
	doubleOp := func(x int) int { return x * 2 }

	res := applyOperation(s, doubleOp)

	if len(res) == 3 && res[0] == 2 && res[1] == 4 && res[2] == 6 {
		fmt.Println("Тест applyOperation: PASS")
	} else {
		fmt.Printf("Тест applyOperation: FAIL (получено %v)\n", res)
	}
}

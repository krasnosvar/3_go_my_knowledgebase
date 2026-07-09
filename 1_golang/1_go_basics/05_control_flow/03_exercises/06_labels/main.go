package main

import (
	"fmt"
)

// findValueInMatrix ищет target в двумерном массиве matrix.
// Если значение найдено, функция должна немедленно прервать ВСЕ циклы с помощью break LABEL
// и вернуть координаты (row, col) и true. Если не найдено — вернуть -1, -1, false.
func findValueInMatrix(matrix [][]int, target int) (int, int, bool) {
	rowIdx, colIdx := -1, -1
	found := false

	// TODO: Реализуйте вложенные циклы с меткой для немедленного выхода при нахождении target.

	return rowIdx, colIdx, found
}

func main() {
	fmt.Println("--- Проверка 06_labels ---")

	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	r, c, ok := findValueInMatrix(m, 5)

	if ok && r == 1 && c == 1 {
		fmt.Println("Тест поиска с меткой: PASS")
	} else {
		fmt.Printf("Тест поиска с меткой: FAIL (row=%d, col=%d, found=%t)\n", r, c, ok)
	}
}

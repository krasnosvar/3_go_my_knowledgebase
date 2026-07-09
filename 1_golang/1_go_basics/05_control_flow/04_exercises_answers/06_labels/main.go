package main

import (
	"fmt"
)

// findValueInMatrix ищет target в двумерном массиве matrix
func findValueInMatrix(matrix [][]int, target int) (int, int, bool) {
	rowIdx, colIdx := -1, -1
	found := false

SearchLabel:
	for r, row := range matrix {
		for c, val := range row {
			if val == target {
				rowIdx = r
				colIdx = c
				found = true
				break SearchLabel
			}
		}
	}

	return rowIdx, colIdx, found
}

func main() {
	fmt.Println("--- Проверка 06_labels (Решение) ---")

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

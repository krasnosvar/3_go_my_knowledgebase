package main

import (
	"fmt"
)

// formatPercentage умножает val на 100 и форматирует с точностью до 2 знаков
func formatPercentage(val float64) string {
	return fmt.Sprintf("%.2f%%", val*100)
}

func main() {
	fmt.Println("--- Проверка 06_width_precision (Решение) ---")

	output := formatPercentage(0.73456)
	expected := "73.46%"

	if output == expected {
		fmt.Println("Тест formatPercentage: PASS")
	} else {
		fmt.Printf("Тест formatPercentage: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}

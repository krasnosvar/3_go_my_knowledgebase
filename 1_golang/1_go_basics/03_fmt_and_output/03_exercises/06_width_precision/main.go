package main

import (
	"fmt"
)

// formatPercentage принимает число вроде 0.85421 и возвращает строку "85.42%"
func formatPercentage(val float64) string {
	// TODO: Умножьте val на 100 и отформатируйте с точностью до 2 знаков после запятой,
	// добавив символ процента в конце. Используйте fmt.Sprintf.
	return ""
}

func main() {
	fmt.Println("--- Проверка 06_width_precision ---")

	output := formatPercentage(0.73456)
	expected := "73.46%"

	if output == expected {
		fmt.Println("Тест formatPercentage: PASS")
	} else {
		fmt.Printf("Тест formatPercentage: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}

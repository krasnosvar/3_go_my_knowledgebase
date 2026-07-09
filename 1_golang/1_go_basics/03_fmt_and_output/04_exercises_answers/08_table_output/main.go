package main

import (
	"fmt"
)

// formatRow форматирует строку таблицы с тремя колонками фиксированной ширины
func formatRow(metric string, cores int, load float64) string {
	return fmt.Sprintf("%-10s%5d%8.2f", metric, cores, load)
}

func main() {
	fmt.Println("--- Проверка 08_table_output (Решение) ---")

	output := formatRow("CPU", 4, 12.345)
	expected := "CPU           4   12.35"

	if output == expected {
		fmt.Println("Тест formatRow: PASS")
	} else {
		fmt.Printf("Тест formatRow: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}

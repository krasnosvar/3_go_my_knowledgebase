package main

import (
	"fmt"
)

// formatRow форматирует строку таблицы с тремя колонками фиксированной ширины:
// 1. Название метрики: ширина 10, выравнивание по левому краю.
// 2. Количество ядер: ширина 5, выравнивание по правому краю.
// 3. Загрузка: ширина 8, выравнивание по правому краю, 2 знака после запятой.
func formatRow(metric string, cores int, load float64) string {
	// TODO: Сформируйте строку с фиксированной шириной колонок с помощью fmt.Sprintf.
	return ""
}

func main() {
	fmt.Println("--- Проверка 08_table_output ---")

	output := formatRow("CPU", 4, 12.345)
	expected := "CPU           4   12.35" // "CPU       " (10 chars), "    4" (5 chars), "   12.35" (8 chars)

	if output == expected {
		fmt.Println("Тест formatRow: PASS")
	} else {
		fmt.Printf("Тест formatRow: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}

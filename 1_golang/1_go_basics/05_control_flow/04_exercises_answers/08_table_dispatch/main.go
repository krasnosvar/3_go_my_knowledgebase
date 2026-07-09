package main

import (
	"fmt"
)

// getDescription возвращает строку описания типа переданного интерфейсного значения
func getDescription(val interface{}) string {
	switch val.(type) {
	case int:
		return "integer"
	case string:
		return "string"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println("--- Проверка 08_table_dispatch (Решение) ---")

	t1 := getDescription(42) == "integer"
	t2 := getDescription("hello") == "string"
	t3 := getDescription(true) == "boolean"
	t4 := getDescription(3.14) == "unknown"

	if t1 && t2 && t3 && t4 {
		fmt.Println("Тест type switch: PASS")
	} else {
		fmt.Println("Тест type switch: FAIL")
	}
}

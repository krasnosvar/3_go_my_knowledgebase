package main

import (
	"fmt"
)

// getDescription возвращает строку описания типа переданного интерфейсного значения:
// int -> "integer"
// string -> "string"
// bool -> "boolean"
// Любые другие типы -> "unknown"
// Используйте type switch (switch v := val.(type)).
func getDescription(val interface{}) string {
	// TODO: Реализуйте type switch.
	return ""
}

func main() {
	fmt.Println("--- Проверка 08_table_dispatch ---")

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

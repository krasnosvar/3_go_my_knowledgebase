package main

import (
	"fmt"
)

// getStatusName возвращает текстовое описание кода состояния HTTP:
// 200 -> "OK"
// 404 -> "Not Found"
// 500 -> "Internal Server Error"
// Все остальные -> "Unknown"
// Используйте switch.
func getStatusName(code int) string {
	// TODO: Реализуйте switch.
	return ""
}

func main() {
	fmt.Println("--- Проверка 02_switch ---")

	t1 := getStatusName(200) == "OK"
	t2 := getStatusName(404) == "Not Found"
	t3 := getStatusName(500) == "Internal Server Error"
	t4 := getStatusName(403) == "Unknown"

	if t1 && t2 && t3 && t4 {
		fmt.Println("Тест getStatusName: PASS")
	} else {
		fmt.Println("Тест getStatusName: FAIL")
	}
}

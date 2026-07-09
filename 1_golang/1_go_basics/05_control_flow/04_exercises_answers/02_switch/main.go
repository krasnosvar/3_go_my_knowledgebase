package main

import (
	"fmt"
)

// getStatusName возвращает текстовое описание кода состояния HTTP
func getStatusName(code int) string {
	switch code {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("--- Проверка 02_switch (Решение) ---")

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

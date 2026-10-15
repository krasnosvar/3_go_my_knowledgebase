package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Проверка 02_short_decl (Решение) ---")

	// Объявляем переменные с помощью :=
	x := 42
	y := "Hello"

	// Тесты
	if x == 42 {
		fmt.Println("Тест x: PASS")
	} else {
		fmt.Printf("Тест x: FAIL (ожидалось 42, получено %v)\n", x)
	}

	if y == "Hello" {
		fmt.Println("Тест y: PASS")
	} else {
		fmt.Printf("Тест y: FAIL (ожидалось \"Hello\", получено %q)\n", y)
	}
}

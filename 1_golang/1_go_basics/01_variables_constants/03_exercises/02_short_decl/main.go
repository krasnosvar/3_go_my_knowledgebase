package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Проверка 02_short_decl ---")

	// TODO: Объявите переменные `x` со значением 42 и `y` со значением "Hello" с помощью краткого объявления `:=`.
	// x := 42
	// y := "Hello"

	// ВНИМАНИЕ: Для прохождения теста замените два объявления ниже на краткую форму из TODO.
	var x int
	var y string

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

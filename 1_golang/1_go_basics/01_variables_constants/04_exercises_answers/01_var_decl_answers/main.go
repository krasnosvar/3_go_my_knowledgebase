package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Проверка 01_var_decl (Решение) ---")

	// Объявляем с явным типом и значением
	var age int = 10

	// Объявляем без указания типа (компилятор выведет тип string)
	var name = "Go"

	// Тесты
	if age == 10 {
		fmt.Println("Тест age: PASS")
	} else {
		fmt.Printf("Тест age: FAIL (ожидалось 10, получено %d)\n", age)
	}

	if name == "Go" {
		fmt.Println("Тест name: PASS")
	} else {
		fmt.Printf("Тест name: FAIL (ожидалось \"Go\", получено %q)\n", name)
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Проверка 01_var_decl ---")

	// TODO: Объявите переменную `age` типа int со значением 10 с помощью ключевого слова `var`.
	// var age int = 10

	// TODO: Объявите переменную `name` типа string со значением "Go" без указания типа (пусть компилятор выведет тип сам).
	// var name = "Go"

	// ВНИМАНИЕ: Для прохождения теста раскомментируйте переменные ниже и присвойте им значения из TODO
	var age int
	var name string

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

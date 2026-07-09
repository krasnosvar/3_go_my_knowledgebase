package main

import "fmt"

/*
	ОТВЕТЫ К ЗАДАНИЮ 1: Основы работы с мапами
*/

func main() {
	fmt.Println("--- Task 1: CRUD ---")
	task1_crud()

	fmt.Println("\n--- Task 2: Проверка наличия ключа ---")
	task2_check_key()
}

func task1_crud() {
	// ================= РЕШЕНИЕ: Task 1 =================
	m := make(map[string]int) // Шаг 1
	m["a"] = 10               // Шаг 2
	m["a"] = 20               // Шаг 3
	m["b"] = 30               // Шаг 4
	delete(m, "a")            // Шаг 5

	fmt.Println(m)
	// ====================================================
}

func task2_check_key() {
	// ================= РЕШЕНИЕ: Task 2 =================
	m := map[string]int{"a": 1}

	if v, ok := m["b"]; ok {
		fmt.Println("Значение:", v)
	} else {
		fmt.Println("Ключ не найден")
	}
	// ====================================================
}

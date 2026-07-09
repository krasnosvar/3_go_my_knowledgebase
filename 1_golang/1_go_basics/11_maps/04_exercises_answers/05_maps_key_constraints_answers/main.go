package main

import "fmt"

/*
	ОТВЕТЫ К ЗАДАНИЮ 5
*/

func main() {
	fmt.Println("--- Task 1: Структура как ключ ---")
	task1_struct_key()

	fmt.Println("\n--- Task 2: Интерфейс как ключ ---")
	task2_interface_key()
}

func task1_struct_key() {
	// ================= РЕШЕНИЕ: Task 1 =================
	type Point struct {
		X int
		Y int
	}

	pointsColors := make(map[Point]string)

	// Структуры сравниваются по значению полей, поэтому их можно использовать как ключи
	p1 := Point{0, 0}
	pointsColors[p1] = "Red"

	// При запросе по той же координате мы получим "Red"
	fmt.Println("Цвет (0,0):", pointsColors[Point{0, 0}])
	// ====================================================
}

func task2_interface_key() {
	// ================= РЕШЕНИЕ: Task 2 =================
	m := make(map[any]string)

	// Это работает
	m[1] = "A"
	m["string_key"] = "C"

	// Это вызовет panic: panic: runtime error: hash of unhashable type []int
	// sl := []int{1, 2, 3}
	// m[sl] = "B"

	fmt.Println("Map with any keys:", m)
	// ====================================================
}

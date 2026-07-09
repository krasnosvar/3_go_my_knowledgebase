package main

import "fmt"

func main() {
	fmt.Println("--- Task 1: Очистка мапы ---")
	task1_clear_map()
}

func task1_clear_map() {
	// ================= РЕШЕНИЕ =================
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i * 2
	}

	fmt.Println("До очистки, длина:", len(m))

	// Способ 1 (Плохой): цикл + delete.
	// Оставляет пустыми "бакеты" внутри мапы, память не отдается ОС!
	// for k := range m {
	// 	delete(m, k)
	// }

	// Способ 2 (Стандартный до Go 1.21): просто создать новую мапу
	// Старая мапа соберется Garbage Collector-ом.
	// m = make(map[int]int)

	// Способ 3 (Лучший для Go 1.21+): встроенная функция clear
	clear(m)

	fmt.Println("После очистки (clear), длина:", len(m))
	// ===========================================
}

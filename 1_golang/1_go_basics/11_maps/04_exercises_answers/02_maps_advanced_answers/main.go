package main

import "fmt"

/*
	ОТВЕТЫ К ЗАДАНИЮ 2: Продвинутые операции с мапами
*/

func main() {
	fmt.Println("--- Task 1: Вложенные мапы ---")
	task1_nested_maps()

	fmt.Println("\n--- Task 2: Мапа со слайсами (группировка) ---")
	task2_slice_values()
}

func task1_nested_maps() {
	// ================= РЕШЕНИЕ: Task 1 =================

	// Шаг 1: Инициализируем внешнюю мапу
	users := make(map[string]map[string]int)

	// Шаг 2: Важно проверить и проинициализировать внутреннюю мапу, прежде чем туда писать!
	if users["Alice"] == nil {
		users["Alice"] = make(map[string]int)
	}
	users["Alice"]["age"] = 25

	// Шаг 3: Обновляем значение (это сработает, так как внутренняя мапа проинициализирована)
	users["Alice"]["age"]++

	fmt.Printf("Alice's age is: %d\n", users["Alice"]["age"])

	// ====================================================
}

type User struct {
	Name string
	Age  int
}

func task2_slice_values() {
	// ================= РЕШЕНИЕ: Task 2 =================

	people := []User{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"Diana", 30},
		{"Eve", 22},
	}

	// Шаг 1: Инициализируем мапу
	groupedByAge := make(map[int][]string)

	// Шаг 2: Проходимся по слайсу и заполняем мапу
	// Подсказка: функция append() корректно работает, даже если слайса (пока что) нет.
	// При первом обращении m[age] вернет нулевое значение (nil слайс), и append создаст новый слайс!
	for _, p := range people {
		groupedByAge[p.Age] = append(groupedByAge[p.Age], p.Name)
	}

	// Выводим результат
	for age, names := range groupedByAge {
		fmt.Printf("Возраст %d: %v\n", age, names)
	}

	// ====================================================
}

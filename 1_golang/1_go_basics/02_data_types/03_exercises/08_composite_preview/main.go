package main

import (
	"fmt"
)

// initStructures должна инициализировать и вернуть:
// 1. Слайс целых чисел со значениями [10, 20, 30]
// 2. Инициализированную мапу map[string]int, содержащую ключи "A": 100 и "B": 200.
func initStructures() ([]int, map[string]int) {
	// TODO: реализуйте функцию
	return nil, nil
}

func main() {
	fmt.Println("--- Проверка 08_composite_preview ---")

	slice, m := initStructures()

	// Проверка слайса
	sliceOk := len(slice) == 3 && slice[0] == 10 && slice[1] == 20 && slice[2] == 30
	if sliceOk {
		fmt.Println("Тест слайса: PASS")
	} else {
		fmt.Printf("Тест слайса: FAIL (ожидалось [10 20 30], получено %v)\n", slice)
	}

	// Проверка мапы
	mapOk := m != nil && m["A"] == 100 && m["B"] == 200 && len(m) == 2
	if mapOk {
		fmt.Println("Тест мапы: PASS")
	} else {
		fmt.Printf("Тест мапы: FAIL (ожидалось map[A:100 B:200], получено %v)\n", m)
	}
}

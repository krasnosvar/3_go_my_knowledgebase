package main

import (
	"fmt"
)

// removeAtIndex удаляет элемент по индексу idx из слайса s с сохранением порядка.
// Используйте встроенные функции copy и слайсинг.
// Если индекс вне границ слайса, верните исходный слайс без изменений.
func removeAtIndex(s []int, idx int) []int {
	// TODO: Реализуйте удаление со сдвигом элементов через copy.
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_copy_delete ---")

	s := []int{10, 20, 30, 40, 50}

	// Удаляем индекс 2 (значение 30)
	res := removeAtIndex(s, 2) // Ожидается [10, 20, 40, 50]

	t1 := len(res) == 4
	t2 := false
	if len(res) == 4 {
		t2 = res[0] == 10 && res[1] == 20 && res[2] == 40 && res[3] == 50
	}

	if t1 && t2 {
		fmt.Println("Тест removeAtIndex: PASS")
	} else {
		fmt.Printf("Тест removeAtIndex: FAIL (получено %v)\n", res)
	}
}

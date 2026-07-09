package main

import (
	"fmt"
)

// removeAtIndex удаляет элемент по индексу
func removeAtIndex(s []int, idx int) []int {
	if idx < 0 || idx >= len(s) {
		return s
	}
	copy(s[idx:], s[idx+1:])
	return s[:len(s)-1]
}

func main() {
	fmt.Println("--- Проверка 03_copy_delete (Решение) ---")

	s := []int{10, 20, 30, 40, 50}
	res := removeAtIndex(s, 2)

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

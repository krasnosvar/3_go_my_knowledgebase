package main

import (
	"fmt"
)

// appendMultiple добавляет все переданные числа nums в конец слайса s с помощью append,
// и возвращает получившийся слайс.
func appendMultiple(s []int, nums ...int) []int {
	// TODO: Реализуйте добавление через append.
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_append_capacity ---")

	base := []int{1, 2}
	res := appendMultiple(base, 3, 4, 5)

	t1 := len(res) == 5
	t2 := false
	if len(res) == 5 {
		t2 = res[0] == 1 && res[1] == 2 && res[2] == 3 && res[3] == 4 && res[4] == 5
	}

	if t1 && t2 {
		fmt.Println("Тест appendMultiple: PASS")
	} else {
		fmt.Printf("Тест appendMultiple: FAIL (len=%d, val=%v)\n", len(res), res)
	}
}

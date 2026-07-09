package main

import (
	"fmt"
)

// sumSlice суммирует все числа в слайсе с помощью for range
func sumSlice(nums []int) int {
	// TODO: Реализуйте суммирование через range.
	return 0
}

// countKeys возвращает количество ключей в мапе с помощью for range (или len, но напишите через range для практики)
func countKeys(m map[string]int) int {
	// TODO: Реализуйте обход через range и инкремент счетчика.
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_range ---")

	s := []int{10, 20, 30}
	t1 := sumSlice(s) == 60

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	t2 := countKeys(m) == 3

	if t1 && t2 {
		fmt.Println("Тест range: PASS")
	} else {
		fmt.Printf("Тест range: FAIL (t1=%t, t2=%t)\n", t1, t2)
	}
}

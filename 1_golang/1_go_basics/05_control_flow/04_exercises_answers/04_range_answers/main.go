package main

import (
	"fmt"
)

// sumSlice суммирует все числа в слайсе с помощью for range
func sumSlice(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

// countKeys возвращает количество ключей в мапе
func countKeys(m map[string]int) int {
	count := 0
	for range m {
		count++
	}
	return count
}

func main() {
	fmt.Println("--- Проверка 04_range (Решение) ---")

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

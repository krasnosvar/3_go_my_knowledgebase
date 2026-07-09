package main

import (
	"fmt"
)

// filterPositive фильтрует слайс на месте
func filterPositive(nums []int) []int {
	res := nums[:0]
	for _, val := range nums {
		if val > 0 {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	fmt.Println("--- Проверка 06_in_place_ops (Решение) ---")

	s := []int{-1, 5, 0, -3, 10}

	origPtr := fmt.Sprintf("%p", s)
	res := filterPositive(s)
	resPtr := fmt.Sprintf("%p", res)

	t1 := len(res) == 2 && res[0] == 5 && res[1] == 10
	t2 := origPtr == resPtr

	if t1 && t2 {
		fmt.Println("Тест filterPositive: PASS")
	} else {
		fmt.Printf("Тест filterPositive: FAIL (len=%d, val=%v, samePtr=%t)\n", len(res), res, t1 && t2)
	}
}

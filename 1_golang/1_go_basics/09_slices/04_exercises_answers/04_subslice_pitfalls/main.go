package main

import (
	"fmt"
)

// runSubsliceDemo выполняет операции с подлайсом
func runSubsliceDemo(s []int) ([]int, []int) {
	sub := s[1:3]
	sub[0] = 99
	sub = append(sub, 100)
	return s, sub
}

func main() {
	fmt.Println("--- Проверка 04_subslice_pitfalls (Решение) ---")

	s := []int{10, 20, 30, 40}
	orig, sub := runSubsliceDemo(s)

	t1 := false
	if len(orig) == 4 {
		t1 = orig[0] == 10 && orig[1] == 99 && orig[2] == 30 && orig[3] == 100
	}
	t2 := false
	if len(sub) == 3 {
		t2 = sub[0] == 99 && sub[1] == 30 && sub[2] == 100
	}

	if t1 && t2 {
		fmt.Println("Тест подлайсов: PASS")
	} else {
		fmt.Printf("Тест подлайсов: FAIL (orig=%v, sub=%v)\n", orig, sub)
	}
}

package main

import (
	"fmt"
)

// createStepCounter возвращает замыкание-счетчик
func createStepCounter(step int) func() int {
	count := 0
	return func() int {
		count += step
		return count
	}
}

func main() {
	fmt.Println("--- Проверка 05_closure_counter (Решение) ---")

	counter1 := createStepCounter(2)
	counter2 := createStepCounter(5)

	t1 := counter1() == 2
	t2 := counter1() == 4
	t3 := counter2() == 5
	t4 := counter2() == 10

	if t1 && t2 && t3 && t4 {
		fmt.Println("Тест замыкания-счетчика: PASS")
	} else {
		fmt.Println("Тест замыкания-счетчика: FAIL")
	}
}

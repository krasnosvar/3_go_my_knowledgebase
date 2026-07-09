package main

import (
	"fmt"
)

// createStepCounter возвращает функцию-замыкание, которая при каждом вызове
// увеличивает внутренний счетчик на значение step и возвращает текущую сумму.
// Счетчик стартует с 0.
func createStepCounter(step int) func() int {
	// TODO: Реализуйте замыкание.
	return func() int { return 0 }
}

func main() {
	fmt.Println("--- Проверка 05_closure_counter ---")

	counter1 := createStepCounter(2)
	counter2 := createStepCounter(5)

	t1 := counter1() == 2  // 0 + 2 = 2
	t2 := counter1() == 4  // 2 + 2 = 4
	t3 := counter2() == 5  // 0 + 5 = 5
	t4 := counter2() == 10 // 5 + 5 = 10

	if t1 && t2 && t3 && t4 {
		fmt.Println("Тест замыкания-счетчика: PASS")
	} else {
		fmt.Println("Тест замыкания-счетчика: FAIL")
	}
}

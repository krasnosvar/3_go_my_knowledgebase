package main

import (
	"fmt"
)

type Counter struct {
	Count int
}

// TODO: Напишите метод Increment() для Counter с pointer receiver,
// который увеличивает Count на 1.

func main() {
	fmt.Println("--- Проверка 04_method_pointer_receiver ---")

	// ВНИМАНИЕ: Раскомментируйте код ниже после написания метода Increment.

	/*
		c := Counter{Count: 0}
		c.Increment()
		c.Increment()

		if c.Count == 2 {
			fmt.Println("Тест Increment: PASS")
		} else {
			fmt.Printf("Тест Increment: FAIL (получено %d)\n", c.Count)
		}
	*/

	fmt.Println("Раскомментируйте проверку в main")
}

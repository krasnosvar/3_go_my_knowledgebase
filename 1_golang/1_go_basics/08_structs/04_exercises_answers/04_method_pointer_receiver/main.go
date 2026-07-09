package main

import (
	"fmt"
)

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}

func main() {
	fmt.Println("--- Проверка 04_method_pointer_receiver (Решение) ---")

	c := Counter{Count: 0}
	c.Increment()
	c.Increment()

	if c.Count == 2 {
		fmt.Println("Тест Increment: PASS")
	} else {
		fmt.Printf("Тест Increment: FAIL (получено %d)\n", c.Count)
	}
}

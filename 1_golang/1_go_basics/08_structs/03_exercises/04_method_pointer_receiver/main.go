package main

import (
	"fmt"
)

type Counter struct {
	Count int
}

// Increment увеличивает Count на 1.
// TODO: Сейчас метод объявлен с value receiver (Counter) и меняет КОПИЮ структуры.
// Исправьте получатель на pointer receiver (*Counter).
func (c Counter) Increment() {
	c.Count++
}

func main() {
	fmt.Println("--- Проверка 04_method_pointer_receiver ---")

	c := Counter{Count: 0}
	c.Increment()
	c.Increment()

	if c.Count == 2 {
		fmt.Println("Тест Increment: PASS")
	} else {
		fmt.Printf("Тест Increment: FAIL (получено %d)\n", c.Count)
	}
}

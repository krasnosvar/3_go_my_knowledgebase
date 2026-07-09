package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	fmt.Println("--- Проверка 03_method_value_receiver (Решение) ---")

	r := Rectangle{Width: 10, Height: 5}
	area := r.Area()

	if area == 50.0 {
		fmt.Println("Тест Area: PASS")
	} else {
		fmt.Printf("Тест Area: FAIL (получено %.2f)\n", area)
	}
}

package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Height float64
}

// TODO: Напишите метод Area() float64 для Rectangle с value receiver.
// Формула: Width * Height.
// ВНИМАНИЕ: ниже заглушка, чтобы программа компилировалась — замените тело метода.
func (r Rectangle) Area() float64 {
	return 0
}

func main() {
	fmt.Println("--- Проверка 03_method_value_receiver ---")

	r := Rectangle{Width: 10, Height: 5}
	area := r.Area()

	if area == 50.0 {
		fmt.Println("Тест Area: PASS")
	} else {
		fmt.Printf("Тест Area: FAIL (получено %.2f)\n", area)
	}
}

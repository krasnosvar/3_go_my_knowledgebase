package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Height float64
}

// TODO: Напишите метод Area() float64 для Rectangle с value receiver.
// Формула: Width * Height.

func main() {
	fmt.Println("--- Проверка 03_method_value_receiver ---")

	// ВНИМАНИЕ: Раскомментируйте код ниже после написания метода Area.

	/*
		r := Rectangle{Width: 10, Height: 5}
		area := r.Area()

		if area == 50.0 {
			fmt.Println("Тест Area: PASS")
		} else {
			fmt.Printf("Тест Area: FAIL (получено %.2f)\n", area)
		}
	*/

	fmt.Println("Раскомментируйте проверку в main")
}

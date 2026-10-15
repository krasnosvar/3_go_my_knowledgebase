// Упражнение 01: методы с value receiver.
//
// Задача: реализуйте методы Area и Perimeter у Rect с value receiver
// (получатель по значению): метод только читает поля и не меняет структуру.
//   - Area() — площадь W*H;
//   - Perimeter() — периметр 2*(W+H).
//
// Запуск: go run ./03_exercises/01_value_receiver/main.go
package main

import (
	"fmt"
)

// Rect — прямоугольник.
type Rect struct {
	W, H int
}

// Area возвращает площадь прямоугольника.
func (r Rect) Area() int {
	// TODO: верните r.W * r.H.
	return 0
}

// Perimeter возвращает периметр прямоугольника.
func (r Rect) Perimeter() int {
	// TODO: верните 2 * (r.W + r.H).
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_value_receiver ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r := Rect{W: 3, H: 4}
	check(r.Area() == 12, fmt.Sprintf("Area 3x4: ожидалось 12, получено %d", r.Area()))
	check(r.Perimeter() == 14, fmt.Sprintf("Perimeter 3x4: ожидалось 14, получено %d", r.Perimeter()))
	check(Rect{}.Area() == 0, "Area нулевого прямоугольника должна быть 0")

	// Value receiver работает и на указателе: Go сам разыменует.
	p := &Rect{W: 2, H: 5}
	check(p.Area() == 10, fmt.Sprintf("Area через указатель: ожидалось 10, получено %d", p.Area()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

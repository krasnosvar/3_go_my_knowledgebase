// Упражнение 01: методы с value receiver (Решение).
package main

import (
	"fmt"
)

// Rect — прямоугольник.
type Rect struct {
	W, H int
}

func (r Rect) Area() int {
	return r.W * r.H
}

func (r Rect) Perimeter() int {
	return 2 * (r.W + r.H)
}

func main() {
	fmt.Println("--- Проверка 01_value_receiver (Решение) ---")

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

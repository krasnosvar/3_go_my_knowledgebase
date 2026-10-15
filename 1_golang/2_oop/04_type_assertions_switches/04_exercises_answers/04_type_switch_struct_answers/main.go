// Упражнение 04: type switch по структурам (Решение).
package main

import (
	"fmt"
	"math"
)

// Circle — круг.
type Circle struct{ R float64 }

// Rect — прямоугольник.
type Rect struct{ W, H float64 }

func area(v any) (float64, bool) {
	switch s := v.(type) {
	case Circle:
		return math.Pi * s.R * s.R, true
	case Rect:
		return s.W * s.H, true
	case *Rect:
		if s == nil {
			return 0, false
		}
		return s.W * s.H, true
	default:
		return 0, false
	}
}

func main() {
	fmt.Println("--- Проверка 04_type_switch_struct (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	near := func(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

	a, ok := area(Circle{R: 1})
	check(ok && near(a, math.Pi), fmt.Sprintf("Circle{1}: ожидалось (Pi, true), получено (%v, %t)", a, ok))

	a, ok = area(Rect{W: 2, H: 3})
	check(ok && near(a, 6), fmt.Sprintf("Rect{2,3}: ожидалось (6, true), получено (%v, %t)", a, ok))

	a, ok = area(&Rect{W: 4, H: 5})
	check(ok && near(a, 20), fmt.Sprintf("*Rect{4,5}: ожидалось (20, true), получено (%v, %t)", a, ok))

	var nilRect *Rect
	_, ok = area(nilRect)
	check(!ok, "nil *Rect: ожидалось ok=false")

	_, ok = area("square")
	check(!ok, "строка не фигура: ожидалось ok=false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// Упражнение 04: type switch по структурам.
//
// Задача: реализуйте area через type switch по конкретным типам фигур:
//
//	Circle -> math.Pi*R*R, Rect -> W*H, *Rect -> тоже W*H (не забудьте nil),
//	иначе -> (0, false). Для известной фигуры второй результат true.
//
// Запуск: go run ./03_exercises/04_type_switch_struct/main.go
package main

import (
	"fmt"
	"math"
)

// Circle — круг.
type Circle struct{ R float64 }

// Rect — прямоугольник.
type Rect struct{ W, H float64 }

// area вычисляет площадь фигуры.
func area(v any) (float64, bool) {
	// TODO: switch s := v.(type) { case Circle: ... case Rect: ... case *Rect: ... }
	_ = math.Pi
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 04_type_switch_struct ---")

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

// Упражнение 04: Open/Closed через интерфейс (Решение).
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	W, H float64
}

func (r Rectangle) Area() float64 { return r.W * r.H }

type Circle struct {
	R float64
}

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

func totalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

func almostEqual(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func main() {
	fmt.Println("--- Проверка 04_ocp_open_closed (Решение) ---")

	shapes := []Shape{
		Rectangle{W: 2, H: 3},
		Circle{R: 1},
	}
	want := 6 + math.Pi

	allOK := true
	if got := totalArea(shapes); !almostEqual(got, want) {
		allOK = false
		fmt.Printf("FAIL: totalArea = %.6f, ожидалось %.6f\n", got, want)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// Упражнение 04: Open/Closed через интерфейс (замена switch по типу).
//
// Задача (Хардкор): посчитать площадь набора фигур так, чтобы добавление
// новой фигуры НЕ требовало правки функции totalArea. Классический анти-
// пример — switch по типу внутри totalArea (его пришлось бы дополнять
// каждой новой фигурой). Вместо этого: интерфейс Shape с методом Area,
// а totalArea работает только с абстракцией.
//
// Реализуйте:
//   - метод Area у Rectangle (W*H) и у Circle (π*R*R, используйте math.Pi);
//   - функцию totalArea, суммирующую площади всех фигур через интерфейс.
//
// Заготовка компилируется и печатает FAIL — доведите до PASS.
//
// Запуск: go run ./03_exercises/04_ocp_open_closed/main.go
package main

import (
	"fmt"
	"math"
)

// Shape — абстракция. Новые фигуры = новые реализации, totalArea не меняется.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	W, H float64
}

// TODO: реализуйте Area для Rectangle (W * H).
func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	R float64
}

// TODO: реализуйте Area для Circle (math.Pi * R * R).
func (c Circle) Area() float64 {
	return 0
}

// totalArea суммирует площади всех фигур, зная только про интерфейс Shape.
func totalArea(shapes []Shape) float64 {
	// TODO: пройдите по shapes и просуммируйте s.Area().
	return 0
}

func almostEqual(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func main() {
	fmt.Println("--- Проверка 04_ocp_open_closed ---")

	shapes := []Shape{
		Rectangle{W: 2, H: 3}, // 6
		Circle{R: 1},          // π
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

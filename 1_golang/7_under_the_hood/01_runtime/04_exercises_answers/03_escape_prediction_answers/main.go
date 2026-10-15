// Упражнение 03: предсказание escape analysis (Решение).
package main

import "fmt"

// newCounter: переменная x.
func newCounter() *int {
	x := 0
	return &x
}

// sumArray: массив a.
func sumArray() int {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

// makeFixed: слайс s с константным размером.
func makeFixed() int {
	s := make([]int, 16)
	return len(s)
}

// sink — глобальная переменная пакета.
var sink []int

// storeGlobal: слайс s, сохраняемый в глобальную переменную.
func storeGlobal() int {
	s := make([]int, 4)
	sink = s
	return len(s)
}

// makeClosure: переменная x, захваченная возвращаемым замыканием.
func makeClosure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// predictions: имя функции -> "stack" или "heap".
var predictions = map[string]string{
	"newCounter":  "heap",  // указатель на локальную переменную возвращается наружу
	"sumArray":    "stack", // массив фиксированного размера не покидает функцию
	"makeFixed":   "stack", // константный размер и слайс не утекает
	"storeGlobal": "heap",  // слайс переживает функцию через глобальную переменную
	"makeClosure": "heap",  // x живёт дольше функции — его захватило замыкание
}

func main() {
	fmt.Println("--- Проверка 03_escape_prediction (Решение) ---")

	// Вызовы, чтобы функции не были «мёртвым кодом».
	_ = *newCounter() + sumArray() + makeFixed() + storeGlobal() + makeClosure()()

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	expected := map[string]string{
		"newCounter":  "heap",
		"sumArray":    "stack",
		"makeFixed":   "stack",
		"storeGlobal": "heap",
		"makeClosure": "heap",
	}
	for _, name := range []string{"newCounter", "sumArray", "makeFixed", "storeGlobal", "makeClosure"} {
		got := predictions[name]
		check(got == expected[name], fmt.Sprintf("%s: ожидалось %q, указано %q", name, expected[name], got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

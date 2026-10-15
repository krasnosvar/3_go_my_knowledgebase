// Упражнение 03: предсказание escape analysis.
//
// Задача: для каждой функции ниже предскажите, где окажется её главная
// переменная — на стеке ("stack") или в куче ("heap"), и заполните карту
// predictions. Проверить себя можно и компилятором:
//
//	go build -gcflags=-m ./03_exercises/03_escape_prediction/
//
// (ищите строки "moved to heap" / "escapes to heap" / "does not escape"
// по строкам самих функций: после инлайнинга в main вывод может отличаться).
//
// Запуск: go run ./03_exercises/03_escape_prediction/main.go
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
	// TODO: заполните для каждой функции.
	"newCounter":  "",
	"sumArray":    "",
	"makeFixed":   "",
	"storeGlobal": "",
	"makeClosure": "",
}

func main() {
	fmt.Println("--- Проверка 03_escape_prediction ---")

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

// Упражнение 07: приоритет срочного канала.
//
// Задача: next(urgent, normal) — неблокирующий выбор с приоритетом:
//  1. если в urgent есть значение — вернуть его;
//  2. иначе, если есть в normal — вернуть его;
//  3. иначе вернуть ("", false).
//
// Обычный select с двумя готовыми ветками выбирает СЛУЧАЙНО, поэтому
// приоритет делается двумя последовательными select с default.
//
// Запуск: go run ./03_exercises/07_priority/main.go
package main

import "fmt"

// next возвращает срочное сообщение, если оно есть, иначе обычное.
func next(urgent, normal <-chan string) (string, bool) {
	// TODO: сначала select по urgent с default, затем по normal с default.
	return "", false
}

func main() {
	fmt.Println("--- Проверка 07_priority ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	urgent := make(chan string, 10)
	normal := make(chan string, 10)
	normal <- "n1"
	normal <- "n2"
	urgent <- "u1"
	urgent <- "u2"

	order := []string{}
	for {
		m, ok := next(urgent, normal)
		if !ok {
			break
		}
		order = append(order, m)
		if len(order) > 10 {
			break
		}
	}
	check(fmt.Sprint(order) == "[u1 u2 n1 n2]", fmt.Sprintf("порядок = %v, ожидалось [u1 u2 n1 n2]", order))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

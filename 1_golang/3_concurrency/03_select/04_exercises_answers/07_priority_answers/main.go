// Упражнение 07: приоритет срочного канала (Решение).
package main

import "fmt"

func next(urgent, normal <-chan string) (string, bool) {
	select {
	case m := <-urgent:
		return m, true
	default:
	}
	select {
	case m := <-normal:
		return m, true
	default:
		return "", false
	}
}

func main() {
	fmt.Println("--- Проверка 07_priority (Решение) ---")

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

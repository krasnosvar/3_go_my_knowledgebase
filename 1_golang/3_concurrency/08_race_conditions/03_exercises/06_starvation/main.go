// Упражнение 06: защита низкого приоритета от голодания.
//
// Задача: fairMerge(high, low, maxStreak) вычитывает два буферизованных
// канала и возвращает порядок обработки. Приоритет у high, но строгий
// приоритет приводит к голоданию (starvation): пока в high есть данные,
// low не обслуживается никогда.
//
// Правило: брать из high, но после maxStreak подряд взятых из high, если в
// low что-то есть, обязательно взять один элемент из low (и сбросить серию).
// Если high пуст — брать из low. Если оба пусты — завершить.
// Неблокирующие чтения — select с default.
//
// Запуск: go run ./03_exercises/06_starvation/main.go
package main

import "fmt"

// fairMerge обслуживает high с приоритетом, но не морит low голодом.
func fairMerge(high, low <-chan string, maxStreak int) []string {
	// TODO
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_starvation ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	high := make(chan string, 10)
	low := make(chan string, 10)
	for i := 1; i <= 7; i++ {
		high <- fmt.Sprint("H", i)
	}
	for i := 1; i <= 3; i++ {
		low <- fmt.Sprint("L", i)
	}

	got := fairMerge(high, low, 3)
	want := "[H1 H2 H3 L1 H4 H5 H6 L2 H7 L3]"
	check(fmt.Sprint(got) == want, fmt.Sprintf("порядок = %v, ожидалось %s", got, want))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

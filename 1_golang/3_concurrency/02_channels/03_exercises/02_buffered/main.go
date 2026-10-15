// Упражнение 02: буферизованный канал.
//
// Задача: fillBuffered(vals) создаёт канал с буфером ровно len(vals),
// отправляет туда все значения БЕЗ отдельной горутины-читателя и
// возвращает канал (не закрывая его).
//
// Это работает, потому что отправка в буферизованный канал блокирует
// только при заполненном буфере.
//
// Запуск: go run ./03_exercises/02_buffered/main.go
package main

import "fmt"

// fillBuffered возвращает канал ёмкостью len(vals), заполненный vals.
func fillBuffered(vals []int) chan int {
	// TODO: создайте канал make(chan int, len(vals)) и отправьте все значения.
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_buffered ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ch := fillBuffered([]int{10, 20, 30})
	if ch == nil {
		check(false, "fillBuffered вернула nil-канал")
	} else {
		check(cap(ch) == 3, fmt.Sprintf("cap = %d, ожидалось 3", cap(ch)))
		check(len(ch) == 3, fmt.Sprintf("len = %d, ожидалось 3 (значения в буфере)", len(ch)))
		got := []int{}
		for len(ch) > 0 {
			got = append(got, <-ch)
		}
		check(fmt.Sprint(got) == "[10 20 30]", fmt.Sprintf("порядок = %v, ожидалось [10 20 30] (FIFO)", got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

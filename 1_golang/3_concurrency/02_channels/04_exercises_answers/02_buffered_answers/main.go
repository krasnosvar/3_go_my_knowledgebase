// Упражнение 02: буферизованный канал (Решение).
package main

import "fmt"

func fillBuffered(vals []int) chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	return ch
}

func main() {
	fmt.Println("--- Проверка 02_buffered (Решение) ---")

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

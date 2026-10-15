// Упражнение 06: защита низкого приоритета от голодания (Решение).
package main

import "fmt"

func fairMerge(high, low <-chan string, maxStreak int) []string {
	tryRecv := func(ch <-chan string) (string, bool) {
		select {
		case v := <-ch:
			return v, true
		default:
			return "", false
		}
	}
	var out []string
	streak := 0
	for {
		if streak >= maxStreak {
			if v, ok := tryRecv(low); ok {
				out = append(out, v)
				streak = 0
				continue
			}
		}
		if v, ok := tryRecv(high); ok {
			out = append(out, v)
			streak++
			continue
		}
		if v, ok := tryRecv(low); ok {
			out = append(out, v)
			streak = 0
			continue
		}
		return out
	}
}

func main() {
	fmt.Println("--- Проверка 06_starvation (Решение) ---")

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

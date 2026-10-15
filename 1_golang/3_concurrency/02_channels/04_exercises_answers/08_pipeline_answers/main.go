// Упражнение 08: pipeline из двух стадий (Решение).
package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func multiply(in <-chan int, k int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * k
		}
	}()
	return out
}

func filterGreater(in <-chan int, min int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v > min {
				out <- v
			}
		}
	}()
	return out
}

// --- вспомогательные функции проверки (не меняйте) ---

// collect читает канал до закрытия; ok=false, если канал не закрылся за d.
func collect(ch <-chan int, d time.Duration) ([]int, bool) {
	res := []int{}
	deadline := time.After(d)
	for {
		select {
		case v, open := <-ch:
			if !open {
				return res, true
			}
			res = append(res, v)
		case <-deadline:
			return res, false
		}
	}
}

func main() {
	fmt.Println("--- Проверка 08_pipeline (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got, closed := collect(filterGreater(multiply(gen(1, 2, 3, 4, 5), 3), 6), 300*time.Millisecond)
	check(closed, "выходной канал конвейера не закрылся")
	check(fmt.Sprint(got) == "[9 12 15]", fmt.Sprintf("конвейер отдал %v, ожидалось [9 12 15]", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

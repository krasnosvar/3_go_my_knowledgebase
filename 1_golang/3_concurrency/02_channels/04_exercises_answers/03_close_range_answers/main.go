// Упражнение 03: закрытие канала и range (Решение).
package main

import (
	"fmt"
	"time"
)

func produce(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func sumAll(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
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
	fmt.Println("--- Проверка 03_close_range (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got, closed := collect(produce(4), 300*time.Millisecond)
	check(closed, "produce не закрыла канал")
	check(fmt.Sprint(got) == "[1 2 3 4]", fmt.Sprintf("produce(4) отдала %v, ожидалось [1 2 3 4]", got))

	ready := make(chan int, 3)
	ready <- 5
	ready <- 6
	ready <- 7
	close(ready)
	check(sumAll(ready) == 18, "sumAll([5 6 7]) должна вернуть 18")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

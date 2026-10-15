// Упражнение 07: горутина без утечки (Решение).
package main

import (
	"fmt"
	"time"
)

func generator(stop <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-stop:
				return
			}
		}
	}()
	return out
}

func main() {
	fmt.Println("--- Проверка 07_leak_prevention (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	stop := make(chan struct{})
	ch := generator(stop)

	first := time.After(300 * time.Millisecond)
	for want := 0; want < 3; want++ {
		select {
		case v := <-ch:
			check(v == want, fmt.Sprintf("получено %d, ожидалось %d", v, want))
		case <-first:
			check(false, fmt.Sprintf("значение %d не пришло", want))
			want = 3
		}
	}

	close(stop)
	closed := false
	deadline := time.After(300 * time.Millisecond)
	for !closed {
		select {
		case _, ok := <-ch:
			if !ok {
				closed = true
			}
		case <-deadline:
			check(false, "после закрытия stop выходной канал не закрылся — горутина утекла")
			closed = true
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

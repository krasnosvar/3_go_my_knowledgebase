// Упражнение 04: конвейер parse → square → sum (Решение).
package main

import (
	"fmt"
	"strconv"
	"time"
)

func parse(in <-chan string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for s := range in {
			n, err := strconv.Atoi(s)
			if err != nil {
				continue
			}
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func sum(in <-chan int) int {
	total := 0
	for v := range in {
		total += v
	}
	return total
}

func main() {
	fmt.Println("--- Проверка 04_pipeline (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	src := make(chan string, 5)
	for _, s := range []string{"1", "x", "2", "", "3"} {
		src <- s
	}
	close(src)

	res := make(chan int, 1)
	go func() {
		sq := square(parse(src))
		if sq == nil {
			res <- -1
			return
		}
		res <- sum(sq)
	}()
	select {
	case got := <-res:
		check(got == 14, fmt.Sprintf("сумма квадратов = %d, ожидалось 14 (1+4+9)", got))
	case <-time.After(300 * time.Millisecond):
		check(false, "конвейер завис: стадии должны закрывать свои выходы")
	}

	ready := make(chan int, 2)
	ready <- 5
	ready <- 6
	close(ready)
	check(sum(ready) == 11, "sum([5 6]) должна вернуть 11")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// Упражнение 04: чтение двух каналов до их закрытия (Решение).
package main

import (
	"fmt"
	"time"
)

func drainBoth(nums <-chan int, words <-chan string) (int, []string) {
	sum := 0
	var ws []string
	for nums != nil || words != nil {
		select {
		case n, ok := <-nums:
			if !ok {
				nums = nil
				continue
			}
			sum += n
		case w, ok := <-words:
			if !ok {
				words = nil
				continue
			}
			ws = append(ws, w)
		}
	}
	return sum, ws
}

func main() {
	fmt.Println("--- Проверка 04_two_channels (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	nums := make(chan int)
	words := make(chan string)
	go func() {
		for i := 1; i <= 4; i++ {
			nums <- i
		}
		close(nums)
	}()
	go func() {
		for _, w := range []string{"go", "select", "nil"} {
			words <- w
		}
		close(words)
	}()

	type result struct {
		sum   int
		words []string
	}
	res := make(chan result, 1)
	go func() {
		s, w := drainBoth(nums, words)
		res <- result{s, w}
	}()

	select {
	case r := <-res:
		check(r.sum == 10, fmt.Sprintf("сумма = %d, ожидалось 10", r.sum))
		check(fmt.Sprint(r.words) == "[go select nil]", fmt.Sprintf("слова = %v, ожидалось [go select nil]", r.words))
	case <-time.After(300 * time.Millisecond):
		check(false, "drainBoth зависла: закрытый канал нужно заменять на nil")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

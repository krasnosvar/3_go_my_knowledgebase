// Упражнение 03: владение состоянием через канал (Решение).
package main

import (
	"fmt"
	"sync"
	"time"
)

func owner(ops <-chan int, result chan<- int) {
	total := 0
	for v := range ops {
		total += v
	}
	result <- total
}

func main() {
	fmt.Println("--- Проверка 03_fix_with_channels (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ops := make(chan int, 1000)
	result := make(chan int, 1)
	go owner(ops, result)

	var wg sync.WaitGroup
	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 10; i++ {
				ops <- i
			}
		}()
	}
	wg.Wait()
	close(ops)

	select {
	case total := <-result:
		check(total == 550, fmt.Sprintf("итог = %d, ожидалось 550", total))
	case <-time.After(300 * time.Millisecond):
		check(false, "owner не отправил итог после закрытия ops")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

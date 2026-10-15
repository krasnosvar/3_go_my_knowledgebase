// Упражнение 06: graceful stop воркера (Решение).
package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context, jobs <-chan int, handle func(int)) int {
	n := 0
	for {
		// select выбирает случайно среди готовых веток, поэтому отмену
		// проверяем явно до попытки взять новую задачу.
		if ctx.Err() != nil {
			return n
		}
		select {
		case <-ctx.Done():
			return n
		case j, ok := <-jobs:
			if !ok {
				return n
			}
			handle(j)
			n++
		}
	}
}

func main() {
	fmt.Println("--- Проверка 06_graceful_stop (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	// 1) закрытый канал: обработать всё и выйти
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)
	check(run(context.Background(), jobs, func(int) {}) == 3, "закрытый канал → обработаны все 3")

	// 2) отмена во время задачи: задача доделывается, новые не берутся
	ctx, cancel := context.WithCancel(context.Background())
	jobs2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		jobs2 <- i
	}
	finished := 0
	done := make(chan int, 1)
	go func() {
		done <- run(ctx, jobs2, func(j int) {
			if j == 1 {
				cancel()
				time.Sleep(20 * time.Millisecond) // долгая задача после отмены
			}
			finished++
		})
	}()
	select {
	case n := <-done:
		check(n == 2 && finished == 2, fmt.Sprintf("после отмены задача 1 доделана, новые не взяты (n=%d)", n))
	case <-time.After(2 * time.Second):
		check(false, "воркер не остановился")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

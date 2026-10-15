// Упражнение 06: простой worker поверх канала задач (Решение).
package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
	}
}

func main() {
	fmt.Println("--- Проверка 06_worker_basic (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	done := make(chan struct{})
	go func() {
		worker(jobs, results)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
		check(false, "worker не завершился после закрытия jobs")
	}

	got := []int{}
	for len(results) > 0 {
		got = append(got, <-results)
	}
	check(fmt.Sprint(got) == "[1 4 9 16 25]", fmt.Sprintf("results = %v, ожидалось [1 4 9 16 25]", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

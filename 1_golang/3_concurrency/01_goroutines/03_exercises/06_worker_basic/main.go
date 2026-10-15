// Упражнение 06: простой worker поверх канала задач.
//
// Задача: реализуйте worker(jobs, results): он читает числа из jobs, пока
// канал не закрыт (for j := range jobs), и для каждого отправляет j*j
// в results. Когда jobs закрыт и вычитан, worker возвращается.
//
// Запуск: go run ./03_exercises/06_worker_basic/main.go
package main

import (
	"fmt"
	"time"
)

// worker читает jobs до закрытия и отправляет квадраты в results.
func worker(jobs <-chan int, results chan<- int) {
	// TODO: for j := range jobs { results <- j * j }
}

func main() {
	fmt.Println("--- Проверка 06_worker_basic ---")

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

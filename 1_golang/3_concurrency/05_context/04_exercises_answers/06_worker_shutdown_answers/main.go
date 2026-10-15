// Упражнение 06: остановка worker'а по ctx.Done() (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

func worker(ctx context.Context, jobs <-chan int, processed *atomic.Int64) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case _, ok := <-jobs:
			if !ok {
				return nil
			}
			processed.Add(1)
		}
	}
}

func main() {
	fmt.Println("--- Проверка 06_worker_shutdown (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)
	var processed atomic.Int64
	res := make(chan error, 1)
	go func() { res <- worker(ctx, jobs, &processed) }()

	for i := 0; i < 3; i++ {
		select {
		case jobs <- i:
		case <-time.After(100 * time.Millisecond):
			check(false, fmt.Sprintf("worker не принял задание %d", i))
		}
	}
	cancel()

	select {
	case err := <-res:
		check(errors.Is(err, context.Canceled), fmt.Sprintf("worker вернул %v, ожидалось context.Canceled", err))
	case <-time.After(300 * time.Millisecond):
		check(false, "worker не остановился после cancel")
	}
	check(processed.Load() == 3, fmt.Sprintf("обработано %d, ожидалось 3", processed.Load()))

	jobs2 := make(chan int)
	close(jobs2)
	err := worker(context.Background(), jobs2, &processed)
	check(err == nil, fmt.Sprintf("при закрытом jobs ожидался nil, получено %v", err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

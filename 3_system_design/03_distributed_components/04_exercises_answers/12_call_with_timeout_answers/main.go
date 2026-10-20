// Упражнение 12: обернуть вызов таймаутом через context (Решение).
package main

import (
	"context"
	"fmt"
	"time"
)

func callWithTimeout(ctx context.Context, work func() int) (int, error) {
	resCh := make(chan int, 1) // буфер 1: горутина не залипнет, даже если мы ушли по таймауту
	go func() {
		resCh <- work()
	}()

	select {
	case r := <-resCh:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {
	fmt.Println("--- Проверка 12_call_with_timeout (Решение) ---")

	allOK := true

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Hour)
	defer cancel1()
	got, err := callWithTimeout(ctx1, func() int { return 42 })
	if err != nil || got != 42 {
		allOK = false
		fmt.Printf("FAIL: быстрая work -> (%d, %v), ожидалось (42, nil)\n", got, err)
	}

	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2()
	never := func() int {
		<-make(chan struct{})
		return 1
	}
	got, err = callWithTimeout(ctx2, never)
	if err == nil {
		allOK = false
		fmt.Printf("FAIL: медленная work -> (%d, nil), ожидалась ошибка таймаута\n", got)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

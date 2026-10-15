// Упражнение 03: владение состоянием через канал.
//
// Задача: вместо мьютекса — «share memory by communicating»: состоянием
// владеет ОДНА горутина, остальные только шлют ей сообщения.
//
// owner(ops, result) читает числа из ops и прибавляет их к локальной
// сумме; когда ops закрыт — отправляет итог в result и завершается.
// Поскольку сумму трогает только owner, гонки нет без всяких блокировок.
//
// Запуск: go run ./03_exercises/03_fix_with_channels/main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// owner единолично владеет суммой и отдаёт итог после закрытия ops.
func owner(ops <-chan int, result chan<- int) {
	// TODO: total := 0; for v := range ops { total += v }; result <- total
}

func main() {
	fmt.Println("--- Проверка 03_fix_with_channels ---")

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

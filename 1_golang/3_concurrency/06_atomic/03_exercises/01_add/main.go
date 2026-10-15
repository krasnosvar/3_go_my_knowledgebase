// Упражнение 01: атомарный инкремент.
//
// Задача: countConcurrently(n) запускает n горутин, каждая 100 раз
// увеличивает общий int64-счётчик через atomic.AddInt64(&counter, 1).
// Дождитесь всех (WaitGroup) и верните atomic.LoadInt64(&counter).
//
// Обычный counter++ здесь — гонка; atomic.Add выполняет
// «прочитать-прибавить-записать» как одну неделимую операцию.
//
// Запуск: go run ./03_exercises/01_add/main.go
package main

import "fmt"

// countConcurrently возвращает n*100, считая атомарно из n горутин.
func countConcurrently(n int) int64 {
	// TODO: var counter int64; n горутин с atomic.AddInt64; wg.Wait().
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_add ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := countConcurrently(50)
	check(got == 5000, fmt.Sprintf("countConcurrently(50) = %d, ожидалось 5000", got))
	check(countConcurrently(0) == 0, "для 0 горутин ожидался 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

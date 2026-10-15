// Упражнение 03: общий счётчик без гонки данных.
//
// Задача: countParallel(workers, perWorker) запускает workers горутин,
// каждая perWorker раз увеличивает общий счётчик; функция возвращает итог.
//
// Наивное решение `counter++` из нескольких горутин — это data race:
// результат будет меньше ожидаемого, а `go run -race` выдаст предупреждение.
// Защитите счётчик через sync.Mutex (или sync/atomic) и дождитесь горутин.
//
// Проверьте себя: go run -race ./03_exercises/03_race_awareness/main.go
//
// Запуск: go run ./03_exercises/03_race_awareness/main.go
package main

import "fmt"

// countParallel возвращает workers*perWorker, считая конкурентно и без гонки.
func countParallel(workers, perWorker int) int {
	// TODO: запустите workers горутин, защитите общий счётчик мьютексом,
	// дождитесь завершения через WaitGroup и верните результат.
	return 0
}

func main() {
	fmt.Println("--- Проверка 03_race_awareness ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := countParallel(50, 1000)
	check(got == 50000, fmt.Sprintf("countParallel(50, 1000) = %d, ожидалось 50000", got))

	got = countParallel(1, 10)
	check(got == 10, fmt.Sprintf("countParallel(1, 10) = %d, ожидалось 10", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

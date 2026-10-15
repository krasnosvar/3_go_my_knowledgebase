// Упражнение 01: найти и устранить data race.
//
// Задача: ниже в комментарии — счётчик с гонкой данных. Сначала
// воспроизведите проблему: скопируйте racyCount в отдельный файл и
// запустите его с флагом -race — детектор напечатает "WARNING: DATA RACE"
// с двумя стеками (запись и конкурентное чтение/запись).
//
//	func racyCount(n int) int {
//	    count := 0
//	    var wg sync.WaitGroup
//	    for i := 0; i < n; i++ {
//	        wg.Add(1)
//	        go func() { defer wg.Done(); count++ }() // гонка!
//	    }
//	    wg.Wait()
//	    return count
//	}
//
// Затем реализуйте safeCount(n) — то же самое, но без гонки, используя
// sync/atomic (atomic.Int64). Проверка: результат точен, а
// go run -race ./03_exercises/01_detect_race/main.go не выдаёт предупреждений.
//
// Запуск: go run ./03_exercises/01_detect_race/main.go
package main

import "fmt"

// safeCount запускает n горутин, каждая увеличивает счётчик на 1.
func safeCount(n int) int {
	// TODO: var count atomic.Int64; n горутин с count.Add(1); wg.Wait().
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_detect_race ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, n := range []int{1, 100, 1000} {
		got := safeCount(n)
		check(got == n, fmt.Sprintf("safeCount(%d) = %d", n, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

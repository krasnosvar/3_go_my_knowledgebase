// Упражнение 02: пул воркеров с сохранением порядка результатов.
//
// Задача: реализуйте process(inputs []int, workers int, fn func(int) int) []int:
//   - запустить ровно workers горутин, читающих индексы из канала задач;
//   - результат fn(inputs[i]) записывать в out[i] (каждая горутина пишет
//     в свой индекс — гонки нет);
//   - дождаться всех через sync.WaitGroup и вернуть out.
//
// Тест проверяет, что одновременно работало не больше workers горутин.
//
// Запуск: go run ./03_exercises/02_worker_pool
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var _ sync.WaitGroup

func process(inputs []int, workers int, fn func(int) int) []int {
	out := make([]int, len(inputs))
	// TODO: канал индексов, workers горутин, WaitGroup.
	return out
}

func main() {
	fmt.Println("--- Проверка 02_worker_pool ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var active, peak atomic.Int32
	square := func(x int) int {
		n := active.Add(1)
		for {
			p := peak.Load()
			if n <= p || peak.CompareAndSwap(p, n) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		active.Add(-1)
		return x * x
	}

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := process(in, 3, square)
	check(fmt.Sprint(got) == "[1 4 9 16 25 36 49 64 81 100]", "результаты по порядку: "+fmt.Sprint(got))
	check(peak.Load() <= 3, fmt.Sprintf("не больше 3 воркеров одновременно (пик %d)", peak.Load()))
	check(peak.Load() >= 2, "работа действительно шла параллельно")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// Упражнение 06: статистика запросов на атомиках.
//
// Задача: RequestStats считает запросы конкурентно без мьютекса:
//   - Record(ok) увеличивает total и, если ok == false, ещё errors;
//   - Snapshot() возвращает (total, errors).
//
// Каждое поле атомарно само по себе. Учтите: Snapshot из двух Load не
// является согласованным «снимком» — между чтениями значения могут
// измениться. Для метрик это обычно допустимо.
//
// Запуск: go run ./03_exercises/06_counter/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// RequestStats — счётчики запросов.
type RequestStats struct {
	total  atomic.Int64
	errors atomic.Int64
}

// Record учитывает один запрос.
func (s *RequestStats) Record(ok bool) {
	// TODO
}

// Snapshot возвращает текущие значения счётчиков.
func (s *RequestStats) Snapshot() (total, errs int64) {
	// TODO
	return 0, 0
}

func main() {
	fmt.Println("--- Проверка 06_counter ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var st RequestStats
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			st.Record(i%4 != 0) // каждый 4-й — ошибка
		}(i)
	}
	wg.Wait()
	total, errs := st.Snapshot()
	check(total == 100, fmt.Sprintf("total = %d, ожидалось 100", total))
	check(errs == 25, fmt.Sprintf("errors = %d, ожидалось 25", errs))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

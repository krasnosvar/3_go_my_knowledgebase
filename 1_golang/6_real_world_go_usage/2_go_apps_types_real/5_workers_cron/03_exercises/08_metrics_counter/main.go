// Упражнение 08: метрики воркера на atomic.
//
// Задача:
//   - (m *metrics) Observe(err error, dur time.Duration): увеличить processed;
//     при err != nil увеличить failed; прибавить dur к totalNanos;
//   - (m *metrics) Snapshot() snapshot: processed, failed и средняя
//     длительность (0, если processed == 0).
//
// Все поля — atomic.Int64, Observe вызывается из многих горутин.
//
// Запуск: go run ./03_exercises/08_metrics_counter
package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type metrics struct {
	processed, failed, totalNanos atomic.Int64
}

type snapshot struct {
	Processed, Failed int64
	Avg               time.Duration
}

func (m *metrics) Observe(err error, dur time.Duration) {
	// TODO: три атомарных Add.
}

func (m *metrics) Snapshot() snapshot {
	// TODO: Load полей, средняя длительность с защитой от деления на 0.
	return snapshot{}
}

func main() {
	fmt.Println("--- Проверка 08_metrics_counter ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var m metrics
	check(m.Snapshot().Avg == 0, "пустые метрики → Avg=0 без паники")

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var err error
			if i%4 == 0 {
				err = errors.New("fail")
			}
			m.Observe(err, 10*time.Millisecond)
		}(i)
	}
	wg.Wait()

	s := m.Snapshot()
	check(s.Processed == 100, "processed = 100")
	check(s.Failed == 25, "failed = 25")
	check(s.Avg == 10*time.Millisecond, "средняя длительность 10ms")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

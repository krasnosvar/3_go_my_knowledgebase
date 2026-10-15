// Упражнение 08: метрики воркера на atomic (Решение).
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
	m.processed.Add(1)
	if err != nil {
		m.failed.Add(1)
	}
	m.totalNanos.Add(int64(dur))
}

func (m *metrics) Snapshot() snapshot {
	p := m.processed.Load()
	s := snapshot{Processed: p, Failed: m.failed.Load()}
	if p > 0 {
		s.Avg = time.Duration(m.totalNanos.Load() / p)
	}
	return s
}

func main() {
	fmt.Println("--- Проверка 08_metrics_counter (Решение) ---")

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

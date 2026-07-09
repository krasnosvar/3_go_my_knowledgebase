// Упражнение 03: подсчёт hit ratio (Решение).
package main

import (
	"fmt"
	"math"
)

var _ = math.Abs

type stats struct {
	present map[string]bool
	hits    int
	misses  int
}

func newStats() *stats {
	return &stats{present: make(map[string]bool)}
}

func (s *stats) Access(key string) bool {
	if s.present[key] {
		s.hits++
		return true
	}
	s.misses++
	s.present[key] = true
	return false
}

func (s *stats) HitRatio() float64 {
	total := s.hits + s.misses
	if total == 0 {
		return 0
	}
	return float64(s.hits) / float64(total)
}

func main() {
	fmt.Println("--- Проверка 03_hit_ratio (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := newStats()

	seq := []string{"a", "b", "a", "a", "c", "b"}
	for _, k := range seq {
		s.Access(k)
	}

	check(s.hits == 3, fmt.Sprintf("ожидалось hits=3, получено %d", s.hits))
	check(s.misses == 3, fmt.Sprintf("ожидалось misses=3, получено %d", s.misses))

	if math.Abs(s.HitRatio()-0.5) > 1e-9 {
		check(false, fmt.Sprintf("ожидался hit ratio 0.5, получено %v", s.HitRatio()))
	}

	empty := newStats()
	if empty.HitRatio() != 0 {
		check(false, "при нуле обращений hit ratio должен быть 0")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

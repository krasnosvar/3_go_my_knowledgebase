// Упражнение 03: подсчёт hit ratio.
//
// Задача: реализуйте кэш-обёртку, которая на каждый запрос ключа считает
// попадания и промахи, и метод HitRatio(), возвращающий долю попаданий:
//
//	hit ratio = hits / (hits + misses)
//
// Правила:
//   - Access(key) ищет ключ в наборе уже загруженных (present);
//     если есть — это hit; если нет — miss, после чего ключ считается
//     загруженным (добавляется в present) и последующие обращения к нему
//     будут hit;
//   - HitRatio() при нуле обращений возвращает 0.
//
// Запуск: go run ./03_exercises/03_hit_ratio/main.go
package main

import (
	"fmt"
	"math"
)

// Держим импорт math живым в заготовке (в решении удобен для сравнения float).
var _ = math.Abs

type stats struct {
	present map[string]bool
	hits    int
	misses  int
}

func newStats() *stats {
	return &stats{present: make(map[string]bool)}
}

// Access обращается к ключу, обновляя счётчики hits/misses.
// Возвращает true, если это было попадание (ключ уже был загружен).
func (s *stats) Access(key string) bool {
	// TODO:
	//  - если key в present -> hits++, вернуть true;
	//  - иначе -> misses++, добавить key в present, вернуть false.
	return false
}

// HitRatio возвращает hits / (hits + misses); при нуле обращений — 0.
func (s *stats) HitRatio() float64 {
	// TODO: посчитайте долю попаданий, аккуратно обработав деление на ноль.
	return 0
}

func main() {
	fmt.Println("--- Проверка 03_hit_ratio ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := newStats()

	// Последовательность обращений: a(miss) b(miss) a(hit) a(hit) c(miss)
	// b(hit)
	// hits=3, misses=3 -> ratio=0.5
	seq := []string{"a", "b", "a", "a", "c", "b"}
	for _, k := range seq {
		s.Access(k)
	}

	check(s.hits == 3, fmt.Sprintf("ожидалось hits=3, получено %d", s.hits))
	check(s.misses == 3, fmt.Sprintf("ожидалось misses=3, получено %d", s.misses))

	if math.Abs(s.HitRatio()-0.5) > 1e-9 {
		check(false, fmt.Sprintf("ожидался hit ratio 0.5, получено %v", s.HitRatio()))
	}

	// Пустая статистика -> 0, без паники.
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

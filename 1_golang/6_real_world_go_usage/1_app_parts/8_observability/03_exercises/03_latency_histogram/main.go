// Упражнение 03: гистограмма латентности (как у Prometheus).
//
// Задача: реализуйте (h *histogram) Observe(ms float64):
//   - buckets — верхние границы по возрастанию (le), counts[i] — число
//     наблюдений <= buckets[i] (кумулятивно, как в Prometheus);
//   - также увеличить h.count и h.sum.
//
// И (h *histogram) Quantile50() float64 — граница первого бакета, в котором
// накопилось >= половины наблюдений (грубая оценка медианы); 0 при count == 0.
//
// Запуск: go run ./03_exercises/03_latency_histogram
package main

import "fmt"

type histogram struct {
	buckets []float64
	counts  []int
	count   int
	sum     float64
}

func newHistogram(b ...float64) *histogram {
	return &histogram{buckets: b, counts: make([]int, len(b))}
}

func (h *histogram) Observe(ms float64) {
	// TODO: увеличьте все бакеты с le >= ms, а также count и sum.
}

func (h *histogram) Quantile50() float64 {
	// TODO: найдите первый бакет с counts[i]*2 >= count.
	return 0
}

func main() {
	fmt.Println("--- Проверка 03_latency_histogram ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	h := newHistogram(10, 50, 100, 500)
	check(h.Quantile50() == 0, "пустая гистограмма → 0")
	for _, v := range []float64{5, 8, 30, 45, 70, 300} {
		h.Observe(v)
	}
	check(fmt.Sprint(h.counts) == "[2 4 5 6]", "кумулятивные бакеты [2 4 5 6], получено "+fmt.Sprint(h.counts))
	check(h.count == 6 && h.sum == 458, "count=6, sum=458")
	check(h.Quantile50() == 50, "медиана попадает в бакет le=50")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

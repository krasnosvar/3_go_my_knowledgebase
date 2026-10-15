// Упражнение 03: гистограмма латентности (Решение).
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
	for i, le := range h.buckets {
		if ms <= le {
			h.counts[i]++
		}
	}
	h.count++
	h.sum += ms
}

func (h *histogram) Quantile50() float64 {
	if h.count == 0 {
		return 0
	}
	for i, c := range h.counts {
		if c*2 >= h.count {
			return h.buckets[i]
		}
	}
	return h.buckets[len(h.buckets)-1]
}

func main() {
	fmt.Println("--- Проверка 03_latency_histogram (Решение) ---")

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

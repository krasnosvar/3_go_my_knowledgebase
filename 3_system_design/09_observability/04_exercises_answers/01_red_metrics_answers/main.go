// Упражнение 01: RED-метрики (Решение).
package main

import "fmt"

type request struct {
	Status   int
	Duration int
}

type red struct {
	Rate      float64
	ErrorRate float64
	AvgMillis float64
}

func computeRED(reqs []request, windowSeconds float64) red {
	n := len(reqs)
	if n == 0 {
		return red{}
	}
	var errs, sum int
	for _, r := range reqs {
		if r.Status >= 500 {
			errs++
		}
		sum += r.Duration
	}
	return red{
		Rate:      float64(n) / windowSeconds,
		ErrorRate: float64(errs) / float64(n),
		AvgMillis: float64(sum) / float64(n),
	}
}

func almostEqual(a, b float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d < 1e-9
}

func main() {
	fmt.Println("--- Проверка 01_red_metrics (Решение) ---")

	reqs := []request{
		{Status: 200, Duration: 100},
		{Status: 200, Duration: 300},
		{Status: 500, Duration: 200},
		{Status: 503, Duration: 400},
	}
	got := computeRED(reqs, 2.0)
	want := red{Rate: 2.0, ErrorRate: 0.5, AvgMillis: 250.0}

	allOK := true
	if !almostEqual(got.Rate, want.Rate) {
		allOK = false
		fmt.Printf("FAIL: Rate=%v, ожидалось %v\n", got.Rate, want.Rate)
	}
	if !almostEqual(got.ErrorRate, want.ErrorRate) {
		allOK = false
		fmt.Printf("FAIL: ErrorRate=%v, ожидалось %v\n", got.ErrorRate, want.ErrorRate)
	}
	if !almostEqual(got.AvgMillis, want.AvgMillis) {
		allOK = false
		fmt.Printf("FAIL: AvgMillis=%v, ожидалось %v\n", got.AvgMillis, want.AvgMillis)
	}

	empty := computeRED(nil, 5.0)
	if empty != (red{}) {
		allOK = false
		fmt.Printf("FAIL: пустой ввод -> %+v, ожидалось нули\n", empty)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

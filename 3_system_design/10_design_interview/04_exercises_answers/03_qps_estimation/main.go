// Упражнение 03: оценка QPS (Решение).
package main

import "fmt"

const secondsPerDay = 86_400

func qps(dau, reqPerUser int, peakFactor float64) (avg, peak int) {
	perDay := float64(dau) * float64(reqPerUser)
	avgF := perDay / secondsPerDay
	return int(avgF), int(avgF * peakFactor)
}

func main() {
	fmt.Println("--- Проверка 03_qps_estimation (Решение) ---")

	cases := []struct {
		dau        int
		reqPerUser int
		peakFactor float64
		wantAvg    int
		wantPeak   int
	}{
		{10_000_000, 20, 2.0, 2314, 4629},
		{100_000_000, 5, 3.0, 5787, 17361},
		{86_400, 1, 2.0, 1, 2},
	}

	allOK := true
	for _, c := range cases {
		avg, peak := qps(c.dau, c.reqPerUser, c.peakFactor)
		if avg != c.wantAvg || peak != c.wantPeak {
			allOK = false
			fmt.Printf("FAIL: qps(%d, %d, %.1f) = (%d, %d), ожидалось (%d, %d)\n",
				c.dau, c.reqPerUser, c.peakFactor, avg, peak, c.wantAvg, c.wantPeak)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

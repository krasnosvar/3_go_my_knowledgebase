// Упражнение 03: оценка QPS.
//
// Задача: посчитайте средний и пиковый QPS по формуле
//
//	средний QPS = (DAU × запросов_на_пользователя) / 86400
//	пиковый QPS = средний × peakFactor
//
// (см. 01_theory/02_estimations.md).
//
// Округление: результат целочисленный (float64 -> int усечением к нулю).
//
// Запуск: go run ./03_exercises/03_qps_estimation/main.go
package main

import "fmt"

const secondsPerDay = 86_400

// qps возвращает средний и пиковый QPS (целые, усечение дробной части).
// dau — дневная аудитория, reqPerUser — запросов на пользователя в сутки,
// peakFactor — во сколько раз пик выше среднего.
func qps(dau, reqPerUser int, peakFactor float64) (avg, peak int) {
	// TODO: посчитайте запросы в сутки, поделите на secondsPerDay,
	// умножьте средний на peakFactor. Верните как int (усечение).
	return 0, 0
}

func main() {
	fmt.Println("--- Проверка 03_qps_estimation ---")

	cases := []struct {
		dau        int
		reqPerUser int
		peakFactor float64
		wantAvg    int
		wantPeak   int
	}{
		{10_000_000, 20, 2.0, 2314, 4629},  // 2e8/86400 ≈ 2314
		{100_000_000, 5, 3.0, 5787, 17361}, // 5e8/86400 ≈ 5787
		{86_400, 1, 2.0, 1, 2},             // ровно 1 QPS
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

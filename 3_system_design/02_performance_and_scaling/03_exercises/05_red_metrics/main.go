// Упражнение 05: RED-метрики (Rate, Errors, Duration).
//
// Задача: по списку записей о запросах посчитать RED-показатели за
// известный интервал наблюдения:
//   - Rate     — запросов в секунду = число запросов / windowSeconds;
//   - ErrorRate — доля ошибок (статус >= 500) от всех запросов, 0..1;
//   - AvgMillis — средняя длительность в миллисекундах.
//
// Реализуйте функцию computeRED. Если запросов нет — вернуть нули
// (без деления на ноль).
//
// Запуск: go run ./03_exercises/05_red_metrics/main.go
package main

import "fmt"

type request struct {
	Status   int // HTTP-статус
	Duration int // длительность в миллисекундах
}

type red struct {
	Rate      float64 // запросов в секунду
	ErrorRate float64 // доля ошибок 0..1
	AvgMillis float64 // средняя длительность, мс
}

// computeRED считает RED по запросам за окно windowSeconds секунд.
func computeRED(reqs []request, windowSeconds float64) red {
	// TODO:
	//  1) если len(reqs)==0 -> вернуть red{} (все нули);
	//  2) Rate = float64(len(reqs)) / windowSeconds;
	//  3) ErrorRate = (число статусов >= 500) / len(reqs);
	//  4) AvgMillis = сумма Duration / len(reqs).
	return red{}
}

func almostEqual(a, b float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d < 1e-9
}

func main() {
	fmt.Println("--- Проверка 05_red_metrics ---")

	reqs := []request{
		{Status: 200, Duration: 100},
		{Status: 200, Duration: 300},
		{Status: 500, Duration: 200},
		{Status: 503, Duration: 400},
	}
	// 4 запроса за 2 секунды -> Rate=2; 2 ошибки из 4 -> 0.5;
	// средняя (100+300+200+400)/4 = 250.
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

	// Пустой ввод -> все нули, без паники.
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

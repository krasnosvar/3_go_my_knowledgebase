// Упражнение 05: fixed-window rate limiter (Хардкор).
//
// Задача: реализуйте Allow(now) у fixed-window лимитера. Лимит — limit
// запросов на окно длиной windowSize (в условных единицах времени).
// Время НЕ реальное: его передают параметром now, поэтому тест
// детерминирован.
//
// Логика:
//   - окно вычисляется как windowStart = now - (now % windowSize);
//   - при входе в НОВОЕ окно (windowStart изменился) счётчик сбрасывается в 0;
//   - если счётчик в текущем окне < limit: увеличить счётчик, вернуть true;
//   - иначе вернуть false (лимит окна исчерпан).
//
// Запуск: go run ./03_exercises/05_fixed_window_limiter/main.go
package main

import "fmt"

type FixedWindow struct {
	limit      int
	windowSize int
	curWindow  int // начало текущего окна
	count      int // запросов в текущем окне
}

func NewFixedWindow(limit, windowSize int) *FixedWindow {
	return &FixedWindow{limit: limit, windowSize: windowSize, curWindow: -1}
}

// Allow решает, пропустить ли запрос в момент now.
func (w *FixedWindow) Allow(now int) bool {
	// TODO: вычислить начало окна, сбросить счётчик при смене окна,
	// затем сравнить count с limit.
	return false
}

var _ = NewFixedWindow

func main() {
	fmt.Println("--- Проверка 05_fixed_window_limiter ---")

	// limit=2 на окно размером 10.
	w := NewFixedWindow(2, 10)

	type step struct {
		now  int
		want bool
	}
	steps := []step{
		{0, true},   // окно [0,10): 1-й
		{3, true},   // окно [0,10): 2-й
		{5, false},  // окно [0,10): лимит исчерпан
		{9, false},  // всё ещё [0,10)
		{10, true},  // новое окно [10,20): счётчик сброшен, 1-й
		{11, true},  // [10,20): 2-й
		{12, false}, // [10,20): исчерпан
		{20, true},  // окно [20,30): 1-й
	}

	allOK := true
	for i, s := range steps {
		got := w.Allow(s.now)
		if got != s.want {
			allOK = false
			fmt.Printf("FAIL: шаг %d Allow(now=%d) = %v, ожидалось %v\n", i, s.now, got, s.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

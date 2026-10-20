// Упражнение 13: fixed-window rate limiter (Решение).
package main

import "fmt"

type FixedWindow struct {
	limit      int
	windowSize int
	curWindow  int
	count      int
}

func NewFixedWindow(limit, windowSize int) *FixedWindow {
	return &FixedWindow{limit: limit, windowSize: windowSize, curWindow: -1}
}

func (w *FixedWindow) Allow(now int) bool {
	windowStart := now - (now % w.windowSize)
	if windowStart != w.curWindow {
		w.curWindow = windowStart
		w.count = 0
	}
	if w.count < w.limit {
		w.count++
		return true
	}
	return false
}

var _ = NewFixedWindow

func main() {
	fmt.Println("--- Проверка 13_fixed_window_limiter (Решение) ---")

	w := NewFixedWindow(2, 10)

	type step struct {
		now  int
		want bool
	}
	steps := []step{
		{0, true},
		{3, true},
		{5, false},
		{9, false},
		{10, true},
		{11, true},
		{12, false},
		{20, true},
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

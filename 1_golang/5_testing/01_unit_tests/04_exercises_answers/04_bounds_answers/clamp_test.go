// Упражнение 04: граничные значения (Решение).
package clamp

import "testing"

func TestClamp(t *testing.T) {
	const lo, hi = 0, 10
	checks := [][2]int{ // {вход, ожидание}
		{-1, 0}, {0, 0}, {1, 1},
		{5, 5},
		{9, 9}, {10, 10}, {11, 10},
	}
	for _, c := range checks {
		if got := Clamp(c[0], lo, hi); got != c[1] {
			t.Errorf("Clamp(%d, %d, %d) = %d, want %d", c[0], lo, hi, got, c[1])
		}
	}
}

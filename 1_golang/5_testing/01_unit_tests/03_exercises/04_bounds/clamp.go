// Package clamp — код под тестом для упражнения 04.
package clamp

// Clamp ограничивает v диапазоном [lo, hi].
func Clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

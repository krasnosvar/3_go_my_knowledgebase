// Package fib — код под тестом для упражнения 01.
package fib

// Fib возвращает n-е число Фибоначчи (итеративно, Fib(0) = 0).
func Fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

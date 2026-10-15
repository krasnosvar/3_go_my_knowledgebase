// Package prealloc — код для упражнения 03 (Решение).
package prealloc

// Squares возвращает [0*0, 1*1, ..., (n-1)*(n-1)].
func Squares(n int) []int {
	out := make([]int, 0, n) // одна аллокация сразу на n элементов
	for i := 0; i < n; i++ {
		out = append(out, i*i)
	}
	return out
}

// Package prealloc — код для упражнения 03.
package prealloc

// Squares возвращает [0*0, 1*1, ..., (n-1)*(n-1)].
func Squares(n int) []int {
	// TODO: выделите ёмкость заранее.
	var out []int
	for i := 0; i < n; i++ {
		out = append(out, i*i)
	}
	return out
}

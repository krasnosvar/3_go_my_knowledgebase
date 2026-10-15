// Package sum — пример: под-бенчмарки по размерам входа.
package sum

// Sum складывает элементы слайса.
func Sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

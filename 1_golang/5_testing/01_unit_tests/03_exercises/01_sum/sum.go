// Package sum — код под тестом для упражнения 01.
package sum

// Sum возвращает сумму элементов nums. Для пустого слайса — 0.
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

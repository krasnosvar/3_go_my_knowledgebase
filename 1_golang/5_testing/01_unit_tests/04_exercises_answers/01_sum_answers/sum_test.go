// Упражнение 01: первый unit-тест (Решение).
package sum

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, want 0", got)
	}
	if got := Sum([]int{7}); got != 7 {
		t.Errorf("Sum([7]) = %d, want 7", got)
	}
	nums := []int{1, 2, -5, 10}
	if got := Sum(nums); got != 8 {
		t.Errorf("Sum(%v) = %d, want 8", nums, got)
	}
}

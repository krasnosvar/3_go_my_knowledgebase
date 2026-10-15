package square

import "testing"

func TestSquare(t *testing.T) {
	t.Parallel()
	for _, n := range []int{-3, 0, 2, 10} {
		t.Run("n", func(t *testing.T) { // одинаковые имена Go дополнит суффиксами #01, #02...
			t.Parallel() // с Go 1.22 n своя на каждой итерации
			if got := Square(n); got != n*n || got < 0 {
				t.Errorf("Square(%d) = %d", n, got)
			}
		})
	}
}

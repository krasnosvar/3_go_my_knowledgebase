// Упражнение 04: медленные тесты и -short (Решение).
package primes

import "testing"

func TestCountPrimesSmall(t *testing.T) {
	cases := map[int]int{0: 0, 2: 0, 3: 1, 10: 4}
	for n, want := range cases {
		if got := CountPrimes(n); got != want {
			t.Errorf("CountPrimes(%d) = %d, want %d", n, got, want)
		}
	}
}

func TestCountPrimesLarge(t *testing.T) {
	if testing.Short() {
		t.Skip("медленный тест: запускайте без -short")
	}
	if got := CountPrimes(50_000_000); got != 3_001_134 {
		t.Errorf("CountPrimes(5e7) = %d, want 3001134", got)
	}
}

// Package primes — код под тестом для упражнения 04.
package primes

// CountPrimes возвращает количество простых чисел < n (решето Эратосфена).
func CountPrimes(n int) int {
	if n < 3 {
		return 0
	}
	composite := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if composite[i] {
			continue
		}
		count++
		for j := i * i; j < n; j += i {
			composite[j] = true
		}
	}
	return count
}

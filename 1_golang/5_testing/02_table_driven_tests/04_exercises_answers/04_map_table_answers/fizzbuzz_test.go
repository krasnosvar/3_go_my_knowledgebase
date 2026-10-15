// Упражнение 04: таблица-map и t.Run (Решение).
package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := map[string]struct {
		in   int
		want string
	}{
		"one":     {1, "1"},
		"three":   {3, "Fizz"},
		"five":    {5, "Buzz"},
		"fifteen": {15, "FizzBuzz"},
		"thirty":  {30, "FizzBuzz"},
		"seven":   {7, "7"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := FizzBuzz(tt.in); got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

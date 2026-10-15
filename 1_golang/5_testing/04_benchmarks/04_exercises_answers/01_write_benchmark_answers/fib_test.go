// Упражнение 01: первый бенчмарк (Решение).
package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	if got := Fib(10); got != 55 {
		t.Errorf("Fib(10) = %d, want 55", got)
	}
}

var sink int // не даёт компилятору выбросить вызов как неиспользуемый

func BenchmarkFib(b *testing.B) {
	for _, n := range []int{10, 20, 40} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sink = Fib(n)
			}
		})
	}
}

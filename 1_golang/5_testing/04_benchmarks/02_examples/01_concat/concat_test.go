package bench

import "testing"

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConcatPlus("go", "lang")
	}
}

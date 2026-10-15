package sum

import (
	"fmt"
	"testing"
)

var sink int

// go test -bench=Sum -benchmem -run='^$' .
func BenchmarkSum(b *testing.B) {
	for _, size := range []int{10, 1_000, 100_000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			nums := make([]int, size) // подготовка не должна попасть в замер
			for i := range nums {
				nums[i] = i
			}
			b.ReportAllocs()
			b.ResetTimer() // сбрасываем время, потраченное на подготовку
			for i := 0; i < b.N; i++ {
				sink = Sum(nums)
			}
		})
	}
}

// Упражнение 05: цикл без аллокаций (Решение).
package main

import (
	"fmt"
	"strconv"
	"testing"
)

// appendInts дописывает nums в dst через запятую и возвращает новый слайс.
func appendInts(dst []byte, nums []int) []byte {
	for i, n := range nums {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = strconv.AppendInt(dst, int64(n), 10) // пишет цифры прямо в dst
	}
	return dst
}

func main() {
	fmt.Println("--- Проверка 05_zero_alloc_loop (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := string(appendInts(nil, []int{1, 22, -3}))
	check(got == "1,22,-3", fmt.Sprintf("ожидалось \"1,22,-3\", получено %q", got))
	check(string(appendInts([]byte("ids="), []int{7})) == "ids=7", "числа должны дописываться к существующему содержимому dst")

	nums := []int{100, 200, 300, 400, 500}
	buf := make([]byte, 0, 64)
	allocs := testing.AllocsPerRun(100, func() { buf = appendInts(buf[:0], nums) })
	check(allocs == 0, fmt.Sprintf("ожидалось 0 аллокаций, получено %.0f", allocs))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

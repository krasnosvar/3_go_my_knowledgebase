// Упражнение 06: выбор партиции по ключу (Решение).
package main

import (
	"fmt"
	"hash/fnv"
)

func partition(key string, n int) int {
	if key == "" {
		return 0
	}
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() % uint32(n))
}

func main() {
	fmt.Println("--- Проверка 06_partition_key (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(partition("order-42", 6) == partition("order-42", 6), "один ключ → одна партиция")
	check(partition("", 6) == 0, "пустой ключ → 0")
	check(partition("a", 4) == int(0xe40c292c%4), "FNV-1a(\"a\") % 4")

	counts := make([]int, 4)
	inRange := true
	for i := 0; i < 1000; i++ {
		p := partition(fmt.Sprint("user-", i), 4)
		if p < 0 || p >= 4 {
			inRange = false
			break
		}
		counts[p]++
	}
	check(inRange, "партиция в диапазоне [0, n)")
	spread := true
	for _, c := range counts {
		if c < 150 {
			spread = false
		}
	}
	check(inRange && spread, "ключи распределены по всем партициям: "+fmt.Sprint(counts))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

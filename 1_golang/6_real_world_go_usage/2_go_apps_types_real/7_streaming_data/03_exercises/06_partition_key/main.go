// Упражнение 06: выбор партиции по ключу.
//
// Порядок в Kafka гарантирован только внутри партиции, поэтому события
// одного заказа должны идти по одному ключу → в одну партицию.
//
// Задача: реализуйте partition(key string, n int) int:
//   - hash FNV-1a 32-bit (hash/fnv: fnv.New32a) от ключа;
//   - результат hash % n (через uint32, чтобы не получить отрицательное число);
//   - пустой ключ → партиция 0 (в реальном продюсере был бы round-robin).
//
// Запуск: go run ./03_exercises/06_partition_key
package main

import (
	"fmt"
	"hash/fnv"
)

var _ = fnv.New32a

func partition(key string, n int) int {
	// TODO: fnv.New32a, Write, Sum32() % uint32(n).
	return -1
}

func main() {
	fmt.Println("--- Проверка 06_partition_key ---")

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

// Упражнение 01: выбор шарда по хеш-ключу (Решение).
package main

import (
	"fmt"
	"hash/fnv"
)

var _ = fnv.New32a

func pickShard(key string, n int) int {
	if n <= 0 {
		return -1
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32() % uint32(n))
}

func main() {
	fmt.Println("--- Проверка 01_pick_shard (Решение) ---")

	want := func(key string, n int) int {
		if n <= 0 {
			return -1
		}
		h := fnv.New32a()
		_, _ = h.Write([]byte(key))
		return int(h.Sum32() % uint32(n))
	}

	cases := []struct {
		key string
		n   int
	}{
		{"alice", 4},
		{"bob", 4},
		{"carol", 8},
		{"", 3},
		{"key-123", 1},
		{"anything", 0},
	}

	allOK := true
	for _, c := range cases {
		got := pickShard(c.key, c.n)
		exp := want(c.key, c.n)
		if got != exp {
			allOK = false
			fmt.Printf("FAIL: pickShard(%q, %d) = %d, ожидалось %d\n", c.key, c.n, got, exp)
		}
		if c.n > 0 && (got < 0 || got >= c.n) {
			allOK = false
			fmt.Printf("FAIL: pickShard(%q, %d) = %d вне диапазона [0,%d)\n", c.key, c.n, got, c.n)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

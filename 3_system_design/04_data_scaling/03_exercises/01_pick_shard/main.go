// Упражнение 01: выбор шарда по хеш-ключу.
//
// Задача: реализуйте pickShard(key, n) — номер шарда для ключа при
// hash-партиционировании: shard = hash(key) % n.
//
// Требования:
//   - используйте hash/fnv (FNV-1a, 32 бита), чтобы результат совпал с проверкой;
//   - результат в диапазоне [0, n);
//   - при n <= 0 верните -1.
//
// Запуск: go run ./03_exercises/01_pick_shard/main.go
package main

import (
	"fmt"
	"hash/fnv"
)

// Держим импорт живым в заготовке (в решении он используется).
var _ = fnv.New32a

// pickShard возвращает номер шарда для ключа: hash(key) % n (FNV-1a, 32 бита).
func pickShard(key string, n int) int {
	// TODO: посчитайте FNV-1a хеш ключа и верните остаток от деления на n.
	// Не забудьте случай n <= 0 -> -1.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_pick_shard ---")

	// Эталонная реализация для сверки (та же формула).
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
		// Проверка диапазона для валидного n.
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

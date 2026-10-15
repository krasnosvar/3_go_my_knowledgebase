// Упражнение 02: доля перемещённых ключей при ре-шардинге (Решение).
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

func movedFraction(keys []string, oldN, newN int) float64 {
	if len(keys) == 0 {
		return 0
	}
	moved := 0
	for _, k := range keys {
		if pickShard(k, oldN) != pickShard(k, newN) {
			moved++
		}
	}
	return float64(moved) / float64(len(keys))
}

func main() {
	fmt.Println("--- Проверка 02_reshard_moved (Решение) ---")

	keys := make([]string, 10000)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	want := func(keys []string, oldN, newN int) float64 {
		if len(keys) == 0 {
			return 0
		}
		moved := 0
		for _, k := range keys {
			if pickShard(k, oldN) != pickShard(k, newN) {
				moved++
			}
		}
		return float64(moved) / float64(len(keys))
	}

	allOK := true

	if got := movedFraction(nil, 3, 4); got != 0.0 {
		allOK = false
		fmt.Printf("FAIL: movedFraction(nil,...) = %v, ожидалось 0.0\n", got)
	}

	if got := movedFraction(keys, 4, 4); got != 0.0 {
		allOK = false
		fmt.Printf("FAIL: movedFraction(keys,4,4) = %v, ожидалось 0.0\n", got)
	}

	got := movedFraction(keys, 3, 4)
	exp := want(keys, 3, 4)
	if got != exp {
		allOK = false
		fmt.Printf("FAIL: movedFraction(keys,3,4) = %v, ожидалось %v\n", got, exp)
	}
	if got <= 0.5 {
		allOK = false
		fmt.Printf("FAIL: доля перемещения %v неожиданно мала (наивный hash%%N двигает почти всё)\n", got)
	}

	fmt.Printf("Инфо: при 3 -> 4 шардов перемещается %.1f%% ключей.\n", got*100)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

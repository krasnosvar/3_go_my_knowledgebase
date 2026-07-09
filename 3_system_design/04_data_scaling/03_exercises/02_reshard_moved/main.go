// Упражнение 02: доля перемещённых ключей при ре-шардинге.
//
// Задача: при наивном hash-партиционировании (shard = hash(key) % n)
// посчитайте, какая ДОЛЯ ключей сменит шард при переходе с oldN на newN.
//
// Реализуйте movedFraction(keys, oldN, newN) float64:
//   - для каждого ключа сравните pickShard(key, oldN) и pickShard(key, newN);
//   - верните долю ключей, у которых шард изменился, в [0.0, 1.0];
//   - если keys пуст — верните 0.0.
//
// Смысл: наивный hash%N при смене N перемещает ~(N-1)/N ключей — почти все.
// Именно поэтому на живых кластерах используют consistent hashing.
//
// Запуск: go run ./03_exercises/02_reshard_moved/main.go
package main

import (
	"fmt"
	"hash/fnv"
)

var _ = fnv.New32a

// pickShard: shard = hash(key) % n (FNV-1a, 32 бита). Уже реализовано.
func pickShard(key string, n int) int {
	if n <= 0 {
		return -1
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32() % uint32(n))
}

// movedFraction возвращает долю ключей, сменивших шард при oldN -> newN.
func movedFraction(keys []string, oldN, newN int) float64 {
	// TODO: пройдите по ключам, сравните pickShard(key, oldN) и
	// pickShard(key, newN), верните долю изменившихся.
	return -1
}

func main() {
	fmt.Println("--- Проверка 02_reshard_moved ---")

	keys := make([]string, 10000)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	// Эталон для сверки.
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

	// 1) Пустой набор -> 0.0
	if got := movedFraction(nil, 3, 4); got != 0.0 {
		allOK = false
		fmt.Printf("FAIL: movedFraction(nil,...) = %v, ожидалось 0.0\n", got)
	}

	// 2) Тот же N -> ничего не двигается.
	if got := movedFraction(keys, 4, 4); got != 0.0 {
		allOK = false
		fmt.Printf("FAIL: movedFraction(keys,4,4) = %v, ожидалось 0.0\n", got)
	}

	// 3) Смена N -> совпадает с эталоном и заметно велика (>0.5).
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

// Пример 02: маршрутизация ключа в шард.
//
// Показывает две базовые стратегии выбора шарда:
//  1. hash-партиционирование: shard = hash(key) % N — равномерно, но
//     диапазонные запросы разлетаются по всем шардам;
//  2. range-партиционирование: ключ по непрерывному диапазону — хорошо
//     для range-запросов, но риск перекоса (hot spot).
//
// Программа маршрутизирует набор ключей обеими стратегиями и печатает,
// в какой шард попадает каждый ключ, плюс итоговое распределение.
//
// Детерминированность: fnv-хеш и фиксированные границы диапазонов дают
// одинаковый вывод при каждом запуске. Горутин нет.
//
// Запуск: go run ./02_examples/02_shard_router/main.go
package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

// numShards — число шардов для hash-стратегии.
const numShards = 4

// hashShard: shard = hash(key) % N.
func hashShard(key string, n int) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32() % uint32(n))
}

// rangeShard: шард по диапазону первой буквы ключа (a..z).
// Границы делят алфавит на len(bounds)+1 диапазонов.
func rangeShard(key string, bounds []byte) int {
	if key == "" {
		return 0
	}
	c := key[0]
	// Находим первый диапазон, верхняя граница которого >= c.
	i := sort.Search(len(bounds), func(i int) bool { return bounds[i] >= c })
	return i
}

func main() {
	fmt.Println("--- Пример: маршрутизация ключа в шард ---")

	keys := []string{"alice", "bob", "carol", "dave", "erin", "frank", "grace", "heidi", "ivan", "mallory", "peggy", "zack"}

	// --- Hash-партиционирование ---
	fmt.Printf("\n[1] hash(key) %% %d:\n", numShards)
	hashDist := make([]int, numShards)
	for _, k := range keys {
		s := hashShard(k, numShards)
		hashDist[s]++
		fmt.Printf("  %-8s -> shard %d\n", k, s)
	}
	fmt.Print("  распределение: ")
	for s, c := range hashDist {
		fmt.Printf("shard%d=%d ", s, c)
	}
	fmt.Println()

	// --- Range-партиционирование ---
	// Границы (включительно верхняя): [a..h]=shard0, [i..p]=shard1,
	// [q..x]=shard2, [y..z]=shard3.
	bounds := []byte{'h', 'p', 'x', 'z'}
	fmt.Println("\n[2] range по первой букве (a-h|i-p|q-x|y-z):")
	rangeDist := make([]int, len(bounds))
	for _, k := range keys {
		s := rangeShard(k, bounds)
		rangeDist[s]++
		fmt.Printf("  %-8s -> shard %d\n", k, s)
	}
	fmt.Print("  распределение: ")
	for s, c := range rangeDist {
		fmt.Printf("shard%d=%d ", s, c)
	}
	fmt.Println()

	fmt.Println("\nВывод: hash даёт равномерность, но range-запрос (напр. 'a'..'c')")
	fmt.Println("разлетится по разным шардам; range хранит соседние ключи рядом,")
	fmt.Println("но легко даёт перекос (hot spot) при неравномерных данных.")
	fmt.Println("Готово, программа завершается.")
}

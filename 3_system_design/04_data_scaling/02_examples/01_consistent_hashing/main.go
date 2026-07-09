// Пример 01: кольцо консистентного хеширования (consistent hashing).
//
// Узлы и ключи хешируются на одно кольцо (uint32 через hash/fnv). Ключ
// принадлежит ближайшему узлу по часовой стрелке. Каждый физический узел
// представлен несколькими виртуальными точками (vnodes) — это выравнивает
// распределение.
//
// Программа:
//  1. раскладывает N ключей по 3 узлам и печатает распределение;
//  2. добавляет 4-й узел и считает, какая доля ключей сменила узел.
//
// Ключевой факт: при добавлении узла переезжает лишь ~1/N ключей,
// а не почти все (как было бы при hash(key) % N).
//
// Детерминированность: fnv-хеш и отсортированное кольцо дают один и тот же
// результат при каждом запуске. Горутин нет.
//
// Запуск: go run ./02_examples/01_consistent_hashing/main.go
package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

// vnodes — сколько виртуальных точек на кольце приходится на один физический узел.
const vnodes = 250

// ring — кольцо консистентного хеширования.
type ring struct {
	points []uint32          // отсортированные позиции vnode на кольце
	owner  map[uint32]string // позиция vnode -> имя физического узла
}

func newRing() *ring {
	return &ring{owner: make(map[uint32]string)}
}

// hash32 отображает строку в точку на кольце [0, 2^32).
func hash32(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// addNode добавляет физический узел как vnodes точек на кольце.
func (r *ring) addNode(node string) {
	for i := 0; i < vnodes; i++ {
		p := hash32(fmt.Sprintf("%s#%d", node, i))
		r.points = append(r.points, p)
		r.owner[p] = node
	}
	sort.Slice(r.points, func(i, j int) bool { return r.points[i] < r.points[j] })
}

// removeNode убирает все vnode-точки узла с кольца.
func (r *ring) removeNode(node string) {
	kept := r.points[:0]
	for _, p := range r.points {
		if r.owner[p] == node {
			delete(r.owner, p)
			continue
		}
		kept = append(kept, p)
	}
	r.points = kept
}

// lookup возвращает физический узел, отвечающий за ключ:
// первая vnode-точка по часовой стрелке (с заворотом кольца).
func (r *ring) lookup(key string) string {
	if len(r.points) == 0 {
		return ""
	}
	h := hash32(key)
	i := sort.Search(len(r.points), func(i int) bool { return r.points[i] >= h })
	if i == len(r.points) {
		i = 0 // заворот кольца
	}
	return r.owner[r.points[i]]
}

func main() {
	fmt.Println("--- Пример: consistent hashing ---")

	// Детерминированный набор ключей.
	const numKeys = 10000
	keys := make([]string, numKeys)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	// Кольцо из трёх узлов.
	r := newRing()
	for _, n := range []string{"node-A", "node-B", "node-C"} {
		r.addNode(n)
	}

	// Раскладка ключей и снимок "ключ -> узел" до изменения топологии.
	before := make(map[string]string, numKeys)
	dist := map[string]int{}
	for _, k := range keys {
		n := r.lookup(k)
		before[k] = n
		dist[n]++
	}

	fmt.Printf("Распределение %d ключей по 3 узлам (vnodes=%d):\n", numKeys, vnodes)
	for _, n := range []string{"node-A", "node-B", "node-C"} {
		fmt.Printf("  %s: %d ключей (%.1f%%)\n", n, dist[n], 100*float64(dist[n])/float64(numKeys))
	}

	// Добавляем 4-й узел и смотрим, сколько ключей переехало.
	r.addNode("node-D")
	moved := 0
	distAfter := map[string]int{}
	for _, k := range keys {
		n := r.lookup(k)
		distAfter[n]++
		if before[k] != n {
			moved++
		}
	}

	fmt.Println("\nПосле добавления node-D:")
	for _, n := range []string{"node-A", "node-B", "node-C", "node-D"} {
		fmt.Printf("  %s: %d ключей\n", n, distAfter[n])
	}
	fmt.Printf("Переехало ключей: %d из %d (%.1f%%)\n", moved, numKeys, 100*float64(moved)/float64(numKeys))
	fmt.Printf("Для сравнения, наивный hash%%N переместил бы ~%.0f%% (3/4).\n", 75.0)

	fmt.Println("\nВывод: consistent hashing перемещает лишь долю ~1/N ключей,")
	fmt.Println("а не почти все, как сделал бы hash(key) % N при смене N.")
	fmt.Println("Готово, программа завершается.")
}

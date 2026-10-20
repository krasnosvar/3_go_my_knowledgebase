// Упражнение 07: распределение по партициям с сохранением порядка (Решение).
package main

import (
	"fmt"
	"hash/fnv"
)

var _ = fnv.New32a

type msg struct {
	key   string
	value int
}

func partitionOf(key string, p int) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32() % uint32(p))
}

func route(msgs []msg, p int) [][]msg {
	parts := make([][]msg, p)
	for _, m := range msgs {
		idx := partitionOf(m.key, p)
		// append сохраняет относительный порядок сообщений одного ключа,
		// т.к. все они попадают в одну и ту же партицию по фиксированному idx.
		parts[idx] = append(parts[idx], m)
	}
	return parts
}

func main() {
	fmt.Println("--- Проверка 07_partition_by_key (Решение) ---")

	const p = 3
	msgs := []msg{
		{"user-1", 1},
		{"user-2", 2},
		{"user-1", 3},
		{"user-3", 4},
		{"user-2", 5},
		{"user-1", 6},
	}

	parts := route(msgs, p)

	allOK := true

	if len(parts) != p {
		allOK = false
		fmt.Printf("FAIL: партиций %d, ожидалось %d\n", len(parts), p)
	}

	if allOK {
		keyPart := map[string]int{}
		lastValForKey := map[string]int{}
		total := 0
		for pi, part := range parts {
			for _, m := range part {
				total++
				if prev, ok := keyPart[m.key]; ok {
					if prev != pi {
						allOK = false
						fmt.Printf("FAIL: ключ %q попал в партиции %d и %d\n", m.key, prev, pi)
					}
				} else {
					keyPart[m.key] = pi
				}
				if lv, ok := lastValForKey[m.key]; ok && m.value <= lv {
					allOK = false
					fmt.Printf("FAIL: порядок для ключа %q нарушен: %d после %d\n", m.key, m.value, lv)
				}
				lastValForKey[m.key] = m.value
			}
		}
		if total != len(msgs) {
			allOK = false
			fmt.Printf("FAIL: распределено %d сообщений, ожидалось %d\n", total, len(msgs))
		}
	}

	if allOK {
		if partitionOf("user-1", p) != partitionOf("user-1", p) {
			allOK = false
			fmt.Println("FAIL: partitionOf недетерминирован")
		}
		idx := partitionOf("user-2", p)
		if idx < 0 || idx >= p {
			allOK = false
			fmt.Printf("FAIL: partitionOf вне диапазона: %d\n", idx)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

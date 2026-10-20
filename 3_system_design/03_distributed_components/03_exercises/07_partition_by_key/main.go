// Упражнение 07: распределение сообщений по партициям с сохранением порядка.
//
// Задача: реализуйте маршрутизацию сообщений по партициям так же, как это
// делают лог-брокеры: partition = hash(key) % P. Все сообщения с одним ключом
// должны попадать в ОДНУ партицию, а их относительный порядок — сохраняться.
//
// partitionOf(key, p) -> индекс партиции [0, p) через хэш ключа по модулю p.
//
//	Используйте fnv.New32a() из hash/fnv для стабильного хэша.
//
// route(msgs, p) -> [][]string: msgs распределены по p партициям; внутри
//
//	каждой партиции сообщения идут в том же относительном порядке, что и во
//	входном срезе (append сохраняет порядок).
//
// Запуск: go run ./03_exercises/07_partition_by_key/main.go
package main

import (
	"fmt"
	"hash/fnv"
)

var _ = fnv.New32a

// msg — сообщение с ключом (по нему выбирается партиция) и порядковым value.
type msg struct {
	key   string
	value int
}

// partitionOf возвращает индекс партиции для ключа: hash(key) % p.
func partitionOf(key string, p int) int {
	// TODO: посчитайте FNV-32a хэш ключа и возьмите остаток по модулю p.
	return 0
}

// route распределяет сообщения по p партициям, сохраняя порядок внутри
// партиции.
func route(msgs []msg, p int) [][]msg {
	// TODO: создайте p партиций и разложите сообщения по partitionOf(key, p).
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_partition_by_key ---")

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

	// 1) Ровно p партиций.
	if len(parts) != p {
		allOK = false
		fmt.Printf("FAIL: партиций %d, ожидалось %d\n", len(parts), p)
	}

	// 2) Все сообщения одного ключа — в одной партиции; порядок value сохранён.
	if allOK {
		keyPart := map[string]int{} // ключ -> в какой партиции его встретили
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

	// 3) partitionOf детерминирован и в диапазоне [0, p).
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

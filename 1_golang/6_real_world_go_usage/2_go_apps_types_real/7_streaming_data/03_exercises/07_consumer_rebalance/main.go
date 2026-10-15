// Упражнение 07: распределение партиций в consumer group (range assignor).
//
// Задача: реализуйте assign(consumers []string, partitions int) map[string][]int:
//   - участники сортируются по имени (порядок входа не важен);
//   - каждому достаётся непрерывный диапазон: partitions / len(consumers),
//     а первые partitions % len(consumers) участников получают на одну больше;
//   - без участников → пустая map;
//   - участник без партиций всё равно присутствует в map (пустой срез).
//
// После ребаланса (новый consumer) распределение пересчитывается целиком.
//
// Запуск: go run ./03_exercises/07_consumer_rebalance
package main

import (
	"fmt"
	"slices"
)

var _ = slices.Sort[[]string]

func assign(consumers []string, partitions int) map[string][]int {
	// TODO: отсортируйте копию consumers и раздайте диапазоны.
	return map[string][]int{}
}

func main() {
	fmt.Println("--- Проверка 07_consumer_rebalance ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := assign([]string{"c2", "c1"}, 5)
	check(fmt.Sprint(a["c1"]) == "[0 1 2]" && fmt.Sprint(a["c2"]) == "[3 4]", "5 партиций на 2: c1=[0 1 2], c2=[3 4]")

	a = assign([]string{"c3", "c1", "c2"}, 5)
	check(fmt.Sprint(a["c1"], a["c2"], a["c3"]) == "[0 1] [2 3] [4]", "ребаланс на 3: "+fmt.Sprint(a["c1"], a["c2"], a["c3"]))

	a = assign([]string{"c1", "c2", "c3"}, 2)
	got, ok := a["c3"]
	check(ok && len(got) == 0, "лишний участник присутствует с пустым списком")
	check(len(assign(nil, 3)) == 0, "без участников → пустая map")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// Упражнение 07: распределение партиций (Решение).
package main

import (
	"fmt"
	"slices"
)

func assign(consumers []string, partitions int) map[string][]int {
	out := map[string][]int{}
	if len(consumers) == 0 {
		return out
	}
	sorted := slices.Clone(consumers)
	slices.Sort(sorted)
	per, extra := partitions/len(sorted), partitions%len(sorted)
	p := 0
	for i, c := range sorted {
		n := per
		if i < extra {
			n++
		}
		out[c] = []int{}
		for j := 0; j < n; j++ {
			out[c] = append(out[c], p)
			p++
		}
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 07_consumer_rebalance (Решение) ---")

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

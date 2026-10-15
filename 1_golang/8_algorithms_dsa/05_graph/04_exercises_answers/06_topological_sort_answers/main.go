// Упражнение 06: топологическая сортировка (Кан) (Решение).
package main

import "fmt"

// respects проверяет, что order — перестановка 0..n-1 и для каждого ребра u идёт раньше v.
func respects(order []int, n int, edges [][2]int) bool {
	if len(order) != n {
		return false
	}
	pos := make(map[int]int, n)
	for i, v := range order {
		pos[v] = i
	}
	if len(pos) != n {
		return false
	}
	for _, e := range edges {
		if pos[e[0]] >= pos[e[1]] {
			return false
		}
	}
	return true
}

// topoSort возвращает топологический порядок или (nil, false) при цикле.
func topoSort(n int, edges [][2]int) ([]int, bool) {
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		indeg[e[1]]++
	}
	var queue []int
	for v := 0; v < n; v++ {
		if indeg[v] == 0 {
			queue = append(queue, v)
		}
	}
	order := make([]int, 0, n)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, w := range adj[v] {
			indeg[w]--
			if indeg[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	if len(order) != n {
		return nil, false
	}
	return order, true
}

func main() {
	fmt.Println("--- Проверка 06_topological_sort (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	// Зависимости сборки: 0=lib, 1=db, 2=api, 3=web, 4=tests.
	edges := [][2]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {2, 4}, {3, 4}}
	order, ok := topoSort(5, edges)
	check(ok && respects(order, 5, edges), fmt.Sprintf("получен некорректный порядок %v (ok=%t)", order, ok))

	_, ok = topoSort(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})
	check(!ok, "граф с циклом: ожидалось ok=false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

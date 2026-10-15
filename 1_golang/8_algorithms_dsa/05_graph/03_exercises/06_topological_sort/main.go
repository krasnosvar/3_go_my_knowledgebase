// Упражнение 06: топологическая сортировка (Кан).
//
// Задача: реализуйте topoSort(n, edges) — топологический порядок вершин
// ориентированного графа (ребро {u, v} значит «u раньше v»).
// Алгоритм Кана: считаем входящие степени, в очередь кладём вершины со
// степенью 0, извлекаем, уменьшаем степени соседей.
// Если в порядок попали не все n вершин — в графе цикл: верните (nil, false).
//
// Запуск: go run ./03_exercises/06_topological_sort/main.go
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
	// TODO: входящие степени + очередь.
	return nil, false
}

func main() {
	fmt.Println("--- Проверка 06_topological_sort ---")

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

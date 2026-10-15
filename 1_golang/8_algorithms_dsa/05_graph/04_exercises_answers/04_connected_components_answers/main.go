// Упражнение 04: подсчёт компонент связности (Решение).
package main

import "fmt"

// countComponents возвращает число компонент связности.
func countComponents(n int, edges [][2]int) int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // сжатие пути
		}
		return parent[x]
	}
	count := n
	for _, e := range edges {
		a, b := find(e[0]), find(e[1])
		if a != b {
			parent[a] = b
			count--
		}
	}
	return count
}

func main() {
	fmt.Println("--- Проверка 04_connected_components (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := countComponents(5, [][2]int{{0, 1}, {1, 2}, {3, 4}})
	check(got == 2, fmt.Sprintf("{0,1,2} и {3,4}: ожидалось 2, получено %d", got))

	got = countComponents(4, nil)
	check(got == 4, fmt.Sprintf("4 вершины без рёбер: ожидалось 4, получено %d", got))

	got = countComponents(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})
	check(got == 1, fmt.Sprintf("треугольник: ожидалось 1, получено %d", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

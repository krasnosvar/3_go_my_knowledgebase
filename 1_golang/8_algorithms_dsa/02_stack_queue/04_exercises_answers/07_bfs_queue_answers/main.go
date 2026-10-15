// Упражнение 07: обход в ширину с очередью (Решение).
package main

import "fmt"

// bfsOrder возвращает порядок посещения вершин при BFS.
func bfsOrder(graph map[int][]int, start int) []int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	var order []int
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, n := range graph[v] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	return order
}

func main() {
	fmt.Println("--- Проверка 07_bfs_queue (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	graph := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4, 5},
		4: {6},
		5: {6},
		6: {1},
	}
	got := bfsOrder(graph, 1)
	check(fmt.Sprint(got) == "[1 2 3 4 5 6]", fmt.Sprintf("ожидалось [1 2 3 4 5 6], получено %v", got))

	got = bfsOrder(graph, 5)
	check(fmt.Sprint(got) == "[5 6 1 2 3 4]", fmt.Sprintf("старт с 5: ожидалось [5 6 1 2 3 4], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

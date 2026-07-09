package main

import "fmt"

func maxSumK(a []int, k int) int {
	if len(a) < k {
		return 0
	}
	s := 0
	for i := 0; i < k; i++ {
		s += a[i]
	}
	best := s
	for r := k; r < len(a); r++ {
		s += a[r] - a[r-k]
		if s > best {
			best = s
		}
	}
	return best
}

func main() { fmt.Println(maxSumK([]int{1, 2, 3, 4, 5}, 2)) }

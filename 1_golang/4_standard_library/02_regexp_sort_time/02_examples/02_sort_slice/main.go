package main

import (
	"fmt"
	"sort"
)

func main() {
	vals := []int{4, 1, 3, 2}
	sort.Ints(vals)
	fmt.Println(vals)
}

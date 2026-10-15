package main

import "fmt"

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}
	for i, v := range nums {
		if j, ok := idx[target-v]; ok {
			return []int{j, i}
		}
		idx[v] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

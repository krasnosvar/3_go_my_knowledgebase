package main

import "fmt"

func main() {
	msgs := []int{1, 2, 3, 4, 5}
	batch := 2
	for i := 0; i < len(msgs); i += batch {
		end := i + batch
		if end > len(msgs) {
			end = len(msgs)
		}
		fmt.Println("batch", msgs[i:end])
	}
}

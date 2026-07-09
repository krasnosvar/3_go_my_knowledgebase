package main

import "fmt"

func toInt(v any) (int, bool) {
	n, ok := v.(int)
	return n, ok
}

func main() {
	fmt.Println(toInt(7))
}

package main

import "fmt"

func main() {
	var v any = 3.14
	switch x := v.(type) {
	case float64:
		fmt.Println(x)
	default:
		fmt.Println("other")
	}
}

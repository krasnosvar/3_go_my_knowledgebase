package main

import "fmt"

func main() {
	var v any = 200
	switch x := v.(type) {
	case int:
		fmt.Println("code", x)
	default:
		fmt.Println("other")
	}
}

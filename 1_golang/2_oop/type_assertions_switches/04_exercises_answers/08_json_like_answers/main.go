package main

import "fmt"

func main() {
	m := map[string]any{"name": "Den", "age": 30}
	if n, ok := m["name"].(string); ok {
		fmt.Println(n)
	}
}

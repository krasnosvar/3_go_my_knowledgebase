package main

import "fmt"

func main() {
	payload := map[string]any{"id": 1}
	if id, ok := payload["id"].(int); ok {
		fmt.Println(id)
	}
}

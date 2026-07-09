package main

import "fmt"

func main() {
	msgs := []string{"m1", "m2", "m3"}
	for _, m := range msgs {
		fmt.Println("consume", m)
	}
}

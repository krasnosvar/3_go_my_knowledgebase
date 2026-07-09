package main

import (
	"fmt"
)

func main() {
	err := fmt.Errorf("temporary")
	for i := 1; i <= 3; i++ {
		if err == nil {
			break
		}
		fmt.Println("attempt", i, "failed")
		if i == 3 {
			fmt.Println("give up")
		}
	}
}

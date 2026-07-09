package main

import "fmt"

func main() {
	maxAttempts := 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Println("attempt", attempt)
	}
}

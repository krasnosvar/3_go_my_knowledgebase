package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 2; i++ {
		<-ticker.C
		fmt.Println("tick")
	}
}

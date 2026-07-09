package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2026-04-17")
	fmt.Println(t.Format(time.RFC3339), (2 * time.Second).String())
}

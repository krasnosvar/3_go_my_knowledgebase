package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "42"
	v, _ := strconv.Atoi(s)
	fmt.Printf("%d -> %d\n", 10, v+10)
}

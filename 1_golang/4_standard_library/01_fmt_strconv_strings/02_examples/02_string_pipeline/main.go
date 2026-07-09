package main

import (
	"fmt"
	"strings"
)

func main() {
	raw := "  go,knowledge,base  "
	parts := strings.Split(strings.TrimSpace(raw), ",")
	fmt.Println(strings.Join(parts, " | "))
}

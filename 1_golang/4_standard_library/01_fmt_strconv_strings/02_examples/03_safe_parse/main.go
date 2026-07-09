package main

import (
	"fmt"
	"strconv"
)

func parsePort(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid port %q: %w", s, err)
	}
	return v, nil
}

func main() {
	fmt.Println(parsePort("8080"))
}

package main

import (
	"errors"
	"fmt"
)

var ErrX = errors.New("x")

func main() {
	err := fmt.Errorf("ctx: %w", ErrX)
	fmt.Println(errors.Is(err, ErrX))
}

package main

import (
	"errors"
	"fmt"
)

var ErrConflict = errors.New("conflict")

func main() {
	fmt.Println(ErrConflict)
}

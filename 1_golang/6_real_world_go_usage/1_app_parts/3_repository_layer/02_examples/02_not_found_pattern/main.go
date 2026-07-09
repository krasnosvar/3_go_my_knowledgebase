package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func main() { fmt.Println(ErrNotFound) }

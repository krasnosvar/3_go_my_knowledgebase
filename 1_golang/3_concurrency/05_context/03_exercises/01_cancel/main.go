package main

import (
	"context"
	"fmt"
)

func main() {
	// TODO: создай context.WithCancel и сразу вызови cancel.
	_ = context.Background
	fmt.Println("expected: context canceled")
}

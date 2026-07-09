package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// TODO: сделай timeout 100ms и проверь ctx.Err() после истечения.
	_ = time.Second
	_ = context.Background
	fmt.Println("implement timeout")
}

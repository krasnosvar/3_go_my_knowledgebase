package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var state int32 = 0
	ok := atomic.CompareAndSwapInt32(&state, 0, 1)
	fmt.Println(ok, state)
}

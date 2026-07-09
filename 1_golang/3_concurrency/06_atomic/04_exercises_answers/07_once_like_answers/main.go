package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var init int32
	if atomic.CompareAndSwapInt32(&init, 0, 1) {
		fmt.Println("init")
	}
	if atomic.CompareAndSwapInt32(&init, 0, 1) {
		fmt.Println("no")
	}
}

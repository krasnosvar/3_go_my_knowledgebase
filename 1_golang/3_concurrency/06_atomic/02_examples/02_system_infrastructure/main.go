package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var healthy int32
	atomic.StoreInt32(&healthy, 1)
	fmt.Println(atomic.LoadInt32(&healthy))
}

package main

import (
	"fmt"
	"sync/atomic"
)

func main() { var v int32 = 1; fmt.Println(atomic.CompareAndSwapInt32(&v, 1, 2), v) }

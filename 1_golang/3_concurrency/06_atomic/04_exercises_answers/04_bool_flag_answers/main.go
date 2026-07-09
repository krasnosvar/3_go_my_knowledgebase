package main

import (
	"fmt"
	"sync/atomic"
)

func main() { var stop int32; atomic.StoreInt32(&stop, 1); fmt.Println(atomic.LoadInt32(&stop) == 1) }

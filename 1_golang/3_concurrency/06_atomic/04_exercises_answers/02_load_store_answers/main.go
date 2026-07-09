package main

import (
	"fmt"
	"sync/atomic"
)

func main() { var v int64; atomic.StoreInt64(&v, 9); fmt.Println(atomic.LoadInt64(&v)) }

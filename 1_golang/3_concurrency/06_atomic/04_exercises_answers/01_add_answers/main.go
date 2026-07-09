package main

import (
	"fmt"
	"sync/atomic"
)

func main() { var v int64; atomic.AddInt64(&v, 5); fmt.Println(v) }

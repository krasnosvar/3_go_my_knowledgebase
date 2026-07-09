package main

import (
	"fmt"
	"runtime"
)

func main() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("alloc_bytes:", ms.Alloc)
	fmt.Println("num_gc:", ms.NumGC)
}

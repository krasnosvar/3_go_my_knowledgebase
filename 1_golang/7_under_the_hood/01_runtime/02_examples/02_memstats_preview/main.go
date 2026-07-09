package main

import (
	"fmt"
	"runtime"
)

func main() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("heap_alloc:", ms.HeapAlloc)
	fmt.Println("goroutines:", runtime.NumGoroutine())
}

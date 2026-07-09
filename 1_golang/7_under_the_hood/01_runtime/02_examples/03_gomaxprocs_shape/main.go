package main

import (
	"fmt"
	"runtime"
)

func main() {
	old := runtime.GOMAXPROCS(2)
	fmt.Println("old:", old, "new:", runtime.GOMAXPROCS(0))
}

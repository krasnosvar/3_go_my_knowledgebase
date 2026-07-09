package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("gomaxprocs:", runtime.GOMAXPROCS(0))
	fmt.Println("numcpu:", runtime.NumCPU())
}

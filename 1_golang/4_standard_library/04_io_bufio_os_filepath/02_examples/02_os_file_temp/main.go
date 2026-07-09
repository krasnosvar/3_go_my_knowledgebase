package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.CreateTemp("", "kb-*.txt")
	defer os.Remove(f.Name())
	defer f.Close()
	fmt.Println(f.Name() != "")
}

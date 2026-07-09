package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: app <build|clean>")
		return
	}
	switch os.Args[1] {
	case "build":
		fmt.Println("build started")
	case "clean":
		fmt.Println("cleanup done")
	default:
		fmt.Println("unknown command")
	}
}

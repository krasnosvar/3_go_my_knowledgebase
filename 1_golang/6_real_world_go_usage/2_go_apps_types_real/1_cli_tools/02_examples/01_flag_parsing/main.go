package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "user name")
	flag.Parse()
	fmt.Println("hello", *name)
}

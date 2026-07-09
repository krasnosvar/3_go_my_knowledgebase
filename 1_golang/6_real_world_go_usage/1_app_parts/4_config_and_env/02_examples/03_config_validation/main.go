package main

import (
	"fmt"
)

func main() { port := 8080; fmt.Println(port > 0 && port < 65536) }

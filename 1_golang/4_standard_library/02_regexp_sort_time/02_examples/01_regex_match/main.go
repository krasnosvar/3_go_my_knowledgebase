package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`^[a-z]+@[a-z]+\.[a-z]+$`)
	fmt.Println(re.MatchString("go@dev.io"))
}

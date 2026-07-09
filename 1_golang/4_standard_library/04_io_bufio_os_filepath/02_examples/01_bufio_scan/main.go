package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := bufio.NewScanner(strings.NewReader("a\nb\n"))
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

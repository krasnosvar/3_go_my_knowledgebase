package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/etc/hosts")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("opened")
}

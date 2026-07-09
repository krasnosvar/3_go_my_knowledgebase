package main

import "fmt"

type A struct{}

func main() {
	var v any = A{}
	switch v.(type) {
	case A:
		fmt.Println("A")
	}
}

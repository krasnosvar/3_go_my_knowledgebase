package main

import "fmt"

type Counter interface {
	Inc()
	Value() int
}

type c struct{ v int }

// TODO: реализуй методы через pointer receiver.

func main() {
	var cnt Counter = &c{}
	cnt.Inc()
	cnt.Inc()
	fmt.Println("expected: 2")
	fmt.Println("actual:", cnt.Value())
}

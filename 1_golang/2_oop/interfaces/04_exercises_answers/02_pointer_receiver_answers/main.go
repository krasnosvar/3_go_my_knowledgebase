package main

import "fmt"

type Counter interface {
	Inc()
	Value() int
}

type c struct{ v int }

func (x *c) Inc()       { x.v++ }
func (x *c) Value() int { return x.v }

func main() {
	var cnt Counter = &c{}
	cnt.Inc()
	cnt.Inc()
	fmt.Println(cnt.Value())
}

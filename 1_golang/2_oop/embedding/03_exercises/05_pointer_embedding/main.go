package main

import "fmt"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

type Wrap struct {
	// TODO: embed *Counter and initialize it safely.
}

func main() {
	w := Wrap{}
	_ = w
	fmt.Println("implement pointer embedding")
}

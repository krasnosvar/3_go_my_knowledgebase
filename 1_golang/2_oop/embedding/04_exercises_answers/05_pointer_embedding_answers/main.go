package main

import "fmt"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

type Wrap struct{ *Counter }

func NewWrap() Wrap { return Wrap{Counter: &Counter{}} }

func main() {
	w := NewWrap()
	w.Inc()
	fmt.Println(w.n)
}

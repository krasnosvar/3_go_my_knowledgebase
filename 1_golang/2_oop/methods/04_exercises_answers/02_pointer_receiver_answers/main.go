package main

import "fmt"

type C struct{ N int }

func (c *C) Inc() { c.N++ }
func main()       { c := &C{}; c.Inc(); fmt.Println(c.N) }

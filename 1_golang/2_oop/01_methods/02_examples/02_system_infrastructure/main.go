package main

import "fmt"

type Counter struct{ N int }

func (c *Counter) Inc() { c.N++ }
func main()             { c := &Counter{}; c.Inc(); fmt.Println(c.N) }

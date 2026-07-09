package main

import "fmt"

type P struct{ X int }

func (p P) Val() int { return p.X }
func main()          { fmt.Println(P{2}.Val()) }

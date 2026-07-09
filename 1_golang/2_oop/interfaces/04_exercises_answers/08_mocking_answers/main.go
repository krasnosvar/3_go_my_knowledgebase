package main

import "fmt"

type Clock interface {
	Now() string
}

type fixedClock struct{}

func (fixedClock) Now() string { return "10:00" }

func greet(c Clock) string { return "time: " + c.Now() }

func main() { fmt.Println(greet(fixedClock{})) }

package main

import "fmt"

type Clock interface {
	Now() string
}

type fixedClock struct{}

// TODO: реализуй Now() string у fixedClock.

func greet(c Clock) string {
	return "time: " + c.Now()
}

func main() {
	fmt.Println(greet(fixedClock{}))
}

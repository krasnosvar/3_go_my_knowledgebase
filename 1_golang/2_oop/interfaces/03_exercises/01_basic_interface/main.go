package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Cat struct{}

// TODO: реализуй метод Speak() string для Cat.

func say(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// expected: meow
	say(Cat{})
}

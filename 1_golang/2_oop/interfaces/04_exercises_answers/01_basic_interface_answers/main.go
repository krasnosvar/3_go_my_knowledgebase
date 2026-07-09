package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Cat struct{}

func (Cat) Speak() string { return "meow" }

func say(s Speaker) { fmt.Println(s.Speak()) }

func main() { say(Cat{}) }

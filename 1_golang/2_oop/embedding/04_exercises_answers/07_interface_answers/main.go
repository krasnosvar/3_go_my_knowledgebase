package main

import "fmt"

type Runner interface{ Run() }

type Base struct{}

func (Base) Run() { fmt.Println("base run") }

type Job struct{ Base }

func main() {
	var r Runner = Job{}
	r.Run()
}

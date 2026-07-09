package main

import "fmt"

type Logger struct{}

func (Logger) Log(msg string) { fmt.Println(msg) }

type Service struct {
	// TODO: embed Logger
}

func main() {
	Service{}.Log("ok")
}

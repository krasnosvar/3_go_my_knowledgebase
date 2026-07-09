package main

import "fmt"

type Logger struct{}

func (Logger) Log(msg string) { fmt.Println(msg) }

type Service struct{ Logger }

func main() { Service{}.Log("ok") }

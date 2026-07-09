package main

import "fmt"

type Auth struct{}

func (Auth) Check() bool { return true }

type API struct {
	// TODO: embed Auth and use Check() in method.
}

func main() { fmt.Println("implement composition") }

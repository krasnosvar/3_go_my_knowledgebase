package main

import "fmt"

type Auth struct{}

func (Auth) Check() bool { return true }

type API struct{ Auth }

func (a API) Allowed() bool { return a.Check() }

func main() { fmt.Println(API{}.Allowed()) }

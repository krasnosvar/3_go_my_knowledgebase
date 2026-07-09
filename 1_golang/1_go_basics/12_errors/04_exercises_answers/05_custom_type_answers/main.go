package main

import "fmt"

type ValidationErr struct{ Field string }

func (v ValidationErr) Error() string { return "validation failed: " + v.Field }

func run() error { return ValidationErr{Field: "email"} }

func main() { fmt.Println(run()) }

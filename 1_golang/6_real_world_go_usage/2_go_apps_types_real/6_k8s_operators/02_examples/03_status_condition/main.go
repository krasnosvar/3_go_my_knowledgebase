package main

import "fmt"

type Condition struct {
	Type   string
	Status string
}

func main() {
	ready := Condition{Type: "Ready", Status: "True"}
	fmt.Printf("%s=%s\n", ready.Type, ready.Status)
}

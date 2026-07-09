package main

import "fmt"

type Request struct{ Name string }

func Reconcile(r Request) (requeue bool) {
	fmt.Println("reconcile", r.Name)
	return false
}

func main() {
	_ = Reconcile(Request{Name: "demo"})
}

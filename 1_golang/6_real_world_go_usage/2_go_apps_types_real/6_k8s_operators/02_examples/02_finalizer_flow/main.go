package main

import "fmt"

type Resource struct {
	Name           string
	DeletionMarked bool
	HasFinalizer   bool
}

func main() {
	r := Resource{Name: "cache", DeletionMarked: true, HasFinalizer: true}
	if r.DeletionMarked && r.HasFinalizer {
		fmt.Println("cleanup external resource and remove finalizer")
	}
}

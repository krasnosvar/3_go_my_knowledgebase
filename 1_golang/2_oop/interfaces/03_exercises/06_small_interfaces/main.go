package main

import "fmt"

type Reader interface {
	Read() string
}

type fileReader struct{}

// TODO: реализуй Read() string.

func handle(r Reader) {
	fmt.Println(r.Read())
}

func main() {
	handle(fileReader{})
}

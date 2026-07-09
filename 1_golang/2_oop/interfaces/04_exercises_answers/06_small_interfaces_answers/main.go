package main

import "fmt"

type Reader interface {
	Read() string
}

type fileReader struct{}

func (fileReader) Read() string { return "ok" }

func handle(r Reader) { fmt.Println(r.Read()) }

func main() { handle(fileReader{}) }

package main

import "fmt"

type errWrap struct{}

func (e *errWrap) Error() string { return "boom" }

func returnsTypedNil() error {
	return nil
}

func main() {
	err := returnsTypedNil()
	fmt.Println(err == nil)
}

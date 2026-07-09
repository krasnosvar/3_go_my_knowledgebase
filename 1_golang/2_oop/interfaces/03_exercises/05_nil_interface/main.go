package main

import "fmt"

type errWrap struct{}

func (e *errWrap) Error() string { return "boom" }

func returnsTypedNil() error {
	// TODO: исправь, чтобы функция возвращала настоящий nil interface.
	var e *errWrap = nil
	return e
}

func main() {
	err := returnsTypedNil()
	fmt.Println("expected: true")
	fmt.Println("actual:", err == nil)
}

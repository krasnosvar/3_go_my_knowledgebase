package main

import "fmt"

func toInt(v any) (int, bool) {
	// TODO: безопасно извлеки int через type assertion.
	return 0, false
}

func main() {
	fmt.Println("expected: 7 true")
	fmt.Println("actual:", toInt(7))
}

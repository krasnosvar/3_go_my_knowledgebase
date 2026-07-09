package main

import "fmt"

func describe(v any) string {
	// TODO: через type switch верни: "int", "string" или "other".
	return ""
}

func main() {
	fmt.Println("expected: int")
	fmt.Println("actual:", describe(1))
}

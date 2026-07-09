package main

import "fmt"

func parse(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int:%d", x)
	case string:
		return "str:" + x
	default:
		return "other"
	}
}
func main() { fmt.Println(parse(7)) }

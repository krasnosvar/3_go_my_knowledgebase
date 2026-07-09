package main

import "fmt"

func main() { var v any = "x"; s, ok := v.(string); fmt.Println(s, ok) }

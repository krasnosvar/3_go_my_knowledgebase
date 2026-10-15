package main

import "fmt"

func main() { var v any = "build"; s, ok := v.(string); fmt.Println(s, ok) }

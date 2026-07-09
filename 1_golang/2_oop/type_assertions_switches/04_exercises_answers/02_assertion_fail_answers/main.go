package main

import "fmt"

func main() { var v any = 1; _, ok := v.(string); fmt.Println(ok) }

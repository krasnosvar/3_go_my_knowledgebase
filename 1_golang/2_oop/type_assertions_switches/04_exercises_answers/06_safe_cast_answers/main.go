package main

import "fmt"

func asInt(v any) (int, bool) { n, ok := v.(int); return n, ok }
func main()                   { fmt.Println(asInt(5)) }

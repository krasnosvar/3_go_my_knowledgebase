package main

import "fmt"

type R interface{ Read([]byte) (int, error) }

func main() { var v any = "x"; _, ok := v.(R); fmt.Println(ok) }

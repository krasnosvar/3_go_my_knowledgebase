package main

import "fmt"

func main() { p := map[string]struct{}{"read": {}}; _, ok := p["read"]; fmt.Println(ok) }

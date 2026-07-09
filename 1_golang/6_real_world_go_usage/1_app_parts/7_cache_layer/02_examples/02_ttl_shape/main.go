package main

import (
	"fmt"
	"time"
)

func main() { exp := time.Now().Add(time.Minute); fmt.Println(time.Now().Before(exp)) }

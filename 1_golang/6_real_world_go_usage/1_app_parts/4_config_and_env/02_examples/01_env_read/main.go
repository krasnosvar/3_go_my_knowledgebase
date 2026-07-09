package main

import (
	"fmt"
	"os"
)

func main() { _ = os.Setenv("APP_PORT", "8080"); fmt.Println(os.Getenv("APP_PORT")) }

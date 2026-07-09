package main

import "fmt"

func step1() error { return fmt.Errorf("db unavailable") }
func step2() error { return fmt.Errorf("service step2: %w", step1()) }

func main() { fmt.Println(step2()) }

package main

import "fmt"

func wrap(err error) error { return fmt.Errorf("service failed: %w", err) }

func main() { fmt.Println(wrap(fmt.Errorf("root")) != nil) }

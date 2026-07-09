package main

import "fmt"

func main() { q := make(chan int, 2); q <- 1; q <- 2; fmt.Println(<-q, <-q) }

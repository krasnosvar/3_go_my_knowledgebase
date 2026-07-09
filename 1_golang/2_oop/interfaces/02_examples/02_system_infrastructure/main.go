package main

import "fmt"

type HealthChecker interface {
	Healthy() bool
}

type Node struct{ alive bool }

func (n Node) Healthy() bool { return n.alive }

func countHealthy(items []HealthChecker) int {
	count := 0
	for _, it := range items {
		if it.Healthy() {
			count++
		}
	}
	return count
}

func main() {
	items := []HealthChecker{Node{true}, Node{false}, Node{true}}
	fmt.Println("healthy:", countHealthy(items))
}

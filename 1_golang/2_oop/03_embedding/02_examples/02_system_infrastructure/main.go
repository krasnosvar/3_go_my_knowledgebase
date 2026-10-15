package main

import "fmt"

type Metrics struct{}

func (Metrics) Emit(name string) { fmt.Println("metric:", name) }

type Worker struct{ Metrics }

func main() {
	w := Worker{}
	w.Emit("jobs_processed")
}

package main

import "fmt"

type Command interface {
	Run(args []string) error
}

type VersionCommand struct{}

func (VersionCommand) Run(args []string) error {
	fmt.Println("v1.0.0")
	return nil
}

func main() {
	registry := map[string]Command{
		"version": VersionCommand{},
	}
	_ = registry["version"].Run(nil)
}

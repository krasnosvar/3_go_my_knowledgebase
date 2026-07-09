package main

import "fmt"

type BaseCmd struct{}

func (BaseCmd) Log(name string) {
	fmt.Println("run:", name)
}

type VersionCmd struct{ BaseCmd }

func (v VersionCmd) Run() {
	v.Log("version")
	fmt.Println("v1.0.0")
}

func main() { VersionCmd{}.Run() }

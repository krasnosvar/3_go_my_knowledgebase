package main

import "fmt"

type user struct{ name string }

func makeUser(name string) *user {
	u := user{name: name}
	return &u // likely escape to heap
}

func main() {
	u := makeUser("den")
	fmt.Println(u.name)
}

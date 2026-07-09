package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type In struct {
	Name string `json:"name"`
}

func main() {
	var in In
	_ = json.NewDecoder(strings.NewReader(`{"name":"den"}`)).Decode(&in)
	fmt.Println(in.Name)
}

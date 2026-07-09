package main

import (
	"errors"
	"fmt"
)

func validate(age int) error {
	if age < 0 {
		return errors.New("age must be >= 0")
	}
	return nil
}

func main() { fmt.Println(validate(-1) != nil) }

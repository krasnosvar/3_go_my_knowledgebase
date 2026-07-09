package main

import (
	"errors"
	"fmt"
)

type MyErr struct{ Msg string }

func (e MyErr) Error() string { return e.Msg }

func main() {
	err := fmt.Errorf("ctx: %w", MyErr{Msg: "bad"})
	var me MyErr
	fmt.Println(errors.As(err, &me), me.Msg)
}

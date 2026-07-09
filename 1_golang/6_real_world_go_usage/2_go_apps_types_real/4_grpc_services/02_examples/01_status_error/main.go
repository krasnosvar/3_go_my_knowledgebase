package main

import "fmt"

type Code string

const (
	CodeOK       Code = "OK"
	CodeNotFound Code = "NOT_FOUND"
)

type RPCStatusError struct {
	Code    Code
	Message string
}

func (e RPCStatusError) Error() string {
	return fmt.Sprintf("rpc error: code=%s message=%q", e.Code, e.Message)
}

func findUser(id int) error {
	if id != 1 {
		return RPCStatusError{Code: CodeNotFound, Message: "user not found"}
	}
	return RPCStatusError{Code: CodeOK, Message: "ok"}
}

func main() {
	fmt.Println(findUser(2))
}

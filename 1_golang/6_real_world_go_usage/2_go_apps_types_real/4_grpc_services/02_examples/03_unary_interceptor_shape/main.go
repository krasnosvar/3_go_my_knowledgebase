package main

import "fmt"

func loggingInterceptor(method string, next func() error) error {
	fmt.Println("rpc:", method)
	return next()
}

func main() {
	_ = loggingInterceptor("/user.v1.UserService/Get", func() error {
		fmt.Println("handler executed")
		return nil
	})
}

package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Service struct{ s Sender }

func (s Service) Notify(msg string) error { return s.s.Send(msg) }

type mockSender struct{}

func (mockSender) Send(msg string) error {
	fmt.Println("sent:", msg)
	return nil
}

func main() {
	svc := Service{s: mockSender{}}
	_ = svc.Notify("hello")
}

package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Service struct {
	s Sender
}

func (s Service) Notify(msg string) error {
	// TODO: делегируй отправку через интерфейс Sender.
	return nil
}

func main() {
	fmt.Println("wire with mock sender")
}

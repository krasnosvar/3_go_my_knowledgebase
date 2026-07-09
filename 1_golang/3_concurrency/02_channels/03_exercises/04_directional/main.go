package main

import "fmt"

func sender(out chan<- int) {
	// TODO: отправь значение и закрой канал.
	_ = out
}

func receiver(in <-chan int) {
	// TODO: прочитай и выведи значения из канала.
	_ = in
}

func main() {
	fmt.Println("implement directional channels")
}

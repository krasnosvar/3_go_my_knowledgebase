package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		time.Sleep(10 * time.Millisecond)
		sig <- syscall.SIGTERM
	}()

	<-sig
	fmt.Println("graceful shutdown")
}

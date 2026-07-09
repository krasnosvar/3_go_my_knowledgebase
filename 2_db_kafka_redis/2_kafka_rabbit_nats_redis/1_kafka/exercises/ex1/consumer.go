package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "group-1",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("received: key=%s value=%s\n", string(m.Key), string(m.Value))
	}
}

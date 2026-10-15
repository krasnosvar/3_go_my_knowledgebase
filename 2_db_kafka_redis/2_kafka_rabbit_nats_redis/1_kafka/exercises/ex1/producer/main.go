package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	// создаём writer для топика "test-topic"
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello Kafka from Go!"),
		},
	)
	if err != nil {
		panic("could not write message " + err.Error())
	}
	fmt.Println("Message sent!")
}

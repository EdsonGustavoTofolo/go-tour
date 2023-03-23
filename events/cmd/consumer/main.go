package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-tour/events/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go func() {
		err := rabbitmq.Consume(ch, msgs)
		if err != nil {
			fmt.Printf("Error in consume: %v\n", err)
		}
	}()

	for msg := range msgs {
		fmt.Printf(string(msg.Body) + "\n")
		if err := msg.Ack(false); err != nil {
			fmt.Printf("Error in ack message: %v\n", err)
		}
	}
}

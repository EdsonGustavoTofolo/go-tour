package main

import "go-tour/events/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	if err := rabbitmq.Publish(ch, "Hello Goland World"); err != nil {
		panic(err)
	}
}

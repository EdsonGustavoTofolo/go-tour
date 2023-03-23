package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return channel
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"minhafila",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	err := ch.PublishWithContext(
		ctx,
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(body),
		})

	if err != nil {
		return err
	}

	return nil
}

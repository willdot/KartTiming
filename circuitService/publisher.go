package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Publisher will publish a message to somewhere
type Publisher interface {
	publishMessage(body []byte, keyName string) error
}

// MessagePublisher uses RabbitMQ to publish messages
type MessagePublisher struct {
	conn  *amqp.Connection
	chann *amqp.Channel
	queue amqp.Queue
}

func newRabbitMqPublisher() *MessagePublisher {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("failed to connected to RabbitMQ")
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}

	result := MessagePublisher{
		conn:  conn,
		chann: channel,
	}

	result.declareQueue()

	return &result
}

func (p *MessagePublisher) declareQueue() error {
	q, err := p.chann.QueueDeclare(
		"StartRace", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		return err
	}

	p.queue = q

	return nil
}

func (p *MessagePublisher) publishMessage(body []byte, keyName string) error {

	err := p.chann.Publish(
		"",
		keyName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	fmt.Println("message sent")
	return err
}

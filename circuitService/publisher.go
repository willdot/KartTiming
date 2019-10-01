package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type messagePublisher struct {
	conn  *amqp.Connection
	chann *amqp.Channel
	queue amqp.Queue
}

func newPublisher() *messagePublisher {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("failed to connected to RabbitMQ")
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}

	return &messagePublisher{
		conn:  conn,
		chann: channel,
	}
}

func (p *messagePublisher) DeclareQueue() error {
	q, err := p.chann.QueueDeclare(
		"race", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)

	if err != nil {
		return err
	}

	p.queue = q

	return nil
}

func (p *messagePublisher) PublishMessage(body []byte, keyName string) error {

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

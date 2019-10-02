package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

// MessageConsumer uses RabbitMQ to get messages about when to start a race session
type MessageConsumer struct {
	conn  *amqp.Connection
	chann *amqp.Channel
	queue amqp.Queue
	msgs  <-chan amqp.Delivery
}

func newRabbitMqConsumer() *MessageConsumer {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("failed to connected to RabbitMQ")
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}

	result := MessageConsumer{
		conn:  conn,
		chann: channel,
	}

	result.declareQueue()

	log.Println(result.queue.Name)

	msgs, err := result.chann.Consume(
		"StartRace", // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	result.msgs = msgs

	return &result
}

func (m *MessageConsumer) declareQueue() error {
	q, err := m.chann.QueueDeclare(
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

	m.queue = q

	return nil
}

func (m *MessageConsumer) getMessages() raceDetails {

	var data []byte

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for d := range m.msgs {
			log.Printf("Received a message: %s", d.Body)

			data = d.Body

			return
		}
	}()

	wg.Wait()

	fmt.Println("got racers")
	var rd raceDetails
	err := json.Unmarshal(data, &rd)

	if err != nil {
		failOnError(err, "failed to decode data")
	}

	return rd
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

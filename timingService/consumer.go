package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

func getRacers() raceDetails {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"StartRace", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var data []byte

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			data = d.Body
			d.Ack(false)

			return
		}
	}()

	wg.Wait()

	fmt.Println("got racers")
	var rd raceDetails
	err = json.Unmarshal(data, &rd)

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

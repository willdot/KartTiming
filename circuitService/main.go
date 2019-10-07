package main

import (
	"log"
	"net/http"
)

func main() {

	publisher := newRabbitMqPublisher()

	server := server{
		publisher: publisher,
	}

	http.HandleFunc("/start", server.StartSessionHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		failOnError(err, "failed")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

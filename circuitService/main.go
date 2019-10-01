package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	publisher := newPublisher()

	server := server{
		publisher: publisher,
	}

	http.HandleFunc("/start", server.StartSessionHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		failOnError(err, "failed")
	}

	racers := make([]Racer, 0)

	racersJSON, err := json.Marshal(racers)

	if err != nil {
		failOnError(err, "Failed to Convert struct to JSON")
	}

	err = publisher.PublishMessage([]byte(racersJSON), "race")

	if err != nil {
		failOnError(err, "Failed to send message")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

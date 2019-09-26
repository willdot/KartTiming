package main

import (
	"encoding/json"
	"log"

	guuid "github.com/google/uuid"
)

// Racer is a single racer that is currently on track
type Racer struct {
	ID         guuid.UUID
	KartNumber int
	Times      []float64
}

func main() {

	publisher := newPublisher()

	racers := createRacers(3)

	racersJSON, err := json.Marshal(racers)

	if err != nil {
		failOnError(err, "Failed to Convert struct to JSON")
	}

	err = publisher.PublishMessage([]byte(racersJSON), "race")

	if err != nil {
		failOnError(err, "Failed to send message")
	}
}

func createRacers(numberOfRacers int) []Racer {
	racers := make([]Racer, numberOfRacers)

	for i := 0; i < numberOfRacers; i++ {
		racer := Racer{
			ID:         guuid.New(),
			KartNumber: i + 1,
		}

		racers[i] = racer
	}

	return racers
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

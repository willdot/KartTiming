package main

import (
	guuid "github.com/google/uuid"
)

// Racer is a single racer that is currently on track
type Racer struct {
	ID         guuid.UUID
	KartNumber int
	Times      []float64
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

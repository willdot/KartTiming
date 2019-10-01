package main

import (
	"errors"

	guuid "github.com/google/uuid"
)

// raceDetails is the data that will be sent to the timing service to start timing
type raceDetails struct {
	SessionTime int
	Racers      []Racer
}

// Racer represents a single person as a racer
type Racer struct {
	ID         guuid.UUID
	KartNumber int
	Times      []float64
}

func (rd *raceDetails) validate() error {

	if rd.SessionTime == 0 {
		return errors.New("Session time must be greater than 0")
	}

	if len(rd.Racers) == 0 {
		return errors.New("There must be racers in the session")
	}
	return nil
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

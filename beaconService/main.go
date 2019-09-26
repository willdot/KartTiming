package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	racers := createRacers(2)

	sessionChannel := make(chan struct{})

	randomTimeGen := RandomTimeGenerator{
		MinLaptime: 5.000,
		MaxLaptime: 10.000,
	}

	raceSession := RaceSession{
		SessionChannel:      sessionChannel,
		Racers:              racers,
		SessionTime:         15,
		RandomTimeGenerator: randomTimeGen,
	}

	raceSession.Start()

}

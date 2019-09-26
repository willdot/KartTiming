package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	racers := createRacers(2)

	sessionChannel := make(chan struct{})

	raceSession := RaceSession{
		SessionChannel: sessionChannel,
		Racers:         racers,
		MinLaptime:     40.000,
		MaxLaptime:     45.000,
		SessionTime:    120,
	}

	raceSession.Start()

}

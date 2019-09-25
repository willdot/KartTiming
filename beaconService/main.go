package main

import (
	"fmt"
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
		MinLaptime:     1.000,
		MaxLaptime:     3.000,
		SessionTime:    5,
	}

	go raceSession.startSession()

	raceSession.startRacing()

	fmt.Println("End of session")

}

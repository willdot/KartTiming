package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var racers []Racer

	useMessageQueue := os.Getenv("MSG")

	if useMessageQueue == "YES" {
		fmt.Println("using message queue")
		racers = getRacers()

	} else {
		fmt.Println("using dummy data")
		racers = createRacers(2)
	}

	rand.Seed(time.Now().UnixNano())

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

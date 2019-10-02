package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var rd raceDetails

	useMessageQueue := os.Getenv("MSG")

	if useMessageQueue == "YES" {
		fmt.Println("using message queue")

		consumer := newRabbitMqConsumer()

		for {
			rd = consumer.getMessages()

			startSession(rd)
		}

	}
	fmt.Println("using dummy data")

	rd = raceDetails{
		SessionTime: 15,
	}

	rd.Racers = createRacers(2)

	startSession(rd)

}

func startSession(rd raceDetails) {

	rand.Seed(time.Now().UnixNano())

	sessionChannel := make(chan struct{})

	randomTimeGen := RandomTimeGenerator{
		MinLaptime: 5.000,
		MaxLaptime: 10.000,
	}

	raceSession := RaceSession{
		SessionChannel:      sessionChannel,
		Racers:              rd.Racers,
		SessionTime:         rd.SessionTime,
		RandomTimeGenerator: randomTimeGen,
	}

	raceSession.Start()

}

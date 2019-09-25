package main

import (
	"fmt"
	"math/rand"
	"time"

	guuid "github.com/google/uuid"
)

var (
	minLaptime = 1.000
	maxLaptime = 3.000

	sessionTime = 10
)

// Racer is a single racer that is currently on track
type Racer struct {
	ID         guuid.UUID
	KartNumber int
	LapCount   int
}

func createRacers(numberOfRacers int) []Racer {
	racers := make([]Racer, numberOfRacers)

	for i := 0; i < numberOfRacers; i++ {
		racer := Racer{
			ID:         guuid.New(),
			KartNumber: i + 1,
			LapCount:   0,
		}

		racers[i] = racer
	}

	return racers
}

func createRandomTime() float64 {

	return minLaptime + rand.Float64()*(maxLaptime-minLaptime)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	racers := createRacers(2)

	sessionFinished := make(chan bool)

	go startSession(sessionFinished)

	for _, racer := range racers {
		go func(r Racer) {
			race(&r)

		}(racer)
	}

	<-sessionFinished
	fmt.Println("End of session")
}

// This will keep going and logging a racers time after a random lap time
func race(racer *Racer) {
	for {
		randomLapTime := createRandomTime()
		time.Sleep(time.Duration(int(randomLapTime)) * time.Second)

		fmt.Printf("Kart: %v did a time of: %.3f\n", racer.KartNumber, randomLapTime)

		racer.LapCount = racer.LapCount + 1
	}

}

func startSession(ch chan bool) {

	fmt.Println("GO GO GO!")
	time.Sleep(time.Duration(sessionTime) * time.Second)

	ch <- true
}

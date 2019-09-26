package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RaceSession will start a session of karters racing, and log their lap times
type RaceSession struct {
	SessionChannel chan struct{}
	Racers         []Racer
	MinLaptime     float64
	MaxLaptime     float64
	SessionTime    int
}

// Start will start a race session
func (rs *RaceSession) Start() {

	fmt.Println(time.Now())

	go rs.startSession()

	rs.startRacing()

	fmt.Println(time.Now())
	fmt.Println("End of session")
}

func (rs *RaceSession) startRacing() {

	var wg sync.WaitGroup
	dataChannel := make(chan Racer, len(rs.Racers))

	wg.Add(len(rs.Racers))

	for _, racer := range rs.Racers {
		go func(r Racer) {

			dataChannel <- rs.race(&r)
			wg.Done()

		}(racer)
	}

	wg.Wait()

	for i := 0; i < cap(dataChannel); i++ {
		racer, ok := <-dataChannel
		if ok {
			fmt.Printf("Kart %v lapcount: %v\n", racer.KartNumber, racer.LapCount)
		}
	}

}

// This will keep going and logging a racers time after a random lap time
func (rs *RaceSession) race(racer *Racer) Racer {
	for {
		select {
		case <-rs.SessionChannel:
			return *racer
		default:
		}
		randomLapTime := rs.createRandomTime()
		time.Sleep(time.Duration(int(randomLapTime)) * time.Second)

		fmt.Printf("Kart: %v did a time of: %.3f\n", racer.KartNumber, randomLapTime)

		racer.LapCount = racer.LapCount + 1
	}
}

func (rs *RaceSession) startSession() {

	fmt.Println("GO GO GO!")
	time.Sleep(time.Duration(rs.SessionTime) * time.Second)

	fmt.Println("🏁 🏁 🏁 🏁 🏁")
	close(rs.SessionChannel)
}

func (rs *RaceSession) createRandomTime() float64 {

	return rs.MinLaptime + rand.Float64()*(rs.MaxLaptime-rs.MinLaptime)
}

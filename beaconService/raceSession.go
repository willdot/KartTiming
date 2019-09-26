package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// RaceSession will start a session of karters racing, and log their lap times
type RaceSession struct {
	SessionChannel      chan struct{}
	Racers              []Racer
	SessionTime         int
	RandomTimeGenerator randomTimeGenerator
	Printer             io.Writer
}

// Start will start a race session
func (rs *RaceSession) Start() {

	// Printer hasn't been set so use the Stdout
	if rs.Printer == nil {
		rs.Printer = os.Stdout
	}

	fmt.Fprintln(rs.Printer, time.Now())

	go rs.startSession()

	rs.startRacing()

	fmt.Fprintln(rs.Printer, time.Now())
	fmt.Fprintln(rs.Printer, "End of session")
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
			fmt.Fprintf(rs.Printer, "Kart %v lapcount: %v\n", racer.KartNumber, len(racer.Times))
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
		randomLapTime := rs.RandomTimeGenerator.CreateRandomTime()
		time.Sleep(time.Duration(int(randomLapTime)) * time.Second)

		fmt.Fprintf(rs.Printer, "Kart: %v did a time of: %.3f\n", racer.KartNumber, randomLapTime)

		racer.Times = append(racer.Times, randomLapTime)
	}
}

func (rs *RaceSession) startSession() {

	fmt.Fprintln(rs.Printer, "GO GO GO!")
	time.Sleep(time.Duration(rs.SessionTime) * time.Second)

	fmt.Fprintln(rs.Printer, "🏁 🏁 🏁 🏁 🏁")
	close(rs.SessionChannel)
}

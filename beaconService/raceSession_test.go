package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

type fakeRandomTimeGenerator struct {
	MinLaptime float64
	MaxLaptime float64
}

func (f fakeRandomTimeGenerator) CreateRandomTime() float64 {

	return 1
}

func TestStartSession(t *testing.T) {

	t.Run("Session starts and finishes on time", func(t *testing.T) {
		sessionChannel := make(chan struct{})

		rs := RaceSession{
			SessionChannel: sessionChannel,
			SessionTime:    5,
			Printer:        os.Stdout,
		}

		go rs.startSession()

		select {
		case <-rs.SessionChannel:
			return
		case <-time.After(time.Second * time.Duration(rs.SessionTime+1)):
			t.Error("Session channel should have stopped")
		}
	})

	t.Run("Session starts and prints the correct data to terminal", func(t *testing.T) {
		sessionChannel := make(chan struct{})

		buffer := &bytes.Buffer{}

		rs := RaceSession{
			SessionChannel: sessionChannel,
			SessionTime:    1,
			Printer:        buffer,
		}

		go rs.startSession()

		<-sessionChannel

		output := buffer.String()

		want := `GO GO GO!
🏁 🏁 🏁 🏁 🏁
`

		if output != want {
			t.Errorf("got '%s' want '%s'", output, want)
		}
	})
}

func TestRace(t *testing.T) {

	sessionChannel := make(chan struct{})

	rtg := fakeRandomTimeGenerator{
		MinLaptime: 1,
		MaxLaptime: 2,
	}

	buffer := &bytes.Buffer{}

	rs := RaceSession{
		SessionChannel:      sessionChannel,
		SessionTime:         5,
		Printer:             buffer,
		RandomTimeGenerator: rtg,
	}

	racer := Racer{
		KartNumber: 1,
	}

	go rs.startSession()

	finalRacer := rs.race(&racer)

	lines := strings.Split(buffer.String(), "\n")

	printedLaptimes := 0

	for _, i := range lines {
		if strings.HasPrefix(i, "Kart: 1") {
			printedLaptimes = printedLaptimes + 1
		}
	}

	if len(finalRacer.Times) != printedLaptimes {
		t.Errorf("racer completed '%v' laps but only '%v' laps were printed", finalRacer.LapCount, printedLaptimes)
	}

	if lines[0] != "GO GO GO!" {
		t.Error("first line is incorrect")
	}

	if lines[len(lines)-3] != "🏁 🏁 🏁 🏁 🏁" {
		t.Error("second to last line printed should be 🏁 🏁 🏁 🏁 🏁")
	}
}

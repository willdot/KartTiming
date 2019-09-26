package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

type fakeRandomTimeGenerator struct {
	MinLaptime float64
	MaxLaptime float64
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

func doSomething(w io.Writer) {
	fmt.Fprintln(w, "hello")
}

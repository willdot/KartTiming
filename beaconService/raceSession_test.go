package main

import (
	"testing"
	"time"
)

func TestCreateRandomTime(t *testing.T) {

	rs := RaceSession{
		MinLaptime:  1.000,
		MaxLaptime:  3.000,
		SessionTime: 5,
	}
	got := rs.createRandomTime()

	if got > rs.MaxLaptime {
		t.Errorf("time `%v` is bigger than the max laptime `%v`", got, rs.MaxLaptime)
	}

	if got < rs.MinLaptime {
		t.Errorf("time `%v` is smaller than the min laptime `%v`", got, rs.MinLaptime)
	}
}

func TestStartSession(t *testing.T) {

	sessionChannel := make(chan struct{})

	rs := RaceSession{
		SessionChannel: sessionChannel,
		SessionTime:    5,
	}

	go rs.startSession()

	select {
	case <-rs.SessionChannel:
		return
	case <-time.After(time.Second * time.Duration(rs.SessionTime+1)):
		t.Error("Session channel should have stopped")
	}

}

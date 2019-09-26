package main

import "testing"

func TestCreateRandomTime(t *testing.T) {

	rtg := RandomTimeGenerator{
		MinLaptime: 10,
		MaxLaptime: 15,
	}

	got := rtg.CreateRandomTime()

	if got < rtg.MinLaptime || got > rtg.MaxLaptime {
		t.Errorf("Random time '%v' is not within the min and max range '%v - %v'", got, rtg.MinLaptime, rtg.MaxLaptime)
	}
}

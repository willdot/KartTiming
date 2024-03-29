package main

import "testing"

func TestCreateRacers(t *testing.T) {

	testCases := []struct {
		Name           string
		NumberOfRacers int
	}{
		{
			Name:           "Create 1 racer",
			NumberOfRacers: 1,
		},
		{
			Name:           "Create 12 racers",
			NumberOfRacers: 12,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			racers := createRacers(test.NumberOfRacers)

			if len(racers) != test.NumberOfRacers {
				t.Errorf("got %v racers but expected %v racers", len(racers), test.NumberOfRacers)
			}

			for _, racer := range racers {
				if racer.KartNumber == 0 {
					t.Error("kart number is 0, which is incorrect")
				}
			}
		})
	}
}

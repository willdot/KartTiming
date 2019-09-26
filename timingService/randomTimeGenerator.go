package main

import "math/rand"

type randomTimeGenerator interface {
	CreateRandomTime() float64
}

// RandomTimeGenerator will generate random lap times
type RandomTimeGenerator struct {
	MinLaptime float64
	MaxLaptime float64
}

// CreateRandomTime will create a random time between a range
func (r RandomTimeGenerator) CreateRandomTime() float64 {
	return r.MinLaptime + rand.Float64()*(r.MaxLaptime-r.MinLaptime)
}

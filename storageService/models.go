package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Racer represents a single racer for MongoDB implementation
type Racer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string
	FastestLap float32
	Sessions   []session
}

// RacerSessions is data received from an HTTP request for storing a racers session
type RacerSession struct {
	ID      string
	Session session
}

type session struct {
	Date       time.Time
	LapTimes   []float32
	KartNumber int
}

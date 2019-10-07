package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type racer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string
	FastestLap float32
	Sessions   []session
}

type session struct {
	Date       time.Time
	LapTimes   []float32
	KartNumber int
}
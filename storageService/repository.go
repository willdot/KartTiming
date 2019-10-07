package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type storage interface {
	AddRacer(ctx context.Context, r *racer) error
}

// StorageService implements the storage interface and will allow data to be saved and read from MongoDB
type StorageService struct {
	client *mongo.Client
}

// AddRacer will add a single racer to the database returning a racer object with an ID or and error
func (s *StorageService) AddRacer(ctx context.Context, r *racer) error {

	collection := s.client.Database("Karting").Collection("Racers")

	insertResult, err := collection.InsertOne(ctx, *r)
	if err != nil {
		log.Println("error inserting racer: ", err)
		return err
	}

	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		r.ID = oid
	}

	return nil
}

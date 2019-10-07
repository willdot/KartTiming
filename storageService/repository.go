package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type storage interface {
	AddRacer(ctx context.Context, r *Racer) error
	AddSessionToRacer(ctx context.Context, ses session, racerID string) error
	GetRacer(ctx context.Context, name string) (Racer, error)
}

// StorageService implements the storage interface and will allow data to be saved and read from MongoDB
type StorageService struct {
	client *mongo.Client
}

// GetRacer will search the MongoDB collection for a racer by name and return that racer
func (s StorageService) GetRacer(ctx context.Context, name string) (Racer, error) {

	var r Racer

	collection := s.client.Database("Karting").Collection("Racers")

	filter := bson.M{"name": name}
	result := collection.FindOne(ctx, filter)

	err := result.Decode(&r)

	if err != nil {
		log.Println("error decoding racer: ", err)
		return r, err
	}

	return r, nil
}

// AddSessionToRacer will add a session to an existing racer
func (s StorageService) AddSessionToRacer(ctx context.Context, ses session, racerID string) error {

	collection := s.client.Database("Karting").Collection("Racers")

	id, err := primitive.ObjectIDFromHex(racerID)

	if err != nil {
		log.Println("error converting racer id to MongoDB id: ", err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$addToSet": bson.M{"Sessions": ses}}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Println("error adding session to racer: ", err)
		return err
	}

	return nil
}

// AddRacer will add a single racer to the database returning a racer object with an ID or and error
func (s StorageService) AddRacer(ctx context.Context, r *Racer) error {

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

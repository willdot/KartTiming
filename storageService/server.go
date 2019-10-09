package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type server struct {
	context        context.Context
	storageService storage
}

func (s *server) CreateNewRacerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newRacer Racer

		err := decodeBody(r, &newRacer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.storageService.AddRacer(s.context, &newRacer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newRacer)
	}
}

func (s *server) AddSessionData() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var data RacerSession

		err := decodeBody(r, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.storageService.AddSessionToRacer(s.context, data.Session, data.ID) 

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func decodeBody(r *http.Request, v interface{}) error {

	return json.NewDecoder(r.Body).Decode(v)
}

func createServer() server {

	ctx := context.Background()
	client := getClient(ctx)

	return server{
		context: ctx,
		storageService: StorageService{
			client: client,
		},
	}
}

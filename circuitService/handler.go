package main

import (
	"encoding/json"
	"net/http"
)

type server struct {
	publisher Publisher
}

// StartSessionHandler is a handler that will start a race session when called
func (s *server) StartSessionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var rd raceDetails

		err := decodeBody(r, &rd)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = rd.validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json, err := json.Marshal(rd)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		s.publisher.publishMessage([]byte(json), "StartRace")
	}
}

func decodeBody(r *http.Request, v interface{}) error {

	return json.NewDecoder(r.Body).Decode(v)
}

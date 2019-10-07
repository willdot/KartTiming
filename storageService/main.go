package main

import (
	"log"
	"net/http"
)

func main() {

	server := createServer()

	http.HandleFunc("/create-racer", server.CreateNewRacerHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

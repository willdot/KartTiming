package main

import (
	"log"
	"net/http"
)

func main() {

	server := createServer()

	http.HandleFunc("/create-racer", server.CreateNewRacerHandler())
	http.HandleFunc("/save-sessions", server.AddSessionDataHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

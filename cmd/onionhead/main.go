package main

import (
	"log"
	"net/http"

	"gitlab.com/EndowTheGreat/onionhead/tor"
	internalHTTP "gitlab.com/EndowTheGreat/onionhead/transport/http"
)

func main() {
	log.Println("Attempting to instantiate chat service...")
	service := tor.NewTorService()
	log.Printf("Instance live on: %s.onion\n", service.ID)
	defer service.Close()
	handler := internalHTTP.NewHandler()
	handler.Service = "http://%s.onion" + service.ID
	handler.SetupRoutes()
	server := &http.Server{
		Handler: handler.Router,
	}
	log.Fatal(server.Serve(service.LocalListener))
}

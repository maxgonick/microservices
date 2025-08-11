package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHelloHandler(logger)
	gh := handlers.NewGoodbyeHandler(logger)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	http.ListenAndServe(":8080", sm)

}

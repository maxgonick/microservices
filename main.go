package main

import (
	"context"
	"log"
	"microservice/product-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProductsHandler(logger)
	sm := http.NewServeMux()
	sm.Handle("/", productHandler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan

	logger.Println("Received signal:", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}

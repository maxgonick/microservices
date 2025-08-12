package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProductsHandler(logger)

	// ROUTING

	router := mux.NewRouter()

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := router.Methods("PUT").Subrouter()
	putRouter.Use(productHandler.MiddlewareProductValidation)
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)

	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

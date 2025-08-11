package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HelloHandler struct {
	logger *log.Logger
}

func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{logger: l}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World!")

	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	h.logger.Printf("Data: %s", d)

	fmt.Fprintf(w, "Hello %s\n", d)
}

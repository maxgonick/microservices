package handlers

import (
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	logger *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *GoodbyeHandler {
	return &GoodbyeHandler{logger: l}
}

func (h *GoodbyeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Goodbye World!")

	w.Write([]byte("Goodbye World!\n"))
}

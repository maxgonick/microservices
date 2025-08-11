package handlers

import (
	"encoding/json"
	"log"
	"microservice/product-api/data"
	"net/http"
)

type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger: l}
}

func (p *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}

	w.Write(data)



}

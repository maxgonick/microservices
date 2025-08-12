package handlers

import (
	"net/http"
	"product-api/data"
)

import _ "product-api/docs"


// @Summary Lists all products
// @Failure 500 {string} string "Internal Server Error"
// @Success 200 {object} productsResponse "List of all products"
// @Router / [get]
func (p *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Request")

	productList := data.GetProducts()

	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

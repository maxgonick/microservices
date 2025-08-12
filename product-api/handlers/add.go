package handlers

import (
	"fmt"
	"net/http"
	"product-api/data"
)

// @Summary Adds a product
// @Description Adds a new product to the list
// @Param product body data.Product true "Product to add"
// @Failure 400 {string} string "Bad Request"
// @Success 200 {string} string "Successfully added product"
// @Router / [post]
func (p *ProductsHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Request")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)

	fmt.Fprintf(w, "Succesfully added product %s\n", prod.Name)
	w.WriteHeader(http.StatusOK)
}

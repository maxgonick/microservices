package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Updates a product
// @Description Updates an existing product in the list
// @Param id path int true "Product ID"
// @Param product body data.Product true "Product data to update"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal Server Error"
// @Success 200
// @Router /{id} [put]
func (p *ProductsHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Request")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Bad URI", http.StatusBadRequest)
		return
	}

	prod, ok := r.Context().Value(productKey{}).(*data.Product)

	if !ok {
		p.logger.Println("Error: ", ok)
		http.Error(w, "Error retrieving product", http.StatusInternalServerError)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Error handling request", http.StatusInternalServerError)
		return
	}
}

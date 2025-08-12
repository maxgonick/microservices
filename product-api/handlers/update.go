package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

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

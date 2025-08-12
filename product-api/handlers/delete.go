package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *ProductsHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle DELETE Request")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Bad URI", http.StatusBadRequest)
		return
	}

	data.DeleteProduct(id)

}

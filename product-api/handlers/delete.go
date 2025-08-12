package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Deletes a product
// @Description Deletes a product from the list by its ID
// @Param id path int true "Product ID"
// @Failure 400 {string} string "Bad Request - Invalid ID"
// @Success 200
// @Router /{id} [delete]
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

package handlers

import (
	"context"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger: l}
}

func (p *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Request")

	productList := data.GetProducts()

	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *ProductsHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Request")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)

}

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

type productKey struct{}

func (p *ProductsHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			p.logger.Println("Error decoding body: ", err)
			http.Error(w, "Unable to decode request body", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), productKey{}, prod)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
)


// A list of products returns in the response
type productsResponse struct {
	// All products in the system
	Body []data.Product
}



type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger: l}
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

		err = prod.Validate()

		if err != nil {
			p.logger.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), productKey{}, prod)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

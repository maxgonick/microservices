package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger: l}
}

func (p *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		// Expect ID in the URI
		
		regex := regexp.MustCompile(`/([0-9]+)`)

		foundMatches := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(foundMatches) != 1 {
			http.Error(w, "Invalid URI more than one id present", http.StatusBadRequest)
			return
		}

		idString := foundMatches[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid URI bad id", http.StatusBadRequest)
		}

		p.logger.Println("ID: ", id)

		p.updateProduct(id, w, r)

	}


	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *ProductsHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Request")

	productList := data.GetProducts()

	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *ProductsHandler) addProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Request")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)

}

func (p *ProductsHandler) updateProduct(id int, w http.ResponseWriter, r *http.Request){
	p.logger.Println("Handle PUT Request")

	prod := &data.Product{}

	err :=	prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to decode request body", http.StatusBadRequest)
		return
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
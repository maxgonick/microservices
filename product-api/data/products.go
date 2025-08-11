package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreateOn:    time.Now().Local().String(),
		UpdatedOn:   time.Now().Local().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Small and intense coffee concentrate without milk",
		Price:       1.99,
		SKU:         "xyz456",
		CreateOn:    time.Now().Local().String(),
		UpdatedOn:   time.Now().Local().String(),
	},
	{
		ID: 1, Name: "tea", Description: "A nice cup of tea", Price: 0, SKU: "", CreateOn: "", UpdatedOn: "", DeletedOn: "",
	},
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return (lastProduct.ID + 1)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

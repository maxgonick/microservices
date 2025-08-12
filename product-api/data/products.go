package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreateOn    string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	//SKU format is abc-abcd-abcdf
	sku_format := `[a-z]{3}-[a-z]{4}-[a-z]{5}`
	re := regexp.MustCompile(sku_format)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
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

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for index, product := range productList {
		if product.ID == id {
			return product, index, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, index, err := findProduct(id)

	if err != nil {
		return err
	}

	productList[index] = p
	return nil
}

func DeleteProduct(id int) error {
	_, index, err := findProduct(id)

	if err != nil {
		return err
	}	

	productList = append(productList[:index], productList[index+1:]...)
	
	return nil
}

func ResetProducts() {
	productList = []*Product{}
}

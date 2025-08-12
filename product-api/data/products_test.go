package data

import "testing"

func TestChecksValidation(t *testing.T) {
	product := &Product{
		Name:  "Koolaid",
		Price: 2.50,
		SKU:   "qwe-qwer-qwert",
	}

	err := product.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteProduct(t *testing.T) {
	
}

func TestMain(m *testing.M) {
	
}
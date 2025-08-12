package data

import (
	"testing"
)

// This test is self-contained and has no side effects, so it's fine as is.
func TestChecksValidation(t *testing.T) {
	product := &Product{
		Name:  "Koolaid",
		Price: 2.50,
		SKU:   "qwe-qwer-qwert",
	}
	if err := product.Validate(); err != nil {
		t.Fatal(err)
	}
}

// TestProductCRUD groups all tests that modify the shared product list.
func TestProductCRUD(t *testing.T) {
	// This helper function resets the state before each subtest.
	// NOTE: You need to create a `ResetProducts()` function in your data package
	// that clears the underlying product list.
	setup := func(t *testing.T) {
		t.Helper()
		ResetProducts()
	}

	t.Run("AddProduct should add a new product to the list", func(t *testing.T) {
		setup(t) // Reset the state for this specific test

		// Arrange
		p := &Product{Name: "Green Tea", Price: 1.99, SKU: "abc-def-ghi"}

		// Act
		AddProduct(p)

		// Assert
		if len(GetProducts()) != 1 {
			t.Errorf("Expected product list to have 1 item, but it has %d", len(GetProducts()))
		}
		if GetProducts()[0].Name != "Green Tea" {
			t.Errorf("Expected the added product to be 'Green Tea', but got '%s'", GetProducts()[0].Name)
		}
	})

	t.Run("DeleteProduct should remove a product from the list", func(t *testing.T) {
		setup(t) // Reset the state for this specific test

		// Arrange: Add a product first so we can delete it.
		p := &Product{Name: "To be Deleted", Price: 10.00, SKU: "del-ete-me"}
		AddProduct(p)
		
        // Assuming the product gets ID 1
		productIDToDelete := 1 
		
		// Act
		err := DeleteProduct(productIDToDelete)

		// Assert
		if err != nil {
			t.Fatalf("Expected no error on delete, but got %v", err)
		}
		if len(GetProducts()) != 0 {
			t.Errorf("Expected product list to be empty after delete, but it has %d items", len(GetProducts()))
		}
	})
    
    t.Run("DeleteProduct should return an error for a non-existent ID", func(t *testing.T) {
        setup(t) // Reset state
        
        // Act
        err := DeleteProduct(999) // An ID that does not exist
        
        // Assert
        if err == nil {
            t.Fatal("Expected an error when deleting a non-existent product, but got nil")
        }
    })
}
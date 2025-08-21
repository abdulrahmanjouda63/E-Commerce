package factory_test

import (
	"ecommerce/factory"
	"ecommerce/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	p := factory.CreateProduct(1, "E001", "Laptop", 1000.0, 10, 24)
	assert.IsType(t, &models.Electronics{}, p, "Should create Electronics")
	assert.Equal(t, "Laptop", p.GetName(), "Name should be set")

	p = factory.CreateProduct(2, "C001", "Shirt", 30.0, 5, []string{"S", "M"})
	assert.IsType(t, &models.Clothing{}, p, "Should create Clothing")

	p = factory.CreateProduct(3, "B001", "Book", 20.0, 10, nil)
	assert.IsType(t, &models.Books{}, p, "Should create Books")

	p = factory.CreateProduct(999, "X001", "Invalid", 0.0, 0, nil)
	assert.Nil(t, p, "Should return nil for invalid type")
}
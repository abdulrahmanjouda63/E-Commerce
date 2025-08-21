package models_test

import (
	"ecommerce/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p := models.NewProduct("P001", "Test Product", 100.0, -5)
	assert.Equal(t, "Test Product", p.GetName(), "Name should be set correctly")
	assert.Equal(t, 100.0, p.GetPrice(), "Price should be set correctly")
	assert.True(t, p.IsInStock(), "Stock should default to 10 for negative input, so IsInStock should be true")
}

func TestProduct_UpdateStock(t *testing.T) {
	p := models.NewProduct("P002", "Test Product", 50.0, 10)
	p.UpdateStock(-5)
	assert.True(t, p.IsInStock(), "Stock should be 5 after update")
	p.UpdateStock(-10)
	assert.False(t, p.IsInStock(), "Stock should be 0 after excessive reduction")
}

func TestProduct_CalculateOrderPrice(t *testing.T) {
	p := models.NewProduct("P003", "Test Product", 20.0, 10)
	assert.Equal(t, 40.0, p.CalculateOrderPrice(2), "Should calculate price correctly")
	assert.Equal(t, 0.0, p.CalculateOrderPrice(0), "Should return 0 for zero quantity")
}
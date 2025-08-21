package models_test

import (
	"ecommerce/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBooks_CalculateOrderPrice(t *testing.T) {
	b := models.NewBooks("B001", "Book", 20.0, 10)
	assert.Equal(t, 30.0, b.CalculateOrderPrice(2), "Should apply buy-one-get-one 50% off")
}

func TestBooks_IsInStock(t *testing.T) {
	b := models.NewBooks("B002", "E-Book", 15.0, 0)
	assert.True(t, b.IsInStock(), "Digital books should always be in stock")
}
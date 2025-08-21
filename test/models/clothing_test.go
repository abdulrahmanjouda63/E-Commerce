package models_test

import (
	"ecommerce/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClothing(t *testing.T) {
	c := models.NewClothing("C001", "Shirt", 30.0, 10, nil)
	assert.Equal(t, []string{"M"}, c.GetSizes(), "Sizes should default to [M]")
}

func TestClothing_CalculateOrderPrice(t *testing.T) {
	c := models.NewClothing("C002", "Jacket", 100.0, 5, []string{"S", "M"})
	assert.Equal(t, 475.0, c.CalculateOrderPrice(5), "Should apply 5% discount for quantity >= 5")
}
package models_test

import (
	"ecommerce/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewElectronics(t *testing.T) {
	e := models.NewElectronics("E001", "Laptop", 1000.0, 10, -1)
	assert.Equal(t, 0, e.GetWarranty(), "Warranty should default to 0")
	assert.Equal(t, 1000.0, e.GetPrice(), "Price should be base price for warranty <= 12")
}

func TestElectronics_GetPrice(t *testing.T) {
	e := models.NewElectronics("E002", "Phone", 500.0, 5, 24)
	assert.Equal(t, 550.0, e.GetPrice(), "Price should have 10% premium for warranty > 12")
}

func TestElectronics_CalculateOrderPrice(t *testing.T) {
	e := models.NewElectronics("E003", "Tablet", 300.0, 5, 6)
	assert.Equal(t, 610.0, e.CalculateOrderPrice(2), "Should include shipping fee for quantity > 1")
}
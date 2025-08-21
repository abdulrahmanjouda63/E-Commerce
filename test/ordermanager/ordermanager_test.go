package ordermanager_test

import (
	"ecommerce/models"
	"ecommerce/ordermanager"
	"ecommerce/payments"
	"ecommerce/reports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderManager_AddToCart(t *testing.T) {
	om := ordermanager.NewOrderManager(&models.Product{}, &models.Product{})
	p := models.NewProduct("P001", "Test", 100.0, 10)
	om.AddToCart(p, 2)
	assert.Equal(t, 200.0, om.CalculateTotal(), "Total should reflect added items")
}

func TestOrderManager_ProcessPayment(t *testing.T) {
	om := ordermanager.NewOrderManager(&models.Product{}, &models.Product{})
	rs := reports.NewReportService()
	p := models.NewBooks("B001", "Book", 20.0, 10)
	om.AddToCart(p, 2)
	success, _ := om.ProcessPayment(&payments.PayPalPayment{}, rs)
	assert.True(t, success, "Payment should succeed")
	assert.Equal(t, 30.0, rs.GetSales()["Book"], "Sale should be recorded")
}
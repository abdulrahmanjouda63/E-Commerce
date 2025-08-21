package ordermanager

import (
	"ecommerce/models"
	"ecommerce/payments"
	"ecommerce/reports"
	"fmt"
)

// OrderManager handles cart and order processing
type OrderManager struct {
	cart     []struct {
		product  models.ProductCategory
		quantity int
	}
	storable models.IStorable
	pricable models.IPricable
}

// NewOrderManager with dependency injection
func NewOrderManager(storable models.IStorable, pricable models.IPricable) *OrderManager {
	return &OrderManager{
		cart:     make([]struct{ product models.ProductCategory; quantity int }, 0),
		storable: storable,
		pricable: pricable,
	}
}

// AddToCart adds product with quantity
func (om *OrderManager) AddToCart(product models.ProductCategory, quantity int) {
	if quantity <= 0 {
		fmt.Println("Error: Quantity must be positive")
		return
	}
	om.cart = append(om.cart, struct {
		product  models.ProductCategory
		quantity int
	}{product, quantity})
}

// CalculateTotal computes cart total
func (om *OrderManager) CalculateTotal() float64 {
	total := 0.0
	for _, item := range om.cart {
		total += item.product.ApplyCategoryRules(item.quantity)
	}
	return total
}

// CalculateTotalWithDiscount applies discount to total
func (om *OrderManager) CalculateTotalWithDiscount(discount interface{ Apply(float64) float64 }) float64 {
	total := om.CalculateTotal()
	return discount.Apply(total)
}

// ProcessPayment processes payment and updates stock on success
func (om *OrderManager) ProcessPayment(payment payments.Payment, reportService *reports.ReportService) (bool, string) {
	total := om.CalculateTotal()
	if total == 0 {
		return false, "Error: Empty cart"
	}
	success, message := payment.Process(total)
	if success {
		for _, item := range om.cart {
			// Update stock if storable
			if storable, ok := item.product.(models.IStorable); ok {
				storable.UpdateStock(-item.quantity)
			}
			// Record sale
			name := item.product.GetName()
			amount := item.product.ApplyCategoryRules(item.quantity)
			reportService.RecordSale(name, amount)
		}
		om.cart = nil // Clear cart
	}
	return success, message
}
package payments

import "fmt"

// Payment interface for processing payments
type Payment interface {
	Process(amount float64) (bool, string)
}

// CreditCardPayment implementation
type CreditCardPayment struct {
	limit float64
}

// NewCreditCardPayment constructor
func NewCreditCardPayment(limit float64) *CreditCardPayment {
	if limit <= 0 {
		limit = 5000 // Default limit
	}
	return &CreditCardPayment{limit: limit}
}

// Process payment with limit check
func (c *CreditCardPayment) Process(amount float64) (bool, string) {
	if amount > c.limit {
		return false, fmt.Sprintf("CreditCard: Amount $%.2f exceeds limit $%.2f", amount, c.limit)
	}
	return true, fmt.Sprintf("CreditCard: Processed $%.2f successfully", amount)
}

// PayPalPayment implementation
type PayPalPayment struct{}

// Process payment with 2% fee
func (p *PayPalPayment) Process(amount float64) (bool, string) {
	fee := amount * 0.02
	return true, fmt.Sprintf("PayPal: Processed $%.2f with $%.2f fee", amount, fee)
}
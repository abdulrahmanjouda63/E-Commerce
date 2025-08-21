package payments_test

import (
	"ecommerce/payments"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardPayment_Process(t *testing.T) {
	p := payments.NewCreditCardPayment(100.0)
	success, msg := p.Process(50.0)
	assert.True(t, success, "Payment should succeed within limit")
	assert.Contains(t, msg, "successfully", "Message should indicate success")

	success, msg = p.Process(150.0)
	assert.False(t, success, "Payment should fail above limit")
	assert.Contains(t, msg, "exceeds limit", "Message should indicate failure")
}

func TestPayPalPayment_Process(t *testing.T) {
	p := &payments.PayPalPayment{}
	success, msg := p.Process(100.0)
	assert.True(t, success, "PayPal should always succeed")
	assert.Contains(t, msg, "2.00 fee", "Message should include fee")
}
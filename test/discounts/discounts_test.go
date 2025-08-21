package discounts_test

import (
	"ecommerce/discounts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoDiscount_Apply(t *testing.T) {
	d := &discounts.NoDiscount{}
	assert.Equal(t, 100.0, d.Apply(100.0), "NoDiscount should return original price")
}

func TestSeasonalDiscount_Apply(t *testing.T) {
	d := &discounts.SeasonalDiscount{}
	assert.Equal(t, 90.0, d.Apply(100.0), "SeasonalDiscount should apply 10% off")
}

func TestLoyaltyDiscount_Apply(t *testing.T) {
	d := discounts.NewLoyaltyDiscount(20.0)
	assert.Equal(t, 80.0, d.Apply(100.0), "LoyaltyDiscount should apply 20% off")
}

func TestGhostDiscount_Apply(t *testing.T) {
	d := discounts.NewGhostDiscount(true)
	assert.Equal(t, 95.0, d.Apply(100.0), "GhostDiscount should apply 5% when condition true")
	d = discounts.NewGhostDiscount(false)
	assert.Equal(t, 90.0, d.Apply(100.0), "GhostDiscount should apply 10% when condition false")
}
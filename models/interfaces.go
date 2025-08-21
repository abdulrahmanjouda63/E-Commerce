package models

// IStorable interface for stock-related methods
type IStorable interface {
	UpdateStock(amount int)
	IsInStock() bool
}

// IPricable interface for price-related methods
type IPricable interface {
	GetPrice() float64
	CalculateOrderPrice(quantity int) float64
}

// BasicPricable interface for simple pricing
type BasicPricable interface {
	GetPrice() float64
}

// ProductCategory interface for extensible pricing rules
type ProductCategory interface {
	ApplyCategoryRules(quantity int) float64
	GetName() string
}
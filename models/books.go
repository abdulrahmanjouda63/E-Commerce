package models

// Books struct with embedded Product
type Books struct {
	*Product
}

// NewBooks constructor
func NewBooks(id, name string, price float64, stock int) *Books {
	return &Books{Product: NewProduct(id, name, price, stock)}
}

// CalculateOrderPrice for Books with buy-one-get-one 50% off
func (b *Books) CalculateOrderPrice(quantity int) float64 {
	if quantity <= 0 {
		return 0
	}
	// Pay full price for floor(quantity/2) pairs, half for remainder
	pairs := quantity / 2
	remainder := quantity % 2
	return (float64(pairs) * b.Product.price * 1.5) + (float64(remainder) * b.Product.price)
}

// ApplyCategoryRules implements ProductCategory for Books
func (b *Books) ApplyCategoryRules(quantity int) float64 {
	return b.CalculateOrderPrice(quantity)
}

// UpdateStock overridden for digital books (no-op)
func (b *Books) UpdateStock(amount int) {
	// No stock change for digital books
}

// IsInStock overridden for digital books
func (b *Books) IsInStock() bool {
	return true // Always available
}
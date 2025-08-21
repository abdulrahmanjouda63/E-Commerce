package models

// Clothing struct with embedded Product
type Clothing struct {
	*Product
	sizes []string
}

// NewClothing constructor
func NewClothing(id, name string, price float64, stock int, sizes []string) *Clothing {
	if len(sizes) == 0 {
		sizes = []string{"M"} // Default size
	}
	return &Clothing{
		Product: NewProduct(id, name, price, stock),
		sizes:   sizes,
	}
}

// GetSizes returns available sizes
func (c *Clothing) GetSizes() []string {
	return c.sizes
}

// CalculateOrderPrice for Clothing with bulk discount
func (c *Clothing) CalculateOrderPrice(quantity int) float64 {
	if quantity <= 0 {
		return 0
	}
	price := c.Product.price * float64(quantity)
	if quantity >= 5 {
		price *= 0.95 // 5% discount
	}
	return price
}

// ApplyCategoryRules implements ProductCategory for Clothing
func (c *Clothing) ApplyCategoryRules(quantity int) float64 {
	return c.CalculateOrderPrice(quantity)
}
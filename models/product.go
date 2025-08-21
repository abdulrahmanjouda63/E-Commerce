package models

// Product represents a basic product with private fields
type Product struct {
	id    string
	name  string
	price float64
	stock int
}

// NewProduct constructor for Product
func NewProduct(id, name string, price float64, stock int) *Product {
	if stock < 0 {
		stock = 10 // Default stock
	}
	return &Product{id: id, name: name, price: price, stock: stock}
}

// GetName returns product name
func (p *Product) GetName() string {
	return p.name
}

// GetPrice returns product price (implements BasicPricable)
func (p *Product) GetPrice() float64 {
	return p.price
}

// UpdateStock adjusts stock, ensuring non-negative (implements IStorable)
func (p *Product) UpdateStock(amount int) {
	p.stock += amount
	if p.stock < 0 {
		p.stock = 0
	}
}

// IsInStock checks if stock > 0 (implements IStorable)
func (p *Product) IsInStock() bool {
	return p.stock > 0
}

// CalculateOrderPrice default implementation (implements IPricable)
func (p *Product) CalculateOrderPrice(quantity int) float64 {
	if quantity <= 0 {
		return 0
	}
	return p.price * float64(quantity)
}

// ApplyCategoryRules default implementation (implements ProductCategory)
func (p *Product) ApplyCategoryRules(quantity int) float64 {
	return p.CalculateOrderPrice(quantity)
}
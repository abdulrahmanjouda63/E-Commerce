package models

// Electronics struct with embedded Product
type Electronics struct {
	*Product
	warrantyMonths int
}

// NewElectronics constructor
func NewElectronics(id, name string, price float64, stock, warrantyMonths int) *Electronics {
	if warrantyMonths < 0 {
		warrantyMonths = 0
	}
	return &Electronics{
		Product:        NewProduct(id, name, price, stock),
		warrantyMonths: warrantyMonths,
	}
}

// GetWarranty returns warranty duration
func (e *Electronics) GetWarranty() int {
	return e.warrantyMonths
}

// GetPrice overrides to apply 10% premium for warranty > 12 months
func (e *Electronics) GetPrice() float64 {
	if e.warrantyMonths > 12 {
		return e.Product.price * 1.10
	}
	return e.Product.price
}

// CalculateOrderPrice for Electronics with shipping fee
func (e *Electronics) CalculateOrderPrice(quantity int) float64 {
	if quantity <= 0 {
		return 0
	}
	price := e.GetPrice() * float64(quantity)
	if quantity > 1 {
		price += 10 // Shipping fee
	}
	return price
}

// ApplyCategoryRules implements ProductCategory for Electronics
func (e *Electronics) ApplyCategoryRules(quantity int) float64 {
	return e.CalculateOrderPrice(quantity)
}
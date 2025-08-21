package discounts

// Discount interface for applying discounts
type Discount interface {
	Apply(price float64) float64
}

// NoDiscount base case
type NoDiscount struct{}

// Apply returns original price
func (n *NoDiscount) Apply(price float64) float64 {
	return price
}

// SeasonalDiscount applies 10% off
type SeasonalDiscount struct{}

// Apply discounts price
func (s *SeasonalDiscount) Apply(price float64) float64 {
	return price * 0.9
}

// LoyaltyDiscount with variable percentage
type LoyaltyDiscount struct {
	percentage float64
}

// NewLoyaltyDiscount constructor
func NewLoyaltyDiscount(percentage float64) *LoyaltyDiscount {
	if percentage < 0 || percentage > 100 {
		percentage = 10 // Default
	}
	return &LoyaltyDiscount{percentage: percentage}
}

// Apply discounts price
func (l *LoyaltyDiscount) Apply(price float64) float64 {
	return price * (1 - l.percentage/100)
}

// GhostDiscount halves discount if conditions met
type GhostDiscount struct {
	condition bool
}

// NewGhostDiscount constructor
func NewGhostDiscount(condition bool) *GhostDiscount {
	return &GhostDiscount{condition: condition}
}

// Apply halves discount if condition true
func (g *GhostDiscount) Apply(price float64) float64 {
	discount := 0.1 // Default 10%
	if g.condition {
		discount /= 2
	}
	return price * (1 - discount)
}
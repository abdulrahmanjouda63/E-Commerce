package reports

import (
	"fmt"
	"sort"
	"strings"
)

// SalesDataProvider interface for sales data access
type SalesDataProvider interface {
	GetSales() map[string]float64
	RecordSale(name string, amount float64)
}

// ReportService generates sales reports and tracks sales
type ReportService struct {
	sales map[string]float64
}

// NewReportService constructor
func NewReportService() *ReportService {
	return &ReportService{sales: make(map[string]float64)}
}

// RecordSale records a sale (implements SalesDataProvider)
func (rs *ReportService) RecordSale(name string, amount float64) {
	rs.sales[name] += amount
}

// GetSales returns sales data (implements SalesDataProvider)
func (rs *ReportService) GetSales() map[string]float64 {
	return rs.sales
}

// GenerateSalesReport returns sorted sales report
func (rs *ReportService) GenerateSalesReport() string {
	if len(rs.sales) == 0 {
		return "No sales data available"
	}
	type sale struct {
		name   string
		amount float64
	}
	var sales []sale
	for name, amount := range rs.sales {
		sales = append(sales, sale{name, amount})
	}
	sort.Slice(sales, func(i, j int) bool {
		return sales[i].amount > sales[j].amount
	})
	var report strings.Builder
	report.WriteString("Sales Report:\n")
	for _, s := range sales {
		report.WriteString(fmt.Sprintf("Product: %s, Total Sales: $%.2f\n", s.name, s.amount))
	}
	return report.String()
}
package reports_test

import (
	"ecommerce/reports"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportService_GenerateSalesReport(t *testing.T) {
	rs := reports.NewReportService()
	rs.RecordSale("Product1", 100.0)
	rs.RecordSale("Product2", 50.0)
	report := rs.GenerateSalesReport()
	assert.Contains(t, report, "Product1, Total Sales: $100.00", "Should include Product1")
	assert.Contains(t, report, "Product2, Total Sales: $50.00", "Should include Product2")
	assert.True(t, strings.Index(report, "Product1") < strings.Index(report, "Product2"), "Should be sorted by amount")
}
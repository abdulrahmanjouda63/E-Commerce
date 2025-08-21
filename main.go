package main

import (
	"ecommerce/models"
	"ecommerce/ordermanager"
	"ecommerce/reports"
	"ecommerce/ui"
)

func main() {
	uiInstance := ui.NewConsoleUI()
	reportService := reports.NewReportService()
	manager := ordermanager.NewOrderManager(&models.Product{}, &models.Product{}) // Default injections
	uiInstance.RunApp(manager, reportService)
}
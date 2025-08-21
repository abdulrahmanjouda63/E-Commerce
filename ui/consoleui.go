package ui

import (
	"bufio"
	"ecommerce/discounts"
	"ecommerce/factory"
	"ecommerce/ordermanager"
	"ecommerce/payments"
	"ecommerce/reports"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConsoleUI handles user input and output
type ConsoleUI struct {
	Scanner *bufio.Scanner
}

// NewConsoleUI constructor
func NewConsoleUI() *ConsoleUI {
	return &ConsoleUI{Scanner: bufio.NewScanner(os.Stdin)}
}

// GetIntInput prompts for integer input
func (ui *ConsoleUI) GetIntInput(prompt string) int {
	fmt.Print(prompt)
	ui.Scanner.Scan()
	val, _ := strconv.Atoi(ui.Scanner.Text())
	return val
}

// GetFloatInput prompts for float input
func (ui *ConsoleUI) GetFloatInput(prompt string) float64 {
	fmt.Print(prompt)
	ui.Scanner.Scan()
	val, _ := strconv.ParseFloat(ui.Scanner.Text(), 64)
	return val
}

// GetStringInput prompts for string input
func (ui *ConsoleUI) GetStringInput(prompt string) string {
	fmt.Print(prompt)
	ui.Scanner.Scan()
	return ui.Scanner.Text()
}

// GetSizesInput prompts for sizes
func (ui *ConsoleUI) GetSizesInput() []string {
	sizesStr := ui.GetStringInput("Sizes (comma-separated): ")
	sizes := strings.Split(sizesStr, ",")
	for i := range sizes {
		sizes[i] = strings.TrimSpace(sizes[i])
	}
	return sizes
}

// GetBoolInput prompts for bool input
func (ui *ConsoleUI) GetBoolInput(prompt string) bool {
	val := ui.GetStringInput(prompt)
	boolVal, _ := strconv.ParseBool(val)
	return boolVal
}

// DisplayMessage displays a message
func (ui *ConsoleUI) DisplayMessage(msg string) {
	fmt.Println(msg)
}

// RunApp runs the application
func (ui *ConsoleUI) RunApp(manager *ordermanager.OrderManager, reportService *reports.ReportService) {
	for {
		ui.DisplayMessage("\n1. Add Product to Cart")
		ui.DisplayMessage("2. Calculate Total")
		ui.DisplayMessage("3. Process Payment")
		ui.DisplayMessage("4. Generate Sales Report")
		ui.DisplayMessage("5. Exit")
		choice := ui.GetIntInput("Choose option: ")

		if choice == 5 {
			break
		}

		switch choice {
		case 1:
			typeChoice := ui.GetIntInput("Product Type (1=Electronics, 2=Clothing, 3=Books): ")
			id := ui.GetStringInput("ID: ")
			name := ui.GetStringInput("Name: ")
			price := ui.GetFloatInput("Price: ")
			stock := ui.GetIntInput("Stock: ")
			quantity := ui.GetIntInput("Quantity: ")

			var extra interface{}
			switch typeChoice {
			case 1:
				extra = ui.GetIntInput("Warranty Months: ")
			case 2:
				extra = ui.GetSizesInput()
			}

			product := factory.CreateProduct(typeChoice, id, name, price, stock, extra)
			if product != nil {
				manager.AddToCart(product, quantity)
				ui.DisplayMessage(fmt.Sprintf("Added %s to cart", name))
			}

		case 2:
			discChoice := ui.GetIntInput("Apply Discount (1=None, 2=Seasonal, 3=Loyalty, 4=Ghost): ")
			var discount discounts.Discount
			switch discChoice {
			case 1:
				discount = &discounts.NoDiscount{}
			case 2:
				discount = &discounts.SeasonalDiscount{}
			case 3:
				percentage := ui.GetFloatInput("Loyalty Discount Percentage: ")
				discount = discounts.NewLoyaltyDiscount(percentage)
			case 4:
				condition := ui.GetBoolInput("Ghost Discount Condition (true/false): ")
				discount = discounts.NewGhostDiscount(condition)
			default:
				discount = &discounts.NoDiscount{}
			}
			total := manager.CalculateTotalWithDiscount(discount)
			ui.DisplayMessage(fmt.Sprintf("Cart Total: $%.2f", total))

		case 3:
			payChoice := ui.GetIntInput("Payment Method (1=CreditCard, 2=PayPal): ")
			var payment payments.Payment
			switch payChoice {
			case 1:
				limit := ui.GetFloatInput("Credit Card Limit: ")
				payment = payments.NewCreditCardPayment(limit)
			case 2:
				payment = &payments.PayPalPayment{}
			default:
				ui.DisplayMessage("Invalid payment method")
				continue
			}
			success, message := manager.ProcessPayment(payment, reportService)
			ui.DisplayMessage(message)
			if success {
				ui.DisplayMessage("Order processed, stock updated")
			}

		case 4:
			ui.DisplayMessage(reportService.GenerateSalesReport())
		}
	}
}
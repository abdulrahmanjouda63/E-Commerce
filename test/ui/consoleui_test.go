package ui_test

import (
	"bytes"
	"ecommerce/models"
	"ecommerce/ordermanager"
	"ecommerce/reports"
	"ecommerce/ui"
	"io"
	"os"
	"strings"
	"testing"

	"bufio"
	"github.com/stretchr/testify/assert"
)

func TestConsoleUI_GetIntInput(t *testing.T) {
	input := "42\n"
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader) // Set scanner for testing
	result := scanner.GetIntInput("Enter number: ")
	assert.Equal(t, 42, result, "Should parse integer input correctly")
}

func TestConsoleUI_GetFloatInput(t *testing.T) {
	input := "99.99\n"
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader)
	result := scanner.GetFloatInput("Enter price: ")
	assert.Equal(t, 99.99, result, "Should parse float input correctly")
}

func TestConsoleUI_GetStringInput(t *testing.T) {
	input := "Test Product\n"
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader)
	result := scanner.GetStringInput("Enter name: ")
	assert.Equal(t, "Test Product", result, "Should parse string input correctly")
}

func TestConsoleUI_GetSizesInput(t *testing.T) {
	input := "S, M, L\n"
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader)
	result := scanner.GetSizesInput()
	assert.Equal(t, []string{"S", "M", "L"}, result, "Should parse comma-separated sizes")
}

func TestConsoleUI_GetBoolInput(t *testing.T) {
	input := "true\n"
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader)
	result := scanner.GetBoolInput("Enter condition: ")
	assert.True(t, result, "Should parse boolean input correctly")
}

func TestConsoleUI_DisplayMessage(t *testing.T) {
	scanner := ui.NewConsoleUI()
	
	// Capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	scanner.DisplayMessage("Hello, World!")
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	assert.Contains(t, buf.String(), "Hello, World!", "Should print message to stdout")
}

func TestConsoleUI_RunApp_AddToCart(t *testing.T) {
	// Mock input for adding a product (Electronics)
	input := strings.Join([]string{
		"1\n",      // Select "Add Product"
		"1\n",      // Product type: Electronics
		"E001\n",   // ID
		"Laptop\n", // Name
		"1000.0\n", // Price
		"10\n",     // Stock
		"2\n",      // Quantity
		"24\n",     // Warranty Months
		"5\n",      // Exit
	}, "")
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader)
	
	// Capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Setup dependencies
	om := ordermanager.NewOrderManager(&models.Product{}, &models.Product{})
	rs := reports.NewReportService()

	// Run in a goroutine to avoid blocking
	go scanner.RunApp(om, rs)
	
	// Wait for output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	w.Close()

	assert.Contains(t, buf.String(), "Added Laptop to cart", "Should confirm product addition")
}

func TestConsoleUI_RunApp_Checkout(t *testing.T) {
	// Mock input for checkout
	input := strings.Join([]string{
		"3\n", // Select "Process Payment"
		"2\n", // Payment Type: PayPal
		"5\n", // Exit
	}, "")
	reader := strings.NewReader(input)
	scanner := ui.NewConsoleUI()
	scanner.Scanner = bufio.NewScanner(reader)
	
	// Capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Setup dependencies with a product in cart
	om := ordermanager.NewOrderManager(&models.Product{}, &models.Product{})
	rs := reports.NewReportService()
	om.AddToCart(models.NewProduct("P001", "Test", 100.0, 10), 1)

	// Run in a goroutine
	go scanner.RunApp(om, rs)
	
	var buf bytes.Buffer
	io.Copy(&buf, r)
	w.Close()

	assert.Contains(t, buf.String(), "PayPal: Processed $100.00", "Should process PayPal payment")
}
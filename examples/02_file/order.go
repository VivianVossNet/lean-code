// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Process an order into an invoice.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	ID       int64
	Product  string
	Quantity int
}

type Invoice struct {
	OrderID  int64
	Product  string
	Quantity int
	Total    float64
}

func ReadOrderFromFile(path string) Order {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "order file not found: %s\n", path)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	id, _ := strconv.ParseInt(lines[0], 10, 64)
	quantity, _ := strconv.Atoi(lines[2])

	return Order{ID: id, Product: lines[1], Quantity: quantity}
}

func ReadStockForProduct(product string) int {
	data, err := os.ReadFile("stock.txt")
	if err != nil {
		return 0
	}

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.SplitN(line, ",", 2)
		if parts[0] == product {
			count, _ := strconv.Atoi(parts[1])
			return count
		}
	}

	return 0
}

func CheckStockAvailability(product string, quantity int) bool {
	stock := ReadStockForProduct(product)
	return stock >= quantity
}

func CreateInvoiceFromOrder(order Order, pricePerUnit float64) Invoice {
	return Invoice{
		OrderID:  order.ID,
		Product:  order.Product,
		Quantity: order.Quantity,
		Total:    float64(order.Quantity) * pricePerUnit,
	}
}

func WriteInvoiceToFile(invoice Invoice, path string) {
	content := fmt.Sprintf("%d\n%s\n%d\n%.2f",
		invoice.OrderID, invoice.Product, invoice.Quantity, invoice.Total)

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write invoice: %s\n", path)
		os.Exit(1)
	}
}

func main() {
	order := ReadOrderFromFile("order.txt")

	if !CheckStockAvailability(order.Product, order.Quantity) {
		fmt.Fprintf(os.Stderr, "out of stock: %s\n", order.Product)
		os.Exit(1)
	}

	invoice := CreateInvoiceFromOrder(order, 29.99)
	WriteInvoiceToFile(invoice, "invoice.txt")
}

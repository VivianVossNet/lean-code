// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Create and write invoices.

package billing

import (
	"fmt"
	"os"
)

type Invoice struct {
	OrderID  int64
	Product  string
	Quantity int
	Total    float64
}

func CreateInvoiceFromOrder(orderID int64, product string, quantity int, pricePerUnit float64) Invoice {
	return Invoice{
		OrderID:  orderID,
		Product:  product,
		Quantity: quantity,
		Total:    float64(quantity) * pricePerUnit,
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

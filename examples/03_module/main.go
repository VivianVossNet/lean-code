// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Orchestrate order to invoice.

package main

import (
	"fmt"
	"os"

	"lean-code/billing"
	"lean-code/order"
	"lean-code/stock"
)

func main() {
	o := order.ReadOrderFromFile("order.txt")

	if !stock.CheckStockAvailability(o.Product, o.Quantity) {
		fmt.Fprintf(os.Stderr, "out of stock: %s\n", o.Product)
		os.Exit(1)
	}

	invoice := billing.CreateInvoiceFromOrder(o.ID, o.Product, o.Quantity, 29.99)
	billing.WriteInvoiceToFile(invoice, "invoice.txt")
}

// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Process all orders into invoices.

package main

import (
	"fmt"
	"os"
	"strconv"

	"lean-code/billing"
	"lean-code/configuration"
	"lean-code/order"
	"lean-code/stock"
)

func main() {
	config := configuration.ReadConfigurationFromFile("config.txt")
	pricePerUnit, _ := strconv.ParseFloat(config.PricePerUnit, 64)

	orderFiles := order.FindOrderFilesInDirectory(config.OrderDirectory)

	for _, file := range orderFiles {
		o := order.ReadOrderFromFile(file)

		if !stock.CheckStockAvailability(o.Product, o.Quantity, config.StockFile) {
			fmt.Fprintf(os.Stderr, "out of stock: %s\n", o.Product)
			continue
		}

		invoice := billing.CreateInvoiceFromOrder(o.ID, o.Product, o.Quantity, pricePerUnit)
		billing.WriteInvoiceToFile(invoice, fmt.Sprintf("invoice_%d.txt", o.ID))
	}
}

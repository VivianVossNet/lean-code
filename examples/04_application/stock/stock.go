// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read and check stock availability.

package stock

import (
	"os"
	"strconv"
	"strings"
)

func ReadStockForProduct(product string, stockPath string) int {
	data, err := os.ReadFile(stockPath)
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

func CheckStockAvailability(product string, quantity int, stockPath string) bool {
	s := ReadStockForProduct(product, stockPath)
	return s >= quantity
}

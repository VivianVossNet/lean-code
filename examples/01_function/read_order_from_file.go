// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read an order from a text file.

package order

import (
	"os"
	"strconv"
	"strings"
)

type Order struct {
	ID       int64
	Product  string
	Quantity int
}

func ReadOrderFromFile(path string) Order {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("order file not found: " + path)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	id, _ := strconv.ParseInt(lines[0], 10, 64)
	quantity, _ := strconv.Atoi(lines[2])

	return Order{ID: id, Product: lines[1], Quantity: quantity}
}

// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read orders from file and find order files.

package order

import (
	"os"
	"path/filepath"
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

func FindOrderFilesInDirectory(directory string) []string {
	matches, _ := filepath.Glob(filepath.Join(directory, "order_*.txt"))
	return matches
}

// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read application configuration.

package configuration

import (
	"os"
	"strings"
)

type Configuration struct {
	PricePerUnit   string
	StockFile      string
	OrderDirectory string
}

func ReadConfigurationFromFile(path string) Configuration {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("configuration file not found: " + path)
	}

	values := map[string]string{}

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.SplitN(line, "=", 2)
		values[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return Configuration{
		PricePerUnit:   values["price_per_unit"],
		StockFile:      values["stock_file"],
		OrderDirectory: values["order_directory"],
	}
}

// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read and check stock availability.

import { readFileSync } from "node:fs";

function readStockForProduct(product) {
  const lines = readFileSync("stock.txt", "utf-8").trim().split("\n");

  for (const line of lines) {
    const [name, count] = line.split(",", 2);
    if (name === product) {
      return parseInt(count, 10);
    }
  }

  return 0;
}

function checkStockAvailability(product, quantity) {
  const stock = readStockForProduct(product);
  return stock >= quantity;
}

export { readStockForProduct, checkStockAvailability };

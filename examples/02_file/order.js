// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Process an order into an invoice.

import { readFileSync, writeFileSync } from "node:fs";

function readOrderFromFile(path) {
  const lines = readFileSync(path, "utf-8").trim().split("\n");

  return {
    id: parseInt(lines[0], 10),
    product: lines[1],
    quantity: parseInt(lines[2], 10),
  };
}

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

function createInvoiceFromOrder(order, pricePerUnit) {
  return {
    orderId: order.id,
    product: order.product,
    quantity: order.quantity,
    total: order.quantity * pricePerUnit,
  };
}

function writeInvoiceToFile(invoice, path) {
  const content = [
    invoice.orderId,
    invoice.product,
    invoice.quantity,
    invoice.total.toFixed(2),
  ].join("\n");

  writeFileSync(path, content);
}

const order = readOrderFromFile("order.txt");

if (!checkStockAvailability(order.product, order.quantity)) {
  process.stderr.write(`out of stock: ${order.product}\n`);
  process.exit(1);
}

const invoice = createInvoiceFromOrder(order, 29.99);
writeInvoiceToFile(invoice, "invoice.txt");

// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Create and write invoices.

import { writeFileSync } from "node:fs";

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

export { createInvoiceFromOrder, writeInvoiceToFile };

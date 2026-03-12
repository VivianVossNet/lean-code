// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Orchestrate order to invoice.

import { readOrderFromFile } from "./order/order.js";
import { checkStockAvailability } from "./stock/stock.js";
import { createInvoiceFromOrder, writeInvoiceToFile } from "./billing/billing.js";

const order = readOrderFromFile("order.txt");

if (!checkStockAvailability(order.product, order.quantity)) {
  process.stderr.write(`out of stock: ${order.product}\n`);
  process.exit(1);
}

const invoice = createInvoiceFromOrder(order, 29.99);
writeInvoiceToFile(invoice, "invoice.txt");

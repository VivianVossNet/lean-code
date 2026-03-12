// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Process all orders into invoices.

import { readConfigurationFromFile } from "./configuration/configuration.js";
import { readOrderFromFile, findOrderFilesInDirectory } from "./order/order.js";
import { checkStockAvailability } from "./stock/stock.js";
import { createInvoiceFromOrder, writeInvoiceToFile } from "./billing/billing.js";

const configuration = readConfigurationFromFile("config.txt");
const pricePerUnit = parseFloat(configuration["price_per_unit"]);

const orderFiles = findOrderFilesInDirectory(configuration["order_directory"]);

for (const file of orderFiles) {
  const order = readOrderFromFile(file);

  if (!checkStockAvailability(order.product, order.quantity, configuration["stock_file"])) {
    process.stderr.write(`out of stock: ${order.product}\n`);
    continue;
  }

  const invoice = createInvoiceFromOrder(order, pricePerUnit);
  writeInvoiceToFile(invoice, `invoice_${order.id}.txt`);
}

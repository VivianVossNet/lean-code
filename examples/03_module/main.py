# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Entry point. Orchestrate order to invoice.

import sys
from order.order import read_order_from_file
from stock.stock import check_stock_availability
from billing.billing import create_invoice_from_order, write_invoice_to_file

order = read_order_from_file("order.txt")

if not check_stock_availability(order["product"], order["quantity"]):
    sys.exit(f"out of stock: {order['product']}")

invoice = create_invoice_from_order(order, 29.99)
write_invoice_to_file(invoice, "invoice.txt")

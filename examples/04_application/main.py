# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Entry point. Process all orders into invoices.

import sys
from configuration.configuration import read_configuration_from_file
from order.order import read_order_from_file, find_order_files_in_directory
from stock.stock import check_stock_availability
from billing.billing import create_invoice_from_order, write_invoice_to_file

configuration = read_configuration_from_file("config.txt")
price_per_unit = float(configuration["price_per_unit"])

order_files = find_order_files_in_directory(configuration["order_directory"])

for file in order_files:
    order = read_order_from_file(file)

    if not check_stock_availability(order["product"], order["quantity"], configuration["stock_file"]):
        print(f"out of stock: {order['product']}", file=sys.stderr)
        continue

    invoice = create_invoice_from_order(order, price_per_unit)
    write_invoice_to_file(invoice, f"invoice_{order['id']}.txt")
